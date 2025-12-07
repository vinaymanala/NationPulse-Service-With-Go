package internals

import (
	"context"
	"net/http"

	"github.com/nationpulse-bff/internal/middlewares"
)

type ServerMux struct {
	UserMux       *http.ServeMux
	AdminMux      *http.ServeMux
	DashboardMux  *http.ServeMux
	PopulationMux *http.ServeMux
	HealthMux     *http.ServeMux
	EconomyMux    *http.ServeMux
	GrowthMux     *http.ServeMux
}

func (sm *ServerMux) NewServerMuxes() *ServerMux {
	return &ServerMux{
		UserMux:       http.NewServeMux(),
		AdminMux:      http.NewServeMux(),
		DashboardMux:  http.NewServeMux(),
		PopulationMux: http.NewServeMux(),
		HealthMux:     http.NewServeMux(),
		EconomyMux:    http.NewServeMux(),
		GrowthMux:     http.NewServeMux(),
	}
}

func groupRoutePrefix(prefix string, mux *http.ServeMux) http.Handler {
	return http.StripPrefix(prefix, mux)
}

func NewServer(ctx context.Context) http.Handler {
	rootMux := http.NewServeMux()

	newMux := &ServerMux{}
	muxes := newMux.NewServerMuxes()

	rootMux.Handle("/api/u/", groupRoutePrefix("/api/u", muxes.UserMux))
	rootMux.Handle("/api/a/",
		middlewares.WithAuthMiddlewares(ctx, groupRoutePrefix("/api/a", muxes.AdminMux)))
	rootMux.Handle("/api/dashboard/",
		middlewares.DefaultMiddlewares(ctx, groupRoutePrefix("/api/dashboard", muxes.DashboardMux)))
	rootMux.Handle("/api/health/",
		middlewares.WithAuthMiddlewares(ctx, groupRoutePrefix("/api/health", muxes.HealthMux)))
	rootMux.Handle("/api/population/",
		middlewares.WithAuthMiddlewares(ctx, groupRoutePrefix("/api/population", muxes.PopulationMux)))
	rootMux.Handle("/api/economy/",
		middlewares.WithAuthMiddlewares(ctx, groupRoutePrefix("/api/economy", muxes.EconomyMux)))
	rootMux.Handle("/api/growth/",
		middlewares.WithAuthMiddlewares(ctx, groupRoutePrefix("/api/growth", muxes.GrowthMux)))

	addRoutes(ctx, muxes)

	return rootMux
}
