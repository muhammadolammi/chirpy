package database

import "fmt"

func (db *DB) CreateUser(email string, pass string) (User, error) {

	db.ensureDB()
	//load old database
	oldDb, err := db.loadDB()

	if err != nil {
		return User{}, fmt.Errorf("error getting old db to create new chip. err : %v", err)
	}
	//chech if old db is empty

	if oldDb.Chirps == nil {
		oldDb.Chirps = make(map[int]Chirp)
	}

	maxId := 0

	for id := range oldDb.Chirps {
		if id > maxId {
			maxId = id
		}
	}
	// create a new chirps with the id and body
	newChirp := Chirp{
		Id:       maxId + 1,
		Email:    email,
		Password: pass,
	}
	newUser := User{
		Id:    maxId + 1,
		Email: email,
	}

	// add the new chirp to the old chirps
	oldDb.Chirps[maxId+1] = newChirp

	// Write the updated chips back to the database
	err = db.writeDB(oldDb)
	if err != nil {
		return User{}, fmt.Errorf(err.Error())
	}

	return newUser, nil

}
