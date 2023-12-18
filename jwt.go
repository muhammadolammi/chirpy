package main

import "github.com/golang-jwt/jwt/v5"

type CustomClaim struct {
	jwt.RegisteredClaims
}
