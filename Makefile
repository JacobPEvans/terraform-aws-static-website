.PHONY: help localstack-start localstack-stop localstack-restart test test-unit test-local fmt validate lint pre-commit-install pre-commit-run clean

help: ## Show this help message
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Available targets:'
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-20s\033[0m %s\n", $$1, $$2}'

localstack-start: ## Start LocalStack for testing
	@echo "Starting LocalStack..."
	docker-compose up -d
	@echo "Waiting for LocalStack to be healthy..."
	@timeout 60 bash -c 'until docker-compose ps | grep -q healthy; do sleep 2; done'
	@echo "LocalStack is ready!"

localstack-stop: ## Stop LocalStack
	@echo "Stopping LocalStack..."
	docker-compose down

localstack-restart: localstack-stop localstack-start ## Restart LocalStack

localstack-logs: ## Show LocalStack logs
	docker-compose logs -f localstack

test: ## Run integration tests with LocalStack
	@echo "Running integration tests..."
	cd tests && go test -v -timeout 30m

test-unit: ## Run unit tests (no LocalStack required)
	@echo "Running unit tests..."
	cd tests && go test -v -run "TestS3BucketResourcesIndividually|TestCloudFrontResources|TestRoute53Resources|TestModuleValidation" -timeout 10m

test-local: localstack-start test ## Start LocalStack and run integration tests

fmt: ## Format Terraform code
	terraform fmt -recursive

validate: fmt ## Validate Terraform code
	terraform init -backend=false
	terraform validate

lint: ## Run tflint
	tflint --init
	tflint --recursive

pre-commit-install: ## Install pre-commit hooks
	pre-commit install

pre-commit-run: ## Run pre-commit hooks on all files
	pre-commit run --all-files

clean: localstack-stop ## Clean up temporary files and stop LocalStack
	rm -rf .terraform .terraform.lock.hcl
	find . -type f -name 'terraform.tfstate*' -delete
