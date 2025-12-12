package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

func WriteJSON(w http.ResponseWriter, status int, data any, success bool, err any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	var response = &ApiResponse{
		Data:      data,
		IsSuccess: success,
		Error:     err,
	}
	log.Printf("Message: %s, isSuccess: %t, Error:%v \n", data, success, err)
	return json.NewEncoder(w).Encode(response)
}
