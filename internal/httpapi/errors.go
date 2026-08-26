package httpapi

import (
	"encoding/json"
	"net/http"
)

type ErrorBody struct {
	Code      string
	Message   string
	RequestID string
}

func writeJSONError(w http.ResponseWriter, status int, code, message, request string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(ErrorBody{Code: code, Message: message, RequestID: request})
}
