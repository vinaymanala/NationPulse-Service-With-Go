package services

import (
	"log"
	"net/http"
)

type EconomyService struct {
	// Add any dependencies like database connections here
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
