package database

import (
	"errors"
)

func (db *DB) UpdateUser(user User) error {

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
	_, ok := databaseStruct.Users[user.Id]
	if !ok {
		return errors.New("no user with that id ")
	}

	databaseStruct.Users[user.Id] = user
	db.writeDB(databaseStruct)
	return nil

}
