package database

func (db *DB) GetUser(email string) (User, error) {
	err := db.ensureDB()
	if err != nil {
		return User{}, err
	}
	chirps, err := db.GetChirps()
	if err != nil {
		return User{}, err
	}
	for _, chirp := range chirps {
		if email == chirp.Email {
			return User{
				Id:    chirp.Id,
				Email: chirp.Email,
			}, nil

		}

	}

	return User{}, nil

}
