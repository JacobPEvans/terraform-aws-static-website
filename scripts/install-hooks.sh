#!/bin/bash
set -e

echo "Installing Git hooks..."

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_ROOT="$(cd "$SCRIPT_DIR/.." && pwd)"
HOOKS_DIR="$PROJECT_ROOT/.git/hooks"

# Create hooks directory if it doesn't exist
mkdir -p "$HOOKS_DIR"

# Pre-commit hook
cat > "$HOOKS_DIR/pre-commit" <<'EOF'
#!/bin/bash
set -e

echo "🔍 Running pre-commit checks..."

# Check if terraform is installed
if ! command -v terraform &> /dev/null; then
    echo "❌ terraform not found. Please install terraform."
    exit 1
fi

# Format check
echo "📝 Checking Terraform format..."
if ! terraform fmt -check -recursive; then
    echo "❌ Terraform files are not formatted. Run 'terraform fmt -recursive' to fix."
    exit 1
fi

# Validate
echo "✅ Validating Terraform..."
terraform init -backend=false > /dev/null 2>&1
if ! terraform validate; then
    echo "❌ Terraform validation failed."
    exit 1
fi

# Check for tflint
if command -v tflint &> /dev/null; then
    echo "🔎 Running tflint..."
    tflint --init > /dev/null 2>&1 || true
    if ! tflint --recursive --config=.tflint.hcl; then
        echo "⚠️  tflint found issues (non-blocking)"
    fi
else
    echo "⚠️  tflint not found. Install with: brew install tflint"
fi

# Run Go tests if available
if command -v go &> /dev/null && [ -d "tests" ]; then
    echo "🧪 Running unit tests..."
    cd tests
    go mod download > /dev/null 2>&1
    if ! go test -run "TestS3BucketResourcesIndividually|TestCloudFrontResources|TestRoute53Resources" -timeout 5m; then
        echo "❌ Unit tests failed."
        cd ..
        exit 1
    fi
    cd ..
else
    echo "⚠️  Go not found or tests directory missing. Skipping unit tests."
fi

echo "✅ All pre-commit checks passed!"
EOF

# Pre-push hook
cat > "$HOOKS_DIR/pre-push" <<'EOF'
#!/bin/bash
set -e

echo "🚀 Running pre-push checks..."

# Run all pre-commit checks first
echo "Running pre-commit checks..."
.git/hooks/pre-commit

# Run validation tests
if command -v go &> /dev/null && [ -d "tests" ]; then
    echo "🧪 Running module validation tests..."
    cd tests
    if ! go test -run TestModuleValidation -timeout 10m; then
        echo "❌ Module validation tests failed."
        cd ..
        exit 1
    fi
    cd ..
fi

echo "✅ All pre-push checks passed! Safe to push."
EOF

# Make hooks executable
chmod +x "$HOOKS_DIR/pre-commit"
chmod +x "$HOOKS_DIR/pre-push"

echo "✅ Git hooks installed successfully!"
echo ""
echo "Hooks installed:"
echo "  - pre-commit: Runs format checks, validation, linting, and unit tests"
echo "  - pre-push: Runs all pre-commit checks plus module validation"
echo ""
echo "To bypass hooks (not recommended):"
echo "  git commit --no-verify"
echo "  git push --no-verify"
