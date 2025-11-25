# Testing Guide

This document provides comprehensive information about testing the terraform-aws-static-website module.

## Table of Contents

- [Prerequisites](#prerequisites)
- [Test Types](#test-types)
- [Quick Start](#quick-start)
- [Unit Tests](#unit-tests)
- [Integration Tests](#integration-tests)
- [LocalStack Setup](#localstack-setup)
- [Continuous Integration](#continuous-integration)
- [Troubleshooting](#troubleshooting)

## Prerequisites

### Required Tools

- **Terraform** >= 1.5.0
- **Go** >= 1.21 (for running tests)
- **Docker** (for LocalStack integration tests)
- **Docker Compose** (for LocalStack orchestration)
- **Make** (optional, for convenient commands)

### Installation

```bash
# Install Terraform
brew install terraform  # macOS
# or download from https://www.terraform.io/downloads

# Install Go
brew install go  # macOS
# or download from https://golang.org/dl/

# Install Docker
# Follow: https://docs.docker.com/get-docker/

# Install Docker Compose
# Follow: https://docs.docker.com/compose/install/
```

## Test Types

### 1. Unit Tests (Fast, No Infrastructure)

Unit tests validate the module structure and configuration without deploying any infrastructure.

**What they test:**
- Terraform syntax and structure
- Presence of required AWS resource types
- Module validation
- Example configurations

**Runtime:** ~2-5 minutes

### 2. Integration Tests (Slow, Requires LocalStack)

Integration tests deploy actual infrastructure to LocalStack and verify resource creation and configuration.

**What they test:**
- S3 bucket creation and configuration
- S3 versioning, encryption, logging
- CloudFront distribution setup
- ACM certificate management
- Route53 DNS records
- Full end-to-end workflows

**Runtime:** ~10-30 minutes

## Quick Start

### Using Make (Recommended)

```bash
# Run unit tests only (fast)
make test-unit

# Start LocalStack and run integration tests
make test-local

# Run integration tests (assumes LocalStack is running)
make test

# Manual LocalStack management
make localstack-start
make test
make localstack-stop
```

### Using Go Directly

```bash
# Run all unit tests
cd tests
go test -v -run "TestS3BucketResourcesIndividually|TestCloudFrontResources|TestRoute53Resources|TestModuleValidation" -timeout 10m

# Run all integration tests
go test -v -run "TestTerraformAwsStaticWebsiteBasic|TestTerraformAwsStaticWebsiteSPA" -timeout 30m

# Run specific test
go test -v -run TestS3BucketResourcesIndividually

# Run all tests
go test -v -timeout 30m
```

## Unit Tests

Unit tests validate the module without deploying infrastructure.

### Running Unit Tests

```bash
make test-unit
```

Or manually:

```bash
cd tests
go test -v -run TestModuleValidation -timeout 10m
```

### Unit Test Coverage

1. **TestS3BucketResourcesIndividually**
   - Validates presence of all S3 resource types
   - Checks for aws_s3_bucket
   - Checks for aws_s3_bucket_versioning
   - Checks for aws_s3_bucket_encryption
   - Checks for aws_s3_bucket_logging
   - Checks for aws_s3_bucket_website_configuration
   - Checks for aws_s3_bucket_acl
   - Checks for aws_s3_bucket_ownership_controls
   - Checks for aws_s3_bucket_public_access_block

2. **TestCloudFrontResources**
   - Validates CloudFront distribution resources
   - Validates ACM certificate resources
   - Validates certificate validation resources

3. **TestRoute53Resources**
   - Validates Route53 record resources
   - Validates Route53 data sources

4. **TestModuleValidation**
   - Validates main module
   - Validates all example configurations
   - Ensures terraform init and validate pass

## Integration Tests

Integration tests deploy infrastructure to LocalStack.

### Running Integration Tests

```bash
# Start LocalStack
make localstack-start

# Run tests
make test

# Stop LocalStack
make localstack-stop
```

Or use the combined command:

```bash
make test-local
```

### Integration Test Coverage

1. **TestTerraformAwsStaticWebsiteBasic**
   - Deploys full basic website configuration
   - Validates all outputs
   - Checks S3 bucket existence
   - Validates S3 versioning
   - Validates S3 encryption (AES256)
   - Validates S3 website configuration
   - Validates S3 logging configuration
   - Validates S3 public access blocks
   - Validates CloudFront distribution
   - Validates HTTPS configuration
   - Validates compression settings

2. **TestTerraformAwsStaticWebsiteSPA**
   - Deploys SPA website configuration
   - Validates SPA-specific settings
   - Checks 404 to index.html redirect
   - Validates all outputs

## LocalStack Setup

### Starting LocalStack

```bash
make localstack-start
```

Or manually:

```bash
docker-compose up -d
```

### Checking LocalStack Health

```bash
# Check container status
docker-compose ps

# View logs
make localstack-logs

# Or
docker-compose logs -f localstack

# Check health endpoint
curl http://localhost:4566/_localstack/health
```

### LocalStack Services

The module requires these LocalStack services:
- S3
- CloudFront
- Route53
- ACM (Certificate Manager)
- STS

### Stopping LocalStack

```bash
make localstack-stop
```

Or manually:

```bash
docker-compose down
```

## Continuous Integration

### GitHub Actions

The CI pipeline runs automatically on:
- Pull requests
- Pushes to main/master branches

### CI Jobs

1. **terraform-validate**: Terraform formatting and validation
2. **tflint**: Static analysis with TFLint
3. **trivy**: Security scanning
4. **unit-tests**: Fast validation tests
5. **localstack-tests**: Integration tests (continue-on-error)

### Running CI Locally

```bash
# Format check
terraform fmt -check -recursive

# Validate
terraform init -backend=false
terraform validate

# Lint
tflint --init
tflint --recursive

# Run tests
make test-unit
```

## Troubleshooting

### Common Issues

#### Issue: Go tests fail with "cannot find package"

**Solution:**
```bash
cd tests
go mod download
go mod tidy
```

#### Issue: LocalStack not starting

**Solution:**
```bash
# Check Docker is running
docker ps

# Check logs
docker-compose logs localstack

# Restart LocalStack
docker-compose down -v
docker-compose up -d
```

#### Issue: LocalStack tests timeout

**Solution:**
LocalStack tests can be slow. Increase timeout:
```bash
go test -v -timeout 60m
```

Or run unit tests only:
```bash
make test-unit
```

#### Issue: Terraform init fails in tests

**Solution:**
```bash
# Clean terraform cache
rm -rf .terraform .terraform.lock.hcl
rm -rf examples/*/.terraform
rm -rf examples/*/.terraform.lock.hcl

# Try again
make test-unit
```

#### Issue: Tests pass locally but fail in CI

**Solution:**
- Check GitHub Actions logs for specific errors
- LocalStack tests are marked as `continue-on-error` in CI
- Unit tests should always pass in CI

### Debug Mode

Enable verbose logging:

```bash
# Go tests
go test -v -run TestName

# Terraform
TF_LOG=DEBUG terraform plan

# LocalStack
# Edit docker-compose.yml and set DEBUG=1
```

## Test Development

### Adding New Tests

1. Create test function in `tests/terraform_aws_static_website_test.go`
2. Follow naming convention: `TestXxx`
3. Use descriptive subtests: `t.Run("TestName", func(t *testing.T) { ... })`
4. Add assertions with `assert` or `require` from testify

### Test Best Practices

- Keep tests independent (no shared state)
- Use `t.Parallel()` when possible
- Clean up resources with `defer`
- Use meaningful assertion messages
- Test both success and failure cases
- Document complex test logic

### Example Test Structure

```go
func TestMyNewFeature(t *testing.T) {
    t.Parallel()

    terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
        TerraformDir: "../examples/basic",
        Vars: map[string]interface{}{
            "my-var": "value",
        },
    })

    defer terraform.Destroy(t, terraformOptions)

    terraform.InitAndApply(t, terraformOptions)

    // Test outputs
    output := terraform.Output(t, terraformOptions, "my_output")
    assert.Equal(t, "expected", output)

    // Test infrastructure
    // ... AWS SDK calls ...
}
```

## Performance

### Test Execution Times

| Test Type | Duration | Infrastructure Required |
|-----------|----------|-------------------------|
| Unit Tests | 2-5 min | None |
| Integration Tests | 10-30 min | LocalStack |
| Full Test Suite | 15-40 min | LocalStack |

### Optimization Tips

1. Run unit tests first (fast feedback)
2. Use `t.Parallel()` for independent tests
3. Cache Go modules in CI
4. Use specific test filters when developing
5. Keep LocalStack running during development

## Resources

- [Terratest Documentation](https://terratest.gruntwork.io/)
- [LocalStack Documentation](https://docs.localstack.cloud/)
- [AWS Provider Documentation](https://registry.terraform.io/providers/hashicorp/aws/latest/docs)
- [Go Testing Documentation](https://golang.org/pkg/testing/)

## Support

For issues related to testing:
1. Check this guide for solutions
2. Review GitHub Actions logs
3. Open an issue with:
   - Test output
   - Go version
   - Terraform version
   - Operating system
   - Steps to reproduce
