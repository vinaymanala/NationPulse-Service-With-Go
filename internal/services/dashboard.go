package services

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/nationpulse-bff/internal/repos"
	. "github.com/nationpulse-bff/internal/utils"
)

type DashboardService struct {
	// Add any dependencies like database connections here
	Configs *Configs
	repo    *repos.DashboardRepo
}

func NewDashboardService(configs *Configs, repo *repos.DashboardRepo) *DashboardService {
	return &DashboardService{
		Configs: configs,
		repo:    repo,
	}
}

func (ds *DashboardService) GetTopCountriesByPopulation(w http.ResponseWriter, r *http.Request) {
	log.Println("fetch top 5 populated countries")
	data, err := ds.repo.GetTopCountriesByPopulationData(2025, 5)
	if err != nil {
		http.Error(w, "failed", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
	w.Write([]byte("Top 5 countries by population"))
}

func (ds *DashboardService) GetTopCountriesByHealth(w http.ResponseWriter, r *http.Request) {
	log.Println("fetch top 5 health related cases in countries")
	data, err := ds.repo.GetTopCountriesByHealthData()
	if err != nil {
		http.Error(w, "failed", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
	w.Write([]byte("Top 5 countries with Health cases"))
}

func (ds *DashboardService) GetTopCountriesByGDP(w http.ResponseWriter, r *http.Request) {
	log.Println("fetch top 5 gdp countries")
	data, err := ds.repo.GetTopCountriesByGDPData(2025, 5)
	if err != nil {
		http.Error(w, "failed", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
	w.Write([]byte("Top 5 countries with highest GDP"))
}
