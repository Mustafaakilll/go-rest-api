package service

import (
	"net/http"
	"src/github.com/mustafaakilll/rest_api/auth"
	"src/github.com/mustafaakilll/rest_api/database"
	"src/github.com/mustafaakilll/rest_api/types"

	"github.com/gin-gonic/gin"
)

type TokenRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type TokenService struct {
	storage database.DatabaseOperations
}

func NewTokenService(storage database.DatabaseOperations) *TokenService {
	return &TokenService{
		storage: storage,
	}
}

func (t TokenService) GenerateToken(ctx *gin.Context) {
	var tokenRequest TokenRequest

	if err := ctx.ShouldBind(&tokenRequest); err != nil {
		ctx.Abort()
		ctx.JSON(types.ErrorResponse(err.Error(), http.StatusBadRequest))
		return
	}

	user, err := t.storage.GetUserByEmail(tokenRequest.Email)
	if err != nil {

		ctx.Abort()
		ctx.JSON(types.ErrorResponse(err.Error(), http.StatusBadRequest))
		return
	}

	credentialError := user.CheckPassword(tokenRequest.Password)
	if credentialError != nil {
		ctx.Abort()
		ctx.JSON(types.ErrorResponse(err.Error(), http.StatusBadRequest))
		return
	}

	token, err := auth.GenerateJWT(user.Email, user.Username)
	if err != nil {
		ctx.Abort()
		ctx.JSON(types.ErrorResponse(err.Error(), http.StatusBadRequest))
		return
	}
	ctx.JSON(types.SuccessResponse(token))

}
