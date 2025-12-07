package services

import (
	"log"
	"net/http"
)

type HealthService struct {
	// Add any dependencies like database connections here
}

func (hs *HealthService) GetHealthByCountryCode(w http.ResponseWriter, r *http.Request) {
	countryCode := r.URL.Query().Get("countryCode")
	log.Printf("fetch health of %s\n", countryCode)
	w.Write([]byte("health fetched"))
}
