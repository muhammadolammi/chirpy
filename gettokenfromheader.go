package main

import (
	"errors"
	"net/http"
	"strings"
)

func getTokenFromHeader(r *http.Request) (string, error) {
	// Get the "Authorization" header
	bearerToken := r.Header.Get("Authorization")

	// Check if the header is empty
	if bearerToken == "" {
		return "", nil // No token present, not an error
	}

	// Split the header value to get the token part
	tokenParts := strings.Split(bearerToken, " ")
	if len(tokenParts) != 2 || strings.ToLower(tokenParts[0]) != "bearer" {
		return "", errors.New("invalid Authorization header format")
	}

	// Return the token part
	return tokenParts[1], nil
}
