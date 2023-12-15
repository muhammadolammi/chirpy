package database

import (
	"testing"
)

func TestCreateDb(t *testing.T) {
	_, err := NewDB("database.json")
	if err != nil {
		t.Error(err)
	}

	// data := struct {
	// 	Body string `json:"body"`
	// }{
	// 	Body: "What about second breakfast?",
	// }
	// strBytes, err := json.Marshal(data)
	// chirp, err := db.CreateChirp(string(strBytes))

}
