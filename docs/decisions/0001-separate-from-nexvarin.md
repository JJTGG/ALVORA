# ADR 0001: Keep ALVORA Separate from NEXVARIN

- **Status:** Accepted
- **Date:** 2026-11-16

## Context

ALVORA is a self-hosted full-stack developer cloud. NEXVARIN is an independent infrastructure and operations project with its own existing architecture, scope and evolution path.

The projects overlap in areas such as durable state, workers, builds, artifacts, runtime execution, routing and reconciliation. However, ALVORA additionally requires a developer-facing resource model, backend services, self-hosting bootstrap, tenant boundaries and platform self-deployment.

## Decision

ALVORA will remain separate from NEXVARIN. It will have its own repository, schemas, APIs, bootstrap process, runtime abstractions, deployment pipeline and recovery tools.

NEXVARIN may provide architectural lessons through read-only comparison, but ALVORA will not depend on its code, database, API, branding or deployment model.

## Consequences

Positive consequences include independent experimentation, clearer ownership, reduced coupling and freedom to choose a resource-oriented product architecture. The cost is possible duplication and the need to re-implement capabilities that appear similar.

If future interoperability is valuable, it must be introduced through explicit, versioned adapters or documented interfaces rather than hidden reuse.

## Reconsideration criteria

This decision may be revisited only if both projects have stable contracts, the integration provides a clear operational benefit, isolation remains intact and the change is approved as a new architecture decision.
