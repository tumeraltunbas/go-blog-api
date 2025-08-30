package handlers

import (
	"encoding/json"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	data := map[string]bool{
		"success": true,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(data)
}
