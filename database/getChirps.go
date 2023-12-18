package database

import (
	"errors"
	"fmt"
)

// GetChirps returns all chirps in the database
func (db *DB) GetChirps() ([]Chirp, error) {
	// db.mux.Lock()
	if db.path != "database/database.json" {
		return nil, errors.New("wrong directory to create chripys")
	}

	dbJson, err := db.loadChirps()
	if err != nil {
		return nil, fmt.Errorf("error loading database: %v", err)
	}

	chirps := make([]Chirp, len(dbJson.Chirps))
	for id, chirp := range dbJson.Chirps {
		chirps[id-1] = chirp
	}

	return chirps, nil
}
