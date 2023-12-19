package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/muhammadolammi/chirpy/database"
	uservalidator "github.com/muhammadolammi/chirpy/user_validator"
)

func (cfg *apiConfig) chirpyPostHandler(w http.ResponseWriter, r *http.Request) {
	tokenString, err := getTokenFromHeader(r, "bearer")
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, err.Error())
		return
	}
	userIdString, err := ValidateJWT(tokenString, cfg.JWT_SECRET, "user-access")
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, err.Error())
		return
	}
	userId, err := strconv.Atoi(userIdString)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	isUserLoggedIn, err := uservalidator.IsUserLoggedIn(userId)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	if !isUserLoggedIn {
		respondWithError(w, http.StatusUnauthorized, err.Error())
		return
	}

	type Parameters struct {
		Body string `json:"body"`
	}

	decoder := json.NewDecoder(r.Body)
	params := Parameters{}
	err = decoder.Decode(&params)

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
	db, err := database.NewChirpsDB("database/database.json")

	//TODO use db to create and get chips

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	resBody, err := db.CreateChirp(string(formattedreqString), userId)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, 201, resBody)

}

// func (cfg *apiConfig) chirpysGetHandler(w http.ResponseWriter, r *http.Request) {
// 	db, err := database.NewChirpsDB("database/database.json")

// 	//TODO use db to create and get chips

// 	if err != nil {

// 		respondWithError(w, 500, "Internal Server Error")
// 		return
// 	}
// 	idS := r.URL.Query().Get("author_id")
// 	// s is a string that contains the value of the author_id query parameter
// 	// if it exists, or an empty string if it doesn't
// 	//if id is empty return all chirp
// 	if idS == "" {
// 		chirps, err := db.GetChirps()
// 		if err != nil {
// 			respondWithError(w, 400, err.Error())
// 			return
// 		}

// 		respondWithJSON(w, 200, chirps)
// 		return
// 	}
// 	//if id is available return that user chirps
// 	userId, err := strconv.Atoi(idS)
// 	if err != nil {
// 		respondWithError(w, 500, err.Error())
// 		return
// 	}
// 	chirps, err := db.GetUserChirps(userId)
// 	if err != nil {
// 		respondWithError(w, 400, err.Error())
// 		return
// 	}
// 	respondWithJSON(w, 200, chirps)

// }
func (cfg *apiConfig) chirpysGetHandler(w http.ResponseWriter, r *http.Request) {
	db, err := database.NewChirpsDB("database/database.json")

	if err != nil {
		respondWithError(w, 500, "Internal Server Error")
		return
	}

	idS := r.URL.Query().Get("author_id")
	order := r.URL.Query().Get("sort")

	// If id is empty, return all chirps
	if idS == "" {
		chirps, err := db.GetChirps()
		if err != nil {
			respondWithError(w, 400, err.Error())
			return
		}

		// Check the order parameter and sort chirps based on IDs
		sort.Slice(chirps, func(i, j int) bool {
			if order == "desc" {
				return chirps[i].Id > chirps[j].Id
			}
			return chirps[i].Id < chirps[j].Id
		})

		respondWithJSON(w, 200, chirps)
		return
	}

	// If id is available, return that user's chirps
	userId, err := strconv.Atoi(idS)
	if err != nil {
		respondWithError(w, 500, err.Error())
		return
	}

	chirps, err := db.GetUserChirps(userId)
	if err != nil {
		respondWithError(w, 400, err.Error())
		return
	}

	// Check the order parameter and sort chirps based on IDs
	sort.Slice(chirps, func(i, j int) bool {
		if order == "desc" {
			return chirps[i].Id > chirps[j].Id
		}
		return chirps[i].Id < chirps[j].Id
	})

	respondWithJSON(w, 200, chirps)
}

func (cfg *apiConfig) chirpGetHandlerWId(w http.ResponseWriter, r *http.Request) {
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

	db, err := database.NewChirpsDB("database/database.json")

	if err != nil {

		respondWithError(w, 500, "Internal Server Error")
		return
	}
	chirp, err := db.GetChirp(chirpID)
	if err != nil {
		respondWithError(w, 400, err.Error())
		return
	}

	respondWithJSON(w, 200, chirp)

}

func (cfg *apiConfig) chirpyDeleteHandler(w http.ResponseWriter, r *http.Request) {

	tokenString, err := getTokenFromHeader(r, "bearer")
	if err != nil {
		respondWithError(w, 403, err.Error())
		return
	}
	userIdString, err := ValidateJWT(tokenString, cfg.JWT_SECRET, "user-access")
	if err != nil {
		respondWithError(w, 403, err.Error())
		return
	}
	userId, err := strconv.Atoi(userIdString)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	isUserLoggedIn, err := uservalidator.IsUserLoggedIn(userId)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	if !isUserLoggedIn {
		respondWithError(w, 403, err.Error())
		return
	}

	//get the chirp id
	chirpIDStr := chi.URLParam(r, "chirpID")

	// Remove the "id:" prefix if it exists
	chirpIDStr = strings.TrimPrefix(chirpIDStr, "id:")

	// Convert chirpID from string to int
	var chirpID int
	if _, err := fmt.Sscanf(chirpIDStr, "%d", &chirpID); err != nil {

		respondWithError(w, 403, err.Error())
		return
	}

	chirpsdb, err := database.NewChirpsDB("database/database.json")
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	chirp, err := chirpsdb.GetChirp(chirpID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	if chirp.AuthorId != userId {
		respondWithError(w, 403, "trying to edit another user post")
		return
	}
	err = chirpsdb.DeleteChirp(chirpID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, 200, "")

}
