.PHONY: build

build:
	go build -o bin/funds ./services/funds/main.go
	go build -o bin/kyc ./services/kyc/main.go

local-deploy:
	docker-compose -f infra/docker-compose.yml up --build -d
	./bin/funds &
	./bin/kyc &