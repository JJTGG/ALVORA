# Architecture Overview

## Layers

ALVORA is organized into three conceptual layers:

1. **Bootstrap layer** — Installs, initializes and recovers the minimum viable platform.
2. **Platform layer** — Manages identity, projects, resources, scheduling, builds, deployments and operations.
3. **Workload layer** — Runs applications, jobs, databases, storage and other project resources.

```text
CLI / Dashboard / SDK
          |
      Public API
          |
 Resource-oriented control plane
          |
 Controllers, scheduler and workers
     /       |        \\
  Builds  Applications  Backend resources
          |
       Gateway, TLS and service discovery
```

## Control plane

The control plane stores desired configuration, permissions, lifecycle state, observed status, events, audit records and usage data. It should remain available even when individual workloads fail.

## Data plane

The data plane executes builds and workloads, stores artifacts, serves application traffic and runs platform-managed databases and storage. It must be isolated from control-plane credentials and administrative interfaces.

## Initial topology

The first implementation targets one operator-controlled node containing the control plane, worker, runtime, database, object storage, registry, gateway and observability components. Components should retain boundaries that permit later multi-node placement.

## Implementation posture

Start as a modular monolith plus isolated workers. Extract services only when a boundary improves reliability, security or scaling. Kubernetes may become an execution adapter, but it is not the public architecture.
