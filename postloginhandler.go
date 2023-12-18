package main

import (
	"encoding/json"
	"net/http"

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
	db, err := database.NewUsersDB("database/users.json")
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
	passValidator, err := uservalidator.ValidatePass(params.Password, params.Email, db)
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

	tokenString, err := createToken(cfg.JWT_SECRET, "user-access", user.Id)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	refreshTokenString, err := createToken(cfg.JWT_SECRET, "user-refresh", user.Id)

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
