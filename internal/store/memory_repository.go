package store

import (
	"context"
	"errors"
	"sync"

	"github.com/JJTGG/ALVORA/internal/resource"
)

// MemoryResourceRepository provides a durable-boundary-compatible repository for local development.
type MemoryResourceRepository struct {
	mu     sync.RWMutex
	values map[string]resource.Resource
}

func NewMemoryResourceRepository() *MemoryResourceRepository {
	return &MemoryResourceRepository{values: make(map[string]resource.Resource)}
}

func (r *MemoryResourceRepository) Create(ctx context.Context, value resource.Resource) error {
	if err := contextErr(ctx); err != nil {
		return err
	}
	if err := value.Validate(); err != nil {
		return err
	}
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, exists := r.values[value.ID]; exists {
		return errors.New("resource already exists")
	}
	r.values[value.ID] = cloneResource(value)
	return nil
}

func (r *MemoryResourceRepository) Get(ctx context.Context, id string) (resource.Resource, error) {
	if err := contextErr(ctx); err != nil {
		return resource.Resource{}, err
	}
	r.mu.RLock()
	defer r.mu.RUnlock()
	value, ok := r.values[id]
	if !ok {
		return resource.Resource{}, ErrResourceNotFound
	}
	return cloneResource(value), nil
}

func (r *MemoryResourceRepository) Update(ctx context.Context, value resource.Resource, expectedGeneration int64) error {
	if err := contextErr(ctx); err != nil {
		return err
	}
	if err := value.Validate(); err != nil {
		return err
	}
	r.mu.Lock()
	defer r.mu.Unlock()
	current, ok := r.values[value.ID]
	if !ok || current.Generation != expectedGeneration {
		return ErrGenerationConflict
	}
	r.values[value.ID] = cloneResource(value)
	return nil
}

func (r *MemoryResourceRepository) ListByScope(ctx context.Context, organizationID, projectID, environmentID string) ([]resource.Resource, error) {
	if err := contextErr(ctx); err != nil {
		return nil, err
	}
	r.mu.RLock()
	defer r.mu.RUnlock()
	values := make([]resource.Resource, 0)
	for _, value := range r.values {
		if value.OrganizationID == organizationID && value.ProjectID == projectID && value.EnvironmentID == environmentID {
			values = append(values, cloneResource(value))
		}
	}
	return values, nil
}

func contextErr(ctx context.Context) error {
	if ctx == nil {
		return nil
	}
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		return nil
	}
}

func cloneResource(value resource.Resource) resource.Resource {
	value.DesiredSpec = append([]byte(nil), value.DesiredSpec...)
	value.ObservedStatus = append([]byte(nil), value.ObservedStatus...)
	return value
}

var _ ResourceRepository = (*MemoryResourceRepository)(nil)
