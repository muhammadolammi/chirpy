package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func createToken(secret, issuer string, id int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, CustomClaim{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    issuer,
			IssuedAt:  jwt.NewNumericDate(time.Now().UTC()),
			ExpiresAt: jwt.NewNumericDate(time.Now().UTC().Add(time.Hour)),
			Subject:   fmt.Sprintf("%v", id),
		},
	})

	tokenString, err := token.SignedString([]byte(secret))

	if err != nil {

		return "", errors.New("error creating token string")
	}

	return tokenString, nil
}
