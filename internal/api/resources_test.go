package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/JJTGG/ALVORA/internal/store"
)

func TestResourceAPICreatesAndGetsResource(t *testing.T) {
	repository := store.NewMemoryResourceRepository()
	api, err := (ResourceAPI{Repository: repository}).Handler()
	if err != nil {
		t.Fatalf("Handler() error = %v", err)
	}

	body := `{"id":"resource-1","type":"application","organization_id":"org-1","project_id":"project-1","desired_spec":{"image":"demo"}}`
	createRequest := httptest.NewRequest(http.MethodPost, "/v1/resources", strings.NewReader(body))
	createResponse := httptest.NewRecorder()
	api.ServeHTTP(createResponse, createRequest)
	if createResponse.Code != http.StatusCreated {
		t.Fatalf("create status = %d, want %d: %s", createResponse.Code, http.StatusCreated, createResponse.Body.String())
	}

	getRequest := httptest.NewRequest(http.MethodGet, "/v1/resources/resource-1", nil)
	getResponse := httptest.NewRecorder()
	api.ServeHTTP(getResponse, getRequest)
	if getResponse.Code != http.StatusOK {
		t.Fatalf("get status = %d, want %d", getResponse.Code, http.StatusOK)
	}
	var result map[string]any
	if err := json.Unmarshal(getResponse.Body.Bytes(), &result); err != nil {
		t.Fatalf("decode response: %v", err)
	}
	if result["state"] != "requested" {
		t.Fatalf("state = %v, want requested", result["state"])
	}
}

func TestResourceAPIReportsNotFound(t *testing.T) {
	api, err := (ResourceAPI{Repository: store.NewMemoryResourceRepository()}).Handler()
	if err != nil {
		t.Fatalf("Handler() error = %v", err)
	}
	request := httptest.NewRequest(http.MethodGet, "/v1/resources/missing", nil)
	response := httptest.NewRecorder()
	api.ServeHTTP(response, request)
	if response.Code != http.StatusNotFound {
		t.Fatalf("status = %d, want %d", response.Code, http.StatusNotFound)
	}
}
