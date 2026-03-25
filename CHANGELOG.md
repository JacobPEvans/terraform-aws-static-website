# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [2.2.0](https://github.com/JacobPEvans/terraform-aws-static-website/compare/v2.1.0...v2.2.0) (2026-03-25)


### Bug Fixes

* **security:** pin trivy-action to SHA, upgrade checkout to v6 ([#18](https://github.com/JacobPEvans/terraform-aws-static-website/issues/18)) ([9f6734d](https://github.com/JacobPEvans/terraform-aws-static-website/commit/9f6734d3a339362bc746c8b4619077bcecccd70a))

## [2.1.0](https://github.com/JacobPEvans/terraform-aws-static-website/compare/v2.0.0...v2.1.0) (2026-03-19)


### Features

* **ci:** implement two-tier integration test strategy ([eedf548](https://github.com/JacobPEvans/terraform-aws-static-website/commit/eedf54867d4e9d87301be5a468e88786c8f1aa16))
* Complete modernization with AWS Provider 5.0 compatibility and comprehensive testing ([4faa999](https://github.com/JacobPEvans/terraform-aws-static-website/commit/4faa9997c7c5833fd5c00e7cef9f11a37da46d79))
* **deps:** add comprehensive Renovate configuration for all dependencies ([c69782a](https://github.com/JacobPEvans/terraform-aws-static-website/commit/c69782ad04c72c5277a7f306fc3a3f5a5d30c6a5))
* **renovate:** migrate to shared preset for common rules ([e524d26](https://github.com/JacobPEvans/terraform-aws-static-website/commit/e524d266bd96574efb88aadaeccbb43ddebaec6f))


### Bug Fixes

* add explicit config-file and manifest-file to release-please action ([374e4bf](https://github.com/JacobPEvans/terraform-aws-static-website/commit/374e4bf18b0963dffd7f6c27fcc637fca2020bac))
* Add go.sum and update .gitignore for Go dependency management ([20c891f](https://github.com/JacobPEvans/terraform-aws-static-website/commit/20c891fd22bcc6866c860c201a44b051a2f82ec9))
* align tag format and add TFLint auth token ([#15](https://github.com/JacobPEvans/terraform-aws-static-website/issues/15)) ([3536028](https://github.com/JacobPEvans/terraform-aws-static-website/commit/35360285047db48486d978dc3050e8613009b447))
* **ci:** add STS service to plan-tests LocalStack config ([faa12b9](https://github.com/JacobPEvans/terraform-aws-static-website/commit/faa12b9aea9010d09ac0910a810bedb68f9d8cbb))
* **ci:** update LocalStack to v4.0 to fix S3 HEAD request handling ([7c53850](https://github.com/JacobPEvans/terraform-aws-static-website/commit/7c53850f34e4d7e0b55b5b9631020c7c1a7d794c))
* **ci:** upgrade LocalStack to 4.1 and add S3 signature validation skip ([b03b5f6](https://github.com/JacobPEvans/terraform-aws-static-website/commit/b03b5f6b749f94ee2cbbb2d7b32daa3ae610f67f))
* **examples:** remove unused aws us-east-1 provider alias ([d0ab109](https://github.com/JacobPEvans/terraform-aws-static-website/commit/d0ab1096d2f964a41c03c9c9062ea9a2448dc4be))
* migrate to reusable release-please workflow ([7c6c630](https://github.com/JacobPEvans/terraform-aws-static-website/commit/7c6c630695492393d17ba6f0ae29de3b7dfd3d2e))
* **release:** add versioning always-bump-minor to prevent major bumps ([b763300](https://github.com/JacobPEvans/terraform-aws-static-website/commit/b7633000a008822c0ff7674aa28b02d5315aa620))
* Resolve all GitHub Actions workflow failures ([1ac181b](https://github.com/JacobPEvans/terraform-aws-static-website/commit/1ac181b907e5d37cc40b526fa9a6db04fcda3aeb))
* Resolve final TFLint warnings and standard module structure ([cfc1b07](https://github.com/JacobPEvans/terraform-aws-static-website/commit/cfc1b07bcf48913b81a85d0489c96606a337d52a))
* Resolve TFLint and provider configuration issues ([44ae42b](https://github.com/JacobPEvans/terraform-aws-static-website/commit/44ae42bcb78259a8c447b987bba52c2335728060))
* **security:** resolve all CodeQL and Dependabot alerts ([8770b25](https://github.com/JacobPEvans/terraform-aws-static-website/commit/8770b25ed05f36b1bafa03a576ecc4cceeb1bd80))
* **tests:** correct SPA module name and workflow improvements ([caaf941](https://github.com/JacobPEvans/terraform-aws-static-website/commit/caaf94181d4dd4125169e81691858e9c64b199ff))
* **workflows:** add actions:write permission for cache operations ([fc58ee8](https://github.com/JacobPEvans/terraform-aws-static-website/commit/fc58ee8d1671b619ab2f0ed3a2a6516bb286f073))

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
