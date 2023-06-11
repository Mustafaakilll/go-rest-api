package middleware

import (
	"net/http"
	"src/github.com/mustafaakilll/rest_api/auth"
	"src/github.com/mustafaakilll/rest_api/types"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(types.ErrorResponse("Authorization header is missing", http.StatusUnauthorized))
			c.Abort()
			return
		}
		err := auth.ValidateJWT(tokenString)
		if err != nil {
			c.JSON(types.ErrorResponse(err.Error(), http.StatusUnauthorized))
			c.Abort()
			return
		}
		c.Next()
	}
}
