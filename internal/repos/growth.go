package repos

import (
	"fmt"
	"log"

	. "github.com/nationpulse-bff/internal/utils"
)

type GrowthRepo struct {
	Configs *Configs
}

func NewGrowthRepo(configs *Configs) *GrowthRepo {
	return &GrowthRepo{
		Configs: configs,
	}
}

func (gr *GrowthRepo) GetGDPGrowthData(countryCode string) (any, error) {
	sqlStatement := `SELECT * FROM get_perfgrowthgdpds_by_country_code($1)`
	rows, err := gr.Configs.Db.Client.Query(gr.Configs.Context, sqlStatement, countryCode)
	if err != nil {
		return nil, err
	}

	var gdpGrowthData []GrowthData
	for rows.Next() {
		var growthgdpDataByCountry GrowthData

		if err := rows.Scan(
			&growthgdpDataByCountry.ID,
			&growthgdpDataByCountry.CountryCode,
			&growthgdpDataByCountry.CountryName,
			&growthgdpDataByCountry.IndicatorCode,
			&growthgdpDataByCountry.Indicator,
			&growthgdpDataByCountry.Year,
			&growthgdpDataByCountry.Value,
			&growthgdpDataByCountry.LastUpdated,
		); err != nil {
			log.Fatalf("Error scanning a row: %v\n", err)
			return nil, err
		}
		fmt.Println(growthgdpDataByCountry)
		gdpGrowthData = append(gdpGrowthData, growthgdpDataByCountry)
	}
	defer rows.Close()
	return gdpGrowthData, nil
}

func (gr *GrowthRepo) GetPopulationGrowth(countryCode string) (any, error) {
	sqlStatement := `SELECT * FROM get_perfgrowthpopulation_by_country_code($1)`
	rows, err := gr.Configs.Db.Client.Query(gr.Configs.Context, sqlStatement, countryCode)
	if err != nil {
		return nil, err
	}

	var populationGrowthData []GrowthData
	for rows.Next() {
		var populationGrowthByCountry GrowthData

		if err := rows.Scan(
			&populationGrowthByCountry.ID,
			&populationGrowthByCountry.CountryCode,
			&populationGrowthByCountry.CountryName,
			&populationGrowthByCountry.IndicatorCode,
			&populationGrowthByCountry.Indicator,
			&populationGrowthByCountry.Year,
			&populationGrowthByCountry.Value,
			&populationGrowthByCountry.LastUpdated,
		); err != nil {
			log.Fatalf("Error scanning a row: %v\n", err)
			return nil, err
		}
		fmt.Println(populationGrowthByCountry)
		populationGrowthData = append(populationGrowthData, populationGrowthByCountry)
	}
	defer rows.Close()
	return populationGrowthData, nil

}
