package helpers

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type CustomError struct{}

func (e CustomError) ApiError(w http.ResponseWriter, status int, message string) {
	error := make(map[string]string)
	
	
	
	error["success"] = "false"
	error["status"] = strconv.Itoa(status)
	error["message"] = message

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(error)

}
