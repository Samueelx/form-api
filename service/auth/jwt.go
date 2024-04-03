package auth

import (
	"strconv"
	"time"

	"github.com/Samueelx/form-api/config"
	"github.com/golang-jwt/jwt/v5"
)

func CreateJWT(sectret []byte, userID int) (string, error) {
	expiration := time.Second * time.Duration(config.Envs.JWTExpirationInSeconds)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID":    strconv.Itoa(userID),
		"expiresAt": time.Now().Add(expiration).Unix(),
	})

	tokenString, err := token.SignedString(sectret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
