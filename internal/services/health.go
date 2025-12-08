package services

import (
	"log"
	"net/http"

	u "github.com/nationpulse-bff/internal/utils"
)

type HealthService struct {
	// Add any dependencies like database connections here
	Configs *u.Configs
}

func (hs *HealthService) GetHealthByCountryCode(w http.ResponseWriter, r *http.Request) {
	countryCode := r.URL.Query().Get("countryCode")
	log.Printf("fetch health of %s\n", countryCode)
	w.Write([]byte("health fetched"))
}
