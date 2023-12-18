package uservalidator

import (
	"github.com/muhammadolammi/chirpy/database"
)

func CheckEmail(email string, db *database.DB) (bool, error) {

	users, err := db.GetUsers()
	if err != nil {

		return false, err

	}

	for _, user := range users {
		if email == user.Email {
			return true, nil
		}
	}

	return false, nil

}
