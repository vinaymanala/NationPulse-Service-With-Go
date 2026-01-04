package services

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/nationpulse-bff/internal/repos"
	. "github.com/nationpulse-bff/internal/utils"
	"github.com/segmentio/kafka-go"
)

type UtilsService struct {
	Configs *Configs
	repo    *repos.UtilsRepo
}

func NewUtilsService(configs *Configs, repo *repos.UtilsRepo) *UtilsService {
	return &UtilsService{
		Configs: configs,
		repo:    repo,
	}
}

func (us *UtilsService) GetUserPermissions(w http.ResponseWriter, r *http.Request) {
	log.Println("Fetching permissions...")
	userID := r.Form.Get("userID")
	fmt.Println("USERID", userID)
	data, err := us.repo.GetPermissions(userID)
	if err != nil {
		http.Error(w, "failed", http.StatusInternalServerError)
		WriteJSON(w, http.StatusInternalServerError, nil, false, err.Error())
		return
	}
	WriteJSON(w, http.StatusOK, data, true, nil)
}

func (us *UtilsService) PublishExportRequest(w http.ResponseWriter, r *http.Request) {
	log.Println("Requesting reporting service to generate csv")
	var request ExportApiRequest

	req, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		fmt.Printf("Error reading from request body: %s", err)
		WriteJSON(w, http.StatusBadRequest, nil, false, err)

	}

	if err := json.Unmarshal(req, &request); err != nil {
		fmt.Println("Error unmarshalling request body", err)
		WriteJSON(w, http.StatusBadRequest, nil, false, err)
	}
	fmt.Println("REQUEST PAYLOAD", request)

	GetQueryAndHeaders(&request)
	// create new writer
	kw := us.Configs.Kafka.NewWriter("message-log")

	// create a timeout using context
	kwriteCtx, cancel := context.WithTimeout(us.Configs.Context, 30*time.Second)

	// marshal the request body to export request payload
	exportPayload, err := json.Marshal(request)
	if err != nil {
		fmt.Println("Error marshalling export payload for kafka message", err)
		WriteJSON(w, http.StatusBadRequest, nil, false, err)
	}

	// publish the message to kafka broker
	publishExportMessage := kafka.Message{
		Key:   []byte(request.ExportID),
		Value: []byte(exportPayload),
	}

	defer kw.Close()
	err = kw.WriteMessages(kwriteCtx, publishExportMessage)
	cancel()

	if err != nil {
		fmt.Println("Error publishing message to kafka broker: ", err)
		WriteJSON(w, http.StatusInternalServerError, nil, false, err)
	}

	var response = struct {
		ExportID string `json:"exportID"`
		Status   string `json:"status"`
	}{
		ExportID: request.ExportID,
		Status:   "processing",
	}

	fmt.Println("Response", response)
	WriteJSON(w, http.StatusOK, response, true, nil)
}

func (us *UtilsService) SubscribeExportResponse(w http.ResponseWriter, r *http.Request) {
	log.Println("Requesting reporting service to receive csv")
}
