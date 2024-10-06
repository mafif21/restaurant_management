package config

import (
	"encoding/json"
	"net/http"
	dto "restaurant_management/internal/models/dtos"
)

func ReadBodyRequest(r *http.Request, w http.ResponseWriter, data any) {
	if err := json.NewDecoder(r.Body).Decode(data); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

func WriteToBodyResponse(w http.ResponseWriter, data *dto.Response) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}
