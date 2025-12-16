# monorepo
Experimenting with monorepo

## Getting started
This is the starting point of this repo. 
* To build run `make build`
* To deploy the local development infra using docker compose, you can run `make local-deploy`. This will spin up all the necessary components and start all the apps.
* You can refer to [infra/LocalSetup.md](infra/LocalSetup.md) for setting up local environment

## Adding a new service
To add a new service to the monorepo, follow these steps:
1. Create a new directory under `services/` for your service.
2. Add your Go code and any other necessary files to this directory.
3. The Makefile is set up to automatically build any service located in the `services/` directory. Ensure your service has a `main.go` file as the entry point.
4. Update the `infra/docker-compose.yml` file to include your new service if it needs to run as part of the local development environment.
5. Run `make build` to build all services, including the new one.
6. Run `make local-deploy` to deploy the updated services.

