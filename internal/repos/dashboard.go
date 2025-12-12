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

	var topCountriesByPopulation []TopPopulationByCountries

	for rows.Next() {
		var topPopulatedCountry TopPopulationByCountries
		if err := rows.Scan(
			&topPopulatedCountry.CountryName,
			&topPopulatedCountry.CountryCode,
			&topPopulatedCountry.Indicator,
			&topPopulatedCountry.Year,
			&topPopulatedCountry.Value); err != nil {
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

	var topCountriesByHealth []TopHealthCasesByCountries

	for rows.Next() {
		var topHealthDatByCountry TopHealthCasesByCountries
		if err := rows.Scan(
			&topHealthDatByCountry.CountryName,
			&topHealthDatByCountry.CountryCode,
			&topHealthDatByCountry.Year,
			&topHealthDatByCountry.Value,
			&topHealthDatByCountry.SexName,
			&topHealthDatByCountry.Cause,
			&topHealthDatByCountry.UnitRange,
		); err != nil {
			log.Fatalf("Error scanning a row: %v\n", err)
			return nil, err
		}
		fmt.Println(topHealthDatByCountry)
		topCountriesByHealth = append(topCountriesByHealth, topHealthDatByCountry)
	}
	defer rows.Close()
	return topCountriesByHealth, nil
}

func (dr *DashboardRepo) GetTopCountriesByGDPData(currentYear int, topNCountries int) (any, error) {
	sqlStatement := `SELECT * FROM get_highest_gdp_countries_dashboard($1, $2)`
	rows, err := dr.configs.Db.Client.Query(dr.configs.Context, sqlStatement, currentYear, topNCountries)
	if err != nil {
		return nil, err
	}

	var highestGDPCountries []HighestGDPCountries

	for rows.Next() {
		var highestGDPCountry HighestGDPCountries
		if err := rows.Scan(
			&highestGDPCountry.CountryCode,
			&highestGDPCountry.Year,
			&highestGDPCountry.Value); err != nil {
			log.Fatalf("Error scanning a row: %v\n", err)
			return nil, err
		}
		fmt.Println(highestGDPCountry)
		highestGDPCountries = append(highestGDPCountries, highestGDPCountry)
	}
	defer rows.Close()
	return highestGDPCountries, nil
}
