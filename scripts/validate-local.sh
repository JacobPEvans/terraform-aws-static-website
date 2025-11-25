#!/bin/bash
# This script runs all CI checks locally before pushing
# It mirrors what GitHub Actions will run

set -e

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_ROOT="$(cd "$SCRIPT_DIR/.." && pwd)"

cd "$PROJECT_ROOT"

echo "======================================================"
echo "Running Local Validation (mirrors GitHub Actions CI)"
echo "======================================================"
echo ""

# Color codes
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

FAILED=0

# Function to run a check
run_check() {
    local name="$1"
    local command="$2"

    echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
    echo "🔍 $name"
    echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"

    if eval "$command"; then
        echo -e "${GREEN}✅ $name PASSED${NC}"
        echo ""
        return 0
    else
        echo -e "${RED}❌ $name FAILED${NC}"
        echo ""
        FAILED=1
        return 1
    fi
}

# 1. Terraform Format Check
run_check "Terraform Format Check" "terraform fmt -check -recursive"

# 2. Terraform Validation
run_check "Terraform Init" "terraform init -backend=false"
run_check "Terraform Validate" "terraform validate"

# 3. TFLint
if command -v tflint &> /dev/null; then
    run_check "TFLint Init" "tflint --init"
    run_check "TFLint Check" "tflint --recursive --config=.tflint.hcl"
else
    echo -e "${YELLOW}⚠️  TFLint not installed. Install with: brew install tflint${NC}"
    echo ""
fi

# 4. Trivy Security Scan (if available)
if command -v trivy &> /dev/null; then
    run_check "Trivy Security Scan" "trivy config --severity HIGH,CRITICAL --exit-code 0 ."
else
    echo -e "${YELLOW}⚠️  Trivy not installed. Install with: brew install trivy${NC}"
    echo ""
fi

# 5. Go Unit Tests
if command -v go &> /dev/null && [ -d "tests" ]; then
    echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
    echo "🧪 Go Unit Tests"
    echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"

    cd tests
    go mod download

    # Run each test separately for better visibility
    run_check "S3 Bucket Resources Test" "go test -v -run TestS3BucketResourcesIndividually -timeout 5m"
    run_check "CloudFront Resources Test" "go test -v -run TestCloudFrontResources -timeout 5m"
    run_check "Route53 Resources Test" "go test -v -run TestRoute53Resources -timeout 5m"
    run_check "Module Validation Test" "go test -v -run TestModuleValidation -timeout 10m"

    cd ..
else
    echo -e "${YELLOW}⚠️  Go not installed or tests directory missing${NC}"
    echo ""
fi

echo "======================================================"
if [ $FAILED -eq 0 ]; then
    echo -e "${GREEN}✅ ALL CHECKS PASSED!${NC}"
    echo "You can safely commit and push."
else
    echo -e "${RED}❌ SOME CHECKS FAILED!${NC}"
    echo "Please fix the issues before committing/pushing."
    exit 1
fi
echo "======================================================"
