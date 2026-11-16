package store

import (
	"context"
	"testing"

	"github.com/JJTGG/ALVORA/internal/resource"
)

func TestNewSQLResourceRepositoryRejectsNilDB(t *testing.T) {
	if _, err := NewSQLResourceRepository(nil); err != ErrNilDB {
		t.Fatalf("NewSQLResourceRepository(nil) error = %v, want %v", err, ErrNilDB)
	}
}

func TestRepositoryInterfaceIsUsable(t *testing.T) {
	var repository ResourceRepository = &fakeResourceRepository{}
	value := resource.Resource{ID: "resource-1"}
	if err := repository.Create(context.Background(), value); err == nil {
		t.Fatal("Create() error = nil, want validation error")
	}
}

type fakeResourceRepository struct{}

func (*fakeResourceRepository) Create(_ context.Context, value resource.Resource) error {
	return value.Validate()
}

func (*fakeResourceRepository) Get(context.Context, string) (resource.Resource, error) {
	return resource.Resource{}, ErrResourceNotFound
}

func (*fakeResourceRepository) Update(context.Context, resource.Resource, int64) error {
	return ErrGenerationConflict
}

func (*fakeResourceRepository) ListByScope(context.Context, string, string, string) ([]resource.Resource, error) {
	return nil, nil
}
