package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/muhammadolammi/chirpy/database"
	"golang.org/x/crypto/bcrypt"
)

func (cfg *apiConfig) putUserHandler(w http.ResponseWriter, r *http.Request) {
	type Parameters struct {
		Id       int    `json:"id"`
		Password string `json:"password"`
		Email    string `json:"email"`
	}
	type Responds struct {
		Id int `json:"id"`

		Email string `json:"email"`
	}

	decoder := json.NewDecoder(r.Body)
	params := Parameters{}
	err := decoder.Decode(&params)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Something went wrong")
		return

	}

	tokS, err := getTokenFromHeader(r)
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

	// claims := CustomClaim{}
	// //pasrse into the claim with the token string
	// _, err = jwt.ParseWithClaims(tokS, &claims, func(token *jwt.Token) (interface{}, error) {
	// 	return []byte(cfg.JWT_SECRET), nil
	// })

	// if err != nil {
	// 	log.Printf("Error parsing token: %v", err)
	// 	respondWithError(w, http.StatusUnauthorized, err.Error())
	// 	return
	// }

	// //get the id
	// userIdS := claims.Subject

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
	err = db.UpdateUser(userid, params.Email, string(encryptedPasss))
	if err != nil {
		respondWithError(w, 400, err.Error())
		return
	}

	resBody := Responds{
		Id:    userid,
		Email: params.Email,
	}

	respondWithJSON(w, 200, resBody)

}
