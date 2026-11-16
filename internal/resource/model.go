package resource

import (
	"errors"
	"strings"
	"time"

	"github.com/JJTGG/ALVORA/internal/control"
)

var (
	ErrInvalidResourceID       = errors.New("resource id is required")
	ErrInvalidResourceType     = errors.New("resource type is required")
	ErrInvalidScope            = errors.New("resource scope is required")
	ErrInvalidStateTransition  = errors.New("invalid resource state transition")
)

// Resource is the durable desired/observed state boundary for a managed object.
type Resource struct {
	ID             string
	Type           string
	OrganizationID string
	ProjectID      string
	EnvironmentID  string
	Generation     int64
	State          control.State
	DesiredSpec    []byte
	ObservedStatus []byte
	ObservedError  string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

// Validate checks the identity and lifecycle invariants required by storage.
func (r Resource) Validate() error {
	if strings.TrimSpace(r.ID) == "" {
		return ErrInvalidResourceID
	}
	if strings.TrimSpace(r.Type) == "" {
		return ErrInvalidResourceType
	}
	if strings.TrimSpace(r.OrganizationID) == "" || strings.TrimSpace(r.ProjectID) == "" {
		return ErrInvalidScope
	}
	if r.Generation < 1 {
		return errors.New("resource generation must be positive")
	}
	if r.State == "" {
		return errors.New("resource state is required")
	}
	return nil
}

// Transition applies a validated lifecycle transition and records the update time.
func (r *Resource) Transition(next control.State, now time.Time) error {
	if r == nil {
		return errors.New("resource is nil")
	}
	if !control.CanTransition(r.State, next) {
		return ErrInvalidStateTransition
	}
	r.State = next
	r.UpdatedAt = now
	return nil
}
