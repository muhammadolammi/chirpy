package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func chirpyValidateHandler(w http.ResponseWriter, r *http.Request) {
	type Parameters struct {
		Body string `json:"body"`
	}

	decoder := json.NewDecoder(r.Body)
	params := Parameters{}
	err := decoder.Decode(&params)

	if err != nil {
		w.WriteHeader(500)
		errBody := struct {
			Error string `json:"error"`
		}{
			Error: "Something went wrong",
		}
		mashalledErr, _ := json.Marshal(errBody)
		//write the mashalled error to the response
		fmt.Fprint(w, string(mashalledErr))
		return
	}
	//lets manage when the decoded param body is less than 140
	if len(params.Body) > 140 {
		w.WriteHeader(400)
		errBody := struct {
			Error string `json:"error"`
		}{
			Error: "Chirp is too long",
		}
		mashalledErr, _ := json.Marshal(errBody)
		//write the mashalled error to the response
		fmt.Fprint(w, string(mashalledErr))
		return
	}

	w.WriteHeader(200)
	resBody := struct {
		Valid bool `json:"valid"`
	}{
		Valid: true,
	}
	mashalledRes, _ := json.Marshal(resBody)
	//write the mashalled error to the response
	fmt.Fprint(w, string(mashalledRes))
}
