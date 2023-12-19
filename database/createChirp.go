package database

import (
	"errors"
	"fmt"
)

// CreateChirp creates a new chirp and saves it to disk

func (db *DB) CreateChirp(body string, userId int) (Chirp, error) {

	db.ensureDB()
	//make sure its chirps database
	if db.path != "database/database.json" {
		return Chirp{}, errors.New("this is not the chrips directory")
	}

	//load old database
	oldChirps, err := db.loadChirps()

	if err != nil {
		return Chirp{}, fmt.Errorf("error getting old db to create new chip. err : %v", err)
	}
	//chech if old cgirps  is empty

	if oldChirps.Chirps == nil {
		oldChirps.Chirps = make(map[int]Chirp)
	}

	maxId := 0

	for id := range oldChirps.Chirps {
		if id > maxId {
			maxId = id
		}
	}
	// create a new chirps with the id and body
	newChirp := Chirp{

		AuthorId: userId,
		Body:     body,
		Id:       maxId + 1,
	}

	// add the new chirp to the old chirps
	oldChirps.Chirps[maxId+1] = newChirp

	// Write the updated chips back to the database
	err = db.writeDB(oldChirps)
	if err != nil {
		return Chirp{}, fmt.Errorf(err.Error())
	}

	return newChirp, nil

}
