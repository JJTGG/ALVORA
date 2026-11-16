package deploy

import (
	"context"
	"fmt"
	"os/exec"
)

// Process is the small process-control seam used by Runner.
type Process interface {
	Wait() error
	Kill() error
}

type process struct {
	cmd *exec.Cmd
}

func (p *process) Wait() error { return p.cmd.Wait() }
func (p *process) Kill() error { return p.cmd.Process.Kill() }

// Runner starts a validated local application process in its source directory.
type Runner struct {
	Start func(context.Context, Spec) (Process, error)
}

// StartProcess is the production process factory.
func StartProcess(ctx context.Context, spec Spec) (Process, error) {
	if err := spec.Validate(); err != nil {
		return nil, fmt.Errorf("validate deployment: %w", err)
	}
	command, err := spec.Command()
	if err != nil {
		return nil, fmt.Errorf("build deployment command: %w", err)
	}
	cmd := exec.CommandContext(ctx, command[0], command[1:]...)
	cmd.Dir = spec.SourceDir
	if err := cmd.Start(); err != nil {
		return nil, fmt.Errorf("start application: %w", err)
	}
	return &process{cmd: cmd}, nil
}

// StartProcess validates the specification and starts the configured local process.
func (r Runner) StartProcess(ctx context.Context, spec Spec) (Process, error) {
	if err := spec.Validate(); err != nil {
		return nil, fmt.Errorf("validate deployment: %w", err)
	}
	factory := r.Start
	if factory == nil {
		factory = StartProcess
	}
	return factory(ctx, spec)
}

var _ = context.Background
