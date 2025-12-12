package repos

import (
	"fmt"
	"log"

	. "github.com/nationpulse-bff/internal/utils"
)

type PopulationRepo struct {
	Configs *Configs
}

func NewPopulationRepo(configs *Configs) *PopulationRepo {
	return &PopulationRepo{
		Configs: configs,
	}
}

func (pr *PopulationRepo) GetPopulationByCountryData(countryCode string) (any, error) {
	sqlStatement := `SELECT * FROM get_population_by_country_code($1)`
	rows, err := pr.Configs.Db.Client.Query(pr.Configs.Context, sqlStatement, countryCode)
	if err != nil {
		return nil, err
	}

	var populationByCountries []PopulationData

	for rows.Next() {
		var populationByCountry PopulationData
		if err := rows.Scan(
			&populationByCountry.ID,
			&populationByCountry.CountryCode,
			&populationByCountry.CountryName,
			&populationByCountry.IndicatorCode,
			&populationByCountry.Indicator,
			&populationByCountry.SexCode,
			&populationByCountry.SexName,
			&populationByCountry.Age,
			&populationByCountry.Year,
			&populationByCountry.Value,
			&populationByCountry.LastUpdated,
		); err != nil {
			log.Fatalf("Error scanning a row: %v\n", err)
			return nil, err
		}
		fmt.Println(populationByCountry)
		populationByCountries = append(populationByCountries, populationByCountry)
	}
	defer rows.Close()
	return populationByCountries, nil
}
