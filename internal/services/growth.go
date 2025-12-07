package services

import (
	"log"
	"net/http"
)

type GrowthService struct {
	// Add any dependencies like database connections here
}

func (gs *GrowthService) GetGDPGrowthByCountryCode(w http.ResponseWriter, r *http.Request) {
	countryCode := r.URL.Query().Get("countryCode")
	log.Printf("fetch Gdp growth of %s\n", countryCode)
	w.Write([]byte("fetch gfp growth"))
}

func (gs *GrowthService) GetPopulationGrowthByCountryCode(w http.ResponseWriter, r *http.Request) {
	countryCode := r.URL.Query().Get("countryCode")
	log.Printf("fetch population growth of %s\n", countryCode)
	w.Write([]byte("fetch population growth"))

}
