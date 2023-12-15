package database

import (
	"encoding/json"
	"fmt"
	"os"
)

// GetChirps returns all chirps in the database
func (db *DB) GetChirps() ([]Chirp, error) {
	db.mux.Lock()
	defer db.mux.Unlock()
	dbByte, err := os.ReadFile(db.path)
	if err != nil {
		return []Chirp{}, fmt.Errorf("error reading db file err: %v", err)
	}

	dbJson := DBStructure{}
	err = json.Unmarshal(dbByte, &dbJson)
	if err != nil {
		return []Chirp{}, fmt.Errorf("error unmashalling db byte. err: %v", err)
	}
	chirps := make([]Chirp, len(dbJson.Chirps))
	for id, chirp := range dbJson.Chirps {
		chirps[id-1] = chirp
	}

	return chirps, nil
}
