package database

func (db *DB) GetUser(email string) (User, error) {
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
				Id:       user.Id,
				Email:    user.Email,
				Password: user.Password,
			}, nil

		}

	}

	return User{}, nil

}
