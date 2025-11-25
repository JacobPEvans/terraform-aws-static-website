package test

import (
	"fmt"
	"os"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestTerraformAwsStaticWebsite(t *testing.T) {
	t.Parallel()

	// Set up environment for LocalStack
	os.Setenv("AWS_ACCESS_KEY_ID", "test")
	os.Setenv("AWS_SECRET_ACCESS_KEY", "test")
	os.Setenv("AWS_DEFAULT_REGION", "us-east-1")

	// Construct the terraform options with default retryable errors to handle the most common
	// retryable errors in terraform testing.
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		// The path to where our Terraform code is located
		TerraformDir: "../examples/basic",

		// Variables to pass to our Terraform code using -var options
		Vars: map[string]interface{}{
			"website-domain-main":     "test.example.com",
			"website-domain-redirect": "www.test.example.com",
			"domains-zone-root":       "example.com",
			"support-spa":             false,
		},

		// Environment variables to set when running Terraform
		EnvVars: map[string]string{
			"AWS_ENDPOINT_URL": "http://localhost:4566",
		},

		// Disable colors in Terraform commands so its easier to parse stdout/stderr
		NoColor: true,
	})

	// At the end of the test, run `terraform destroy` to clean up any resources that were created
	defer terraform.Destroy(t, terraformOptions)

	// Run `terraform init` and `terraform apply`. Fail the test if there are any errors.
	terraform.InitAndApply(t, terraformOptions)

	// Run `terraform output` to get the values of output variables
	websiteCdnRootId := terraform.Output(t, terraformOptions, "website_cdn_root_id")
	websiteRootBucket := terraform.Output(t, terraformOptions, "website_root_s3_bucket")
	websiteLogsBucket := terraform.Output(t, terraformOptions, "website_logs_s3_bucket")
	websiteRedirectBucket := terraform.Output(t, terraformOptions, "website_redirect_s3_bucket")

	// Verify we're getting back the outputs we expect
	assert.NotEmpty(t, websiteCdnRootId, "CloudFront distribution ID should not be empty")
	assert.Equal(t, "test.example.com-root", websiteRootBucket)
	assert.Equal(t, "test.example.com-logs", websiteLogsBucket)
	assert.Equal(t, "test.example.com-redirect", websiteRedirectBucket)

	fmt.Println("All tests passed!")
}

func TestTerraformAwsStaticWebsiteSPA(t *testing.T) {
	t.Parallel()

	// Set up environment for LocalStack
	os.Setenv("AWS_ACCESS_KEY_ID", "test")
	os.Setenv("AWS_SECRET_ACCESS_KEY", "test")
	os.Setenv("AWS_DEFAULT_REGION", "us-east-1")

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../examples/spa",

		Vars: map[string]interface{}{
			"website-domain-main": "spa.example.com",
			"domains-zone-root":   "example.com",
			"support-spa":         true,
		},

		EnvVars: map[string]string{
			"AWS_ENDPOINT_URL": "http://localhost:4566",
		},

		NoColor: true,
	})

	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	websiteCdnRootId := terraform.Output(t, terraformOptions, "website_cdn_root_id")
	websiteRootBucket := terraform.Output(t, terraformOptions, "website_root_s3_bucket")

	assert.NotEmpty(t, websiteCdnRootId)
	assert.Equal(t, "spa.example.com-root", websiteRootBucket)

	fmt.Println("SPA tests passed!")
}
