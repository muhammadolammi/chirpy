package database

import (
	"errors"
)

func (db *DB) UpdateUser(id int, email, pass string) error {

	//make sure its the user database
	if db.path != "database/users.json" {
		return errors.New("this is not the chrips directory")
	}
	//get the chrips map from database
	databaseStruct, err := db.loaUsers()
	if err != nil {

		return errors.New("error gettings database struct map")
	}

	//get the current user with the user id.
	curChirp, ok := databaseStruct.Users[id]
	if !ok {
		return errors.New("no user with that id ")
	}
	updtatedChrip := curChirp
	updtatedChrip.Email = email
	updtatedChrip.Password = pass
	databaseStruct.Users[id] = updtatedChrip
	db.writeDB(databaseStruct)
	return nil

}
