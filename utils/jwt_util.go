package utils

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var accessSecretKey = []byte(os.Getenv("ACCESS_SECRET_KEY"))
var refreshSecretKey = []byte(os.Getenv("REFRESH_SECRET_KEY"))

type Payload struct {
	ID   string `json:"userId"`
	Name string `json:"username"`
}

type authClaims struct {
	jwt.StandardClaims
	UserID string `json:userId"`
}

func generateToken(user *Payload, key []byte) (string, error) {
	expiresAt := time.Now().Add(24 * time.Hour).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodES512, authClaims{
		StandardClaims: jwt.StandardClaims{
			Subject:   user.Name,
			ExpiresAt: expiresAt,
		},
		UserID: user.ID,
	})

	tokenString, err := token.SignedString(key)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func generateAccessToken(user *Payload) (string, error) {
	return generateToken(user, accessSecretKey)
}

func generateRefreshToken(user *Payload) (string, error) {
	return generateToken(user, refreshSecretKey)
}
