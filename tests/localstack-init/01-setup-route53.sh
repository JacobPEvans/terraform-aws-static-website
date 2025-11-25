#!/bin/bash

# Create a Route53 hosted zone for testing
echo "Creating Route53 hosted zone for test domain..."

awslocal route53 create-hosted-zone \
  --name "example.com" \
  --caller-reference "$(date +%s)" \
  --hosted-zone-config Comment="Test zone for static website module"

echo "Route53 hosted zone created successfully"
