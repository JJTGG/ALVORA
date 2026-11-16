# Control Plane

## Responsibility

The control plane is the authoritative management system for ALVORA resources. It accepts authenticated intent, persists desired configuration, schedules work and reports observed status.

## Core responsibilities

- identity and organization membership;
- projects and environments;
- resource definitions and lifecycle state;
- deployment and build coordination;
- secrets and access policies;
- domains and routing intent;
- audit events and usage records;
- reconciliation and recovery requests.

## Resource lifecycle

Every managed resource should expose a stable identifier, owner, project and environment association, desired configuration, observed status, lifecycle state, timestamps, version or generation, and deletion behavior.

A controller compares desired and observed state, creates durable work, performs idempotent changes and records the result. Reconciliation must be safe to repeat after crashes.

## API direction

The public API should be resource-oriented. Operations, jobs and events remain implementation mechanisms. API contracts must validate input, enforce tenant scope, return stable error shapes and avoid exposing internal scheduler details.

## State model

The control plane must distinguish requested, provisioning, ready, updating, failed, deleting and deleted states. A failed operation must preserve enough information for diagnosis and retry without silently losing intent.

## Initial implementation

Use a single deployable service with clear modules, PostgreSQL-backed state and a worker process. Add a queue and event mechanism before introducing distributed controllers.
