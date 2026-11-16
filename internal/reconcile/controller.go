package reconcile

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/JJTGG/ALVORA/internal/control"
	"github.com/JJTGG/ALVORA/internal/resource"
	"github.com/JJTGG/ALVORA/internal/store"
)

var ErrNilReconciler = errors.New("reconciler dependencies are required")

// Action applies the external side effects required to make a resource ready.
type Action interface {
	Apply(context.Context, resource.Resource) ([]byte, error)
}

// Reconciler drives one resource toward its desired state.
type Reconciler struct {
	Repository store.ResourceRepository
	Action     Action
	Now        func() time.Time
}

// Reconcile performs one idempotent reconciliation attempt for a resource ID.
func (r *Reconciler) Reconcile(ctx context.Context, id string) error {
	if r == nil || r.Repository == nil || r.Action == nil {
		return ErrNilReconciler
	}
	now := time.Now
	if r.Now != nil {
		now = r.Now
	}

	current, err := r.Repository.Get(ctx, id)
	if err != nil {
		return fmt.Errorf("load resource: %w", err)
	}
	if current.State == control.StateDeleted {
		return nil
	}
	if current.State == control.StateReady {
		return nil
	}

	originalGeneration := current.Generation
	current.Generation++
	if err := current.Transition(control.StateProvisioning, now()); err != nil {
		return fmt.Errorf("start reconciliation: %w", err)
	}
	if err := r.Repository.Update(ctx, current, originalGeneration); err != nil {
		return fmt.Errorf("persist provisioning state: %w", err)
	}

	observed, err := r.Action.Apply(ctx, current)
	if err != nil {
		failed := current
		failed.Generation++
		failed.ObservedError = err.Error()
		if transitionErr := failed.Transition(control.StateFailed, now()); transitionErr == nil {
			_ = r.Repository.Update(ctx, failed, current.Generation)
		}
		return fmt.Errorf("apply resource: %w", err)
	}

	ready := current
	ready.Generation++
	ready.ObservedStatus = observed
	ready.ObservedError = ""
	if err := ready.Transition(control.StateReady, now()); err != nil {
		return fmt.Errorf("complete reconciliation: %w", err)
	}
	if err := r.Repository.Update(ctx, ready, current.Generation); err != nil {
		return fmt.Errorf("persist ready state: %w", err)
	}
	return nil
}
