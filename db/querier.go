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
	GetCountriesList(ctx context.Context) ([]models.Country, error)
	AddState(ctx context.Context, arg AddStateParams) (models.State, error)
	GetStateList(ctx context.Context) ([]models.State, error)
	GetStateListByCountry(ctx context.Context, arg GetStateListByCountryParams) ([]models.State, error)
}

var _ Querier = (*Queries)(nil)
