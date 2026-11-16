# Security Architecture

## Security objectives

ALVORA must protect control-plane state, isolate projects, limit workloads, protect secrets, authenticate administrative actions and preserve audit evidence.

## Boundaries

- user to public API;
- organization to organization;
- control plane to data plane;
- build worker to host;
- application to platform services;
- workload to workload;
- operator to recovery material.

## Initial controls

Use authenticated sessions or tokens, tenant-scoped authorization, schema validation, rate limits, least-privilege service credentials, encrypted secrets, secure defaults, resource quotas, network restrictions, TLS, audit logging and explicit deletion confirmation.

## Workload safety

Arbitrary customer code must not run with host privileges or control-plane credentials. Build and runtime execution require resource limits, restricted capabilities, filesystem boundaries, timeouts and a documented network egress policy. Stronger sandboxing is required before accepting hostile multi-tenant workloads.

## Secrets

Secrets must be encrypted at rest, redacted from logs, scoped to environments, rotatable and unavailable to unrelated projects. Recovery keys require separate handling and should not be stored in ordinary application configuration.

## Security posture

Threat modeling, dependency review, backup encryption, incident runbooks and security testing are release requirements. This document defines direction; it does not claim that these controls are implemented.
