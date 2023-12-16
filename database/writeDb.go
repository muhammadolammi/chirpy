package database

import (
	"encoding/json"
	"fmt"
	"os"
)

// writeDB writes the database file to disk
func (db *DB) writeDB(dbStructure DBStructure) error {
	db.mux.Lock()
	defer db.mux.Unlock()
	newChipsBytes, err := json.MarshalIndent(dbStructure, "", "    ")
	if err != nil {
		return fmt.Errorf("error marshalling updated chips to JSON. err: %v", err)
	}

	err = os.WriteFile(db.path, newChipsBytes, 0644)
	if err != nil {
		return fmt.Errorf("error writing updated chips to file. err: %v", err)
	}

	return nil
}
