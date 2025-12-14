package repos

import (
	"encoding/json"
	"fmt"
	"log"

	. "github.com/nationpulse-bff/internal/utils"
)

type DashboardRepo struct {
	Configs *Configs
}

var dashboardID = "dashboard:"

func NewDashboardRepo(configs *Configs) *DashboardRepo {
	return &DashboardRepo{
		Configs: configs,
	}
}

func (dr *DashboardRepo) GetTopCountriesByPopulationData(currentYear int, top_countries int) (any, error) {
	var topCountriesByPopulation []TopPopulationByCountries
	data, err := GetDataFromCache(dr.Configs, dashboardID+"population", &topCountriesByPopulation)
	if err != nil {
		log.Println("Cache Get Failed. Trying DB.")
	} else {
		return *data, nil
	}
	// Cache nil, checking DB
	sqlStatement := `SELECT * FROM get_perfgrowthpopulation_dashboard($1, $2);`
	rows, err := dr.Configs.Db.Client.Query(dr.Configs.Context, sqlStatement, currentYear, top_countries)
	if err != nil {
		return nil, err
	}

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
		//fmt.Println(topPopulatedCountry)
		topCountriesByPopulation = append(topCountriesByPopulation, topPopulatedCountry)
	}
	defer rows.Close()
	marshalledData, err := json.Marshal(topCountriesByPopulation)
	if err != nil {
		log.Println("Error marshaling data to cache")
	}
	if err := dr.Configs.Cache.SetData(dr.Configs.Context, dashboardID+"population", marshalledData); err != nil {
		log.Println("Error set cache data:", err)
	}
	fmt.Println("Call Successfull")
	return topCountriesByPopulation, nil
}

func (dr *DashboardRepo) GetTopCountriesByHealthData() (any, error) {
	if data, err := dr.Configs.Cache.GetData(dr.Configs.Context, dashboardID+"healthdata"); err != nil {
		log.Printf("Error fetching data from cache. Trying DB...\n")
	} else {
		return data, nil
	}
	sqlStatement := `SELECT * FROM get_healthstatus_dashboard()`
	rows, err := dr.Configs.Db.Client.Query(dr.Configs.Context, sqlStatement)
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
		//fmt.Println(topHealthDatByCountry)
		topCountriesByHealth = append(topCountriesByHealth, topHealthDatByCountry)
	}
	if err := dr.Configs.Cache.SetData(dr.Configs.Context, dashboardID+"healthdata", topCountriesByHealth); err != nil {
		log.Println("Error Set Cache data", err)
	}
	defer rows.Close()
	return topCountriesByHealth, nil
}

func (dr *DashboardRepo) GetTopCountriesByGDPData(currentYear int, topNCountries int) (any, error) {
	if data, err := dr.Configs.Cache.GetData(dr.Configs.Context, dashboardID+"GDPdata"); err != nil {
		log.Printf("Error fetching data from cache. Trying DB...\n")
	} else {
		return data, nil
	}

	sqlStatement := `SELECT * FROM get_highest_gdp_countries_dashboard($1, $2)`
	rows, err := dr.Configs.Db.Client.Query(dr.Configs.Context, sqlStatement, currentYear, topNCountries)
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
	if err := dr.Configs.Cache.SetData(dr.Configs.Context, dashboardID+"GDPdata", highestGDPCountries); err != nil {
		log.Println("Error Set Cache Data", err)
	}
	return highestGDPCountries, nil
}
