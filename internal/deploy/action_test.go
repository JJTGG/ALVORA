package deploy

import (
	"context"
	"encoding/json"
	"errors"
	"testing"

	"github.com/JJTGG/ALVORA/internal/control"
	"github.com/JJTGG/ALVORA/internal/resource"
)

func TestProcessActionRequiresRunner(t *testing.T) {
	action := &ProcessAction{}
	value := actionResource(t)
	if _, err := action.Apply(context.Background(), value); err == nil {
		t.Fatal("Apply() error = nil, want runner error")
	}
}

func TestProcessActionStartsProcessAndReusesIt(t *testing.T) {
	starts := 0
	action := &ProcessAction{Runner: Runner{Start: func(context.Context, Spec) (Process, error) {
		starts++
		return fakeProcess{}, nil
	}}}
	value := actionResource(t)

	first, err := action.Apply(context.Background(), value)
	if err != nil {
		t.Fatalf("first Apply() error = %v", err)
	}
	second, err := action.Apply(context.Background(), value)
	if err != nil {
		t.Fatalf("second Apply() error = %v", err)
	}
	if starts != 1 {
		t.Fatalf("process starts = %d, want 1", starts)
	}
	for _, observed := range [][]byte{first, second} {
		var status map[string]string
		if err := json.Unmarshal(observed, &status); err != nil {
			t.Fatalf("decode status: %v", err)
		}
		if status["state"] != "running" {
			t.Fatalf("status = %#v, want running", status)
		}
	}
}

func actionResource(t *testing.T) resource.Resource {
	t.Helper()
	desired, err := json.Marshal(Spec{Application: "demo", SourceDir: "/tmp/demo", StartCommand: "./demo", Port: 8080})
	if err != nil {
		t.Fatalf("marshal spec: %v", err)
	}
	return resource.Resource{ID: "resource-1", Type: "application", OrganizationID: "org-1", ProjectID: "project-1", Generation: 1, State: control.StateProvisioning, DesiredSpec: desired}
}

var _ = errors.New
