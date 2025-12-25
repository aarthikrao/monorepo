# Setting up your local environment

You need to install the following 

1. **Docker:** Follow the instructions [here](https://docs.docker.com/get-docker/) to install Docker on your machine. You can use Colima as your Docker backend for better performance on macOS.
2. **Go:** Make sure you have Go installed. You can download it from [here](https://golang.org/dl/). Ensure that your Go version is compatible with the project requirements.
3. **Make:** Ensure you have Make installed on your system. Most Unix-based systems come with Make pre-installed. You can check by running `make --version` in your terminal.
4. **Docker Compose:** Docker Compose is included with Docker Desktop. If you're using Colima, ensure you have the Docker CLI plugins installed. You can refer to the [Docker CLI plugins documentation](https://docs.docker.com/engine/reference/commandline/cli_plugins/) for more details.
Note: Recent versions of Docker use `docker compose` (without a hyphen) as the command for Docker Compose.
You can verify the installation by running `docker compose version` in your terminal.
For Docker to find the plugin, add "cliPluginsExtraDirs" to `~/.docker/config.json`:
    ```bash 
    "cliPluginsExtraDirs": [
        "/opt/homebrew/lib/docker/cli-plugins"
    ]
    ```

### Building the services
Before deploying the local development environment, you need to build the services. You can do this by running the following command from the root of the repository:
```bash
make build
```

### Deploying the local development environment
Once you have all the prerequisites installed, you can set up your local development environment by running the following command from the root of the repository:
```bash
make local-deploy
```
This command will use Docker Compose to spin up all the necessary components and start all the apps defined in the `infra/local_setup/docker-compose.yml` file.

### Checking the status of docker containers
You can check the status of the running containers using:
```bash
docker compose -f infra/local_setup/docker-compose.yml ps
```
This will list all the services that are currently running as part of your local development environment.

### Stopping the local development environment
You can stop the local development environment by running:
```bash
docker compose -f infra/local_setup/docker-compose.yml down
```

### Additional commands
Refer to the `Makefile` for additional commands that can help you manage your local setup, such as building the services or tidying up Go modules.