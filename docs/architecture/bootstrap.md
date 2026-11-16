# Bootstrap and Genesis

## Problem

A self-hosting platform cannot deploy its first version from nothing. ALVORA therefore needs an explicit genesis mechanism that installs the smallest trusted runtime and then hands normal releases to ALVORA itself.

## Bootstrap stages

1. Operator obtains a signed release, source bundle or offline distribution.
2. Installer validates prerequisites and initializes local state.
3. Minimal runtime, database, registry and worker are started.
4. Initial administrator and recovery material are created.
5. Platform services are built or imported and registered.
6. Health checks confirm that ALVORA can manage itself.
7. Subsequent releases use the normal self-deployment lifecycle.

## Bootstrap properties

The installer should be deterministic, idempotent, resumable and safe to run without internet access. It must document filesystem layout, ports, credentials, encryption material and recovery steps.

## Minimum bootstrap set

- container or workload runtime;
- control-plane database;
- minimal API;
- first worker;
- internal artifact registry;
- gateway;
- local configuration and recovery tools.

## Trust boundary

Bootstrap credentials are highly sensitive. They must be generated or supplied explicitly, stored securely, rotated after initialization and excluded from application workloads. A recovery path must exist if the normal control plane is unavailable.

Bootstrap is a proposed design until an installer and clean-machine acceptance test exist.
