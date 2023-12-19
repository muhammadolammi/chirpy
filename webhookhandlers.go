package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/muhammadolammi/chirpy/database"
)

func (cfg *apiConfig) webhookPostHandler(w http.ResponseWriter, r *http.Request) {
	apiToken, err := getTokenFromHeader(r, "apikey")
	if err != nil {
		respondWithError(w, 401, err.Error())
		return
	}
	if apiToken == "" {
		respondWithError(w, 401, "Invalid API token")
		return
	}
	if cfg.POLKA_API_KEY != apiToken {
		respondWithError(w, 401, "Invalid API token")
		return
	}
	type Parameters struct {
		Event string `json:"event"`
		Data  struct {
			UserId int `json:"user_id"`
		} `json:"data"`
	}
	params := Parameters{}
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	log.Println(params.Event)
	if params.Event != "user.upgraded" {

		respondWithJSON(w, 200, "")
		return
	}
	if params.Event == "user.upgraded" {

		userId := params.Data.UserId
		usersDb, err := database.NewUsersDB("database/users.json")
		if err != nil {
			respondWithError(w, http.StatusBadRequest, err.Error())
			return
		}
		user, err := usersDb.GetUserById(userId)
		if err != nil {
			respondWithError(w, 404, err.Error())
		}
		user.IsChirpyRed = true
		err = usersDb.UpdateUser(user)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, err.Error())
		}
		respondWithJSON(w, 200, "")
	}
}
