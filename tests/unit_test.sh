#!/bin/bash
set -e

echo "Running unit tests (validation only, no LocalStack required)..."

# Run resource structure validation tests
echo "Testing S3 bucket resources..."
go test -v -run TestS3BucketResourcesIndividually -timeout 5m

echo "Testing CloudFront resources..."
go test -v -run TestCloudFrontResources -timeout 5m

echo "Testing Route53 resources..."
go test -v -run TestRoute53Resources -timeout 5m

echo "Testing module validation..."
go test -v -run TestModuleValidation -timeout 10m

echo "✅ All unit tests passed!"
