package db

import (
	"context"
	"github.com/amallick86/country-api/models"
)

const addCountry = `
INSERT INTO country (
  id, name , country_short_name 
) VALUES (
  $1, $2, $3
)
RETURNING id, name, country_short_name, created_at
`

type AddCountryParams struct {
	ID               int    `json:"id"`
	Name             string `json:"name" `
	CountryShortName string `json:"country_short_name"`
}

func (q *Queries) AddCountry(ctx context.Context, arg AddCountryParams) (models.Country, error) {
	row := q.db.QueryRowContext(ctx, addCountry, arg.ID, arg.Name, arg.CountryShortName)
	var i models.Country
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.CountryShortName,
		&i.CreatedAt,
	)
	return i, err
}

const getCountriesList = `
SELECT id,name,country_short_name,created_at FROM country `

func (q *Queries) GetCountriesList(ctx context.Context) ([]models.Country, error) {
	rows, err := q.db.QueryContext(ctx, getCountriesList)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []models.Country{}
	for rows.Next() {
		var i models.Country
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.CountryShortName,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
