package app

import (
	"encoding/json"
	"net/http"

	"izzaalfiansyah/learn_gocrud/config"
)

func AppController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(map[string]any{
		"app": map[string]string{
			"name": config.Env.AppName,
		},
		"endpoints": map[string][]string{
			"/categories":      {"GET", "POST"},
			"/categories/{id}": {"GET", "PUT", "DELETE"},
			"/products":        {"GET", "POST"},
			"/products/{id}":   {"GET", "PUT", "DELETE"},
		},
	})
}
