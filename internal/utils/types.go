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

type ApiResponse struct {
	IsSuccess bool `json:"isSuccess"`
	Data      any  `json:"data"`
	Error     any  `json:"error"`
}
type TopPopulationByCountries struct {
	CountryCode string  `json:"country_code"`
	CountryName string  `json:"country_name"`
	Indicator   string  `json:"indicator"`
	Year        int     `json:"year"`
	Value       float64 `json:"value"`
}

type TopHealthCasesByCountries struct {
	CountryCode string  `json:"country_code"`
	CountryName string  `json:"country_name"`
	Year        int     `json:"year"`
	Value       float64 `json:"value"`
	SexName     string  `json:"sex_name"`
	Cause       string  `json:"cause"`
	UnitRange   string  `json:"unitRange"`
}

type PopulationData struct {
	ID            int       `json:"id"`
	CountryCode   string    `json:"country_code"`
	CountryName   string    `json:"country_name"`
	IndicatorCode string    `json:"indicator_code"`
	Indicator     string    `json:"indicator"`
	SexCode       string    `json:"sex_code"`
	SexName       string    `json:"sex_name"`
	Age           string    `json:"age"`
	Year          int       `json:"year"`
	Value         float64   `json:"value"`
	LastUpdated   time.Time `json:"last_updated"`
}

type HealthData struct {
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

type HighestGDPCountries struct {
	CountryCode string  `json:"country_code"`
	Year        string  `json:"year"`
	Value       float64 `json:"value"`
}

type EconomyData struct {
	ID            int       `json:"id"`
	CountryCode   string    `json:"country_code"`
	CountryName   string    `json:"country_name"`
	IndicatorCode string    `json:"indicator_code"`
	Indicator     string    `json:"indicator"`
	Year          string    `json:"year"`
	Value         float64   `json:"value"`
	LastUpdated   time.Time `json:"last_updated"`
}

type GrowthData struct {
	ID            int       `json:"id"`
	CountryCode   string    `json:"country_code"`
	CountryName   string    `json:"country_name"`
	IndicatorCode string    `json:"indicator_code"`
	Indicator     string    `json:"indicator"`
	Year          string    `json:"year"`
	Value         float64   `json:"value"`
	LastUpdated   time.Time `json:"last_updated"`
}
