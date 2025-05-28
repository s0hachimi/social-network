package utils

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"errors"
	"net/http"
)

func GenerateSessionToken() (string, error) {
	tokne := make([]byte, 32)
	_, err := rand.Read(tokne)
	if err != nil {
		return "", errors.New("creation sissiontoken")
	}

	tokn := base64.URLEncoding.EncodeToString(tokne)
	return tokn, nil
}

func SendData(w http.ResponseWriter, code int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(data)
}
