package repos

import (
	"fmt"
	"log"

	. "github.com/nationpulse-bff/internal/utils"
)

type HealthRepo struct {
	Configs *Configs
}

func NewHealthRepo(configs *Configs) *HealthRepo {
	return &HealthRepo{
		Configs: configs,
	}
}

func (hr *HealthRepo) GetHealthData(countryCode string) (any, error) {
	sqlStatement := `SELECT * FROM get_healthstatus_by_country_code($1)`
	rows, err := hr.Configs.Db.Client.Query(hr.Configs.Context, sqlStatement, countryCode)
	if err != nil {
		return nil, err
	}

	var healthData []HealthData
	for rows.Next() {
		var healthDataByCountry HealthData

		if err := rows.Scan(
			&healthDataByCountry.ID,
			&healthDataByCountry.CountryCode,
			&healthDataByCountry.CountryName,
			&healthDataByCountry.IndicatorCode,
			&healthDataByCountry.Indicator,
			&healthDataByCountry.SexCode,
			&healthDataByCountry.SexName,
			&healthDataByCountry.Cause,
			&healthDataByCountry.UnitRange,
			&healthDataByCountry.Year,
			&healthDataByCountry.Value,
			&healthDataByCountry.LastUpdated,
		); err != nil {
			log.Fatalf("Error scanning a row: %v\n", err)
			return nil, err
		}
		fmt.Println(healthDataByCountry)
		healthData = append(healthData, healthDataByCountry)
	}
	defer rows.Close()
	return healthData, nil
}
