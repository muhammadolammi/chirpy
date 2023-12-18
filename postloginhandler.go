package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"

	"github.com/muhammadolammi/chirpy/database"
	uservalidator "github.com/muhammadolammi/chirpy/user_validator"
)

func (cfg *apiConfig) postLoginHandler(w http.ResponseWriter, r *http.Request) {

	//TODO implement func
	type Parameters struct {
		Password           string `json:"password"`
		Email              string `json:"email"`
		Expires_In_Seconds *int   `json:"expires_in_seconds"`
	}
	type Responds struct {
		Id           int    `json:"id"`
		Email        string `json:"email"`
		Token        string `json:"token"`
		RefreshToken string `json:"refresh_token"`
	}

	decoder := json.NewDecoder(r.Body)
	params := Parameters{}
	err := decoder.Decode(&params)

	if err != nil {
		respondWithError(w, 500, "Something went wrong")
		return

	}
	db, err := database.NewDB("database/database.json")
	if err != nil {
		respondWithError(w, 500, err.Error())
		return
	}

	isEmail, err := uservalidator.CheckEmail(params.Email, db)
	if !isEmail {
		respondWithError(w, 500, "no user with that mail")
		return
	}
	if err != nil {
		respondWithError(w, 500, err.Error())
		return

	}
	passValidator, err := uservalidator.ValidatePass(params.Password, db)
	if !passValidator {
		respondWithError(w, 401, "wrong password")
		return
	}
	if err != nil {
		respondWithError(w, 500, err.Error())
		return

	}

	user, err := db.GetUser(params.Email)
	if err != nil {
		respondWithError(w, 500, err.Error())
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, CustomClaim{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "chirpy-access",
			IssuedAt:  jwt.NewNumericDate(time.Now().UTC()),
			ExpiresAt: jwt.NewNumericDate(time.Now().UTC().Add(time.Hour)),
			Subject:   fmt.Sprintf("%v", user.Id),
		},
	})

	tokenString, err := token.SignedString([]byte(cfg.JWT_SECRET))

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, CustomClaim{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "chirpy-refresh",
			IssuedAt:  jwt.NewNumericDate(time.Now().UTC()),
			ExpiresAt: jwt.NewNumericDate(time.Now().UTC().Add(60 * 24 * time.Hour)),
			Subject:   fmt.Sprintf("%v", user.Id),
		},
	})

	refreshTokenString, err := refreshToken.SignedString([]byte(cfg.JWT_SECRET))

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	resBody := Responds{
		Id:           user.Id,
		Email:        user.Email,
		Token:        tokenString,
		RefreshToken: refreshTokenString,
	}

	respondWithJSON(w, 200, resBody)

}
