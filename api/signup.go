package api

import (
	"github.com/amallick86/country-api/db"
	"github.com/amallick86/country-api/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

type createAccountRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6"`
}

type createAccountResponse struct {
	Id        int32     `json:"id"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"created_at" `
}

// CreateUser handles request for user creation
// @Summary add a new user
// @Tags User
// @ID CreateUser
// @Accept json
// @Produce json
// @Param data body createAccountRequest true "create user"
// @Success 201 {object} createAccountResponse
// @Failure 400 {object} Err
// @Failure 500 {object} Err
// @Router /user/create [post]
func (server *Server) createUser(ctx *gin.Context) {
	var req createAccountRequest
	var res createAccountResponse
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	arg := db.CreateUserParams{
		Username: req.Username,
		Password: hashedPassword,
	}
	user, err := server.store.CreateUser(ctx, arg)
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	res = createAccountResponse{
		Id:        user.ID,
		Username:  user.Username,
		CreatedAt: user.CreatedAt,
	}
	ctx.JSON(http.StatusCreated, res)
}
