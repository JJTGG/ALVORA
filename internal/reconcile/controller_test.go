package reconcile

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/JJTGG/ALVORA/internal/control"
	"github.com/JJTGG/ALVORA/internal/resource"
)

func TestReconcilerRequiresDependencies(t *testing.T) {
	if err := (&Reconciler{}).Reconcile(context.Background(), "resource-1"); !errors.Is(err, ErrNilReconciler) {
		t.Fatalf("Reconcile() error = %v, want %v", err, ErrNilReconciler)
	}
}

func TestReconcilerTransitionsResourceToReady(t *testing.T) {
	value := testResource()
	repository := &memoryRepository{value: value}
	action := &fakeAction{observed: []byte(`{"healthy":true}`)}
	reconciler := &Reconciler{
		Repository: repository,
		Action:     action,
		Now:        func() time.Time { return time.Unix(100, 0).UTC() },
	}

	if err := reconciler.Reconcile(context.Background(), value.ID); err != nil {
		t.Fatalf("Reconcile() error = %v", err)
	}
	if repository.value.State != control.StateReady {
		t.Fatalf("state = %q, want %q", repository.value.State, control.StateReady)
	}
	if string(repository.value.ObservedStatus) != string(action.observed) {
		t.Fatalf("observed status = %s, want %s", repository.value.ObservedStatus, action.observed)
	}
	if repository.value.Generation != 3 {
		t.Fatalf("generation = %d, want 3", repository.value.Generation)
	}
}

func TestReconcilerRecordsFailure(t *testing.T) {
	value := testResource()
	repository := &memoryRepository{value: value}
	reconciler := &Reconciler{
		Repository: repository,
		Action:     &fakeAction{err: errors.New("build failed")},
		Now:        time.Now,
	}

	if err := reconciler.Reconcile(context.Background(), value.ID); err == nil {
		t.Fatal("Reconcile() error = nil, want action error")
	}
	if repository.value.State != control.StateFailed {
		t.Fatalf("state = %q, want %q", repository.value.State, control.StateFailed)
	}
	if repository.value.ObservedError != "build failed" {
		t.Fatalf("observed error = %q, want build failed", repository.value.ObservedError)
	}
}

func testResource() resource.Resource {
	return resource.Resource{ID: "resource-1", Type: "application", OrganizationID: "org-1", ProjectID: "project-1", Generation: 1, State: control.StateRequested}
}

type fakeAction struct {
	observed []byte
	err      error
}

func (a *fakeAction) Apply(context.Context, resource.Resource) ([]byte, error) {
	return a.observed, a.err
}

type memoryRepository struct{ value resource.Resource }

func (m *memoryRepository) Create(context.Context, resource.Resource) error { return nil }
func (m *memoryRepository) Get(context.Context, string) (resource.Resource, error) { return m.value, nil }
func (m *memoryRepository) Update(_ context.Context, value resource.Resource, _ int64) error { m.value = value; return nil }
func (m *memoryRepository) ListByScope(context.Context, string, string, string) ([]resource.Resource, error) { return []resource.Resource{m.value}, nil }
