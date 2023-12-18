package database

import "sync"

type User struct {
	Id    int    `json:"id"`
	Email string `json:"email"`
	Token string `json:"token"`
}
type DB struct {
	path string
	mux  *sync.RWMutex
}
type DBStructure struct {
	Chirps map[int]Chirp `json:"chirps"`
}

type Chirp struct {
	Id       int    `json:"id"`
	Body     string `json:"body"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
