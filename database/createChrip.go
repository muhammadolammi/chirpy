package database

import (
	"encoding/json"
	"fmt"
	"os"
)

// CreateChirp creates a new chirp and saves it to disk

func (db *DB) CreateChirp(body string) (Chirp, error) {

	db.mux.Lock()
	defer db.mux.Unlock()

	//get old db bytes and manage errors
	oldChipsBytes, err := os.ReadFile(db.path)
	if err != nil {
		return Chirp{}, fmt.Errorf("error reading old files to create new chip. err : %v", err)
	}
	// create a new oldChirps DBStructure and unmarshal the bytes in it
	oldChirps := DBStructure{}
	err = json.Unmarshal(oldChipsBytes, &oldChirps)
	if err != nil {
		return Chirp{}, fmt.Errorf("error Unmarshalling old chips byte. err : %v", err)
	}
	//get the current id

	maxId := 0

	for id := range oldChirps.Chirps {
		if id > maxId {
			maxId = id
		}
	}
	fmt.Println(maxId)

	// create a new chirps with the id and body
	newChirp := Chirp{
		Id:   maxId + 1,
		Body: body,
	}

	// Initialize the Chirps map if it's nil
	if oldChirps.Chirps == nil {
		oldChirps.Chirps = make(map[int]Chirp)
	}

	// add the new chirp to the old chirps
	oldChirps.Chirps[maxId+1] = newChirp
	//add the new chirps to the old chirps
	oldChirps.Chirps[maxId+1] = newChirp
	//TODO add the new chip to other chip with new id
	// Write the updated chips back to the database
	newChipsBytes, err := json.MarshalIndent(oldChirps, "", "    ")
	if err != nil {
		return Chirp{}, fmt.Errorf("error marshalling updated chips to JSON. err: %v", err)
	}

	err = os.WriteFile(db.path, newChipsBytes, 0644)
	if err != nil {
		return Chirp{}, fmt.Errorf("error writing updated chips to file. err: %v", err)
	}

	return newChirp, nil

}
