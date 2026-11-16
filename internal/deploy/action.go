package deploy

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"sync"

	"github.com/JJTGG/ALVORA/internal/resource"
)

var ErrNilActionRunner = errors.New("deployment action runner is required")

// ProcessAction adapts the local process runner to the reconciliation action boundary.
type ProcessAction struct {
	Runner Runner
	mu     sync.Mutex
	live   map[string]Process
}

// Apply starts or reuses the process described by the resource desired spec.
func (a *ProcessAction) Apply(ctx context.Context, value resource.Resource) ([]byte, error) {
	if a == nil {
		return nil, ErrNilActionRunner
	}
	if err := value.Validate(); err != nil {
		return nil, fmt.Errorf("validate resource: %w", err)
	}

	var spec Spec
	if err := json.Unmarshal(value.DesiredSpec, &spec); err != nil {
		return nil, fmt.Errorf("decode desired deployment: %w", err)
	}
	if err := spec.Validate(); err != nil {
		return nil, fmt.Errorf("validate desired deployment: %w", err)
	}

	a.mu.Lock()
	defer a.mu.Unlock()
	if a.live == nil {
		a.live = make(map[string]Process)
	}
	if _, ok := a.live[value.ID]; ok {
		return statusJSON(value.ID, "running")
	}

	process, err := a.Runner.StartProcess(ctx, spec)
	if err != nil {
		return nil, fmt.Errorf("start resource process: %w", err)
	}
	a.live[value.ID] = process
	return statusJSON(value.ID, "running")
}

func statusJSON(id, state string) ([]byte, error) {
	return json.Marshal(map[string]string{"resource_id": id, "state": state})
}
