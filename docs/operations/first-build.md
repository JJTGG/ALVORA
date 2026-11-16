# ALVORA First-Build Acceptance Plan

This plan defines the first implementation gate. It is a test plan, not a claim that the platform already passes it.

## Clean-node prerequisites

- supported Linux distribution;
- container runtime selected by the implementation;
- local persistent storage;
- a local or operator-provided PostgreSQL instance during the earliest bootstrap, unless ALVORA provisions it itself;
- no Vercel, Render, Supabase, Cloudflare or other managed-service dependency.

## Acceptance sequence

1. Run the bootstrap command on a clean node.
2. Confirm that configuration and secrets are created with safe permissions.
3. Confirm that the health endpoint reports control plane and worker readiness.
4. Create an organization, project and environment.
5. Register a sample application from local source.
6. Build an immutable artifact.
7. Deploy the artifact and verify its health endpoint.
8. Read status and logs through the API or CLI.
9. Deploy a second release and verify replacement.
10. Force a failed release and verify that the previous healthy release remains available or is restored.
11. Restart the worker and confirm reconciliation resumes.
12. Back up state, restore it to a clean node and verify resource visibility.

## Required evidence

Each step should produce machine-readable test output, versioned configuration, relevant logs and a clear pass/fail result. Any manual step must be documented and repeatable.

## Gate

The MVP is not ready for external users until the sequence passes on a clean node without hidden third-party control-plane dependencies.
