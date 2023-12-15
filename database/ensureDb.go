// ensureDB creates a new database file if it doesn't exist
package database

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

func (db *DB) ensureDB() error {
	if db.mux == nil {
		db.mux = &sync.RWMutex{}
	}
	// Ensure the directory structure exists
	if err := os.MkdirAll(filepath.Dir(db.path), 0755); err != nil {
		return fmt.Errorf("error creating directory structure: %v\n", err)

	}

	return nil
}
