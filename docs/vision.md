# ALVORA Vision

## Purpose

ALVORA is a self-hosted full-stack developer cloud. It should let developers deploy applications and provision the backend capabilities those applications need from one coherent platform.

The platform is intended to operate on infrastructure controlled by its operator. External providers may be supported as optional adapters, but they must not be required by the core architecture.

## Desired experience

A developer should be able to bootstrap ALVORA, create a project and environment, provision PostgreSQL and object storage, deploy an application, inject secrets, expose a domain, inspect logs, and roll back a release.

## Scope

The initial scope includes application runtimes, builds, PostgreSQL, object storage, authentication foundations, background jobs, domains, secrets, logs, recovery and self-deployment.

## Non-goals for the initial phase

ALVORA will not initially attempt to provide a global CDN, a worldwide edge network, arbitrary untrusted code execution, multi-region database consensus, or every feature of established cloud providers.

## Success condition

The first meaningful release is a self-hosted single-node installation that can deploy a complete application and can build, upgrade and recover its own platform services.

All capabilities in this document are goals or proposals unless implementation and tests prove otherwise.
