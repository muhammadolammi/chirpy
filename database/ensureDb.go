package database

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

// ensureDB creates a new database file if it doesn't exist and ensures the directory structure exists.
func (db *DB) ensureDB() error {
	// Initialize mutex if it's not already set
	if db.mux == nil {
		db.mux = &sync.RWMutex{}
	}
	homeDir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return err
	}

	dirPath := filepath.Join(homeDir, "database")
	err = os.MkdirAll(dirPath, 0777)
	if err != nil {
		fmt.Println(err)
		return err
	}
	// fmt.Println("Directory hierarchy created successfully")

	// Check if the directory actually exists
	_, err = os.Stat(dirPath)
	if err != nil {
		// fmt.Println("Error checking directory existence:", err)
		return err
	}

	// fmt.Println("Directory exists:", info.IsDir())

	// Ensure the directory structure exists
	// file, err := os.Create(filepath.Dir(db.path))
	// if err != nil {
	// 	return fmt.Errorf("error creating directory structure: %v", err)
	// }
	// defer file.Close()

	return nil
}
