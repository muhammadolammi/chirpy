package uservalidator

import (
	"errors"

	"github.com/muhammadolammi/chirpy/database"
	"golang.org/x/crypto/bcrypt"
)

func ValidatePass(pass, email string, db *database.DB) (bool, error) {
	user, err := db.GetUserByEmail(email)
	if err != nil {
		return false, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pass))
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
	return false, nil
}
