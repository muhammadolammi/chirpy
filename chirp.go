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

	_, err := database.NewDB("/database/database.json")
	//TODO use db to create and get chips

	if err != nil {
		log.Panicf("error creating db. err: %v", err)
	}

	decoder := json.NewDecoder(r.Body)
	params := Parameters{}
	err = decoder.Decode(&params)

	if err != nil {
		respondWithError(w, 500, "Something went wrong")
		// w.WriteHeader(500)
		// errBody := struct {
		// 	Error string `json:"error"`
		// }{
		// 	Error: "Something went wrong",
		// }
		// mashalledErr, _ := json.Marshal(errBody)
		// //write the mashalled error to the response
		// fmt.Fprint(w, string(mashalledErr))
		// return
	}
	//lets manage when the decoded param body is less than 140
	if len(params.Body) > 140 {
		respondWithError(w, 400, "Chirp is too long")
		// w.WriteHeader(400)
		// errBody := struct {
		// 	Error string `json:"error"`
		// }{
		// 	Error: "Chirp is too long",
		// }
		// mashalledErr, _ := json.Marshal(errBody)
		// //write the mashalled error to the response
		// fmt.Fprint(w, string(mashalledErr))
		// return
	}

	reqString := params.Body
	formattedreqString := formatString(reqString)

	resBody := struct {
		Id   int    `json:"id"`
		Body string `json:"body"`
	}{

		Body: formattedreqString,
	}
	respondWithJSON(w, 200, resBody)

}

func chirpyGetHandler(w http.ResponseWriter, r *http.Request) {

}
