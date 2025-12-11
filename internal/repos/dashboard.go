package repos

import (
	"fmt"
	"log"

	. "github.com/nationpulse-bff/internal/utils"
)

type DashboardRepo struct {
	configs *Configs
}

func NewDashboardRepo(configs *Configs) *DashboardRepo {
	return &DashboardRepo{
		configs: configs,
	}
}

func (dr *DashboardRepo) GetTopCountriesByPopulationData(currentYear int, top_countries int) (any, error) {
	sqlStatement := `SELECT * FROM get_perfgrowthpopulation_dashboard($1, $2);`
	rows, err := dr.configs.Db.Client.Query(dr.configs.Context, sqlStatement, currentYear, top_countries)
	if err != nil {
		return nil, err
	}

	var topCountriesByPopulation []TopCountriesByPopulation

	for rows.Next() {
		var topPopulatedCountry TopCountriesByPopulation
		if err := rows.Scan(
			&topPopulatedCountry.ID,
			&topPopulatedCountry.CountryName,
			&topPopulatedCountry.CountryCode,
			&topPopulatedCountry.Indicator,
			&topPopulatedCountry.IndicatorCode,
			&topPopulatedCountry.Year,
			&topPopulatedCountry.Value,
			&topPopulatedCountry.LastUpdated); err != nil {
			log.Fatalf("Error scanning a row: %v\n", err)
			return nil, err
		}
		fmt.Println(topPopulatedCountry)
		topCountriesByPopulation = append(topCountriesByPopulation, topPopulatedCountry)
	}
	defer rows.Close()
	return topCountriesByPopulation, nil
}

func (dr *DashboardRepo) GetTopCountriesByHealthData() (any, error) {
	sqlStatement := `SELECT * FROM get_healthstatus_dashboard()`
	rows, err := dr.configs.Db.Client.Query(dr.configs.Context, sqlStatement)
	if err != nil {
		return nil, err
	}

	var topCountriesByHealth []TopCountriesByHealth

	for rows.Next() {
		var topHealthDatByCountry TopCountriesByHealth
		if err := rows.Scan(
			&topHealthDatByCountry.ID,
			&topHealthDatByCountry.CountryName,
			&topHealthDatByCountry.CountryCode,
			&topHealthDatByCountry.Indicator,
			&topHealthDatByCountry.IndicatorCode,
			&topHealthDatByCountry.Cause,
			&topHealthDatByCountry.SexCode,
			&topHealthDatByCountry.SexName,
			&topHealthDatByCountry.UnitRange,
			&topHealthDatByCountry.Year,
			&topHealthDatByCountry.Value,
			&topHealthDatByCountry.LastUpdated); err != nil {
			log.Fatalf("Error scanning a row: %v\n", err)
			return nil, err
		}
		fmt.Println(topHealthDatByCountry)
		topCountriesByHealth = append(topCountriesByHealth, topHealthDatByCountry)
	}
	defer rows.Close()
	return topCountriesByHealth, nil
}

// TODO: Stored function need to be added in db
func (dr *DashboardRepo) GetTopCountriesByGDPData() (any, error) {
	sqlStatement := `SELECT * FROM get_healthstatus_dashboard()`
	rows, err := dr.configs.Db.Client.Query(dr.configs.Context, sqlStatement)
	if err != nil {
		return nil, err
	}

	var topCountriesByHealth []TopCountriesByHealth

	for rows.Next() {
		var topHealthDatByCountry TopCountriesByHealth
		if err := rows.Scan(
			&topHealthDatByCountry.ID,
			&topHealthDatByCountry.CountryName,
			&topHealthDatByCountry.CountryCode,
			&topHealthDatByCountry.Indicator,
			&topHealthDatByCountry.IndicatorCode,
			&topHealthDatByCountry.Cause,
			&topHealthDatByCountry.SexCode,
			&topHealthDatByCountry.SexName,
			&topHealthDatByCountry.UnitRange,
			&topHealthDatByCountry.Year,
			&topHealthDatByCountry.Value,
			&topHealthDatByCountry.LastUpdated); err != nil {
			log.Fatalf("Error scanning a row: %v\n", err)
			return nil, err
		}
		fmt.Println(topHealthDatByCountry)
		topCountriesByHealth = append(topCountriesByHealth, topHealthDatByCountry)
	}
	defer rows.Close()
	return topCountriesByHealth, nil
}
