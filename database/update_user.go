package database

import (
	"errors"
)

func (db *DB) UpdateUser(id int, email, pass string) error {

	//get the chrips map from database
	databaseStruct, err := db.loadDB()
	if err != nil {

		return errors.New("error gettings database struct map")
	}

	//get the current user with the user id.
	curChirp, ok := databaseStruct.Chirps[id]
	if !ok {
		return errors.New("no user with that id ")
	}
	updtatedChrip := curChirp
	updtatedChrip.Email = email
	updtatedChrip.Password = pass
	databaseStruct.Chirps[id] = updtatedChrip
	db.writeDB(databaseStruct)
	return nil

}
