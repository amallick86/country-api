package db

import (
	"context"
	"github.com/amallick86/country-api/models"
)

const addState = `
INSERT INTO state (
  id, state_name , country_id 
) VALUES (
  $1, $2, $3
)
RETURNING id,state_name,country_id,created_at
`

type AddStateParams struct {
	ID        int    `json:"id"`
	StateName string `json:"state_name" `
	CountryId int    `json:"country_id"`
}

func (q *Queries) AddState(ctx context.Context, arg AddStateParams) (models.State, error) {
	row := q.db.QueryRowContext(ctx, addState, arg.ID, arg.StateName, arg.CountryId)
	var i models.State
	err := row.Scan(
		&i.ID,
		&i.StateName,
		&i.CountryId,
		&i.CreatedAt,
	)
	return i, err
}

const getTotalStateCount = `
SELECT COUNT(*) FROM state
`

func (q *Queries) GetTotalStateCount(ctx context.Context) (models.StateCount, error) {
	row := q.db.QueryRowContext(ctx, getTotalStateCount)
	var i models.StateCount
	err := row.Scan(
		&i.TotalStateCount,
	)
	return i, err
}

const getStateList = `
SELECT state.id,state_name,country_id,state.created_at,country.name FROM state INNER JOIN country ON country.id= state.country_id where state.id >= $1 LIMIT $2`

type GetStateListParams struct {
	FromId int `json:"fromId"`
	Limit  int `json:"limit"`
}

func (q *Queries) GetStateList(ctx context.Context, arg GetStateListParams) ([]models.State, error) {
	rows, err := q.db.QueryContext(ctx, getStateList, arg.FromId, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []models.State{}
	for rows.Next() {
		var i models.State
		if err := rows.Scan(
			&i.ID,
			&i.StateName,
			&i.CountryId,
			&i.CreatedAt,
			&i.Name,
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

const getStateListByCountry = `
SELECT state.id,state_name,country_id,state.created_at,country.name FROM state INNER JOIN country ON country.id= state.country_id WHERE lower(country.name)=$1`

type GetStateListByCountryParams struct {
	Name string `json:"name" `
}

func (q *Queries) GetStateListByCountry(ctx context.Context, arg GetStateListByCountryParams) ([]models.State, error) {
	rows, err := q.db.QueryContext(ctx, getStateListByCountry, arg.Name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []models.State{}
	for rows.Next() {
		var i models.State
		if err := rows.Scan(
			&i.ID,
			&i.StateName,
			&i.CountryId,
			&i.CreatedAt,
			&i.Name,
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
