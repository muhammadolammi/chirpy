package database

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func (db *DB) loadDB() (DBStructure, error) {
	db.mux.Lock()
	defer db.mux.Unlock()
	err := db.ensureDB()
	if err != nil {
		log.Print(err.Error())
	}

	dbByte, _ := os.ReadFile(db.path)

	dbJson := DBStructure{}
	err = json.Unmarshal(dbByte, &dbJson)
	if err != nil {
		return DBStructure{}, fmt.Errorf("error unmarshalling db byte. err: %v", err)
	}
	return dbJson, nil
}
