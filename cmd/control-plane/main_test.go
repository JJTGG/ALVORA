package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthz(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"status":"ok"}`))
	})

	req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
	res := httptest.NewRecorder()
	mux.ServeHTTP(res, req)

	if res.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", res.Code, http.StatusOK)
	}
	if got := res.Body.String(); got != `{"status":"ok"}` {
		t.Fatalf("body = %q, want %q", got, `{"status":"ok"}`)
	}
}

func TestEnvInt(t *testing.T) {
	t.Setenv("ALVORA_TEST_PORT", "9090")
	if got := envInt("ALVORA_TEST_PORT", 8080); got != 9090 {
		t.Fatalf("envInt() = %d, want 9090", got)
	}

	t.Setenv("ALVORA_TEST_PORT", "invalid")
	if got := envInt("ALVORA_TEST_PORT", 8080); got != 8080 {
		t.Fatalf("invalid envInt() = %d, want 8080", got)
	}
}
