# Renovate Configuration

This repository uses [Renovate Bot](https://docs.renovatebot.com/) to
automatically keep all dependencies up to date.

## What Renovate Manages

### GitHub Actions (`.github/workflows/*.yml`)

- All `uses:` actions (e.g., `actions/checkout@v4`, `actions/setup-go@v5`)
- Docker images in `services:` sections (e.g., `localstack/localstack:4.1`)
- `go-version` parameter in setup-go
- `terraform_version` parameter in setup-terraform

**Current Actions:**

- actions/checkout@v4
- actions/setup-go@v5
- actions/cache@v4
- hashicorp/setup-terraform@v3
- terraform-linters/setup-tflint@v4
- aquasecurity/trivy-action@master
- github/codeql-action/upload-sarif@v3
- softprops/action-gh-release@v2

### Pre-commit Hooks (`.pre-commit-config.yaml`)

- All `rev:` tags for pre-commit repos

**Current Hooks:**

- pre-commit/pre-commit-hooks @ v4.5.0
- antonbabenko/pre-commit-terraform @ v1.86.0
- gitleaks/gitleaks @ v8.18.1

### Go Dependencies (`tests/go.mod`)

- Go version (`go 1.23`)
- All direct dependencies:
  - github.com/aws/aws-sdk-go
  - github.com/gruntwork-io/terratest
  - github.com/stretchr/testify
- All indirect dependencies (auto-updated by Renovate)

### Terraform Providers (`.tf` files)

- AWS provider (`hashicorp/aws ~> 5.0`)
- Archive provider (`hashicorp/archive ~> 2.0` in with-lambda example)
- Terraform version constraints (`>= 1.5.0`)

**Locations:**

- `main.tf`
- `modules/*/main.tf`
- `examples/*/main.tf`

## Auto-merge Rules

Renovate will automatically merge the following updates after CI passes:

- ✅ **GitHub Actions**: Patch and minor updates
- ✅ **Pre-commit hooks**: Patch and minor updates
- ✅ **Go dependencies**: Patch updates only
- ✅ **Lock files**: Weekly lock file maintenance
- ❌ **Major updates**: Always require manual review
- ❌ **Terraform providers**: Always require manual review
- ❌ **Docker images**: Always require manual review

## Schedule

- **Regular updates**: Weeknights (9pm-6am) and weekends
- **Lock file maintenance**: Monday mornings before 6am
- **Concurrent PRs**: Maximum 3 at a time

## Security Alerts

Renovate is configured to create PRs for security vulnerabilities
immediately, labeled with `security`. These PRs:

- Are **not** auto-merged (require manual review)
- Take priority over the regular schedule
- Are labeled for easy identification

## Dependency Dashboard

Renovate creates a
[Dependency Dashboard](../../issues?q=is%3Aissue+is%3Aopen+title%3A%22Dependency+Dashboard%22)
issue that shows:

- All pending updates
- Rate-limited updates
- Updates with errors
- Ignored dependencies

## Testing Changes

All Renovate PRs trigger the full CI pipeline:

- Unit tests
- Integration tests (LocalStack)
- Terraform validation
- TFLint
- Trivy security scanning
- Pre-commit hooks

## Overriding Renovate

To ignore a specific dependency, add it to `.github/renovate.json`:

```json
{
  "ignoreDeps": ["package-name"]
}
```

To pin a version and prevent updates:

```json
{
  "packageRules": [
    {
      "matchPackageNames": ["package-name"],
      "enabled": false
    }
  ]
}
```

## Troubleshooting

If Renovate stops working:

1. Check the
   [Dependency Dashboard](../../issues?q=is%3Aissue+is%3Aopen+title%3A%22Dependency+Dashboard%22)
   for errors
2. Validate the config: `renovate-config-validator .github/renovate.json`
3. Check [Renovate logs](https://app.renovatebot.com/dashboard)

## References

- [Renovate Documentation](https://docs.renovatebot.com/)
- [Configuration Options](https://docs.renovatebot.com/configuration-options/)
- [Terraform Manager](https://docs.renovatebot.com/modules/manager/terraform/)
- [Go Modules Manager](https://docs.renovatebot.com/modules/manager/gomod/)
