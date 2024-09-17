package common

import (
	"time"

	"chris.com/models"
	"github.com/golang-jwt/jwt/v4"
)

var jwtSecret = []byte("secret_key")

type Cliams struct {
	ID uint
	jwt.StandardClaims
}

func ReleaseToken(user models.User) (string, error) {
	expirationTime := time.Now().Add(30 * 24 * time.Hour)
	claims := Cliams{
		ID: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "chris.com",
			Subject:   "user token",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func ParseToken(tokenString string) (*Cliams, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Cliams{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if token != nil {
		if cliams, ok := token.Claims.(*Cliams); ok && token.Valid {
			return cliams, nil
		}
	}
	return nil, err
}
