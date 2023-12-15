package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/muhammadolammi/chirpy/database"
)

type apiConfig struct {
	fileserverHits int
}
type Data struct {
	Body string `json:"body"`
}

func main() {
	db, err := database.NewDB("db.json")
	if err != nil {
		fmt.Println("Error creating database:", err)
		return
	}

	// body1 := Data{
	// 	Body: "hi good morning",
	// }
	// body2 := Data{
	// 	Body: "hi good afternoon",
	// }
	// jsonData1, _ := json.Marshal(body1)
	// jsonData2, _ := json.Marshal(body2)

	chirps, _ := db.GetChirps()
	// if err != nil {
	// 	fmt.Println("Error creating database:", err)
	// }
	fmt.Println(chirps)

	wait := true
	if wait {
		return
	}

	const port = "8080"
	cfg := apiConfig{}
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
	apiRouter.Get("/chirps", chirpyGetHandler)
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
