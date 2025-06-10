package utils

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func GenerateSessionToken() (string, time.Time) {
	sessionToken := uuid.New().String()
	expiration := time.Now().Add(12 * time.Hour)
	return sessionToken, expiration
}

func SendData(w http.ResponseWriter, code int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(data)
}
