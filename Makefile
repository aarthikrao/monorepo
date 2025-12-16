.PHONY: build

.SILENT: build tidy

build:
	go build -o bin/funds ./services/funds/main.go
	go build -o bin/kyc ./services/kyc/main.go
	echo "📦 Buit binaries:"
	@for file in bin/*; do \
		echo "$$(du -h $$file)"; \
	done

tidy:
	go work use -r ./
	go work sync
	go mod tidy -v

local-deploy:
	docker compose -f infra/docker-compose.yml up --build -d
	docker compose -f infra/docker-compose.yml ps
	./bin/funds &
	./bin/kyc &