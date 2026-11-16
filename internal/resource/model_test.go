package resource

import (
	"testing"
	"time"

	"github.com/JJTGG/ALVORA/internal/control"
)

func validResource() Resource {
	return Resource{
		ID: "resource-1", Type: "application", OrganizationID: "org-1", ProjectID: "project-1",
		Generation: 1, State: control.StateRequested,
	}
}

func TestResourceValidate(t *testing.T) {
	if err := validResource().Validate(); err != nil {
		t.Fatalf("Validate() error = %v", err)
	}

	tests := []struct {
		name string
		edit func(*Resource)
	}{
		{"missing id", func(r *Resource) { r.ID = "" }},
		{"missing type", func(r *Resource) { r.Type = "" }},
		{"missing scope", func(r *Resource) { r.ProjectID = "" }},
		{"invalid generation", func(r *Resource) { r.Generation = 0 }},
		{"missing state", func(r *Resource) { r.State = "" }},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := validResource()
			tt.edit(&r)
			if err := r.Validate(); err == nil {
				t.Fatal("Validate() error = nil, want validation error")
			}
		})
	}
}

func TestResourceTransition(t *testing.T) {
	now := time.Date(2026, time.November, 16, 0, 0, 0, 0, time.UTC)
	r := validResource()
	if err := r.Transition(control.StateProvisioning, now); err != nil {
		t.Fatalf("Transition() error = %v", err)
	}
	if r.State != control.StateProvisioning || !r.UpdatedAt.Equal(now) {
		t.Fatalf("resource after transition = %+v", r)
	}
	if err := r.Transition(control.StateDeleted, now); err == nil {
		t.Fatal("Transition() error = nil, want invalid transition")
	}
}
