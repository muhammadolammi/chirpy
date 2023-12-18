package main

import (
	"html/template"
	"log"
	"os"

	"net/http"
)

func (cfg *apiConfig) getHitsHandler(w http.ResponseWriter, r *http.Request) {
	htmlTemplate, err := os.ReadFile("metrics.html")
	if err != nil {
		log.Fatal(err)
	}

	// Dummy variable for demonstration

	// Parse the HTML template
	tmpl, err := template.New("index").Parse(string(htmlTemplate))
	if err != nil {
		log.Fatal(err)
	}
	variables := PageVariables{Hits: cfg.FileserverHits}
	err = tmpl.Execute(w, variables)
	if err != nil {
		log.Fatal(err)
	}
}

type PageVariables struct {
	Hits int
}
