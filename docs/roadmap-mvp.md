# ALVORA MVP Work Sequence

## Milestone 1 — Executable skeleton

Create the implementation foundation, CLI entrypoint, control-plane process, worker process, health endpoint, configuration loader and structured logging.

## Milestone 2 — Durable state

Add migrations, database access, resource tables, lifecycle states, audit events and transaction boundaries. Add characterization tests before changing behavior.

## Milestone 3 — Reconciliation

Implement desired-state persistence, durable tasks, worker leases, retry policy, idempotent reconciliation and restart recovery.

## Milestone 4 — Local application deployment

Implement source intake, build adapter, immutable artifact records, runtime adapter, health checks, logs and release replacement.

## Milestone 5 — Operator workflow

Implement bootstrap, administrator setup, organization/project/environment creation, CLI commands, configuration validation and recovery output.

## Milestone 6 — Self-deployment proof

Build ALVORA from source, deploy it through its own lifecycle, run compatibility checks, perform a controlled upgrade and verify offline rollback.

## Engineering rule

Do not introduce PostgreSQL provisioning, authentication-as-a-service, object storage, multi-node scheduling or public multi-tenancy until the single-node application lifecycle is reliable and tested.
