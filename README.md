<p align="center">
  <img src="https://github.com/user-attachments/assets/962dfe26-126c-4f93-b279-052031f7c3f0" alt="kfix">
</p>

<p align="center">
  <strong>Opinionated Kubernetes YAML formatter for clean, consistent manifests</strong>
</p>

<p align="center">
  <a href="https://github.com/amaanx86/kfix/releases"><img src="https://img.shields.io/github/v/release/amaanx86/kfix" alt="Release"></a>
  <a href="https://kfix.readthedocs.io"><img src="https://img.shields.io/badge/docs-readthedocs-blue" alt="Documentation"></a>
  <a href="LICENSE"><img src="https://img.shields.io/github/license/amaanx86/kfix" alt="License"></a>
</p>

---

## Overview

kfix is a command-line tool that formats Kubernetes YAML files with context-aware field ordering and consistent styling. It understands Kubernetes resource structure and automatically applies best-practice formatting rules.

> **Note**: kfix formats **valid** YAML manifests. It does not fix malformed or syntactically incorrect YAML files.

## Features

- **K8s-Aware Formatting** - Understands Kubernetes resource structure
- **Field Reordering** - Automatically orders fields following K8s conventions  
- **Multi-Document Support** - Handles multiple manifests in a single file
- **Configurable Indentation** - Customize spacing (default: 2 spaces)
- **In-Place Editing** - Update files directly or output to stdout
- **Context-Aware** - Properly formats nested structures (containers, volumes, etc.)

## Installation

### Binary Installation (Recommended)

Download pre-built binaries from [GitHub Releases](https://github.com/amaanx86/kfix/releases):

> **Note**: Replace `latest` with a specific version (e.g., `v1.0.0`) if needed.

**macOS**:
```bash
# Apple Silicon
curl -LO https://github.com/amaanx86/kfix/releases/latest/download/kfix-darwin-arm64
sudo mv kfix-darwin-arm64 /usr/local/bin/kfix
sudo chmod +x /usr/local/bin/kfix

# Intel
curl -LO https://github.com/amaanx86/kfix/releases/latest/download/kfix-darwin-amd64
sudo mv kfix-darwin-amd64 /usr/local/bin/kfix
sudo chmod +x /usr/local/bin/kfix
```

**Linux**:
```bash
# x86_64
curl -LO https://github.com/amaanx86/kfix/releases/latest/download/kfix-linux-amd64
sudo mv kfix-linux-amd64 /usr/local/bin/kfix
sudo chmod +x /usr/local/bin/kfix

# ARM64
curl -LO https://github.com/amaanx86/kfix/releases/latest/download/kfix-linux-arm64
sudo mv kfix-linux-arm64 /usr/local/bin/kfix
sudo chmod +x /usr/local/bin/kfix
```

**Alternatively**, visit the [releases page](https://github.com/amaanx86/kfix/releases) and download manually.

### Build from Source

**Using Makefile** (Recommended):
```bash
git clone https://github.com/amaanx86/kfix.git
cd kfix
make build
```

**Using go build**:
```bash
go build -o kfix
```

**Using go install**:
```bash
go install github.com/amaanx86/kfix@latest
```

## Quick Start

Format a manifest:
```bash
kfix format deployment.yaml
```

Format multiple files:
```bash
kfix format deployment.yaml service.yaml configmap.yaml
```

Format all YAML files:
```bash
kfix format *.yaml
```

Format from stdin:
```bash
cat deployment.yaml | kfix format
kubectl get deployment nginx -o yaml | kfix format
```

Format in-place:
```bash
kfix format -w deployment.yaml
```

Custom indentation:
```bash
kfix format -i 4 deployment.yaml
```

## Usage Examples

### Format kubectl Output

```bash
kubectl get deployment nginx -o yaml | kfix format > nginx-deployment.yaml
```

### Format Before Commit

```bash
kfix format -w k8s/*.yaml
git add k8s/
git commit -m "style: format Kubernetes manifests"
```

### Format Helm Templates

```bash
helm template myapp ./chart | kfix format > rendered.yaml
```

### Format Kustomize Output

```bash
kustomize build . | kfix format > manifest.yaml
```

## Field Ordering

kfix applies Kubernetes-aware field ordering:

**Top-level fields**:
```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: myapp
  namespace: production
spec:
  replicas: 3
status:
  availableReplicas: 3
```

**Container fields**:
```yaml
containers:
- name: app
  image: myapp:1.0.0
  command: ["/app"]
  args: ["--config=/etc/app/config.yaml"]
  ports:
  - containerPort: 8080
  env:
  - name: LOG_LEVEL
    value: info
  resources:
    limits:
      memory: 512Mi
      cpu: 500m
  volumeMounts:
  - name: config
    mountPath: /etc/app
  livenessProbe:
    httpGet:
      path: /healthz
      port: 8080
```

## What kfix Does

- Reorders fields according to K8s conventions  
- Applies consistent indentation  
- Formats nested structures (containers, volumes, etc.)  
- Handles multi-document YAML files  
- Validates input as Kubernetes resources  

## What kfix Does NOT Do

- Fix malformed YAML syntax  
- Validate resource schema  
- Add missing required fields  
- Modify resource logic or behavior  
- Convert between API versions  

## Documentation

Comprehensive documentation available at **[kfix.readthedocs.io](https://kfix.readthedocs.io)**

- [Installation Guide](https://kfix.readthedocs.io/en/latest/installation.html)
- [Usage Guide](https://kfix.readthedocs.io/en/latest/usage.html)
- [Configuration](https://kfix.readthedocs.io/en/latest/configuration.html)

## Contributing

Contributions are welcome! Please feel free to submit issues or pull requests.

## License

See [LICENSE](LICENSE)
