package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
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

func chirpysGetHandler(w http.ResponseWriter, r *http.Request) {
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

func chirpGetHandlerWId(w http.ResponseWriter, r *http.Request) {
	// Extract the chirpID from the URL using chi
	chirpIDStr := chi.URLParam(r, "chirpID")

	// Remove the "id:" prefix if it exists
	chirpIDStr = strings.TrimPrefix(chirpIDStr, "id:")

	// Convert chirpID from string to int
	var chirpID int
	if _, err := fmt.Sscanf(chirpIDStr, "%d", &chirpID); err != nil {

		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	db, err := database.NewDB("database/database.json")

	if err != nil {

		respondWithError(w, 500, "Internal Server Error")
		return
	}
	chirps, err := db.GetChirps()
	if err != nil {
		respondWithError(w, 400, err.Error())
		return
	}

	if chirpID > len(chirps) {

		respondWithError(w, 404, "doesnt exit")
		return
	}
	respondWithJSON(w, 200, chirps[chirpID-1])

}
