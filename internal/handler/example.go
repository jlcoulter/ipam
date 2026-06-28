package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// Example handlers — delete and replace with your own domain handlers.

type Example struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func ListExamples(w http.ResponseWriter, r *http.Request) {
	// TODO: Fetch from store
	examples := []Example{}
	json.NewEncoder(w).Encode(examples)
}

func CreateExample(w http.ResponseWriter, r *http.Request) {
	var req Example
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	// TODO: Save to store
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(req)
}

func GetExample(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	// TODO: Fetch from store
	_ = id
	http.Error(w, "not found", http.StatusNotFound)
}

func DeleteExample(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	// TODO: Delete from store
	_ = id
	w.WriteHeader(http.StatusNoContent)
}