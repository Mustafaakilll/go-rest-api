package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

var jwtSecret = []byte("MyVeryVerySecretKey")

type JWTClaims struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	jwt.StandardClaims
}

func GenerateJWT(email, username string) (string, error) {
	expireTime := time.Now().Add(1 * time.Hour)
	claims := &JWTClaims{
		Email:    email,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func ValidateJWT(jwtToken string) error {
	token, err := jwt.ParseWithClaims(
		jwtToken,
		&JWTClaims{},
		func(token *jwt.Token) (interface{}, error) { return []byte(jwtSecret), nil },
	)

	if err != nil {
		return err
	}

	claims, ok := token.Claims.(*JWTClaims)
	if !ok {
		return errors.New("invalid JWT claims")
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		return errors.New("token expired")
	}
	return err
}
