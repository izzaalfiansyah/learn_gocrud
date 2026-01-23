package app

import (
	"encoding/json"
	"net/http"
)

func AppController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(map[string]any{
		"message": "CRUD Category",
		"endpoints": map[string][]string{
			"/categories":      {"GET", "POST"},
			"/categories/{id}": {"GET", "PUT", "DELETE"},
		},
	})
}
