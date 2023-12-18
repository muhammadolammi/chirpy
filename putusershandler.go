package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
	"strings"

	jwt "github.com/golang-jwt/jwt/v5"
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
		respondWithError(w, 500, "Something went wrong")
		return

	}
	db, err := database.NewDB("database/database.json")
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	tokS, err := getTokenFromHeader(r)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	log.Println(tokS)
	claims := CustomClaim{}
	//pasrse into the claim with the token string
	_, err = jwt.ParseWithClaims(tokS, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(cfg.JWT_SECRET), nil
	})

	if err != nil {

		respondWithError(w, http.StatusUnauthorized, err.Error())
		return

	}
	//get the id
	userIdS := claims.Subject
	userid, err := strconv.Atoi(userIdS)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	// encrypt the password
	encryptedPasss, err := bcrypt.GenerateFromPassword([]byte(params.Password), 10)
	if err != nil {
		respondWithError(w, 400, err.Error())
		return
	}
	//lets update the user
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
