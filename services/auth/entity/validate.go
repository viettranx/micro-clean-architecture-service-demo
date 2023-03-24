package entity

import "net/mail"

func emailIsValid(s string) bool {
	_, err := mail.ParseAddress(s)
	return err == nil
}

func checkPassword(s string) error {
	if len(s) < 8 || len(s) > 30 {
		return ErrPasswordIsNotValid
	}

	return nil
}

func checkFirstName(s string) error {
	if s == "" {
		return ErrFirstNameIsEmpty
	}

	if len(s) > 30 {
		return ErrFirstNameTooLong
	}

	return nil
}

func checkLastName(s string) error {
	if s == "" {
		return ErrLastNameIsEmpty
	}

	if len(s) > 30 {
		return ErrLastNameTooLong
	}

	return nil
}
