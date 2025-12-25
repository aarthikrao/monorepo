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
4. Update the `infra/local_setup/docker-compose.yml` file to include your new service if it needs to run as part of the local development environment.
5. Run `make build` to build all services, including the new one.
6. Run `make local-deploy` to deploy the updated services.

## Load testing
Add all the scripts and documentation related to load testing `infra/loadtest/` directory.
See [infra/loadtest/README.md](infra/loadtest/README.md)

## Why mono repo?
* Easier to share code between services.
* Easier to manage dependencies
* Single source of thruth for infra, load testing, local setup, code etc.
* Easier to refactor code across services

## Change in coding mindset
* Each service should 
    * be independent and should not depend on other services directly. They should communicate via APIs or messaging queues.
    * have its own database/schema to avoid tight coupling.
    * have its own configuration files.
* Before you develop a new feature, 
    * Think how this feature can be made generic enough to be used by other services in future.
    * Services should not depend on other services. Services can depend on common packages, but common packages should not depend on services. Avoid circular dependencies between services and packages.
    * Check if the feature can be implemented as a shared library or package. If it doesnt exist, create a new package under `common/` directory.
* Write tests for your packages and services.
    * Each service and common package should have its own tests.
    * Assume the next developer is a newbie and is bound to screw up. Write tests to cover all edge cases.

## Assumptions
* This kind of monorepo setup can help small to medium sized teams (up to 20-30 developers) working on related services.
* All services are written in Go. Common packages are also in Go.
* Services are loosely coupled and communicate via APIs or messaging queues.
* Each service has its own database/schema.
* We adhere to good coding standards and practices. We follow strict code review process before merging any code to main branch. 
* We follow the uber style guide for Go code. You can find it [here](https://github.com/uber-go/guide/blob/master/style.md)

## What all a monorepo can contain?
* Services, each in its own directory under `services/`. Common packages under `common/`.
* Local setup scripts and docker compose files under `infra/local_setup/`
* Load testing scripts and documentation under `infra/loadtest/`
* CI/CD pipelines and scripts under `infra/ci_cd/`
* Database index creation and migration scripts under `infra/db/`. Redis lua scripts, Postgres functions etc.
* Documentation related to architecture, design decisions, coding guidelines etc. under `docs/`
* Configuration files for different environments under `config/`
* Runbooks and operational documentation under `ops/`
* Lots of documentation in `README.md` files in respective directories.
* Common `Makefile` at the root to build, test, deploy services and manage local setup.
* Grafana, Prometheus, Loki setup for monitoring and logging under `infra/monitoring/`

## Branching strategy
* `main` branch is the production branch. This branch should always be stable and deployable.
* Create feature branches from `main` for new features or bug fixes. Use descriptive names for feature branches, e.g., `feature/user-authentication`, `bugfix/login-issue`.
* Once the feature or bug fix is complete, create a pull request to merge the changes back into `main`. Ensure that the code is reviewed and tested before merging. You will need at least one approval from another developer.
* After merging, delete the feature branch to keep the repository clean.
* If the feature
    * is small - Create a single PR for the feature branch.
    * is large - Create multiple PRs for the feature branch. Each PR should be small. You should create a seperate release branch from main for the large feature. Merge small PRs into the release branch. Once the feature is complete, create a PR to merge the release branch into main.
* Regularly sync your feature branches with `main` to avoid merge conflicts.

