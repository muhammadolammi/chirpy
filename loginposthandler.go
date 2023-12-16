package main

import (
	"encoding/json"
	"net/http"

	"github.com/muhammadolammi/chirpy/database"
	uservalidator "github.com/muhammadolammi/chirpy/user_validator"
)

func loginPostHandler(w http.ResponseWriter, r *http.Request) {

	//TODO implement func
	type Parameters struct {
		Password string `json:"password"`
		Email    string `json:"email"`
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

	resBody, err := db.GetUser(params.Email)
	if err != nil {
		respondWithError(w, 500, err.Error())
	}
	respondWithJSON(w, 200, resBody)

}
