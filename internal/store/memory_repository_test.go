package store

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/JJTGG/ALVORA/internal/control"
	"github.com/JJTGG/ALVORA/internal/resource"
)

func TestMemoryResourceRepositoryCRUDAndGeneration(t *testing.T) {
	repository := NewMemoryResourceRepository()
	value := memoryResource()
	ctx := context.Background()

	if err := repository.Create(ctx, value); err != nil {
		t.Fatalf("Create() error = %v", err)
	}
	if _, err := repository.Get(ctx, value.ID); err != nil {
		t.Fatalf("Get() error = %v", err)
	}

	value.Generation = 2
	if err := repository.Update(ctx, value, 1); err != nil {
		t.Fatalf("Update() error = %v", err)
	}
	if err := repository.Update(ctx, value, 1); !errors.Is(err, ErrGenerationConflict) {
		t.Fatalf("stale Update() error = %v, want %v", err, ErrGenerationConflict)
	}
}

func TestMemoryResourceRepositoryClonesBytes(t *testing.T) {
	repository := NewMemoryResourceRepository()
	value := memoryResource()
	if err := repository.Create(context.Background(), value); err != nil {
		t.Fatalf("Create() error = %v", err)
	}
	value.DesiredSpec[0] = 'x'
	stored, err := repository.Get(context.Background(), value.ID)
	if err != nil {
		t.Fatalf("Get() error = %v", err)
	}
	if stored.DesiredSpec[0] == 'x' {
		t.Fatal("repository did not clone desired specification")
	}
}

func memoryResource() resource.Resource {
	return resource.Resource{
		ID: "memory-1", Type: "application", OrganizationID: "org-1", ProjectID: "project-1",
		Generation: 1, State: control.StateRequested, DesiredSpec: []byte("{}"), CreatedAt: time.Now(), UpdatedAt: time.Now(),
	}
}
