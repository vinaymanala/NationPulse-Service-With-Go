package services

import (
	"log"
	"net/http"
)

type DashboardService struct {
	// Add any dependencies like database connections here

}

func (ds *DashboardService) GetTopCountriesByPopulation(w http.ResponseWriter, r *http.Request) {
	log.Println("fetch top 5 populated countries")
	w.Write([]byte("Top 5 countries by population"))
}

func (ds *DashboardService) GetTopCountriesByHealth(w http.ResponseWriter, r *http.Request) {
	log.Println("fetch top 5 health related cases in countries")
	w.Write([]byte("Top 5 countries with Health cases"))
}

func (ds *DashboardService) GetTopCountriesByGDP(w http.ResponseWriter, r *http.Request) {
	log.Println("fetch top 5 gdp countries")
	w.Write([]byte("Top 5 countries with highest GDP"))
}
