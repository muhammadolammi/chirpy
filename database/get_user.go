package database

func (db *DB) GetUserByEmail(email string) (User, error) {
	err := db.ensureDB()
	if err != nil {
		return User{}, err
	}
	users, err := db.GetUsers()
	if err != nil {
		return User{}, err
	}
	for _, user := range users {
		if email == user.Email {
			return User{
				Id:          user.Id,
				Email:       user.Email,
				Password:    user.Password,
				IsLoggedIn:  user.IsLoggedIn,
				IsChirpyRed: user.IsChirpyRed,
			}, nil

		}

	}

	return User{}, nil

}

func (db *DB) GetUserById(userId int) (User, error) {
	err := db.ensureDB()
	if err != nil {
		return User{}, err
	}
	users, err := db.GetUsers()
	if err != nil {
		return User{}, err
	}
	for _, user := range users {
		if userId == user.Id {
			return User{
				Id:          user.Id,
				Email:       user.Email,
				Password:    user.Password,
				IsLoggedIn:  user.IsLoggedIn,
				IsChirpyRed: user.IsChirpyRed,
			}, nil

		}

	}

	return User{}, nil

}
