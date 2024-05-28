package auth

import (
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/golang-jwt/jwt/v5"
)

func createJWT(secret []byte, userID int) (string, error) {
	expiration := time.Second * time.Duration(config)
	
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID" : strconv.Itoa(userID),
		"expiredAt" : time.Now().Add(expiration).Unix(),
	})

	tokenString, err := toke
}