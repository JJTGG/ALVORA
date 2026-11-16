# ALVORA Principles

1. **Self-hosted core** — Core operation must work on operator-controlled infrastructure.
2. **External services are adapters** — Git providers, cloud compute, DNS and payment systems are optional integrations.
3. **Bootstrap is explicit** — Installation and recovery must have documented, reproducible paths.
4. **ALVORA deploys ALVORA** — Platform services should use the normal build and deployment lifecycle.
5. **Control plane and data plane are separate** — Customer workloads must not compromise platform state.
6. **Desired state is durable** — Resources have declared configuration and observable status.
7. **Recovery is a feature** — Export, backup, rollback and offline recovery are design requirements.
8. **Portable state** — Applications, artifacts, configuration and data should have documented export paths.
9. **Secure defaults** — Isolation, least privilege, validation, quotas and auditability precede convenience.
10. **Evidence over claims** — Documentation distinguishes implemented behavior from proposals.
11. **Minimal first system** — Begin with a modular, testable single-node platform before distributed expansion.
12. **NEXVARIN remains independent** — Similar lessons do not justify a runtime or code dependency.
