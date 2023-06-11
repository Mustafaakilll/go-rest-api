package service

import (
	"net/http"
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

	ctx.JSON(http.StatusCreated, user)
}
