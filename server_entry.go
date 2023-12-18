package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func serverEnrty(cfg *apiConfig) {
	const port = "8080"

	// mux := http.NewServeMux()
	mainRouter := chi.NewRouter()
	apiRouter := chi.NewRouter()
	adminRouter := chi.NewRouter()
	appHandler := cfg.middlewareMetricsInc(http.StripPrefix("/app", http.FileServer(http.Dir("."))))
	mainRouter.Handle("/app/*", appHandler)
	mainRouter.Handle("/app", appHandler)

	apiRouter.Get("/healthz", readinessHandler)
	apiRouter.HandleFunc("/reset", cfg.resetHitsHandler)
	apiRouter.Post("/chirps", chirpyPostHandler)
	apiRouter.Get("/chirps", chirpysGetHandler)
	apiRouter.Get("/chirps/{chirpID:[0-9]+}", chirpGetHandlerWId)
	apiRouter.Post("/users", postUsersHandler)
	apiRouter.Put("/users", cfg.putUserHandler)
	apiRouter.Post("/login", cfg.postLoginHandler)
	mainRouter.Mount("/api", apiRouter)
	// Mount the apiRouter at the root path

	adminRouter.Get("/metrics", cfg.getHitsHandler)
	// corsMux := middlewareCors(mux)
	// Mount the mainRouter at /api

	mainRouter.Mount("/admin", adminRouter)

	corsRouter := middlewareCors(mainRouter)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: corsRouter,
	}

	log.Printf("Serving on port: %s\n", port)

	log.Fatal(srv.ListenAndServe())
}
