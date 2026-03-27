# IDP Project

This is a modular Internal Developer Platform (IDP) CLI tool built in Go.

## Folder Structure

```
idp-project/
├── cli/           # Contains all Go code for the idp tool
├── sample-app/    # A sample Python Flask application and its Dockerfile
└── k8s/           # Kubernetes configuration files (deployment, service)
```

## Creating the Executable

To compile the CLI tool for yourself:
1. Open up a terminal and navigate to the `cli/` folder (`cd cli`).
2. Run the build command:
   ```bash
   go build -o idp.exe
   ```

## Usage

Make sure Docker and Minikube are running. Since the core commands depend on relative paths properly targeting your codebase, you should always run the `deploy` tool from **inside the `cli/` folder**:

```bash
cd cli
.\idp.exe deploy
```

The automated pipeline will run through the following structured check:
1. Validate Docker is running
2. Build the Docker Image for the code inside `../sample-app`
3. Load the new Image directly into the running Minikube cluster
4. Automatically apply Kubernetes configurations living in `../k8s/deployment.yaml`
