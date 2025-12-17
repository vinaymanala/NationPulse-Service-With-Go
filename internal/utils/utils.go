package utils

import (
	"encoding/json"
	"errors"
	"fmt"
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
	log.Printf("Message: %s, isSuccess: %t, Error:%v \n", "Data received", success, err)
	return json.NewEncoder(w).Encode(response)
}

func GetUserDetailsFromCache(r *http.Request, configs *Configs) {
	fmt.Printf("Form Details: %s\n", r.Form.Get("userID"))
}

func GetDataFromCache[T any](configs *Configs, key string, mappedStruct T) (*T, error) {
	//var zero T
	data, err := configs.Cache.GetData(configs.Context, key)
	fmt.Println("CACHE DATA:", data, err)
	if err != nil {
		log.Printf("Error fetching data from cache %s\n", err)
		return nil, errors.New("error fetching data from cache")
	}
	if err := json.Unmarshal([]byte(data), &mappedStruct); err != nil {
		log.Println("Error unmarshalling data from cache.")
		return nil, errors.New("error unmarshalling data from cache")
	}
	fmt.Println("==================================")
	// fmt.Println("Unmarshal data", &mappedStruct)
	fmt.Println("Fetched Data from Cache!!")
	return &mappedStruct, nil

}
