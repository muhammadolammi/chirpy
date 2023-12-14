package main

import (
	"fmt"
	"net/http"
)

func (cfg *apiConfig) resetHitsHandler(w http.ResponseWriter, r *http.Request) {
	cfg.fileserverHits = 0
	fmt.Fprintf(w, "Hits: %v	", cfg.fileserverHits)
}
