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
		Id    int    `json:"id"`
		Email string `json:"email"`
		Token string `json:"token"`
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
	exp := getExpiringTime(params.Expires_In_Seconds)
	expirationTime := time.Now().UTC().Add(time.Duration(exp) * time.Second)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, CustomClaim{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "chirpy",
			IssuedAt:  jwt.NewNumericDate(time.Now().UTC()),
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			Subject:   fmt.Sprintf("%v", user.Id),
		},
	})

	tokenString, err := token.SignedString([]byte(cfg.JWT_SECRET))

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	resBody := Responds{
		Id:    user.Id,
		Email: user.Email,
		Token: tokenString,
	}

	respondWithJSON(w, 200, resBody)

}

func getExpiringTime(exp *int) int {
	hour := int(time.Hour.Seconds())
	if exp == nil {
		return hour
	}
	if *exp <= 0 {
		return hour
	}
	if *exp > hour {
		return hour
	}
	return *exp
}
