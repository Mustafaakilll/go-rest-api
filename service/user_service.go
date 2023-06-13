package service

import (
	"net/http"
	"src/github.com/mustafaakilll/rest_api/auth"
	"src/github.com/mustafaakilll/rest_api/database"
	"src/github.com/mustafaakilll/rest_api/types"

	"github.com/gin-gonic/gin"
)

type UserService struct {
	storage database.DatabaseOperations
}

func NewUserService(storage database.DatabaseOperations) *UserService {
	return &UserService{
		storage: storage,
	}
}

func (u UserService) HandleRegisterUser(ctx *gin.Context) {
	var user types.User
	if err := ctx.ShouldBind(&user); err != nil {
		ctx.Abort()
		ctx.JSON(types.ErrorResponse(err.Error(), http.StatusBadRequest))
		return
	}

	if err := user.HashPassword(user.Password); err != nil {
		ctx.Abort()
		ctx.JSON(types.ErrorResponse(err.Error(), http.StatusInternalServerError))
		return
	}

	err := u.storage.RegisterUser(&user)
	if err != nil {
		ctx.Abort()
		ctx.JSON(types.ErrorResponse(err.Error(), http.StatusInternalServerError))
		return
	}

	token, err := auth.GenerateJWT(user.Email, user.Username)
	if err != nil {
		ctx.Abort()
		ctx.JSON(types.ErrorResponse(err.Error(), http.StatusBadRequest))
		return
	}

	ctx.JSON(types.SuccessResponse(map[string]any{
		"user":  user,
		"token": token,
	}))
}

func (u UserService) HandleLoginUser(ctx *gin.Context) {
	var incomingUser types.User
	if err := ctx.ShouldBind(&incomingUser); err != nil {
		ctx.Abort()
		ctx.JSON(types.ErrorResponse(err.Error(), http.StatusBadRequest))
		return
	}

	user, err := u.storage.GetUserByEmail(incomingUser.Email)
	if err != nil {
		ctx.Abort()
		ctx.JSON(types.ErrorResponse(err.Error(), http.StatusInternalServerError))
		return
	}

	credentialError := user.CheckPassword(incomingUser.Password)
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

	ctx.JSON(types.SuccessResponse(map[string]any{
		"user":  user,
		"token": token,
	}))
}
