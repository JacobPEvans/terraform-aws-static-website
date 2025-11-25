# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [2.0.0] - 2025-11-25

### Added
- ✨ LocalStack integration for local testing
- ✨ Terratest integration tests in Go
- ✨ Pre-commit hooks configuration (.pre-commit-config.yaml)
- ✨ GitHub Actions CI/CD pipeline
- ✨ TFLint configuration (.tflint.hcl)
- ✨ Makefile for common operations
- ✨ Three complete examples (basic, SPA, Lambda@Edge)
- ✨ Docker Compose configuration for LocalStack
- ✨ S3 bucket versioning on all buckets
- ✨ S3 bucket encryption (AES256) on all buckets
- ✨ S3 public access block configurations
- ✨ S3 bucket ownership controls
- ✨ Comprehensive README with badges and architecture diagram
- ✨ CONTRIBUTING.md guide
- ✨ CHANGELOG.md (this file)

### Changed
- 🔄 **BREAKING**: Updated to AWS Provider 5.0 compatibility
- 🔄 **BREAKING**: Minimum Terraform version now 1.5.0
- 🔄 **BREAKING**: Split S3 bucket configurations into separate resources:
  - `aws_s3_bucket_acl` for ACL configuration
  - `aws_s3_bucket_logging` for logging configuration
  - `aws_s3_bucket_website_configuration` for website configuration
  - `aws_s3_bucket_versioning` for versioning
  - `aws_s3_bucket_server_side_encryption_configuration` for encryption
  - `aws_s3_bucket_ownership_controls` for ownership
  - `aws_s3_bucket_public_access_block` for public access settings
- 🔄 Updated CloudFront to use TLS 1.2 only (removed TLS 1.0 and 1.1)
- 🔄 Improved security posture across all resources
- 🔄 Complete README overhaul with modern documentation

### Security
- 🔒 Enforced TLS 1.2 minimum for all HTTPS connections
- 🔒 Added encryption at rest for all S3 buckets
- 🔒 Configured appropriate public access blocks
- 🔒 Added Trivy security scanning in CI/CD
- 🔒 Added gitleaks for secrets detection in pre-commit hooks

### Fixed
- 🐛 Fixed deprecated S3 bucket arguments for AWS Provider 5.0
- 🐛 Corrected website endpoint references to use new resource

### Removed
- ❌ Support for TLS 1.0 and 1.1 (security improvement)
- ❌ Direct ACL, logging, and website configuration in S3 bucket resource

## [1.0.0] - 2020-XX-XX

### Added
- Initial release
- Basic static website hosting with S3 and CloudFront
- ACM certificate management
- Route53 DNS configuration
- Lambda@Edge support
- SPA support

[2.0.0]: https://github.com/JacobPEvans/terraform-aws-static-website/compare/v1.0.0...v2.0.0
[1.0.0]: https://github.com/JacobPEvans/terraform-aws-static-website/releases/tag/v1.0.0
