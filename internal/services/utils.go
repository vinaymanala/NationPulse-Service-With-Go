package services

import (
	"fmt"
	"log"
	"net/http"

	"github.com/nationpulse-bff/internal/repos"
	. "github.com/nationpulse-bff/internal/utils"
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
