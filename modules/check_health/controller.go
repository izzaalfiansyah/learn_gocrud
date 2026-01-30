package check_health

import (
	"encoding/json"
	"net/http"
)

func CheckHealthController(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]any{
		"status":  "ok",
		"message": "Application is running",
	})
}
