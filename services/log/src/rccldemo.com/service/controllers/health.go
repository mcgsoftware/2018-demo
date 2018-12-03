package controllers

import (
	"encoding/json"
	"net/http"
)

// Handles /health
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	values := map[string]string{"health": "Up"}
	json, _ := json.Marshal(values)
	w.Write(json)

}