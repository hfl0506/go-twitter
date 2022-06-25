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

type AuthClaims struct {
	jwt.StandardClaims
	UserID string `json:userId"`
}

func GenerateToken(user *Payload, key []byte) (string, error) {
	expiresAt := time.Now().Add(24 * time.Hour).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodES512, AuthClaims{
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

func ValidateToken(tokenString string) (bool, error) {
	var claims AuthClaims
	_, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		return jwt.ParseRSAPrivateKeyFromPEM([]byte(claims.Subject))
	})

	if err != nil {
		return false, err
	}

	return true, nil
}

func GenerateAccessToken(user *Payload) (string, error) {
	return GenerateToken(user, accessSecretKey)
}

func GenerateRefreshToken(user *Payload) (string, error) {
	return GenerateToken(user, refreshSecretKey)
}
