# Lambda@Edge Example

This example demonstrates advanced usage with Lambda@Edge for adding security headers to all responses.

## Features

- Lambda@Edge function for security headers
- Automatic HSTS, X-Frame-Options, CSP headers
- CloudFront CDN distribution with HTTPS
- ACM certificate with automatic DNS validation
- Route53 DNS records
- S3 bucket logging and encryption

## Security Headers Added

The Lambda@Edge function adds the following security headers:

- `Strict-Transport-Security`: Enforces HTTPS
- `X-Content-Type-Options`: Prevents MIME-type sniffing
- `X-Frame-Options`: Prevents clickjacking
- `X-XSS-Protection`: Enables browser XSS protection

## Usage

```hcl
module "static_website_with_lambda" {
  source = "github.com/jacobpevans/terraform-aws-static-website"

  website-domain-main                   = "example.com"
  website-domain-redirect               = "www.example.com"
  domains-zone-root                     = "example.com"
  support-spa                           = false
  cloudfront_lambda_function_arn        = aws_lambda_function.security_headers.qualified_arn
  cloudfront_lambda_function_event_type = "viewer-response"

  tags = {
    Environment = "production"
  }
}
```

## Lambda@Edge Events

The module supports the following Lambda@Edge event types:

- `viewer-request`: Before CloudFront forwards request to origin
- `origin-request`: Before CloudFront forwards request to origin (after cache miss)
- `origin-response`: After CloudFront receives response from origin
- `viewer-response`: Before CloudFront returns response to viewer (used in this example)

## Prerequisites

1. AWS account with appropriate credentials configured
2. Route53 hosted zone for your domain
3. Terraform >= 1.5.0
4. Understanding of Lambda@Edge limitations and best practices

## Important Notes

- Lambda@Edge functions must be created in us-east-1
- Function must be published (not $LATEST)
- Execution time limits are stricter than regular Lambda
- Function will be replicated to edge locations

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

**Note**: Lambda@Edge replicas at edge locations may take several hours to be deleted after the function is removed.
