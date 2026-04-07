package utils

import (
	"encoding/json"
	"net/http"
)

func WriteJsonResponse(w http.ResponseWriter, data any) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(data)
}

func WriteJsonSuccessResponse(w http.ResponseWriter, data any) error {
	response := map[string]any{}
	response["status"] = "success"
	response["error"] = false
	response["data"] = data
	return WriteJsonResponse(w, response)
}

func WriteJsonErrorResponse(w http.ResponseWriter, err error) error {
	response := map[string]any{}
	response["status"] = "failure"
	response["error"] = true
	response["data"] = err
	return WriteJsonResponse(w, response)
}
