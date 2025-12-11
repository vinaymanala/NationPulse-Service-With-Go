package utils

import (
	"context"
	"time"

	"github.com/nationpulse-bff/internal/store"
)

type Configs struct {
	Db      *store.PgClient
	Cache   *store.Redis
	Context context.Context
}

type TopCountriesByPopulation struct {
	ID            int       `json:"id"`
	CountryCode   string    `json:"country_code"`
	CountryName   string    `json:"country_name"`
	IndicatorCode string    `json:"indicator_code"`
	Indicator     string    `json:"indicator"`
	Year          int       `json:"year"`
	Value         float64   `json:"value"`
	LastUpdated   time.Time `json:"last_updated"`
}

type TopCountriesByHealth struct {
	ID            int       `json:"id"`
	CountryCode   string    `json:"country_code"`
	CountryName   string    `json:"country_name"`
	IndicatorCode string    `json:"indicator_code"`
	Indicator     string    `json:"indicator"`
	SexCode       string    `json:"sex_code"`
	SexName       string    `json:"sex_name"`
	Cause         string    `json:"cause"`
	UnitRange     string    `json:"unitRange"`
	Year          int       `json:"year"`
	Value         float64   `json:"value"`
	LastUpdated   time.Time `json:"last_updated"`
}
