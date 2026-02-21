Usage Guide
===========

This guide covers common usage patterns for kfix.

.. important::
   kfix formats **valid** YAML only. It will not fix malformed or syntactically
   incorrect YAML. Use a YAML validator first if you encounter issues.

Basic Usage
-----------

**Format a single file**::

    kfix format deployment.yaml

Output:

* Formatted YAML is printed to stdout
* Original file is unchanged
* Exit code 0 on success, 1 on error

**Format multiple files**::

    kfix format deployment.yaml service.yaml configmap.yaml

**Format all YAML files in a directory**::

    kfix format *.yaml
    kfix format manifests/*.yaml

**Format from stdin**::

    cat deployment.yaml | kfix format
    kubectl get pod nginx -o yaml | kfix format

In-Place Editing
----------------

Use the ``-w`` (or ``--in-place``) flag to modify files directly::

    kfix format -w deployment.yaml

.. warning::
   In-place editing modifies files directly. Consider version control or backups.

**Format all manifests in place**::

    kfix format -w manifests/*.yaml

**Safe in-place update pattern**::

    # Preview changes first
    kfix format deployment.yaml
    
    # If satisfied, apply in-place
    kfix format -w deployment.yaml

Indentation Control
-------------------

Customize indentation with the ``-i`` flag (default: 2 spaces)::

    kfix format -i 4 deployment.yaml

**Standard indentations**:

* 2 spaces (default, recommended for K8s)
* 4 spaces (alternative style)

Commands Reference
------------------

kfix format
~~~~~~~~~~~

Format Kubernetes YAML files with consistent styling.

**Syntax**::

    kfix format [flags] [file...]

**Arguments**:

* ``file...`` - One or more YAML files to format (optional)
* If no files provided, reads from stdin

**Flags**:

``-w, --in-place``
    Write formatted output back to the file instead of stdout

``-i, --indent int``
    Number of spaces for indentation (default: 2)

**What the formatter does**:

1. Validates input is a Kubernetes resource (has ``apiVersion`` and ``kind``)
2. Reorders top-level fields: apiVersion → kind → metadata → spec → data/status
3. Reorders metadata fields: name → namespace → labels → annotations
4. Reorders container fields: name → image → command → args → ports → env → resources
5. Applies consistent indentation throughout
6. Handles multi-document YAML (separated by ``---``)

kfix version
~~~~~~~~~~~~

Display version information::

    kfix version

Shows:

* ASCII art banner
* Version number
* Source repository link
* Documentation link

kfix help
~~~~~~~~~

Get help on any command::

    kfix help
    kfix help format
    kfix format --help

Common Workflows
----------------

Format Before Commit
~~~~~~~~~~~~~~~~~~~~

Integrate into your Git workflow::

    # Format all manifests before commit
    kfix format -w k8s/*.yaml
    git add k8s/
    git commit -m "style: format Kubernetes manifests"

Format CI/CD Pipeline
~~~~~~~~~~~~~~~~~~~~~

Check formatting in CI::

    #!/bin/bash
    # Check if manifests are formatted
    
    for file in k8s/*.yaml; do
        if ! diff -u "$file" <(kfix format "$file"); then
            echo "ERROR: $file is not formatted"
            echo "Run: kfix format -w $file"
            exit 1
        fi
    done

Format kubectl Output
~~~~~~~~~~~~~~~~~~~~~

Pipe kubectl output through kfix::

    # Get and format a deployment
    kubectl get deployment nginx -o yaml | kfix format > nginx-deployment.yaml
    
    # Get all resources in a namespace
    kubectl get all -n production -o yaml | kfix format > production-backup.yaml

Format Helm Output
~~~~~~~~~~~~~~~~~~

Format Helm template output::

    helm template myapp ./chart | kfix format > rendered.yaml

Format Kustomize Output
~~~~~~~~~~~~~~~~~~~~~~~

Format Kustomize build output::

    kustomize build . | kfix format > manifest.yaml

Integration with Editors
------------------------

Vim/Neovim
~~~~~~~~~~

Format current buffer::

    :%!kfix format

Or add to ``.vimrc``::

    autocmd FileType yaml nmap <leader>k :%!kfix format<CR>

VS Code
~~~~~~~

Add to tasks.json::

    {
        "label": "Format with kfix",
        "type": "shell",
        "command": "kfix format -w ${file}"
    }

Pre-commit Hook
~~~~~~~~~~~~~~~

Add to ``.git/hooks/pre-commit``::

    #!/bin/bash
    
    # Format staged YAML files
    git diff --cached --name-only --diff-filter=ACMR | grep '\.ya\?ml$' | while read file; do
        kfix format -w "$file"
        git add "$file"
    done

Troubleshooting
---------------

Error: "file does not appear to be valid Kubernetes resource(s)"
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

**Cause**: Input YAML is either:

* Malformed/invalid YAML syntax
* Missing required fields (``apiVersion`` and ``kind``)
* Not a Kubernetes resource

**Solution**:

1. Validate YAML syntax::

    yamllint file.yaml

2. Ensure resource has both ``apiVersion`` and ``kind`` fields

3. kfix only formats Kubernetes resources, not generic YAML

Error: "failed to parse YAML"
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

**Cause**: YAML syntax error (malformed structure)

**Solution**:

1. Fix YAML syntax errors first
2. kfix formats valid YAML, it does not fix broken YAML
3. Use a linter: ``yamllint file.yaml``

Multi-Document Files
~~~~~~~~~~~~~~~~~~~~

kfix handles multiple documents separated by ``---``::

    # Input
    kind: Service
    apiVersion: v1
    ---
    kind: Deployment
    apiVersion: apps/v1
    
    # Output
    apiVersion: v1
    kind: Service
    ---
    apiVersion: apps/v1
    kind: Deployment

Each document is formatted independently.

Examples
--------

Basic Deployment
~~~~~~~~~~~~~~~~

**Before**::

    kind: Deployment
    spec:
      replicas: 3
      selector:
        matchLabels:
          app: nginx
    metadata:
      name: nginx
    apiVersion: apps/v1

**After**::

    apiVersion: apps/v1
    kind: Deployment
    metadata:
      name: nginx
    spec:
      replicas: 3
      selector:
        matchLabels:
          app: nginx

Complex Container Spec
~~~~~~~~~~~~~~~~~~~~~~

**Before**::

    spec:
      containers:
      - resources:
          limits:
            memory: 256Mi
        ports:
        - containerPort: 80
        image: nginx:latest
        name: nginx
        env:
        - name: DEBUG
          value: "true"

**After**::

    spec:
      containers:
      - name: nginx
        image: nginx:latest
        ports:
        - containerPort: 80
        env:
        - name: DEBUG
          value: "true"
        resources:
          limits:
            memory: 256Mi
