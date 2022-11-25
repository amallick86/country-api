package api

import (
	"database/sql"
	"github.com/amallick86/country-api/db"
	"github.com/amallick86/country-api/models"
	"github.com/amallick86/country-api/util"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"time"
)

type loginRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6"`
}

func newUserResponse(user models.User) createAccountResponse {
	return createAccountResponse{
		Id:        user.ID,
		Username:  user.Username,
		CreatedAt: user.CreatedAt,
	}
}

type loginResponse struct {
	SessionId             uuid.UUID             `json:"session_id"`
	AccessToken           string                `json:"accessToken"`
	AccessTokenExpiresAt  time.Time             `json:"accessTokenExpiresAt"`
	RefreshToken          string                `json:"refreshToken"`
	RefreshTokenExpiresAt time.Time             `json:"refreshTokenExpiresAt"`
	User                  createAccountResponse `json:"user"`
}

// Login handles request for user creation
// @Summary login and generate token with JWT
// @Tags User
// @ID Login
// @Accept json
// @Produce json
// @Param data body loginRequest true "Login request"
// @Success 200 {object} loginResponse
// @Failure 400 {object} Err
// @Failure 500 {object} Err
// @Router /user/login [post]
func (server *Server) login(ctx *gin.Context) {
	var req loginRequest
	var res loginResponse
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	user, err := server.store.GetUser(ctx, req.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	err = util.CheckPassword(req.Password, user.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}
	accessToken, acessPayload, err := server.tokenMaker.CreateToken(
		user.Username,
		server.config.AccessTokenDuration)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	refreshToken, refreshPayload, err := server.tokenMaker.CreateToken(
		user.Username,
		server.config.RefreshTokenDuration)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	session, err := server.store.CreateSession(ctx, db.CreateSessionParams{
		ID:           refreshPayload.ID,
		UserId:       user.ID,
		RefreshToken: refreshToken,
		UserAgent:    ctx.Request.UserAgent(),
		ClientIp:     ctx.ClientIP(),
		IsBlocked:    false,
		ExpiresAt:    refreshPayload.ExpiredAt,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	res = loginResponse{
		SessionId:             session.ID,
		AccessToken:           accessToken,
		AccessTokenExpiresAt:  acessPayload.ExpiredAt,
		RefreshToken:          refreshToken,
		RefreshTokenExpiresAt: refreshPayload.ExpiredAt,
		User:                  newUserResponse(user),
	}
	ctx.JSON(http.StatusOK, res)
}
