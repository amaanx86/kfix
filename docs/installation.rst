Installation
============

kfix can be installed via pre-built binaries or compiled from source.

Binary Installation (Recommended)
----------------------------------

Download the latest release for your platform from:
https://github.com/amaanx86/kfix/releases

.. note::
   The ``/releases/latest/download/`` URL automatically redirects to the newest release.
   You can also replace ``latest`` with a specific version like ``v1.0.0``.

**macOS (Apple Silicon)**::

    curl -LO https://github.com/amaanx86/kfix/releases/latest/download/kfix-darwin-arm64
    sudo mv kfix-darwin-arm64 /usr/local/bin/kfix
    sudo chmod +x /usr/local/bin/kfix

**macOS (Intel)**::

    curl -LO https://github.com/amaanx86/kfix/releases/latest/download/kfix-darwin-amd64
    sudo mv kfix-darwin-amd64 /usr/local/bin/kfix
    sudo chmod +x /usr/local/bin/kfix

**Linux (x86_64)**::

    curl -LO https://github.com/amaanx86/kfix/releases/latest/download/kfix-linux-amd64
    sudo mv kfix-linux-amd64 /usr/local/bin/kfix
    sudo chmod +x /usr/local/bin/kfix

**Linux (ARM64)**::

    curl -LO https://github.com/amaanx86/kfix/releases/latest/download/kfix-linux-arm64
    sudo mv kfix-linux-arm64 /usr/local/bin/kfix
    sudo chmod +x /usr/local/bin/kfix

**Windows**::

    # Download kfix-windows-amd64.exe from the releases page
    # Add to your PATH

Verify Installation
-------------------

Check that kfix is installed correctly::

    kfix version

You should see output similar to::

     _   __  _____   _____   __   __
    | | / / |  ___| |_   _|  \ \ / /
    | |/ /  | |_      | |     \ V / 
    |    \  |  _|     | |     /   \ 
    | |\  \ | |      _| |_   / / \ \
    \_| \_/ \_|     |_____| \/   \/

    Version: v1.0.0
    Source: https://github.com/amaanx86/kfix
    Docs: https://kfix.readthedocs.io

Build from Source
-----------------

**Requirements**

* Go 1.25 or later
* Git
* Make (optional, but recommended)

**Using Makefile (Recommended)**

Clone the repository::

    git clone https://github.com/amaanx86/kfix.git
    cd kfix

Build the binary::

    make build

This creates a ``kfix`` binary in the current directory with version information from git tags.

Install to system::

    make install

**Using go build**

For a simple build without version information::

    go build -o kfix

For a build with version information::

    VERSION=$(git describe --tags --always --dirty 2>/dev/null || echo "dev")
    go build -ldflags "-X github.com/amaanx86/kfix/cmd.version=$VERSION" -o kfix

**Using go install**

Install the latest version directly::

    go install github.com/amaanx86/kfix@latest

Building Release Binaries
-------------------------

For maintainers building multi-platform releases::

    make release

This builds binaries for all supported platforms in the ``dist/`` directory:

* Linux (amd64, arm64)
* macOS (amd64, arm64)
* Windows (amd64)

Version Management
------------------

When building with the Makefile, version information is automatically extracted from git tags.

To create a new release::

    git tag v1.0.0
    git push origin v1.0.0
    make build

The binary will be built with version ``v1.0.0`` embedded.

Creating GitHub Releases
-------------------------

For maintainers publishing releases:

1. **Build release binaries**::

    git tag v1.0.0
    make release

2. **Create GitHub release**:

   * Go to https://github.com/amaanx86/kfix/releases
   * Click "Create a new release"
   * Choose tag ``v1.0.0``
   * Upload binaries from ``dist/`` directory

3. **Binary names must match**:

   * ``kfix-linux-amd64``
   * ``kfix-linux-arm64``
   * ``kfix-darwin-amd64``
   * ``kfix-darwin-arm64``
   * ``kfix-windows-amd64.exe``

Once published, the ``/releases/latest/download/`` URLs will automatically work.
