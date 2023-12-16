package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/muhammadolammi/chirpy/database"
)

func chirpyPostHandler(w http.ResponseWriter, r *http.Request) {
	type Parameters struct {
		Body string `json:"body"`
	}

	decoder := json.NewDecoder(r.Body)
	params := Parameters{}
	err := decoder.Decode(&params)

	if err != nil {
		respondWithError(w, 500, "Something went wrong")
		return

	}
	//lets manage when the decoded param body is less than 140
	if len(params.Body) > 140 {
		respondWithError(w, 400, "Chirp is too long")
		return

	}

	reqString := params.Body
	formattedreqString := formatString(reqString)
	db, err := database.NewDB("database/database.json")

	//TODO use db to create and get chips

	if err != nil {
		log.Printf("error creating db. err: %v", err)
	}

	resBody, err := db.CreateChirp(string(formattedreqString))
	if err != nil {
		respondWithError(w, 400, err.Error())
	}

	respondWithJSON(w, 201, resBody)

}

func chirpyGetHandler(w http.ResponseWriter, r *http.Request) {
	db, err := database.NewDB("database/database.json")

	//TODO use db to create and get chips

	if err != nil {

		respondWithError(w, 500, "Internal Server Error")
		return
	}
	chirps, err := db.GetChirps()
	if err != nil {
		respondWithError(w, 400, err.Error())
		return
	}
	respondWithJSON(w, 200, chirps)

}
