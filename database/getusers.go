package database

import (
	"errors"
	"fmt"
)

// GetChirps returns all chirps in the database
func (db *DB) GetUsers() ([]User, error) {
	// db.mux.Lock()
	if db.path != "database/users.json" {
		return nil, errors.New("wrong directory to create chripys")
	}

	dbJson, err := db.loaUsers()
	if err != nil {
		return nil, fmt.Errorf("error loading database: %v", err)
	}

	users := make([]User, len(dbJson.Users))
	for id, user := range dbJson.Users {
		users[id-1] = user
	}

	return users, nil
}
