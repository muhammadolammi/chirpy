package database

import (
	"encoding/json"
	"fmt"
	"os"
)

func NewUsersDB(path string) (*DB, error) {
	db := &DB{
		path: path,
	}
	expectedPath := "database/users.json"
	if path != expectedPath {
		return db, fmt.Errorf("wrong directory to create users database, expected %s", expectedPath)
	}
	//make sure directory exist
	err := db.ensureDB()
	if err != nil {
		return db, fmt.Errorf(err.Error())
	}
	//check if the file exit,
	_, err = os.Stat(db.path)
	if os.IsNotExist(err) {
		//create new dile and an empty json structure if it doesnt exit
		file, err := os.Create(db.path)
		if err != nil {
			return nil, fmt.Errorf("error creating data file: %v", err)
		}
		defer file.Close()

		emptyDb := Userstructure{Users: make(map[int]User)}
		emptyDBjson, err := json.MarshalIndent(emptyDb, "", "    ")
		if err != nil {
			return nil, fmt.Errorf("error marshalling db to json: %v", err)
		}

		err = os.WriteFile(db.path, emptyDBjson, 0644)
		if err != nil {
			return nil, fmt.Errorf("error writing json db to file: %v", err)
		}
	}

	//it existed or was created return it

	return db, nil
}

func NewChirpsDB(path string) (*DB, error) {

	db := &DB{
		path: path,
	}
	expectedPath := "database/database.json"
	if path != expectedPath {
		return db, fmt.Errorf("wrong directory to create users database, expected %s", expectedPath)
	}
	//make sure directory exist
	err := db.ensureDB()
	if err != nil {
		return db, fmt.Errorf(err.Error())
	}
	//check if the file exit,
	_, err = os.Stat(db.path)
	if os.IsNotExist(err) {
		//create new dile and an empty json structure if it doesnt exit
		file, err := os.Create(db.path)
		if err != nil {
			return nil, fmt.Errorf("error creating data file: %v", err)
		}
		defer file.Close()

		emptyDb := Chirpstructure{Chirps: make(map[int]Chirp)}
		emptyDBjson, err := json.MarshalIndent(emptyDb, "", "    ")
		if err != nil {
			return nil, fmt.Errorf("error marshalling db to json: %v", err)
		}

		err = os.WriteFile(db.path, emptyDBjson, 0655)
		if err != nil {
			return nil, fmt.Errorf("error writing json db to file: %v", err)
		}
	}

	//it existed or was created return it

	return db, nil
}

func NewSessionsDB(path string) (*DB, error) {

	db := &DB{
		path: path,
	}
	expectedPath := "database/sessions.json"
	if path != expectedPath {
		return db, fmt.Errorf("wrong directory to create users database, expected %s", expectedPath)
	}
	//make sure directory exist
	err := db.ensureDB()
	if err != nil {
		return db, fmt.Errorf(err.Error())
	}
	//check if the file exit,
	_, err = os.Stat(db.path)
	if os.IsNotExist(err) {
		//create new dile and an empty json structure if it doesnt exit
		file, err := os.Create(db.path)
		if err != nil {
			return nil, fmt.Errorf("error creating data file: %v", err)
		}
		defer file.Close()

		emptyDb := RevokedSesssionstructure{RevokedSessions: make(map[string]RevokedSessionDetail)}
		emptyDBjson, err := json.MarshalIndent(emptyDb, "", "    ")
		if err != nil {
			return nil, fmt.Errorf("error marshalling db to json: %v", err)
		}

		err = os.WriteFile(db.path, emptyDBjson, 0655)
		if err != nil {
			return nil, fmt.Errorf("error writing json db to file: %v", err)
		}
	}

	//it existed or was created return it

	return db, nil
}
