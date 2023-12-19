package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/muhammadolammi/chirpy/database"
	uservalidator "github.com/muhammadolammi/chirpy/user_validator"
	"golang.org/x/crypto/bcrypt"
)

func (cfg *apiConfig) postUsersHandler(w http.ResponseWriter, r *http.Request) {
	type Parameters struct {
		Password string `json:"password"`
		Email    string `json:"email"`
	}
	type Responds struct {
		Id          int    `json:"id"`
		Email       string `json:"email"`
		IsChirpyRed bool   `json:"is_chirpy_red"`
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

	user, err := db.CreateUser(string(reqEmail), string(encryptedPasss), false)
	if err != nil {
		respondWithError(w, 400, err.Error())
	}
	resBody := Responds{
		Id:    user.Id,
		Email: params.Email,
	}

	respondWithJSON(w, 201, resBody)

}

func (cfg *apiConfig) putUserHandler(w http.ResponseWriter, r *http.Request) {
	type Parameters struct {
		Id       int    `json:"id"`
		Password string `json:"password"`
		Email    string `json:"email"`
	}
	type Responds struct {
		Id int `json:"id"`

		Email       string `json:"email"`
		IsChirpyRed bool   `json:"is_chirpy_red"`
	}

	decoder := json.NewDecoder(r.Body)
	params := Parameters{}
	err := decoder.Decode(&params)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Something went wrong")
		return

	}

	tokS, err := getTokenFromHeader(r, "bearer")
	if err != nil {
		respondWithError(w, 401, err.Error())
		return
	}
	userIdString, err := ValidateJWT(tokS, cfg.JWT_SECRET, "user-access")
	if err != nil {
		respondWithError(w, 401, err.Error())
		return
	}

	log.Println(tokS)

	//parse string id to int
	userid, err := strconv.Atoi(userIdString)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// encrypt the password
	encryptedPasss, err := bcrypt.GenerateFromPassword([]byte(params.Password), 10)
	if err != nil {
		respondWithError(w, 400, err.Error())
		return
	} //
	//lets update the user
	db, err := database.NewUsersDB("database/users.json")
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	user, err := db.GetUserById(userid)
	if err != nil {
		respondWithError(w, 400, err.Error())
		return
	}
	user.Password = string(encryptedPasss)
	err = db.UpdateUser(user)
	if err != nil {
		respondWithError(w, 400, err.Error())
		return
	}

	resBody := Responds{
		Id:          user.Id,
		Email:       user.Email,
		IsChirpyRed: user.IsChirpyRed,
	}

	respondWithJSON(w, 200, resBody)

}

func getUsersHandler(w http.ResponseWriter, r *http.Request) {
	//TODO implement func

}
