package uservalidator

import (
	"errors"
	"unicode"
)

func ValidatePassStrenght(pass string) (bool, error) {
	if len(pass) < 15 {
		return false, errors.New("your password is less than 15 characters")
	}
	hasSpecial := false
	hasCapital := false

	for _, char := range pass {
		if unicode.IsUpper(char) {
			hasCapital = true
		} else if !unicode.IsLetter(char) && !unicode.IsDigit(char) {
			hasSpecial = true
		}
	}

	if !hasSpecial {
		return false, errors.New("your password must contain at least one special character")
	}

	if !hasCapital {
		return false, errors.New("your password must contain at least one capital letter")
	}

	return true, nil

}
