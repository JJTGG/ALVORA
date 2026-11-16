# Recovery and Portability

## Recovery objective

An operator must be able to recover ALVORA after service failure, node loss, corrupted state, failed upgrade or accidental deletion without relying on the unavailable control plane.

## Required recovery materials

- platform release artifacts;
- control-plane database backups;
- application and configuration exports;
- database backups and restore metadata;
- object-storage data and manifests;
- encryption and recovery keys;
- deployment and resource state;
- documented bootstrap and restore procedures.

## Backup policy

Backups should be encrypted, versioned, integrity-checked and tested through scheduled restoration. Backup retention and deletion must be explicit. A backup that has never been restored is not proven recoverable.

## Offline recovery

A recovery tool should inspect local state, restore a compatible release, recover metadata, re-register workers and bring services back under management. It must not require the failed control plane to execute the first recovery steps.

## Portability

Users should be able to export source references, images or build artifacts, environment configuration, database dumps, object data and domain configuration. Export formats and ownership boundaries should be documented before production claims are made.

## Failure handling

Recovery procedures must distinguish transient failure, degraded service, data loss and unrecoverable state. Every destructive action should define confirmation, retention and rollback behavior.
