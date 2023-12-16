package uservalidator

import (
	"errors"

	"github.com/muhammadolammi/chirpy/database"
	"golang.org/x/crypto/bcrypt"
)

func ValidatePass(pass string, db *database.DB) (bool, error) {
	chirps, err := db.GetChirps()
	if err != nil {
		return false, err
	}
	for _, chirp := range chirps {
		err := bcrypt.CompareHashAndPassword([]byte(chirp.Password), []byte(pass))
		if err == nil {
			// Passwords match
			return true, nil
		} else if err == bcrypt.ErrMismatchedHashAndPassword {
			// Passwords do not match
			return false, errors.New("Wrong password")
		} else {
			// An error occurred
			return false, err
		}
	}

	return false, nil
}
