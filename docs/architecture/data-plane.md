# Data Plane

## Responsibility

The data plane builds and runs customer workloads and platform-managed services. It should continue to enforce isolation and resource limits even when the control plane is degraded.

## Initial components

- build workers;
- OCI artifact registry and cache;
- application runtime;
- scheduled and background job workers;
- PostgreSQL service;
- object storage service;
- gateway and service discovery;
- logs, metrics and traces.

## Workload contract

A workload receives an immutable artifact or declared image, validated configuration, scoped environment variables, resource limits, network policy and a health-check contract. It must not receive control-plane administration credentials.

## Runtime progression

Start with a single-node container runtime and explicit worker agents. Add node registration, capacity tracking, placement, draining and failure reconciliation before multi-node scheduling. MicroVMs or stronger sandboxing may be evaluated for higher-risk workloads.

## Storage

Artifacts, application data, database data, logs and backups require separate lifecycle and access policies. Data-plane services must expose health and capacity signals to the control plane without allowing arbitrary state mutation.

## External dependencies

Self-hosted implementations are the default direction. External registries, object stores, compute providers or managed databases may be adapters, not assumptions in the core API.
