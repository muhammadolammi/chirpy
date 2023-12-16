package uservalidator

import (
	"errors"

	"github.com/muhammadolammi/chirpy/database"
)

func CheckEmail(email string, db *database.DB) (bool, error) {

	chirps, err := db.GetChirps()
	if err != nil {

		return false, errors.New("error getting db")

	}

	for _, chirp := range chirps {
		if email == chirp.Email {
			return true, nil
		}
	}

	return false, nil

}
