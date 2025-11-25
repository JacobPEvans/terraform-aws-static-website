# Contributing to terraform-aws-static-website

Thank you for considering contributing to this project! We welcome contributions from the community.

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

   # Install pre-commit hooks
   make pre-commit-install

   # Install Go (for tests)
   # Follow: https://golang.org/doc/install

   # Install Docker (for LocalStack)
   # Follow: https://docs.docker.com/get-docker/
   ```

3. **Make your changes**
   - Follow Terraform best practices
   - Add tests for new functionality
   - Update documentation as needed
   - Ensure code is properly formatted

4. **Run quality checks**
   ```bash
   # Format code
   make fmt

   # Validate Terraform
   make validate

   # Run linter
   make lint

   # Run pre-commit hooks
   make pre-commit-run
   ```

5. **Run tests**
   ```bash
   # Start LocalStack
   make localstack-start

   # Run integration tests
   cd tests && go test -v

   # Or use the combined command
   make test-local
   ```

6. **Commit your changes**
   ```bash
   git add .
   git commit -m "Add my awesome feature"
   ```

   Follow these commit message guidelines:
   - Use present tense ("Add feature" not "Added feature")
   - Use imperative mood ("Move cursor to..." not "Moves cursor to...")
   - Limit first line to 72 characters
   - Reference issues and pull requests when relevant

7. **Push to your fork**
   ```bash
   git push origin feature/my-awesome-feature
   ```

8. **Open a Pull Request**
   - Provide a clear description of the changes
   - Reference any related issues
   - Include test results if applicable
   - Update CHANGELOG.md under "Unreleased" section

## Development Setup

### Prerequisites

- Terraform >= 1.5.0
- Go >= 1.21 (for tests)
- Docker (for LocalStack)
- Python >= 3.11 (for pre-commit)
- AWS CLI (optional, for manual testing)

### Project Structure

```
.
├── .github/
│   └── workflows/        # GitHub Actions CI/CD
├── examples/             # Usage examples
│   ├── basic/
│   ├── spa/
│   └── with-lambda/
├── tests/                # Integration tests
│   ├── localstack-init/  # LocalStack setup scripts
│   └── *.go              # Terratest files
├── main.tf               # Main module code
├── variables.tf          # Input variables
├── output.tf             # Output values
├── .pre-commit-config.yaml
├── .tflint.hcl
├── docker-compose.yml
└── Makefile
```

### Testing Philosophy

- All new features should include tests
- Tests should use LocalStack for local execution
- Integration tests use Terratest (Go)
- Pre-commit hooks ensure code quality
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

## Getting Help

- Check [existing issues](https://github.com/JacobPEvans/terraform-aws-static-website/issues)
- Start a [discussion](https://github.com/JacobPEvans/terraform-aws-static-website/discussions)
- Review [examples](./examples/)
- Read the [README](./README.md)

## License

By contributing, you agree that your contributions will be licensed under the MIT License.
