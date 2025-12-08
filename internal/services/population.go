package services

import (
	"log"
	"net/http"

	u "github.com/nationpulse-bff/internal/utils"
)

type PopulationService struct {
	// Add any dependencies like database connections here
	Configs *u.Configs
}

func (ps *PopulationService) GetPopulationByCountryCode(w http.ResponseWriter, r *http.Request) {
	countryCode := r.URL.Query().Get("countryCode")
	log.Printf("fetch population of %s\n", countryCode)
	w.Write([]byte("Population fetched.."))
}
