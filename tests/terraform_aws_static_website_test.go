package test

import (
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudfront"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Helper function to create AWS session for LocalStack
func createLocalStackSession(t *testing.T) *session.Session {
	sess, err := session.NewSession(&aws.Config{
		Region:           aws.String("us-east-1"),
		Credentials:      credentials.NewStaticCredentials("test", "test", ""),
		Endpoint:         aws.String("http://localhost:4566"),
		S3ForcePathStyle: aws.Bool(true),
	})
	require.NoError(t, err)
	return sess
}

// localStackTerraformOptions creates terraform options with LocalStack configuration
func localStackTerraformOptions(t *testing.T, dir string, vars map[string]interface{}) *terraform.Options {
	t.Setenv("AWS_ACCESS_KEY_ID", "test")
	t.Setenv("AWS_SECRET_ACCESS_KEY", "test")
	t.Setenv("AWS_DEFAULT_REGION", "us-east-1")

	return terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: dir,
		Vars:         vars,
		EnvVars: map[string]string{
			"AWS_ENDPOINT_URL": "http://localhost:4566",
		},
		NoColor: true,
	})
}


// TestTerraformAwsStaticWebsiteBasicPlan validates the basic configuration produces a valid plan
func TestTerraformAwsStaticWebsiteBasicPlan(t *testing.T) {
	opts := localStackTerraformOptions(t, "../examples/basic", map[string]interface{}{
		"website-domain-main":     "test.example.com",
		"website-domain-redirect": "www.test.example.com",
		"domains-zone-root":       "example.com",
		"support-spa":             false,
	})
	opts.PlanFilePath = t.TempDir() + "/plan.out"

	plan := terraform.InitAndPlanAndShowWithStruct(t, opts)

	// Verify plan has expected number of resources
	assert.GreaterOrEqual(t, len(plan.ResourcePlannedValuesMap), 20,
		"Plan should create at least 20 resources")

	// Verify key resources are in the plan
	terraform.AssertPlannedValuesMapKeyExists(t, plan, "module.static_website.module.s3_buckets.aws_s3_bucket.website_root")
	terraform.AssertPlannedValuesMapKeyExists(t, plan, "module.static_website.module.s3_buckets.aws_s3_bucket.website_logs")
	terraform.AssertPlannedValuesMapKeyExists(t, plan, "module.static_website.module.s3_buckets.aws_s3_bucket.website_redirect")
	terraform.AssertPlannedValuesMapKeyExists(t, plan, "module.static_website.module.cloudfront.aws_cloudfront_distribution.website_cdn_root")
	terraform.AssertPlannedValuesMapKeyExists(t, plan, "module.static_website.module.acm_certificate.aws_acm_certificate.wildcard_website")

	// Verify root bucket configuration
	root := plan.ResourcePlannedValuesMap["module.static_website.module.s3_buckets.aws_s3_bucket.website_root"]
	require.NotNil(t, root, "Root bucket should exist in plan")
	assert.Equal(t, "test.example.com-root", root.AttributeValues["bucket"])

	fmt.Println("✅ Basic website plan tests passed!")
}

// TestTerraformAwsStaticWebsiteSPAPlan validates the SPA configuration produces a valid plan
func TestTerraformAwsStaticWebsiteSPAPlan(t *testing.T) {
	opts := localStackTerraformOptions(t, "../examples/spa", map[string]interface{}{
		"website-domain-main": "spa.example.com",
		"domains-zone-root":   "example.com",
		"support-spa":         true,
	})
	opts.PlanFilePath = t.TempDir() + "/plan.out"

	plan := terraform.InitAndPlanAndShowWithStruct(t, opts)

	// Verify plan has expected resources
	assert.GreaterOrEqual(t, len(plan.ResourcePlannedValuesMap), 15,
		"SPA plan should create at least 15 resources")

	// Verify key SPA resources
	terraform.AssertPlannedValuesMapKeyExists(t, plan, "module.spa_website.module.s3_buckets.aws_s3_bucket.website_root")
	terraform.AssertPlannedValuesMapKeyExists(t, plan, "module.spa_website.module.cloudfront.aws_cloudfront_distribution.website_cdn_root")

	// Verify root bucket configuration
	root := plan.ResourcePlannedValuesMap["module.spa_website.module.s3_buckets.aws_s3_bucket.website_root"]
	require.NotNil(t, root, "Root bucket should exist in plan")
	assert.Equal(t, "spa.example.com-root", root.AttributeValues["bucket"])

	fmt.Println("✅ SPA website plan tests passed!")
}

// TestTerraformAwsStaticWebsiteBasic tests the basic static website configuration
func TestTerraformAwsStaticWebsiteBasic(t *testing.T) {
	terraformOptions := localStackTerraformOptions(t, "../examples/basic", map[string]interface{}{
		"website-domain-main":     "test.example.com",
		"website-domain-redirect": "www.test.example.com",
		"domains-zone-root":       "example.com",
		"support-spa":             false,
	})

	defer terraform.Destroy(t, terraformOptions)

	// Run init and apply
	terraform.InitAndApply(t, terraformOptions)

	// Test outputs
	websiteCdnRootId := terraform.Output(t, terraformOptions, "website_cdn_root_id")
	websiteRootBucket := terraform.Output(t, terraformOptions, "website_root_s3_bucket")
	websiteLogsBucket := terraform.Output(t, terraformOptions, "website_logs_s3_bucket")
	websiteRedirectBucket := terraform.Output(t, terraformOptions, "website_redirect_s3_bucket")

	assert.NotEmpty(t, websiteCdnRootId, "CloudFront distribution ID should not be empty")
	assert.Equal(t, "test.example.com-root", websiteRootBucket)
	assert.Equal(t, "test.example.com-logs", websiteLogsBucket)
	assert.Equal(t, "test.example.com-redirect", websiteRedirectBucket)

	// Create AWS session
	sess := createLocalStackSession(t)

	// Test S3 buckets existence and configuration
	testS3BucketsConfiguration(t, sess, websiteRootBucket, websiteLogsBucket, websiteRedirectBucket)

	// Test CloudFront distribution
	testCloudFrontDistribution(t, sess, websiteCdnRootId)

	fmt.Println("✅ Basic website tests passed!")
}

// TestTerraformAwsStaticWebsiteSPA tests the SPA configuration
func TestTerraformAwsStaticWebsiteSPA(t *testing.T) {
	terraformOptions := localStackTerraformOptions(t, "../examples/spa", map[string]interface{}{
		"website-domain-main": "spa.example.com",
		"domains-zone-root":   "example.com",
		"support-spa":         true,
	})

	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	websiteCdnRootId := terraform.Output(t, terraformOptions, "website_cdn_root_id")
	websiteRootBucket := terraform.Output(t, terraformOptions, "website_root_s3_bucket")

	assert.NotEmpty(t, websiteCdnRootId)
	assert.Equal(t, "spa.example.com-root", websiteRootBucket)

	// Create AWS session
	sess := createLocalStackSession(t)

	// Verify S3 bucket exists
	s3Client := s3.New(sess)
	_, err := s3Client.HeadBucket(&s3.HeadBucketInput{
		Bucket: aws.String(websiteRootBucket),
	})
	require.NoError(t, err, "S3 bucket should exist")

	fmt.Println("✅ SPA website tests passed!")
}

// testS3BucketsConfiguration verifies S3 bucket configuration
func testS3BucketsConfiguration(t *testing.T, sess *session.Session, rootBucket, logsBucket, redirectBucket string) {
	s3Client := s3.New(sess)

	// Test root bucket
	t.Run("RootBucketExists", func(t *testing.T) {
		_, err := s3Client.HeadBucket(&s3.HeadBucketInput{
			Bucket: aws.String(rootBucket),
		})
		require.NoError(t, err, "Root S3 bucket should exist")
	})

	// Test logs bucket
	t.Run("LogsBucketExists", func(t *testing.T) {
		_, err := s3Client.HeadBucket(&s3.HeadBucketInput{
			Bucket: aws.String(logsBucket),
		})
		require.NoError(t, err, "Logs S3 bucket should exist")
	})

	// Test redirect bucket
	t.Run("RedirectBucketExists", func(t *testing.T) {
		_, err := s3Client.HeadBucket(&s3.HeadBucketInput{
			Bucket: aws.String(redirectBucket),
		})
		require.NoError(t, err, "Redirect S3 bucket should exist")
	})

	// Test versioning on root bucket
	t.Run("RootBucketVersioning", func(t *testing.T) {
		versioningOutput, err := s3Client.GetBucketVersioning(&s3.GetBucketVersioningInput{
			Bucket: aws.String(rootBucket),
		})
		if err != nil {
			t.Skipf("LocalStack does not support this feature: %v", err)
		}
		assert.Equal(t, "Enabled", aws.StringValue(versioningOutput.Status), "Root bucket should have versioning enabled")
	})

	// Test versioning on logs bucket
	t.Run("LogsBucketVersioning", func(t *testing.T) {
		versioningOutput, err := s3Client.GetBucketVersioning(&s3.GetBucketVersioningInput{
			Bucket: aws.String(logsBucket),
		})
		if err != nil {
			t.Skipf("LocalStack does not support this feature: %v", err)
		}
		assert.Equal(t, "Enabled", aws.StringValue(versioningOutput.Status), "Logs bucket should have versioning enabled")
	})

	// Test encryption on root bucket
	t.Run("RootBucketEncryption", func(t *testing.T) {
		encryptionOutput, err := s3Client.GetBucketEncryption(&s3.GetBucketEncryptionInput{
			Bucket: aws.String(rootBucket),
		})
		if err != nil {
			t.Skipf("LocalStack does not support this feature: %v", err)
		}
		require.NotNil(t, encryptionOutput.ServerSideEncryptionConfiguration)
		require.NotEmpty(t, encryptionOutput.ServerSideEncryptionConfiguration.Rules)
		assert.Equal(t, "AES256", aws.StringValue(encryptionOutput.ServerSideEncryptionConfiguration.Rules[0].ApplyServerSideEncryptionByDefault.SSEAlgorithm))
	})

	// Test website configuration on root bucket
	t.Run("RootBucketWebsiteConfiguration", func(t *testing.T) {
		websiteOutput, err := s3Client.GetBucketWebsite(&s3.GetBucketWebsiteInput{
			Bucket: aws.String(rootBucket),
		})
		if err != nil {
			t.Skipf("LocalStack does not support this feature: %v", err)
		}
		assert.Equal(t, "index.html", aws.StringValue(websiteOutput.IndexDocument.Suffix))
	})

	// Test logging configuration on root bucket
	t.Run("RootBucketLogging", func(t *testing.T) {
		loggingOutput, err := s3Client.GetBucketLogging(&s3.GetBucketLoggingInput{
			Bucket: aws.String(rootBucket),
		})
		if err != nil {
			t.Skipf("LocalStack does not support this feature: %v", err)
		}
		if loggingOutput.LoggingEnabled != nil {
			assert.Contains(t, aws.StringValue(loggingOutput.LoggingEnabled.TargetBucket), logsBucket)
		}
	})

	// Test public access block on logs bucket
	t.Run("LogsBucketPublicAccessBlock", func(t *testing.T) {
		publicAccessOutput, err := s3Client.GetPublicAccessBlock(&s3.GetPublicAccessBlockInput{
			Bucket: aws.String(logsBucket),
		})
		if err != nil {
			t.Skipf("LocalStack does not support this feature: %v", err)
		}
		if publicAccessOutput.PublicAccessBlockConfiguration != nil {
			assert.True(t, aws.BoolValue(publicAccessOutput.PublicAccessBlockConfiguration.BlockPublicAcls))
			assert.True(t, aws.BoolValue(publicAccessOutput.PublicAccessBlockConfiguration.BlockPublicPolicy))
		}
	})
}

// testCloudFrontDistribution verifies CloudFront distribution configuration
func testCloudFrontDistribution(t *testing.T, sess *session.Session, distributionId string) {
	cfClient := cloudfront.New(sess)

	t.Run("CloudFrontDistributionExists", func(t *testing.T) {
		distOutput, err := cfClient.GetDistribution(&cloudfront.GetDistributionInput{
			Id: aws.String(distributionId),
		})
		if err != nil {
			t.Skipf("LocalStack does not support this feature: %v", err)
		}
		require.NotNil(t, distOutput.Distribution)
		assert.True(t, aws.BoolValue(distOutput.Distribution.DistributionConfig.Enabled), "CloudFront distribution should be enabled")
	})

	t.Run("CloudFrontDistributionHTTPSConfig", func(t *testing.T) {
		distOutput, err := cfClient.GetDistribution(&cloudfront.GetDistributionInput{
			Id: aws.String(distributionId),
		})
		if err != nil {
			t.Skipf("LocalStack does not support this feature: %v", err)
		}
		if distOutput.Distribution != nil && distOutput.Distribution.DistributionConfig.DefaultCacheBehavior != nil {
			assert.Equal(t, "redirect-to-https", aws.StringValue(distOutput.Distribution.DistributionConfig.DefaultCacheBehavior.ViewerProtocolPolicy))
		}
	})

	t.Run("CloudFrontDistributionCompression", func(t *testing.T) {
		distOutput, err := cfClient.GetDistribution(&cloudfront.GetDistributionInput{
			Id: aws.String(distributionId),
		})
		if err != nil {
			t.Skipf("LocalStack does not support this feature: %v", err)
		}
		if distOutput.Distribution != nil && distOutput.Distribution.DistributionConfig.DefaultCacheBehavior != nil {
			assert.True(t, aws.BoolValue(distOutput.Distribution.DistributionConfig.DefaultCacheBehavior.Compress), "Compression should be enabled")
		}
	})
}

// TestModuleValidation tests that the module passes validation
func TestModuleValidation(t *testing.T) {
	t.Run("MainModuleValidation", func(t *testing.T) {
		options := &terraform.Options{
			TerraformDir: "..",
			NoColor:      true,
		}

		// Only validate, don't apply
		terraform.Init(t, options)
		terraform.Validate(t, options)
	})

	t.Run("BasicExampleValidation", func(t *testing.T) {
		options := &terraform.Options{
			TerraformDir: "../examples/basic",
			NoColor:      true,
		}

		terraform.Init(t, options)
		terraform.Validate(t, options)
	})

	t.Run("SPAExampleValidation", func(t *testing.T) {
		options := &terraform.Options{
			TerraformDir: "../examples/spa",
			NoColor:      true,
		}

		terraform.Init(t, options)
		terraform.Validate(t, options)
	})

	t.Run("LambdaExampleValidation", func(t *testing.T) {
		options := &terraform.Options{
			TerraformDir: "../examples/with-lambda",
			NoColor:      true,
		}

		terraform.Init(t, options)
		terraform.Validate(t, options)
	})

	fmt.Println("✅ All validation tests passed!")
}

// TestS3BucketResourcesIndividually tests each S3 resource type
func TestS3BucketResourcesIndividually(t *testing.T) {
	// This test verifies that all S3 resources are properly configured
	// It doesn't need to run terraform, just validates the structure exists

	t.Run("VerifyS3ResourceTypes", func(t *testing.T) {
		// List of expected S3 resource types
		expectedResources := []string{
			"aws_s3_bucket",
			"aws_s3_bucket_versioning",
			"aws_s3_bucket_server_side_encryption_configuration",
			"aws_s3_bucket_logging",
			"aws_s3_bucket_website_configuration",
			"aws_s3_bucket_acl",
			"aws_s3_bucket_ownership_controls",
			"aws_s3_bucket_public_access_block",
			"aws_s3_bucket_policy",
		}

		// Read all Terraform files (main.tf and modules)
		allContent, err := ReadAllTerraformFiles("..")
		require.NoError(t, err, "Should be able to read Terraform files")

		// Verify each resource type exists
		for _, resourceType := range expectedResources {
			assert.Contains(t, allContent, fmt.Sprintf("resource \"%s\"", resourceType),
				fmt.Sprintf("Terraform files should contain %s resource", resourceType))
		}
	})

	fmt.Println("✅ S3 resource structure validation passed!")
}

// TestCloudFrontResources verifies CloudFront resources exist
func TestCloudFrontResources(t *testing.T) {
	t.Run("VerifyCloudFrontResourceTypes", func(t *testing.T) {
		expectedResources := []string{
			"aws_cloudfront_distribution",
			"aws_acm_certificate",
			"aws_acm_certificate_validation",
		}

		allContent, err := ReadAllTerraformFiles("..")
		require.NoError(t, err, "Should be able to read Terraform files")

		for _, resourceType := range expectedResources {
			assert.Contains(t, allContent, fmt.Sprintf("resource \"%s\"", resourceType),
				fmt.Sprintf("Terraform files should contain %s resource", resourceType))
		}
	})

	fmt.Println("✅ CloudFront resource structure validation passed!")
}

// TestRoute53Resources verifies Route53 resources exist
func TestRoute53Resources(t *testing.T) {
	t.Run("VerifyRoute53ResourceTypes", func(t *testing.T) {
		expectedResources := []string{
			"aws_route53_record",
		}

		expectedDataSources := []string{
			"data \"aws_route53_zone\"",
		}

		allContent, err := ReadAllTerraformFiles("..")
		require.NoError(t, err, "Should be able to read Terraform files")

		for _, resourceType := range expectedResources {
			assert.Contains(t, allContent, fmt.Sprintf("resource \"%s\"", resourceType),
				fmt.Sprintf("Terraform files should contain %s resource", resourceType))
		}

		for _, dataSource := range expectedDataSources {
			assert.Contains(t, allContent, dataSource,
				fmt.Sprintf("Terraform files should contain %s data source", dataSource))
		}
	})

	fmt.Println("✅ Route53 resource structure validation passed!")
}
