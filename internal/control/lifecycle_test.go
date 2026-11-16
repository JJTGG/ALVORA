package control

import "testing"

func TestCanTransition(t *testing.T) {
	tests := []struct {
		name string
		from State
		to   State
		want bool
	}{
		{"request to provision", StateRequested, StateProvisioning, true},
		{"provision to ready", StateProvisioning, StateReady, true},
		{"ready to update", StateReady, StateUpdating, true},
		{"ready to delete", StateReady, StateDeleting, true},
		{"failed to ready", StateFailed, StateReady, false},
		{"deleted to ready", StateDeleted, StateReady, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CanTransition(tt.from, tt.to); got != tt.want {
				t.Fatalf("CanTransition(%q, %q) = %v, want %v", tt.from, tt.to, got, tt.want)
			}
		})
	}
}
