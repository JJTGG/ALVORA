package api

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/JJTGG/ALVORA/internal/control"
	"github.com/JJTGG/ALVORA/internal/resource"
	"github.com/JJTGG/ALVORA/internal/store"
)

var ErrNilResourceAPI = errors.New("resource API dependencies are required")

type ResourceAPI struct {
	Repository store.ResourceRepository
	Now        func() time.Time
}

type resourceRequest struct {
	ID             string          `json:"id"`
	Type           string          `json:"type"`
	OrganizationID string          `json:"organization_id"`
	ProjectID      string          `json:"project_id"`
	EnvironmentID  string          `json:"environment_id"`
	DesiredSpec    json.RawMessage `json:"desired_spec"`
}

type resourceResponse struct {
	ID             string          `json:"id"`
	Type           string          `json:"type"`
	OrganizationID string          `json:"organization_id"`
	ProjectID      string          `json:"project_id"`
	EnvironmentID  string          `json:"environment_id,omitempty"`
	Generation     int64           `json:"generation"`
	State          control.State   `json:"state"`
	DesiredSpec    json.RawMessage `json:"desired_spec,omitempty"`
	ObservedStatus json.RawMessage `json:"observed_status,omitempty"`
	ObservedError  string          `json:"observed_error,omitempty"`
}

func (a ResourceAPI) Handler() (http.Handler, error) {
	if a.Repository == nil {
		return nil, ErrNilResourceAPI
	}
	if a.Now == nil {
		a.Now = time.Now
	}
	mux := http.NewServeMux()
	mux.HandleFunc("POST /v1/resources", a.create)
	mux.HandleFunc("GET /v1/resources/{id}", a.get)
	return mux, nil
}

func (a ResourceAPI) create(w http.ResponseWriter, r *http.Request) {
	var input resourceRequest
	if err := decodeJSON(r, &input); err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}
	value := resource.Resource{ID: input.ID, Type: input.Type, OrganizationID: input.OrganizationID, ProjectID: input.ProjectID, EnvironmentID: input.EnvironmentID, Generation: 1, State: control.StateRequested, DesiredSpec: append([]byte(nil), input.DesiredSpec...), CreatedAt: a.Now(), UpdatedAt: a.Now()}
	if err := a.Repository.Create(r.Context(), value); err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}
	writeJSON(w, http.StatusCreated, toResponse(value))
}

func (a ResourceAPI) get(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimSpace(r.PathValue("id"))
	value, err := a.Repository.Get(r.Context(), id)
	if errors.Is(err, store.ErrResourceNotFound) {
		writeError(w, http.StatusNotFound, err.Error())
		return
	}
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, toResponse(value))
}

func decodeJSON(r *http.Request, target any) error {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	return decoder.Decode(target)
}

func toResponse(value resource.Resource) resourceResponse {
	return resourceResponse{ID: value.ID, Type: value.Type, OrganizationID: value.OrganizationID, ProjectID: value.ProjectID, EnvironmentID: value.EnvironmentID, Generation: value.Generation, State: value.State, DesiredSpec: value.DesiredSpec, ObservedStatus: value.ObservedStatus, ObservedError: value.ObservedError}
}

func writeJSON(w http.ResponseWriter, status int, value any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(value)
}

func writeError(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, map[string]string{"error": message})
}

var _ = context.Background
