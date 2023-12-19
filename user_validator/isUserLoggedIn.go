package uservalidator

import (
	"errors"

	"github.com/muhammadolammi/chirpy/database"
)

func IsUserLoggedIn(id int) (bool, error) {

	usersDb, err := database.NewUsersDB("database/users.json")
	if err != nil {

		return false, err

	}
	user, err := usersDb.GetUserById(id)
	if err != nil {

		return false, err

	}

	if !user.IsLoggedIn {

		return false, errors.New("user not logged in")
	}
	return true, nil
}
