# Resource Model

## Hierarchy

```text
Organization
└── Project
    ├── Environment
    │   ├── Application
    │   ├── Database
    │   ├── Storage bucket
    │   ├── Function or job
    │   ├── Domain
    │   └── Secret set
    └── Repository connection
```

## Resource identity

Each resource needs an immutable identifier, resource type, organization and project scope, environment scope where applicable, region or placement, desired specification, observed status, generation, lifecycle state and timestamps.

## Environments

Development, preview and production environments should be isolated through configuration, credentials, resource quotas and deployment policies. Preview environments may be created and destroyed automatically, while production deletion requires stronger safeguards.

## Desired versus observed state

Desired state expresses what the user requested. Observed state expresses what the platform currently sees. Controllers reconcile the difference and publish conditions explaining readiness, failure and pending work.

## Resource relationships

Dependencies must be explicit. An application may depend on a database, bucket, secret set and domain. Deletion must account for dependents, retention policy and recoverability rather than silently destroying data.

## Initial resource types

The first vertical slice should include organization, project, environment, application, deployment, database, bucket, secret set and domain. Additional types should be introduced only with lifecycle and recovery semantics.
