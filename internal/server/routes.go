package internals

import (
	"context"
	"fmt"
	"net/http"

	"github.com/nationpulse-bff/internal/services"
	"github.com/nationpulse-bff/internal/store"
	u "github.com/nationpulse-bff/internal/utils"
)

func handleDashboardRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Handling Dashboard Route..")
}
func handleHealthRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Handling Health Route..")
}
func handlePopulationRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Handling Population Route..")
}
func handleEconomyRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Handling Economy Route..")
}
func handleGrowthRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Handling PerformanceGrowth Route..")
}

func addRoutes(ctx context.Context, muxes *ServerMux) {
	configs := &u.Configs{
		Db:      ctx.Value("db").(*store.PgClient).Client,
		Cache:   ctx.Value("redisClient").(*store.Redis),
		Context: ctx,
	}

	userService := &services.UserService{Configs: configs}
	adminService := &services.AdminService{Configs: configs}
	dashboardService := &services.DashboardService{Configs: configs}
	populationService := &services.PopulationService{Configs: configs}
	healthService := &services.HealthService{Configs: configs}
	economyService := &services.EconomyService{Configs: configs}
	growthService := &services.GrowthService{Configs: configs}

	muxes.UserMux.HandleFunc("POST /login", userService.HandleLogin)
	muxes.UserMux.HandleFunc("POST /logout", userService.HandleLogout)
	muxes.UserMux.HandleFunc("GET /token/refresh", userService.HandleRefreshToken)
	muxes.AdminMux.HandleFunc("GET /permissions", adminService.GetAllPermissions)

	muxes.DashboardMux.HandleFunc("GET /", handleDashboardRoute)
	muxes.DashboardMux.HandleFunc("GET /population/topCountriesByPopulation", dashboardService.GetTopCountriesByPopulation)
	muxes.DashboardMux.HandleFunc("GET /api/health/topCountriesByHealth", dashboardService.GetTopCountriesByHealth)
	muxes.DashboardMux.HandleFunc("GET /api/gdp/topCountriesByGDP", dashboardService.GetTopCountriesByGDP)

	muxes.PopulationMux.HandleFunc("GET /", handlePopulationRoute)
	muxes.PopulationMux.HandleFunc("GET /api/country/{countryCode}", populationService.GetPopulationByCountryCode)

	muxes.HealthMux.HandleFunc("GET /", handleHealthRoute)
	muxes.HealthMux.HandleFunc("GET /api/country/{countryCode}", healthService.GetHealthByCountryCode)

	muxes.EconomyMux.HandleFunc("GET /", handleEconomyRoute)
	muxes.EconomyMux.HandleFunc("GET /api/governmentdata/country/{countryCode}", economyService.GetEconomyGovernmentDataByCountryCode)
	muxes.EconomyMux.HandleFunc("GET /api/gdp/country/{countryCode}", economyService.GetEconomyGDPByCountryCode)

	muxes.GrowthMux.HandleFunc("GET /", handleGrowthRoute)
	muxes.GrowthMux.HandleFunc("GET /api/gdpgrowth/country/{countryCode}", growthService.GetGDPGrowthByCountryCode)
	muxes.GrowthMux.HandleFunc("GET /api/populationgrowth/country/{countryCode}", growthService.GetPopulationGrowthByCountryCode)
}
