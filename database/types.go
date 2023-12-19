package database

import (
	"sync"
	"time"
)

type User struct {
	Id          int    `json:"id"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	IsLoggedIn  bool   `json:"is_logged_in"`
	IsChirpyRed bool   `json:"is_chirpy_red"`
}
type DB struct {
	path string
	mux  *sync.RWMutex
}
type Chirpstructure struct {
	Chirps map[int]Chirp `json:"chirps"`
}

type RevokedSessionDetail struct {
	Time time.Time `json:"time"`
}
type RevokedSesssionstructure struct {
	RevokedSessions map[string]RevokedSessionDetail `json:"revoked_sessions"`
}
type Userstructure struct {
	Users map[int]User `json:"users"`
}

type Chirp struct {
	AuthorId int    `json:"author_id"`
	Body     string `json:"body"`
	Id       int    `json:"id"`
}
