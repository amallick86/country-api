package db

import (
	"context"
	"github.com/amallick86/country-api/models"
	"github.com/google/uuid"
)

type Querier interface {
	CreateUser(ctx context.Context, arg CreateUserParams) (models.User, error)
	GetUser(ctx context.Context, username string) (models.User, error)
	GetUserById(ctx context.Context, userid int32) (models.User, error)
	CreateSession(ctx context.Context, arg CreateSessionParams) (models.Sessions, error)
	GetSession(ctx context.Context, id uuid.UUID) (models.Sessions, error)
	AddCountry(ctx context.Context, arg AddCountryParams) (models.Country, error)
	GetTotalCountryCount(ctx context.Context) (models.CountryCount, error)
	GetCountriesList(ctx context.Context, arg GetCountriesListParams) ([]models.Country, error)
	AddState(ctx context.Context, arg AddStateParams) (models.State, error)
	GetTotalStateCount(ctx context.Context) (models.StateCount, error)
	GetStateList(ctx context.Context, arg GetStateListParams) ([]models.State, error)
	GetStateListByCountry(ctx context.Context, arg GetStateListByCountryParams) ([]models.State, error)
}

var _ Querier = (*Queries)(nil)
