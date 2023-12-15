package database

import (
	"encoding/json"
	"fmt"
	"os"
)

func NewDB(path string) (*DB, error) {
	db := &DB{
		path: path,
	}
	//make sure directory exist
	db.ensureDB()
	//check if the file exit,
	_, err := os.Stat(db.path)
	if os.IsNotExist(err) {
		//create new dile and an empty json structure if it doesnt exit
		file, err := os.Create(db.path)
		if err != nil {
			return nil, fmt.Errorf("error creating data file: %v", err)
		}
		defer file.Close()

		emptyDb := DBStructure{Chirps: make(map[int]Chirp)}
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
