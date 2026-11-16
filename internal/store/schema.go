package store

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"
)

var ErrNilDB = errors.New("database handle is required")

// Schema contains the first durable control-plane schema.
const Schema = `
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
`

// ApplySchema executes the initial schema as one transaction.
func ApplySchema(db *sql.DB) error {
	if db == nil {
		return ErrNilDB
	}
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("begin schema transaction: %w", err)
	}
	defer func() { _ = tx.Rollback() }()

	for _, statement := range strings.Split(Schema, ";") {
		statement = strings.TrimSpace(statement)
		if statement == "" {
			continue
		}
		if _, err := tx.Exec(statement); err != nil {
			return fmt.Errorf("apply schema statement: %w", err)
		}
	}
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("commit schema transaction: %w", err)
	}
	return nil
}
