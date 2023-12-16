package database

import (
	"fmt"
	"log"
)

// GetChirps returns all chirps in the database
func (db *DB) GetChirps() ([]Chirp, error) {
	// db.mux.Lock()
	log.Println("worrlsss")
	dbJson, err := db.loadDB()
	if err != nil {
		return nil, fmt.Errorf("error loading database: %v", err)
	}

	chirps := make([]Chirp, len(dbJson.Chirps))
	for id, chirp := range dbJson.Chirps {
		chirps[id-1] = chirp
	}

	return chirps, nil
}
