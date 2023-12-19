package database

import (
	"encoding/json"
	"os"
	"testing"
)

func TestGetChips(t *testing.T) {
	db, err := NewChirpsDB("db_test.json")
	if err != nil {
		t.Errorf("theres is an err : %v", err)
	}

	chirps, err := db.GetChirps()
	if err != nil {
		t.Errorf("theres is an err : %v", err)
	}
	localDbyte, err := os.ReadFile(db.path)
	if err != nil {
		t.Errorf("theres is an err : %v", err)
	}
	localDb := Chirpstructure{}
	err = json.Unmarshal(localDbyte, &localDb)
	if err != nil {
		t.Errorf("theres is an err : %v", err)
	}
	localDbArr := MapToArray(localDb.Chirps)
	for id, ch := range localDbArr {
		ch2 := chirps[id]
		if ch.AuthorId != ch2.AuthorId {
			t.Error("something went wrong")

		}
	}

}
