.PHONY: build
.SILENT: build tidy local-deploy clean help
.DEFAULT_GOAL := help


help: ## Show this help. You canb refer infra/LocalSetup.md for setting up local environment.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make <target>\n\nTargets:\n"} \
	/^[a-zA-Z_-]+:.*##/ { printf "  %-15s %s\n", $$1, $$2 }' $(MAKEFILE_LIST)

build: tidy ## Build all services. Ensure all your services are in services/ directory. All the binaries will be placed in bin/ directory.
	@for dir in services/*/; do \
		service=$$(basename $$dir); \
		go build -o bin/$$service ./$$dir/.; \
	done
	echo "📦 Buit binaries:"
	@for file in bin/*; do \
		echo "$$(du -h $$file)"; \
	done

tidy: ## Tidy up go.mod file. Run this after adding a new module.
	go mod tidy

clean: ## Clean up all built binaries.
	rm -rf bin/*

local-deploy: ## Deploy the local development infra using docker compose and start all the apps.
	docker compose -f infra/local_setup/docker-compose.yml up --build -d
	echo "To fetch latest status of running docker containers use \n>  docker compose -f infra/local_setup/docker-compose.yml ps" 
	mkdir -p logs
	./bin/funds >> logs/funds.log 2>&1 &
	./bin/kyc >> logs/kyc.log 2>&1 &
	./bin/notifier >> logs/notifier.log 2>&1 &
	echo "🚀 Started all services in background. Logs are being written to logs/ directory."