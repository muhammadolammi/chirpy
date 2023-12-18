package database

import (
	"fmt"
)

// GetChirps returns all chirps in the database
func (db *DB) GetChirps() ([]Chirp, error) {
	// db.mux.Lock()

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

func (db *DB) GetChirpsMap() (map[int]Chirp, error) {
	// db.mux.Lock()

	dbJson, err := db.loadDB()
	if err != nil {
		return nil, fmt.Errorf("error loading database: %v", err)
	}

	return dbJson.Chirps, nil
}
