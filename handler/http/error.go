package http

import (
	"encoding/json"
	"net/http"
)

func parseToErrorMsg(w http.ResponseWriter, httpStatus int, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatus)
	json.NewEncoder(w).Encode(map[string]string{
		"message": err.Error(),
	})
}
