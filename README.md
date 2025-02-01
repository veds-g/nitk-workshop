# nitk-workshop
word count using numaflow

## Installation

The following steps are to install the pipeline in your Kubernetes cluster and run it to show how it works.

### Prerequisites

- [Docker Engine](https://docs.docker.com/desktop/install/windows-install/)
- [Chocolatey](https://chocolatey.org/install)
```bash
Set-ExecutionPolicy Bypass -Scope Process -Force; [System.Net.ServicePointManager]::SecurityProtocol = [System.Net.ServicePointManager]::SecurityProtocol -bor 3072; iex ((New-Object System.Net.WebClient).DownloadString('https://community.chocolatey.org/install.ps1'))
```
- [`kubectl`](https://kubernetes.io/docs/tasks/tools/#kubectl)
```bash
choco install kubernetes-cli
```
- [`k3d`](https://k3d.io/v5.6.0/#install-script)
```bash
choco install k3d
```
- `git`
```bash
choco install git
```

### Installation Steps

1. Creating a local Kubernetes cluster using k3d

```bash
k3d cluster create ci-cd-workshop-cluster --api-port 6550 -p "8081:80@loadbalancer" --agents 2
```

2. Install Numaflow

```bash
kubectl create ns numaflow-system
kubectl apply -n numaflow-system -f https://raw.githubusercontent.com/numaproj/numaflow/stable/config/install.yaml
kubectl apply -f https://raw.githubusercontent.com/numaproj/numaflow/stable/examples/0-isbsvc-jetstream.yaml
```

3. Create the word count pipeline using Numaflow

```bash
kubectl apply -f https://raw.githubusercontent.com/veds-g/nitk-workshop/master/pipeline.yaml
```

4. View the pipeline

```bash
kubectl port-forward svc/numaflow-server 8443 -n numaflow-system
```

Open the browser "https://localhost:8443/", then go to Numaflow UI, select `default` namespace, and click the `word-count-pl` pipeline.
