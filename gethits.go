package main

import (
	"fmt"
	"net/http"
)

func (cfg *apiConfig) getHitsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hits: %v	", cfg.fileserverHits)
}
