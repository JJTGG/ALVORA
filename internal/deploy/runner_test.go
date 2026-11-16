package deploy

import (
	"context"
	"errors"
	"testing"
)

type fakeProcess struct{}

func (fakeProcess) Wait() error { return nil }
func (fakeProcess) Kill() error { return nil }

func TestRunnerDelegatesToFactory(t *testing.T) {
	called := false
	runner := Runner{Start: func(context.Context, Spec) (Process, error) {
		called = true
		return fakeProcess{}, nil
	}}
	spec := Spec{Application: "demo", SourceDir: "/tmp/demo", StartCommand: "./demo", Port: 8080}
	if _, err := runner.StartProcess(context.Background(), spec); err != nil {
		t.Fatalf("StartProcess() error = %v", err)
	}
	if !called {
		t.Fatal("runner factory was not called")
	}
}

func TestRunnerRejectsInvalidSpecBeforeFactory(t *testing.T) {
	runner := Runner{Start: func(context.Context, Spec) (Process, error) {
		t.Fatal("factory called for invalid specification")
		return nil, errors.New("unreachable")
	}}
	if _, err := runner.StartProcess(context.Background(), Spec{}); err == nil {
		t.Fatal("StartProcess() error = nil, want validation error")
	}
}
