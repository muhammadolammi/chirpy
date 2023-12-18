package main

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

// ValidateJWT -
func ValidateJWT(tokenString, tokenSecret, issuerS string) (string, error) {
	claimsStruct := jwt.RegisteredClaims{}
	token, err := jwt.ParseWithClaims(
		tokenString,
		&claimsStruct,
		func(token *jwt.Token) (interface{}, error) { return []byte(tokenSecret), nil },
	)
	if err != nil {
		return "", err
	}
	issuer, err := claimsStruct.GetIssuer()
	if err != nil {

		return "", err

	}
	if issuer != issuerS {

		return "", fmt.Errorf("%v token not allowed", issuer)

	}

	userIDString, err := token.Claims.GetSubject()
	if err != nil {
		return "", err
	}

	return userIDString, nil
}
