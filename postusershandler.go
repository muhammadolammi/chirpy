package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/muhammadolammi/chirpy/database"
	uservalidator "github.com/muhammadolammi/chirpy/user_validator"
	"golang.org/x/crypto/bcrypt"
)

func postUsersHandler(w http.ResponseWriter, r *http.Request) {
	type Parameters struct {
		Password string `json:"password"`
		Email    string `json:"email"`
	}
	type Responds struct {
		Id    int    `json:"id"`
		Email string `json:"email"`
	}
	decoder := json.NewDecoder(r.Body)
	params := Parameters{}
	err := decoder.Decode(&params)

	if err != nil {
		respondWithError(w, 500, "Something went wrong")
		return

	}

	//check pass strenght
	// passStrong, err := validatePass(params.Password)

	// if !passStrong {
	// 	respondWithError(w, 500, err.Error())
	// 	return

	// }
	encryptedPasss, err := bcrypt.GenerateFromPassword([]byte(params.Password), 10)
	if err != nil {
		respondWithError(w, 400, err.Error())
		return
	}

	reqEmail := params.Email

	db, err := database.NewUsersDB("database/users.json")
	if err != nil {
		log.Printf("error creating db. err: %v", err)
	}

	//check if email is already in database
	isEmail, err := uservalidator.CheckEmail(params.Email, db)
	if isEmail {
		respondWithError(w, 500, "Email already exists")
		return
	}
	if err != nil {
		respondWithError(w, 500, err.Error())
		return
	}

	user, err := db.CreateUser(string(reqEmail), string(encryptedPasss))
	if err != nil {
		respondWithError(w, 400, err.Error())
	}
	resBody := Responds{
		Id:    user.Id,
		Email: params.Email,
	}

	respondWithJSON(w, 201, resBody)

}
