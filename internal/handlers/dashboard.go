package handlers

import (
	"net/http"

	"github.com/nationpulse-bff/internal/services"
)

type DashboardHandler struct {
	// Add any dependencies like services here
	mux     *http.ServeMux
	service *services.DashboardService
}

func NewDashboardHandler(mux *http.ServeMux, service *services.DashboardService) *DashboardHandler {
	return &DashboardHandler{
		mux:     mux,
		service: service,
	}
}

func (dh *DashboardHandler) RegisterRoutes() {
	// Register dashboard-related routes here
	//dh.mux.HandleFunc("GET /", handleDashboardRoute)
	dh.mux.HandleFunc("GET /population/topCountriesByPopulation", dh.service.GetTopCountriesByPopulation)
	dh.mux.HandleFunc("GET /api/health/topCountriesByHealth", dh.service.GetTopCountriesByHealth)
	dh.mux.HandleFunc("GET /api/gdp/topCountriesByGDP", dh.service.GetTopCountriesByGDP)

}
