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

func (db *DB) GetChirp(id int) (Chirp, error) {
	// db.mux.Lock()
	if db.path != "database/database.json" {
		return Chirp{}, errors.New("wrong directory to create chripys")
	}

	dbJson, err := db.loadChirps()
	if err != nil {
		return Chirp{}, fmt.Errorf("error loading database: %v", err)
	}
	chirp, ok := dbJson.Chirps[id]
	if !ok {
		return Chirp{}, errors.New("Chirp not in database")
	}

	return chirp, nil
}

func (db *DB) GetUserChirps(userId int) ([]Chirp, error) {
	// db.mux.Lock()
	if db.path != "database/database.json" {
		return nil, errors.New("wrong directory to create chripys")
	}

	dbJson, err := db.loadChirps()
	if err != nil {
		return nil, fmt.Errorf("error loading database: %v", err)
	}
	chirps := make([]Chirp, 0)
	for _, chirp := range dbJson.Chirps {
		if chirp.AuthorId == userId {
			chirps = append(chirps, chirp)
		}

	}

	return chirps, nil
}
func (db *DB) DeleteChirp(id int) error {
	// db.mux.Lock()
	if db.path != "database/database.json" {
		return errors.New("wrong directory to create chripys")
	}

	dbJson, err := db.loadChirps()
	if err != nil {
		return fmt.Errorf("error loading database: %v", err)
	}
	chirp, ok := dbJson.Chirps[id]
	if !ok {
		return errors.New("Chirp not in database")
	}
	updatedDbJson := dbJson
	delete(updatedDbJson.Chirps, chirp.Id)
	err = db.writeDB(updatedDbJson)
	if err != nil {
		return err
	}
	return nil
}
