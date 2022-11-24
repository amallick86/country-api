package models

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID        int32     `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdAt"`
}

type Sessions struct {
	ID           uuid.UUID `json:"id"`
	UserId       int32     `json:"userId"`
	RefreshToken string    `json:"refreshToken"`
	UserAgent    string    `json:"userAgent"`
	ClientIp     string    `json:"clientIp"`
	IsBlocked    bool      `json:"isBlocked"`
	ExpiresAt    time.Time `json:"expiresAt"`
	CreatedAt    time.Time `json:"createdAt"`
}

type Country struct {
	ID               int       `json:"id"`
	Name             string    `json:"name" `
	CountryShortName string    `json:"country_short_name"`
	CreatedAt        time.Time `json:"createdAt" `
}

type State struct {
	ID        int       `json:"id"`
	CountryId int       `json:"countryId"`
	StateName string    `json:"stateName" `
	CreatedAt time.Time `json:"createdAt" `
}
