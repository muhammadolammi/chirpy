package database

import (
	"errors"
	"fmt"
)

func (db *DB) GetRevokedSessions() (map[string]RevokedSessionDetail, error) {
	if db.path != "database/sessions.json" {
		return nil, errors.New("wrong directory to get revoked sessions")
	}

	dbJson, err := db.loadSessions()
	if err != nil {
		return nil, fmt.Errorf("error loading database: %v", err)
	}
	sessions := dbJson.RevokedSessions
	return sessions, nil

}

func (db *DB) GetSessionDetail(token string) (RevokedSessionDetail, error) {

	sessions, err := db.GetRevokedSessions()
	if err != nil {
		return RevokedSessionDetail{}, err
	}
	sessionDetail, ok := sessions[token]
	if !ok {
		return RevokedSessionDetail{}, errors.New("no session")
	}
	return sessionDetail, nil

}

func (db *DB) IsSession(token string) (bool, error) {
	sessionsMap, err := db.GetRevokedSessions()
	if err != nil {
		return false, err
	}
	_, ok := sessionsMap[token]
	if ok {
		return true, nil
	}
	return false, nil

}
