package main

import (
	"net/http"

	"github.com/muhammadolammi/chirpy/database"
)

func (cfg *apiConfig) postRevokeHandler(w http.ResponseWriter, r *http.Request) {

	refreshTokenString, err := getTokenFromHeader(r)
	if err != nil {
		respondWithError(w, 401, err.Error())
		return
	}

	db, err := database.NewSessionsDB("database/sessions.json")
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	//todo make sure to solve this commendted out issession
	err = db.UpdateSession(refreshTokenString)
	respondWithJSON(w, 200, "")
}
