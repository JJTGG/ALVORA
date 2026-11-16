package store

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/JJTGG/ALVORA/internal/control"
	"github.com/JJTGG/ALVORA/internal/resource"
)

var (
	ErrResourceNotFound = errors.New("resource not found")
	ErrGenerationConflict = errors.New("resource generation conflict")
)

// ResourceRepository persists desired and observed resource state.
type ResourceRepository interface {
	Create(ctx context.Context, value resource.Resource) error
	Get(ctx context.Context, id string) (resource.Resource, error)
	Update(ctx context.Context, value resource.Resource, expectedGeneration int64) error
	ListByScope(ctx context.Context, organizationID, projectID, environmentID string) ([]resource.Resource, error)
}

// SQLResourceRepository implements ResourceRepository using PostgreSQL-compatible SQL.
type SQLResourceRepository struct {
	db *sql.DB
}

func NewSQLResourceRepository(db *sql.DB) (*SQLResourceRepository, error) {
	if db == nil {
		return nil, ErrNilDB
	}
	return &SQLResourceRepository{db: db}, nil
}

func (r *SQLResourceRepository) Create(ctx context.Context, value resource.Resource) error {
	if err := value.Validate(); err != nil {
		return fmt.Errorf("validate resource: %w", err)
	}
	_, err := r.db.ExecContext(ctx, `
		INSERT INTO resources (
			id, resource_type, organization_id, project_id, environment_id,
			generation, state, desired_spec, observed_status, observed_error,
			created_at, updated_at
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
	`, value.ID, value.Type, value.OrganizationID, value.ProjectID, value.EnvironmentID,
		value.Generation, value.State, value.DesiredSpec, value.ObservedStatus,
		value.ObservedError, value.CreatedAt, value.UpdatedAt)
	if err != nil {
		return fmt.Errorf("create resource: %w", err)
	}
	return nil
}

func (r *SQLResourceRepository) Get(ctx context.Context, id string) (resource.Resource, error) {
	var value resource.Resource
	err := r.db.QueryRowContext(ctx, `
		SELECT id, resource_type, organization_id, project_id, environment_id,
			generation, state, desired_spec, observed_status, observed_error,
			created_at, updated_at
		FROM resources WHERE id = $1
	`, id).Scan(
		&value.ID, &value.Type, &value.OrganizationID, &value.ProjectID, &value.EnvironmentID,
		&value.Generation, &value.State, &value.DesiredSpec, &value.ObservedStatus,
		&value.ObservedError, &value.CreatedAt, &value.UpdatedAt)
	if errors.Is(err, sql.ErrNoRows) {
		return resource.Resource{}, ErrResourceNotFound
	}
	if err != nil {
		return resource.Resource{}, fmt.Errorf("get resource: %w", err)
	}
	return value, nil
}

func (r *SQLResourceRepository) Update(ctx context.Context, value resource.Resource, expectedGeneration int64) error {
	if err := value.Validate(); err != nil {
		return fmt.Errorf("validate resource: %w", err)
	}
	if expectedGeneration < 1 {
		return errors.New("expected generation must be positive")
	}
	result, err := r.db.ExecContext(ctx, `
		UPDATE resources SET
			resource_type = $2, organization_id = $3, project_id = $4,
			environment_id = $5, generation = $6, state = $7, desired_spec = $8,
			observed_status = $9, observed_error = $10, updated_at = $11
		WHERE id = $1 AND generation = $12
	`, value.ID, value.Type, value.OrganizationID, value.ProjectID, value.EnvironmentID,
		value.Generation, value.State, value.DesiredSpec, value.ObservedStatus,
		value.ObservedError, value.UpdatedAt, expectedGeneration)
	if err != nil {
		return fmt.Errorf("update resource: %w", err)
	}
	count, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("inspect resource update: %w", err)
	}
	if count == 0 {
		return ErrGenerationConflict
	}
	return nil
}

func (r *SQLResourceRepository) ListByScope(ctx context.Context, organizationID, projectID, environmentID string) ([]resource.Resource, error) {
	rows, err := r.db.QueryContext(ctx, `
		SELECT id, resource_type, organization_id, project_id, environment_id,
			generation, state, desired_spec, observed_status, observed_error,
			created_at, updated_at
		FROM resources
		WHERE organization_id = $1 AND project_id = $2
			AND environment_id IS NOT DISTINCT FROM $3
		ORDER BY id
	`, organizationID, projectID, environmentID)
	if err != nil {
		return nil, fmt.Errorf("list resources: %w", err)
	}
	defer rows.Close()

	var values []resource.Resource
	for rows.Next() {
		var value resource.Resource
		if err := rows.Scan(
			&value.ID, &value.Type, &value.OrganizationID, &value.ProjectID, &value.EnvironmentID,
			&value.Generation, &value.State, &value.DesiredSpec, &value.ObservedStatus,
			&value.ObservedError, &value.CreatedAt, &value.UpdatedAt); err != nil {
			return nil, fmt.Errorf("scan resource: %w", err)
		}
		values = append(values, value)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate resources: %w", err)
	}
	return values, nil
}

var _ ResourceRepository = (*SQLResourceRepository)(nil)
var _ = control.State("")
var _ = time.Time{}
