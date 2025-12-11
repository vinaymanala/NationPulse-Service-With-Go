package handlers

import (
	"net/http"

	"github.com/nationpulse-bff/internal/services"
)

type EconomyHandler struct {
	mux     *http.ServeMux
	service *services.EconomyService
}

func NewEconomyHandler(mux *http.ServeMux, service *services.EconomyService) *EconomyHandler {
	return &EconomyHandler{
		mux:     mux,
		service: service,
	}
}

func (eh *EconomyHandler) RegisterRoutes() {
	eh.mux.HandleFunc("GET /api/governmentdata/country/{countryCode}", eh.service.GetEconomyGovernmentDataByCountryCode)
	eh.mux.HandleFunc("GET /api/gdp/country/{countryCode}", eh.service.GetEconomyGDPByCountryCode)

}
