package encrypt

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func Hash(passwort string) (string, error) {

	if passwort == "" {
		return "", fmt.Errorf("no input")
	} else {
		GenPassword, err := bcrypt.GenerateFromPassword([]byte(passwort), bcrypt.DefaultCost)

		if err != nil {
			return "cannot generate encrypted password", err
		}

		hashedString := string(GenPassword)
		return hashedString, nil
	}

}

func Verify(passwort, hashedPasswort string) (bool, error) {
	if passwort == "" || hashedPasswort == "" {
		return false, nil
	}

	err := bcrypt.CompareHashAndPassword([]byte(hashedPasswort), []byte(passwort))
	if err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return false, fmt.Errorf("invalid string comparison: %v", err)
		}
		return false, err
	}
	return true, nil
}
