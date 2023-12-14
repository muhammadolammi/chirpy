package main

import (
	"log"
	"net/http"
)

type apiConfig struct {
	fileserverHits int
}

func main() {
	const port = "8080"
	cfg := apiConfig{}
	mux := http.NewServeMux()
	appHandler := http.StripPrefix("/app/", http.FileServer(http.Dir(".")))
	mux.Handle("/app/", cfg.middlewareMetricsInc(appHandler))
	mux.HandleFunc("/healthz", readinessHandler)
	mux.HandleFunc("/metrics", cfg.getHitsHandler)
	mux.HandleFunc("/reset", cfg.resetHitsHandler)
	corsMux := middlewareCors(mux)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: corsMux,
	}

	log.Printf("Serving on port: %s\n", port)
	middlewareLog(srv.Handler)
	log.Fatal(srv.ListenAndServe())
}
