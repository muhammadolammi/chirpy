package database

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
)

func (db *DB) loadSessions() (RevokedSesssionstructure, error) {
	//make sure its chirps database
	if db.path != "database/sessions.json" {
		return RevokedSesssionstructure{}, errors.New("this is not the chrips directory")
	}
	db.mux.Lock()
	defer db.mux.Unlock()
	err := db.ensureDB()
	if err != nil {
		log.Print(err.Error())
	}

	dbByte, _ := os.ReadFile(db.path)

	dbJson := RevokedSesssionstructure{}
	err = json.Unmarshal(dbByte, &dbJson)
	if err != nil {
		return RevokedSesssionstructure{}, fmt.Errorf("error unmarshalling db byte. err: %v", err)
	}
	return dbJson, nil
}

func (db *DB) loadChirps() (Chirpstructure, error) {
	//make sure its chirps database
	if db.path != "database/database.json" {
		return Chirpstructure{}, errors.New("this is not the chrips directory")
	}
	db.mux.Lock()
	defer db.mux.Unlock()
	err := db.ensureDB()
	if err != nil {
		log.Print(err.Error())
	}

	dbByte, _ := os.ReadFile(db.path)

	dbJson := Chirpstructure{}
	err = json.Unmarshal(dbByte, &dbJson)
	if err != nil {
		return Chirpstructure{}, fmt.Errorf("error unmarshalling db byte. err: %v", err)
	}
	return dbJson, nil
}

func (db *DB) loaUsers() (Userstructure, error) {
	//make sure its the user database
	if db.path != "database/users.json" {
		return Userstructure{}, errors.New("this is not the chrips directory")
	}
	db.mux.Lock()
	defer db.mux.Unlock()
	err := db.ensureDB()
	if err != nil {
		log.Print(err.Error())
	}

	dbByte, _ := os.ReadFile(db.path)

	dbJson := Userstructure{}
	err = json.Unmarshal(dbByte, &dbJson)
	if err != nil {
		return Userstructure{}, fmt.Errorf("error unmarshalling db byte. err: %v", err)
	}
	return dbJson, nil
}
