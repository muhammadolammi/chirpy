package main

import (
	"net/http"
	"strconv"

	"github.com/muhammadolammi/chirpy/database"
)

func (cfg *apiConfig) postRefreshHandler(w http.ResponseWriter, r *http.Request) {

	refreshTokenString, err := getTokenFromHeader(r)
	if err != nil {
		respondWithError(w, 401, err.Error())
		return
	}
	userIdString, err := ValidateJWT(refreshTokenString, cfg.JWT_SECRET, "user-refresh")
	if err != nil {
		respondWithError(w, 401, err.Error())
		return
	}
	userid, err := strconv.Atoi(userIdString)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	db, err := database.NewSessionsDB("database/sessions.json")
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	//todo make sure to solve this commendted out issession
	isSession, err := db.IsSession(refreshTokenString)
	if err != nil {
		respondWithError(w, 401, err.Error())
		return
	}

	//manage if no session ie session not revoked
	if isSession {
		respondWithError(w, 401, "session revoked")
		return
	}
	token, err := createToken(cfg.JWT_SECRET, "user-access", userid)
	if err != nil {
		respondWithError(w, 401, err.Error())
		return
	}
	resBody := struct {
		Token string `json:"token"`
	}{
		Token: token,
	}
	respondWithJSON(w, 200, resBody)
}
