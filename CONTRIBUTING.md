# Contributing to terraform-aws-static-website

Thank you for considering contributing to this project! We welcome contributions from the community.

## ⚠️ CRITICAL: Local-First Testing Policy

**ALL checks that run in GitHub Actions MUST be run locally BEFORE committing or pushing.**

This policy exists to:
- ✅ Reduce CI/CD usage and costs
- ✅ Provide immediate feedback (seconds vs minutes)
- ✅ Catch issues before they reach GitHub
- ✅ Significantly decrease development time
- ✅ Prevent wasted CI runs
- ✅ Maintain repository quality

### Mandatory Local Checks

Before **EVERY** commit/push, you MUST run:

```bash
# Option 1: Install Git hooks (RECOMMENDED - automatic)
make install-hooks

# Option 2: Run checks manually before pushing
make validate-local
```

### What Gets Checked Locally

The `make validate-local` command runs:
1. ✅ Terraform format check (`terraform fmt`)
2. ✅ Terraform validation (`terraform validate`)
3. ✅ TFLint static analysis
4. ✅ Trivy security scanning
5. ✅ Go unit tests (all resource validation tests)

**These are the EXACT SAME checks that GitHub Actions runs.**

### Consequences of Skipping Local Checks

- ❌ Your PR may be closed
- ❌ Wasted time waiting for CI to fail
- ❌ Multiple push cycles to fix simple issues
- ❌ Reviewer frustration
- ❌ Increased CI costs

### Git Hooks

Installing hooks ensures checks run automatically:

```bash
# Install hooks once
make install-hooks

# Now checks run automatically on:
# - Every commit (format, validate, lint, unit tests)
# - Every push (all of the above + module validation)
```

To bypass hooks (NOT RECOMMENDED):
```bash
git commit --no-verify  # Skip pre-commit checks
git push --no-verify     # Skip pre-push checks
```

## Code of Conduct

Be respectful, constructive, and professional in all interactions.

## How to Contribute

### Reporting Bugs

1. Check if the bug has already been reported in [Issues](https://github.com/JacobPEvans/terraform-aws-static-website/issues)
2. If not, create a new issue with:
   - Clear title and description
   - Steps to reproduce
   - Expected vs actual behavior
   - Terraform and AWS provider versions
   - Relevant code snippets or configurations

### Suggesting Enhancements

1. Check [existing issues](https://github.com/JacobPEvans/terraform-aws-static-website/issues) for similar suggestions
2. Create a new issue describing:
   - The enhancement and its benefits
   - Possible implementation approach
   - Any breaking changes it might introduce

### Pull Requests

1. **Fork the repository** and create your branch from `main`
   ```bash
   git checkout -b feature/my-awesome-feature
   ```

2. **Install development dependencies**
   ```bash
   # Install pre-commit
   pip install pre-commit

   # Install Git hooks (REQUIRED)
   make install-hooks

   # Install Go (for tests)
   # Follow: https://golang.org/doc/install

   # Install Docker (for LocalStack)
   # Follow: https://docs.docker.com/get-docker/

   # Install TFLint
   brew install tflint  # macOS
   # or: https://github.com/terraform-linters/tflint

   # Install Trivy
   brew install trivy  # macOS
   # or: https://aquasecurity.github.io/trivy/latest/getting-started/installation/
   ```

3. **Make your changes**
   - Follow Terraform best practices
   - Add tests for new functionality
   - Update documentation as needed
   - Ensure code is properly formatted

4. **Run quality checks BEFORE committing** (hooks do this automatically)
   ```bash
   # Run ALL local checks (mirrors CI)
   make validate-local

   # Or run individual checks
   make fmt         # Format code
   make validate    # Validate Terraform
   make lint        # Run linter
   make test-unit   # Run unit tests
   ```

5. **Commit your changes**
   ```bash
   git add .
   git commit -m "Add my awesome feature"
   # Hooks will run automatically if installed
   ```

   Follow these commit message guidelines:
   - Use present tense ("Add feature" not "Added feature")
   - Use imperative mood ("Move cursor to..." not "Moves cursor to...")
   - Limit first line to 72 characters
   - Reference issues and pull requests when relevant

6. **Verify everything passes locally**
   ```bash
   # This is what CI will run
   make validate-local
   ```

7. **Push to your fork**
   ```bash
   git push origin feature/my-awesome-feature
   # Pre-push hooks will run if installed
   ```

8. **Open a Pull Request**
   - Provide a clear description of the changes
   - Reference any related issues
   - Include test results if applicable
   - Update CHANGELOG.md under "Unreleased" section
   - Confirm you ran `make validate-local` successfully

## Development Setup

### Prerequisites

- **Terraform** >= 1.5.0
- **Go** >= 1.21 (for tests)
- **Docker** (for LocalStack)
- **Python** >= 3.11 (for pre-commit)
- **TFLint** (for linting)
- **Trivy** (for security scanning)
- **AWS CLI** (optional, for manual testing)

### Project Structure

```
.
├── .github/
│   └── workflows/        # GitHub Actions CI/CD
├── examples/             # Usage examples
│   ├── basic/
│   ├── spa/
│   └── with-lambda/
├── scripts/              # Helper scripts
│   ├── install-hooks.sh  # Git hooks installer
│   └── validate-local.sh # Local CI validation
├── tests/                # Integration tests
│   ├── localstack-init/  # LocalStack setup scripts
│   └── *.go              # Terratest files
├── main.tf               # Main module code
├── variables.tf          # Input variables
├── output.tf             # Output values
├── .pre-commit-config.yaml
├── .tflint.hcl
├── docker-compose.yml
├── Makefile
└── TESTING.md            # Comprehensive testing guide
```

### Testing Philosophy

- All new features should include tests
- Tests should use LocalStack for local execution
- Integration tests use Terratest (Go)
- **Local checks MUST pass before pushing**
- CI/CD validates all changes

### Code Style

- Follow [Terraform Style Guide](https://www.terraform.io/docs/language/syntax/style.html)
- Use `terraform fmt` for formatting
- Run `tflint` for linting
- Use meaningful variable and resource names
- Add comments for complex logic
- Keep lines under 120 characters when possible

### Documentation

- Update README.md for user-facing changes
- Update examples if behavior changes
- Add inline comments for complex Terraform code
- Update CHANGELOG.md for all changes
- Use terraform-docs format for variable/output docs

## Release Process

1. Update CHANGELOG.md with version and date
2. Create a git tag: `git tag -a v2.1.0 -m "Release v2.1.0"`
3. Push tag: `git push origin v2.1.0`
4. GitHub Actions will create the release automatically

## Common Development Workflows

### Before Starting Work

```bash
# 1. Install hooks (one time)
make install-hooks

# 2. Sync with upstream
git pull origin main

# 3. Create feature branch
git checkout -b feature/my-feature
```

### During Development

```bash
# Format code
make fmt

# Validate locally
make validate-local

# Run unit tests
make test-unit

# Commit (hooks run automatically)
git commit -m "feat: add my feature"
```

### Before Pushing

```bash
# CRITICAL: Run full validation
make validate-local

# If all passes, push
git push origin feature/my-feature
```

### Running Tests

```bash
# Unit tests only (fast)
make test-unit

# Integration tests
make test-local

# All checks (CI mirror)
make validate-local
```

## Getting Help

- Check [existing issues](https://github.com/JacobPEvans/terraform-aws-static-website/issues)
- Start a [discussion](https://github.com/JacobPEvans/terraform-aws-static-website/discussions)
- Review [examples](./examples/)
- Read the [README](./README.md)
- Read the [TESTING guide](./TESTING.md)

## License

By contributing, you agree that your contributions will be licensed under the MIT License.
