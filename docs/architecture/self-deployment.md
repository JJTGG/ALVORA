# Self-Deployment

## Goal

ALVORA must be able to build, release, upgrade and recover its own platform services through the same deployment model offered to users.

## Release flow

```text
Source revision
  -> reproducible build
  -> tests and security checks
  -> signed artifacts
  -> internal registry
  -> migration plan
  -> canary deployment
  -> health verification
  -> traffic promotion
  -> retained rollback version
```

## Required safeguards

- immutable source revision and artifacts;
- compatibility checks for database migrations;
- staged rollout and health checks;
- readiness and liveness signals;
- automatic rollback for failed promotion;
- operator approval for dangerous migrations;
- preserved recovery version;
- audit trail for every release step.

## Control-plane upgrades

The platform should update components in an order that preserves API, database and worker compatibility. Schema changes should support an expand-and-contract approach where practical. An offline recovery tool must be able to restore a previous release without depending on the new control plane.

## Self-hosting test

A clean installation should be able to build ALVORA from a declared source revision, deploy it through ALVORA, verify health, perform a controlled update and roll back. Until this test passes, self-deployment remains a goal rather than a capability.
