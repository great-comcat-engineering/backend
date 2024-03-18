package handler

import (
	"encoding/json"
	"net/http"
	"sync"
)

// Test represents a test item structure.
type Test struct {
	ID        string `json:"id"`
	Item      string `json:"item"`
	Completed bool   `json:"completed"`
}

var (
	// mutex to avoid race conditions
	mu sync.Mutex
	// tests simulates a database table of Test items.
	tests = make(map[string]Test)
)

// Handler is the main entry point for the serverless function.
func AuthLoginHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		handleGet(w, r)
	case "POST":
		handlePost(w, r)
	case "PUT":
		handlePut(w, r)
	case "DELETE":
		handleDelete(w, r)
	default:
		http.Error(w, "Unsupported HTTP method", http.StatusMethodNotAllowed)
	}
}

func handleGet(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	items := make([]Test, 0, len(tests))
	for _, item := range tests {
		items = append(items, item)
	}

	json.NewEncoder(w).Encode(items)
}

func handlePost(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	var t Test
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	tests[t.ID] = t
	json.NewEncoder(w).Encode(t)
}

func handlePut(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	var t Test
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if _, exists := tests[t.ID]; !exists {
		http.Error(w, "Item not found", http.StatusNotFound)
		return
	}

	tests[t.ID] = t
	json.NewEncoder(w).Encode(t)
}

func handleDelete(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "ID query parameter is required", http.StatusBadRequest)
		return
	}

	if _, exists := tests[id]; !exists {
		http.Error(w, "Item not found", http.StatusNotFound)
		return
	}

	delete(tests, id)
	w.WriteHeader(http.StatusNoContent)
}
