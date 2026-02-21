Configuration
=============

Field Ordering Rules
--------------------

kfix applies Kubernetes-aware field ordering to make manifests more readable
and consistent. Fields are ordered according to their importance and common usage patterns.

Top-Level Fields
~~~~~~~~~~~~~~~~

Resource-level fields are ordered:

1. ``apiVersion`` - API version of the resource
2. ``kind`` - Type of Kubernetes resource
3. ``metadata`` - Resource metadata
4. ``spec`` - Resource specification
5. ``data`` - ConfigMap/Secret data
6. ``status`` - Resource status (typically read-only)

**Example**::

    apiVersion: apps/v1
    kind: Deployment
    metadata:
      name: myapp
    spec:
      replicas: 3
    status:
      availableReplicas: 3

Metadata Fields
~~~~~~~~~~~~~~~

Within ``metadata``, fields are ordered:

1. ``name`` - Resource name (required)
2. ``namespace`` - Namespace location
3. ``labels`` - Label selectors
4. ``annotations`` - Additional metadata
5. Other fields (alphabetically)

**Example**::

    metadata:
      name: myapp
      namespace: production
      labels:
        app: myapp
        version: v1.0.0
      annotations:
        description: "My application"

Container Spec Fields
~~~~~~~~~~~~~~~~~~~~~

Within container specifications, fields are ordered:

1. ``name`` - Container name
2. ``image`` - Container image
3. ``imagePullPolicy`` - Image pull strategy
4. ``command`` - Entrypoint override
5. ``args`` - Arguments to command
6. ``workingDir`` - Working directory
7. ``ports`` - Exposed ports
8. ``env`` - Environment variables
9. ``resources`` - Resource requests/limits
10. ``volumeMounts`` - Volume mount points
11. ``livenessProbe`` - Liveness health check
12. ``readinessProbe`` - Readiness health check
13. ``startupProbe`` - Startup health check
14. ``lifecycle`` - Lifecycle hooks
15. ``securityContext`` - Security settings

**Example**::

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
        requests:
          memory: 256Mi
          cpu: 250m
      volumeMounts:
      - name: config
        mountPath: /etc/app
      livenessProbe:
        httpGet:
          path: /healthz
          port: 8080

Pod Spec Fields
~~~~~~~~~~~~~~~

Within pod specifications, fields are ordered:

1. ``replicas`` - Number of replicas (for Deployments/StatefulSets)
2. ``selector`` - Label selector
3. ``template`` - Pod template
4. ``serviceName`` - Service name (for StatefulSets)
5. ``serviceAccountName`` - Service account
6. ``serviceAccount`` - Legacy service account
7. ``automountServiceAccountToken`` - Token mounting
8. ``nodeSelector`` - Node selection
9. ``affinity`` - Pod affinity rules
10. ``tolerations`` - Node tolerations
11. ``initContainers`` - Init containers
12. ``containers`` - Main containers
13. ``volumes`` - Volume definitions
14. ``restartPolicy`` - Restart behavior
15. ``terminationGracePeriodSeconds`` - Graceful shutdown time

**Example**::

    spec:
      replicas: 3
      selector:
        matchLabels:
          app: myapp
      template:
        metadata:
          labels:
            app: myapp
        spec:
          serviceAccountName: myapp
          containers:
          - name: app
            image: myapp:1.0.0
          volumes:
          - name: config
            configMap:
              name: myapp-config
          restartPolicy: Always

Indentation Settings
--------------------

Default: 2 Spaces
~~~~~~~~~~~~~~~~~

kfix uses 2-space indentation by default, which is the Kubernetes standard::

    apiVersion: v1
    kind: Pod
    metadata:
      name: myapp
    spec:
      containers:
      - name: app
        image: myapp:1.0.0

Custom Indentation
~~~~~~~~~~~~~~~~~~

Use the ``--indent`` flag to customize::

    kfix format --indent 4 deployment.yaml

Output with 4 spaces::

    apiVersion: v1
    kind: Pod
    metadata:
        name: myapp
    spec:
        containers:
        - name: app
          image: myapp:1.0.0

Formatting Behavior
-------------------

What kfix Does
~~~~~~~~~~~~~~

✓ Reorders fields according to K8s conventions

✓ Applies consistent indentation

✓ Formats nested structures (containers, volumes, etc.)

✓ Handles multi-document YAML files

✓ Preserves comments (where possible)

✓ Maintains data types (strings, numbers, booleans)

What kfix Does NOT Do
~~~~~~~~~~~~~~~~~~~~~

✗ Fix malformed YAML syntax

✗ Validate resource schema

✗ Add missing required fields

✗ Modify resource logic or behavior

✗ Convert between API versions

✗ Apply cluster-specific policies

Limitations
-----------

Valid YAML Required
~~~~~~~~~~~~~~~~~~~

kfix requires syntactically valid YAML. It will reject::

    # Invalid indentation
    metadata:
      name: test
       namespace: default  # Incorrect indent
    
    # Missing colons
    metadata
      name: test
    
    # Unclosed quotes
    metadata:
      name: "test

Kubernetes Resources Only
~~~~~~~~~~~~~~~~~~~~~~~~~

kfix only formats files with ``apiVersion`` and ``kind`` fields::

    # Will be formatted
    apiVersion: v1
    kind: ConfigMap
    data:
      key: value
    
    # Will be rejected (no apiVersion/kind)
    data:
      key: value

Future Configuration
--------------------

Configuration File Support
~~~~~~~~~~~~~~~~~~~~~~~~~~

Planned support for ``.kfix.yaml`` or ``.kfix.yml`` configuration files.

Future configuration options:

* Custom field ordering preferences
* Field exclusion rules
* Resource-type-specific formatting
* Indentation per resource type
* Comment preservation strategies

**Example future config**::

    # .kfix.yaml
    indent: 2
    
    fieldOrder:
      container:
        - name
        - image
        - ports
        - env
    
    excludeFields:
      - status
      - managedFields

Stay tuned for updates at: https://github.com/amaanx86/kfix
