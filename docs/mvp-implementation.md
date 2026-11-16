# ALVORA MVP Implementation Specification

- **Status:** Proposed implementation plan
- **Scope:** Self-hosted single-node foundation

## Outcome

The first implementation must install on one operator-controlled Linux node, start a minimal ALVORA control plane and worker, persist resource state, and demonstrate one complete application lifecycle.

## First vertical slice

```text
bootstrap
  -> initialize control-plane database
  -> create administrator
  -> create organization, project and environment
  -> register application
  -> accept a local source directory or archive
  -> build an immutable application artifact
  -> start the workload
  -> expose a local route
  -> report status and logs
  -> stop, replace or roll back the release
```

PostgreSQL provisioning, object storage, authentication service, jobs and domains remain planned resource types. They should not block proving the application lifecycle.

## Suggested repository shape

```text
cmd/alvora/              # installer and operator CLI
cmd/control-plane/       # API and reconciliation process
cmd/worker/               # build and runtime worker
internal/api/             # request validation and handlers
internal/auth/            # bootstrap identity and sessions
internal/control/         # resource state and reconciliation
internal/build/           # source intake and artifact creation
internal/runtime/         # workload runtime adapter
internal/gateway/         # route adapter
internal/store/           # database access and migrations
internal/events/          # durable events and audit records
migrations/               # versioned schema
bootstrap/                # manifests and recovery assets
docs/                     # architecture and runbooks
```

This is a starting boundary, not a commitment to many deployable services.

## Initial resource records

The first schema should support organizations, memberships, projects, environments, applications, application releases, resource specifications and observed status, reconciliation tasks, audit events, and encrypted configuration references.

Every record needs an immutable ID, scope, generation or version, lifecycle state, timestamps and an explicit deletion policy.

## Bootstrap contract

The installer must be idempotent, validate prerequisites, initialize local directories, create or migrate the control-plane database, generate or accept administrator credentials, start the minimum services, run health checks and print recovery information. It must fail safely and never log secrets.

## Self-deployment proof

Self-deployment is not complete until a clean-node test can install a released ALVORA artifact, build ALVORA from a declared source revision, deploy the new control-plane and worker artifacts through ALVORA, verify readiness and a sample application deployment, perform a compatible schema upgrade, and roll back the platform release using offline recovery tooling.

## Explicit exclusions

The MVP does not include multi-node scheduling, public multi-tenancy, multi-region replication, global edge networking, billing, arbitrary hostile-code execution or managed-service dependencies.
