package repos

import (
	"fmt"
	"log"

	. "github.com/nationpulse-bff/internal/utils"
)

type EconomyRepo struct {
	Configs *Configs
}

var economyId = "economy:"

func NewEconomyRepo(configs *Configs) *EconomyRepo {
	return &EconomyRepo{
		Configs: configs,
	}
}

func (er *EconomyRepo) GetGovernmentData(countryCode string) (any, error) {
	var governmentData []EconomyData
	data, err := GetDataFromCache(er.Configs, economyId+"government", &governmentData)
	if err != nil {
		log.Println("Cache Get Failed. Trying DB.")
	} else {
		return *data, nil
	}
	sqlStatement := `SELECT * FROM get_publicgovernmentyearly_by_country_code($1)`
	rows, err := er.Configs.Db.Client.Query(er.Configs.Context, sqlStatement, countryCode)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var economygovernmentByCountry EconomyData

		if err := rows.Scan(
			&economygovernmentByCountry.ID,
			&economygovernmentByCountry.CountryCode,
			&economygovernmentByCountry.CountryName,
			&economygovernmentByCountry.IndicatorCode,
			&economygovernmentByCountry.Indicator,
			&economygovernmentByCountry.Year,
			&economygovernmentByCountry.Value,
			&economygovernmentByCountry.LastUpdated,
		); err != nil {
			log.Fatalf("Error scanning a row: %v\n", err)
			return nil, err
		}
		fmt.Println(economygovernmentByCountry)
		governmentData = append(governmentData, economygovernmentByCountry)
	}
	if err := er.Configs.Cache.SetData(er.Configs.Context, economyId+"government", governmentData); err != nil {
		log.Println("Error Set Cache Data", err)
	}
	defer rows.Close()
	return governmentData, nil
}

func (er *EconomyRepo) GetGDPData(countryCode string) (any, error) {
	var gdpData []EconomyData
	data, err := GetDataFromCache(er.Configs, populationID+"GDP", &gdpData)
	if err != nil {
		log.Println("Cache Get Failed. Trying DB.")
	} else {
		return *data, nil
	}
	sqlStatement := `SELECT * FROM get_gdppercapita_by_country_code($1)`
	rows, err := er.Configs.Db.Client.Query(er.Configs.Context, sqlStatement, countryCode)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var economyGDPByCountry EconomyData

		if err := rows.Scan(
			&economyGDPByCountry.ID,
			&economyGDPByCountry.CountryCode,
			&economyGDPByCountry.CountryName,
			&economyGDPByCountry.IndicatorCode,
			&economyGDPByCountry.Indicator,
			&economyGDPByCountry.Year,
			&economyGDPByCountry.Value,
			&economyGDPByCountry.LastUpdated,
		); err != nil {
			log.Fatalf("Error scanning a row: %v\n", err)
			return nil, err
		}
		fmt.Println(economyGDPByCountry)
		gdpData = append(gdpData, economyGDPByCountry)
	}
	if err := er.Configs.Cache.SetData(er.Configs.Context, economyId+"GDP", gdpData); err != nil {
		log.Println("Error Set Cache Data", err)
	}
	defer rows.Close()
	return gdpData, nil

}
