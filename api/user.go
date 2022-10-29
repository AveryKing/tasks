package api

import (
	"errors"
	db "github.com/AveryKing/tasks/db/sqlc"
	"github.com/gin-gonic/gin"
	"net/http"
)

type createUserRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6"`
}

func (server *Server) createUser(ctx *gin.Context) {
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	arg := db.CreateUserParams{
		Username: req.Username,
		Password: req.Password,
	}

	user, err := server.queries.CreateUser(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errors.New("There was an error creating your account"+err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func (server *Server) getUser(ctx *gin.Context) {

}
