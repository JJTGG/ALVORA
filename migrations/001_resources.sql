-- ALVORA migration 001: control-plane resource state
-- This migration mirrors internal/store.Schema and is suitable for
-- operators who run migrations outside the application process.

CREATE TABLE IF NOT EXISTS resources (
    id TEXT PRIMARY KEY,
    resource_type TEXT NOT NULL,
    organization_id TEXT NOT NULL,
    project_id TEXT NOT NULL,
    environment_id TEXT,
    generation BIGINT NOT NULL CHECK (generation > 0),
    state TEXT NOT NULL,
    desired_spec BYTEA NOT NULL DEFAULT '{}',
    observed_status BYTEA NOT NULL DEFAULT '{}',
    observed_error TEXT NOT NULL DEFAULT '',
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL
);

CREATE INDEX IF NOT EXISTS resources_project_idx
    ON resources (organization_id, project_id, environment_id);

CREATE INDEX IF NOT EXISTS resources_state_idx
    ON resources (state);
