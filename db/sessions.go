package db

import (
	"context"
	"github.com/amallick86/country-api/models"
	"github.com/google/uuid"
	"time"
)

// save user session in database
const createSession = `
INSERT INTO sessions (
  id,
 user_id,
 refresh_token,
 user_agent,
 client_ip,
 is_blocked,
 expires_at 
) VALUES (
  $1, $2, $3, $4, $5, $6, $7
)
RETURNING  id, user_id, refresh_token, user_agent, client_ip, is_blocked, expires_at, created_at
`

type CreateSessionParams struct {
	ID           uuid.UUID `json:"id"`
	UserId       int32     `json:"userId"`
	RefreshToken string    `json:"refreshToken"`
	UserAgent    string    `json:"userAgent"`
	ClientIp     string    `json:"clientIp"`
	IsBlocked    bool      `json:"isBlocked"`
	ExpiresAt    time.Time `json:"expiresAt"`
}

func (q *Queries) CreateSession(ctx context.Context, arg CreateSessionParams) (models.Sessions, error) {
	row := q.db.QueryRowContext(ctx, createSession,
		arg.ID,
		arg.UserId,
		arg.RefreshToken,
		arg.UserAgent,
		arg.ClientIp,
		arg.IsBlocked,
		arg.ExpiresAt,
	)
	var i models.Sessions
	err := row.Scan(
		&i.ID,
		&i.UserId,
		&i.RefreshToken,
		&i.UserAgent,
		&i.ClientIp,
		&i.IsBlocked,
		&i.ExpiresAt,
		&i.CreatedAt,
	)
	return i, err
}

// Get user session

const getSession = `
SELECT  id, user_id, refresh_token, user_agent, client_ip, is_blocked, expires_at, created_at FROM sessions
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetSession(ctx context.Context, id uuid.UUID) (models.Sessions, error) {
	row := q.db.QueryRowContext(ctx, getSession, id)
	var i models.Sessions
	err := row.Scan(
		&i.ID,
		&i.UserId,
		&i.RefreshToken,
		&i.UserAgent,
		&i.ClientIp,
		&i.IsBlocked,
		&i.ExpiresAt,
		&i.CreatedAt,
	)
	return i, err
}
