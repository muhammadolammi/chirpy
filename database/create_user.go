package database

import (
	"errors"
	"fmt"
)

func (db *DB) CreateUser(email, encryptedPass string) (User, error) {

	db.ensureDB()

	//make sure its the user database
	if db.path != "database/users.json" {
		return User{}, errors.New("this is not the chrips directory")
	}
	//load old database
	oldUsers, err := db.loaUsers()

	if err != nil {
		return User{}, fmt.Errorf("error getting old users to new new chip. err : %v", err)
	}
	//chech if old db is empty

	if oldUsers.Users == nil {
		oldUsers.Users = make(map[int]User)
	}

	maxId := 0

	for id := range oldUsers.Users {
		if id > maxId {
			maxId = id
		}
	}
	// create a new chirps with the id and body

	newUser := User{
		Id:       maxId + 1,
		Email:    email,
		Password: encryptedPass,
	}

	// add the new chirp to the old chirps
	oldUsers.Users[maxId+1] = newUser

	// Write the updated chips back to the database
	err = db.writeDB(oldUsers)
	if err != nil {
		return User{}, fmt.Errorf(err.Error())
	}

	return newUser, nil

}
