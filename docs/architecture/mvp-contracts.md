# ALVORA MVP Contracts

## Resource lifecycle

Resources use these states:

```text
requested -> provisioning -> ready
                         \-> failed
ready -> updating -> ready
ready -> deleting -> deleted
```

A failed resource retains its desired specification, observed error, generation and retry information. Reconciliation must be idempotent.

## Application specification

An application specification should identify the project and environment, source input, build configuration, artifact or release policy, runtime command, environment and secret references, resource limits, health checks and route intent.

## Release status

A release should expose source revision, immutable artifact reference, build status, runtime status, health status, timestamps and rollback eligibility. Mutable tags must not be the sole deployment identity.

## Controller boundary

The API persists desired state. A reconciler observes resources and creates durable tasks. Workers claim tasks with leases, perform idempotent actions and publish observed status. The API must not depend on an in-memory task queue for correctness.

## Error contract

Errors should include a stable code, human-readable message, retryability and a correlation identifier. Secrets, tokens and internal host credentials must never appear in responses or logs.

## Versioning

Public resource fields require compatibility rules before release. Database migrations should use an expand-and-contract approach where a rolling upgrade may occur.
