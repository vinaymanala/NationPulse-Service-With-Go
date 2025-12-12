package services

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/nationpulse-bff/internal/repos"
	. "github.com/nationpulse-bff/internal/utils"
)

type GrowthService struct {
	// Add any dependencies like database connections here
	Configs *Configs
	repo    *repos.GrowthRepo
}

func NewGrowthService(configs *Configs, repo *repos.GrowthRepo) *GrowthService {
	return &GrowthService{
		Configs: configs,
		repo:    repo,
	}
}

func (gs *GrowthService) GetGDPGrowthByCountryCode(w http.ResponseWriter, r *http.Request) {
	countryCode := r.URL.Query().Get("countryCode")
	log.Printf("fetch Gdp growth of %s\n", countryCode)
	data, err := gs.repo.GetGDPGrowthData(countryCode)
	if err != nil {
		http.Error(w, "failed", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
	w.Write([]byte("fetch gfp growth"))
}

func (gs *GrowthService) GetPopulationGrowthByCountryCode(w http.ResponseWriter, r *http.Request) {
	countryCode := r.URL.Query().Get("countryCode")
	data, err := gs.repo.GetPopulationGrowth(countryCode)
	if err != nil {
		http.Error(w, "failed", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
	log.Printf("fetch population growth of %s\n", countryCode)
	w.Write([]byte("fetch population growth"))

}
