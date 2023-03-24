package entity

import (
	"net/mail"
	"regexp"
)

func emailIsValid(s string) bool {
	_, err := mail.ParseAddress(s)
	return err == nil
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

func checkPhoneNumber(s string) error {
	re := regexp.MustCompile(`^(?:(?:\(?(?:00|\+)([1-4]\d\d|[1-9]\d?)\)?)?[\-\.\ \\\/]?)?((?:\(?\d{1,}\)?[\-\.\ \\\/]?){0,})(?:[\-\.\ \\\/]?(?:#|ext\.?|extension|x)[\-\.\ \\\/]?(\d+))?$`)

	if !re.MatchString(s) {
		return ErrPhoneIsNotValid
	}

	return nil
}

func checkGender(g Gender) error {
	if g != GenderUnknown && g != GenderMale && g != GenderFemale {
		return ErrGenderIsNotValid
	}

	return nil
}

func checkRole(r SystemRole) error {
	if r != RoleUser && r != RoleSuperAdmin && r != RoleAdmin {
		return ErrRoleIsNotValid
	}

	return nil
}

func checkStatus(s Status) error {

	if s != StatusPendingVerify && s != StatusActive && s != StatusBanned {
		return ErrStatusIsNotValid
	}

	return nil
}
