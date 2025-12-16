.PHONY: build
.SILENT: build tidy
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

tidy: ## Tidy up the go.work and go.mod files. Run this after adding a new module.
	go work use -r ./
	go work sync
	go mod tidy -v

local-deploy: ## Deploy the local development infra using docker compose and start all the apps.
	docker compose -f infra/docker-compose.yml up --build -d
	docker compose -f infra/docker-compose.yml ps
	./bin/funds &
	./bin/kyc &