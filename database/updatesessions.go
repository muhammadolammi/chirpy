package database

import (
	"errors"
	"time"
)

func (db *DB) UpdateSession(token string) error {

	//make sure its the user database
	if db.path != "database/sessions.json" {
		return errors.New("this is not the session directory")
	}
	//get the chrips map from database
	databaseStruct, err := db.loadSessions()
	if err != nil {

		return errors.New("error gettings database struct map")
	}

	//get the current user with the user id.
	curSession, ok := databaseStruct.RevokedSessions[token]
	if ok {
		return errors.New("session already revoked ")
	}
	updtatedSession := curSession
	updtatedSession.Time = time.Now().UTC()
	databaseStruct.RevokedSessions[token] = updtatedSession
	db.writeDB(databaseStruct)
	return nil

}
