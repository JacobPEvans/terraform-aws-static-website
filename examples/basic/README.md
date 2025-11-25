# Basic Static Website Example

This example demonstrates the basic usage of the terraform-aws-static-website module.

## Features

- Creates a static website hosted on S3
- CloudFront CDN distribution with HTTPS
- ACM certificate with automatic DNS validation
- Route53 DNS records
- Optional redirect from www to apex domain
- S3 bucket logging

## Usage

```hcl
module "static_website" {
  source = "github.com/jacobpevans/terraform-aws-static-website"

  website-domain-main     = "example.com"
  website-domain-redirect = "www.example.com"
  domains-zone-root       = "example.com"
  support-spa             = false

  tags = {
    Environment = "production"
  }
}
```

## Prerequisites

1. AWS account with appropriate credentials configured
2. Route53 hosted zone for your domain
3. Terraform >= 1.5.0

## Deployment

```bash
terraform init
terraform plan
terraform apply
```

## Cleanup

```bash
terraform destroy
```

## Testing with LocalStack

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
