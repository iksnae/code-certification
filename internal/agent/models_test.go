package agent

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestListModels_OpenAIFormat(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/models" && r.URL.Path != "/models" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		if r.Method != "GET" {
			t.Errorf("unexpected method: %s", r.Method)
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{
			"object": "list",
			"data": [
				{"id": "gpt-4o", "object": "model", "owned_by": "openai", "created": 1700000000},
				{"id": "llama-3.3-70b", "object": "model", "owned_by": "meta", "created": 1700000001}
			]
		}`)
	}))
	defer srv.Close()

	models, err := ListModels(context.Background(), srv.URL+"/v1", "test-key")
	if err != nil {
		t.Fatalf("ListModels() error: %v", err)
	}
	if len(models) != 2 {
		t.Fatalf("got %d models, want 2", len(models))
	}
	if models[0].ID != "gpt-4o" {
		t.Errorf("models[0].ID = %q, want gpt-4o", models[0].ID)
	}
	if models[0].OwnedBy != "openai" {
		t.Errorf("models[0].OwnedBy = %q, want openai", models[0].OwnedBy)
	}
	if models[1].ID != "llama-3.3-70b" {
		t.Errorf("models[1].ID = %q, want llama-3.3-70b", models[1].ID)
	}
}

func TestListModels_OllamaFormat(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Ollama returns 404 on /v1/models but has /api/tags
		if r.URL.Path == "/v1/models" || r.URL.Path == "/models" {
			w.WriteHeader(404)
			return
		}
		if r.URL.Path == "/api/tags" {
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, `{
				"models": [
					{"name": "qwen2.5-coder:7b", "size": 4700000000, "modified_at": "2025-01-01T00:00:00Z"},
					{"name": "llama3.2:3b", "size": 2000000000, "modified_at": "2025-01-02T00:00:00Z"}
				]
			}`)
			return
		}
		w.WriteHeader(404)
	}))
	defer srv.Close()

	models, err := ListModels(context.Background(), srv.URL+"/v1", "")
	if err != nil {
		t.Fatalf("ListModels() error: %v", err)
	}
	if len(models) != 2 {
		t.Fatalf("got %d models, want 2", len(models))
	}
	if models[0].ID != "qwen2.5-coder:7b" {
		t.Errorf("models[0].ID = %q, want qwen2.5-coder:7b", models[0].ID)
	}
	if models[0].OwnedBy != "ollama" {
		t.Errorf("models[0].OwnedBy = %q, want ollama", models[0].OwnedBy)
	}
}

func TestListModels_Error(t *testing.T) {
	// Connect to a port that's not listening
	_, err := ListModels(context.Background(), "http://127.0.0.1:1", "")
	if err == nil {
		t.Fatal("ListModels() should return error on connection failure")
	}
}

func TestListModels_EmptyResponse(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"object": "list", "data": []}`)
	}))
	defer srv.Close()

	models, err := ListModels(context.Background(), srv.URL+"/v1", "test-key")
	if err != nil {
		t.Fatalf("ListModels() error: %v", err)
	}
	if len(models) != 0 {
		t.Errorf("got %d models, want 0", len(models))
	}
}

func TestListModels_NoAuth(t *testing.T) {
	var gotAuth string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotAuth = r.Header.Get("Authorization")
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"object": "list", "data": [{"id": "local-model", "object": "model"}]}`)
	}))
	defer srv.Close()

	models, err := ListModels(context.Background(), srv.URL+"/v1", "")
	if err != nil {
		t.Fatalf("ListModels() error: %v", err)
	}
	if gotAuth != "" {
		t.Errorf("Authorization header sent for local provider: %q", gotAuth)
	}
	if len(models) != 1 {
		t.Fatalf("got %d models, want 1", len(models))
	}
	if models[0].ID != "local-model" {
		t.Errorf("models[0].ID = %q", models[0].ID)
	}
}

func TestListModels_WithAuth(t *testing.T) {
	var gotAuth string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotAuth = r.Header.Get("Authorization")
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"object": "list", "data": [{"id": "cloud-model", "object": "model"}]}`)
	}))
	defer srv.Close()

	_, err := ListModels(context.Background(), srv.URL+"/v1", "sk-test-key")
	if err != nil {
		t.Fatalf("ListModels() error: %v", err)
	}
	if gotAuth != "Bearer sk-test-key" {
		t.Errorf("Authorization = %q, want 'Bearer sk-test-key'", gotAuth)
	}
}
