# Single Page Application (SPA) Example

This example demonstrates hosting a Single Page Application (React, Vue, Angular, etc.) using the terraform-aws-static-website module.

## Features

- SPA support with 404 → index.html redirect
- CloudFront CDN distribution with HTTPS
- ACM certificate with automatic DNS validation
- Route53 DNS records
- S3 bucket logging and encryption
- S3 versioning enabled

## Usage

```hcl
module "spa_website" {
  source = "github.com/jacobpevans/terraform-aws-static-website"

  website-domain-main = "app.example.com"
  domains-zone-root   = "example.com"
  support-spa         = true

  tags = {
    Environment = "production"
    Application = "MyReactApp"
  }
}
```

## SPA Support

When `support-spa = true`, the module configures CloudFront to:
- Return `index.html` for 404 errors with a 200 status code
- Allow client-side routing to work properly
- Support deep linking to any route in your SPA

## Prerequisites

1. AWS account with appropriate credentials configured
2. Route53 hosted zone for your domain
3. Terraform >= 1.5.0
4. Built SPA application ready to upload to S3

## Deployment

```bash
terraform init
terraform plan
terraform apply

# Upload your SPA build to S3
aws s3 sync ./build s3://app.example.com-root
```

## Testing Locally with LocalStack

```bash
# Start LocalStack
docker-compose up -d

# Run terraform with LocalStack
export AWS_ENDPOINT_URL=http://localhost:4566
terraform init
terraform apply

# Clean up
terraform destroy
docker-compose down
```
