# Roadmap

This roadmap is directional. A milestone is complete only when its behavior is implemented, tested and documented.

## Phase 0 — Documentation and decisions

- establish scope and non-goals;
- define bootstrap, control-plane and data-plane boundaries;
- define resource lifecycle and recovery expectations;
- keep ALVORA independent from NEXVARIN.

## Phase 1 — Single-node foundation

- repository and build system;
- bootstrap installer;
- control-plane database;
- identity, organization, project and environment model;
- minimal API and dashboard;
- worker and local runtime;
- health, logs and audit events.

## Phase 2 — Application deployment

- source input and repository adapters;
- reproducible builds;
- internal OCI registry;
- application releases;
- gateway, domains and TLS;
- environment variables and rollback.

## Phase 3 — Backend resources

- self-hosted PostgreSQL provisioning;
- object storage;
- authentication service foundations;
- background jobs and schedules;
- backup and restoration workflows.

## Phase 4 — Self-deployment and hardening

- ALVORA builds and deploys itself;
- upgrade compatibility checks;
- canary and rollback automation;
- security isolation, quotas and rate limits;
- clean-install and disaster-recovery tests.

## Phase 5 — Multi-node operation

- node registration and capacity;
- scheduling and placement;
- workload draining;
- replicated control-plane state;
- regional and failure-domain awareness.

## Deferred exploration

Multi-region data services, edge networking, Anycast, CDN capabilities and advanced sandboxing require separate feasibility and operations work.
