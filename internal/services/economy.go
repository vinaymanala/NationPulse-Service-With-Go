package services

import (
	"log"
	"net/http"

	. "github.com/nationpulse-bff/internal/utils"
)

type EconomyService struct {
	// Add any dependencies like database connections here
	Configs *Configs
}

func NewEconomyService(configs *Configs) *EconomyService {
	return &EconomyService{
		Configs: configs,
	}
}

func (es *EconomyService) GetEconomyGovernmentDataByCountryCode(w http.ResponseWriter, r *http.Request) {
	countryCode := r.URL.Query().Get("countryCode")
	log.Printf("Economy GovermentData of %s\n", countryCode)
	w.Write([]byte("economy fetch"))
}

func (es *EconomyService) GetEconomyGDPByCountryCode(w http.ResponseWriter, r *http.Request) {
	countryCode := r.URL.Query().Get("countryCode")
	log.Printf("Economy GDP of %s\n", countryCode)
	w.Write([]byte("economy gdp fetch"))
}
