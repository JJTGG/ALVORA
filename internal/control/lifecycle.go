package control

// State describes the lifecycle of a managed ALVORA resource.
type State string

const (
	StateRequested    State = "requested"
	StateProvisioning State = "provisioning"
	StateReady        State = "ready"
	StateUpdating     State = "updating"
	StateFailed       State = "failed"
	StateDeleting     State = "deleting"
	StateDeleted      State = "deleted"
)

// CanTransition reports whether a resource may move from one state to another.
func CanTransition(from, to State) bool {
	if from == to {
		return true
	}
	switch from {
	case StateRequested:
		return to == StateProvisioning || to == StateFailed
	case StateProvisioning:
		return to == StateReady || to == StateFailed
	case StateReady:
		return to == StateUpdating || to == StateDeleting || to == StateFailed
	case StateUpdating:
		return to == StateReady || to == StateFailed
	case StateFailed:
		return to == StateProvisioning || to == StateDeleting
	case StateDeleting:
		return to == StateDeleted || to == StateFailed
	default:
		return false
	}
}
