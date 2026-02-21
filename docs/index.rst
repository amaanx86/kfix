kfix Documentation
==================

**Opinionated Kubernetes YAML formatter for clean, consistent manifests.**

kfix is a command-line tool that formats Kubernetes YAML files with context-aware
field ordering and consistent styling. It understands Kubernetes resource structure
and applies best-practice formatting rules automatically.

.. note::
   kfix formats valid YAML manifests. It does not fix malformed or syntactically
   incorrect YAML files. Ensure your YAML is valid before formatting.

Key Features
------------

* **K8s-Aware Formatting**: Understands Kubernetes resource structure
* **Field Ordering**: Automatically reorders fields following K8s conventions
* **Multi-Document Support**: Handles multiple manifests in a single file
* **Configurable Indentation**: Customize spacing to match your style
* **In-Place Editing**: Update files directly or output to stdout
* **Context-Aware**: Nested structures (containers, volumes, etc.) are properly ordered

Quick Start
-----------

Install from GitHub releases::

    # Download the latest release for your platform
    # Visit https://github.com/amaanx86/kfix/releases
    
    # macOS/Linux
    sudo mv kfix /usr/local/bin/
    chmod +x /usr/local/bin/kfix

Format a manifest::

    kfix format deployment.yaml

Format in-place::

    kfix format -w deployment.yaml

Documentation Index
-------------------

.. toctree::
   :maxdepth: 2
   :caption: Contents:

   installation
   usage
   configuration

Project Links
-------------

* **Source Code**: https://github.com/amaanx86/kfix
* **Releases**: https://github.com/amaanx86/kfix/releases
* **Issues**: https://github.com/amaanx86/kfix/issues
