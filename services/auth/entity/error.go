package entity

import "errors"

var (
	ErrPasswordIsNotValid = errors.New("password must have from 8 to 30 characters")
	ErrEmailIsNotValid    = errors.New("email is not valid")
	ErrEmailHasExisted    = errors.New("email has existed")
	ErrLoginFailed        = errors.New("email and password are not valid")
	ErrFirstNameIsEmpty   = errors.New("first name can not be blank")
	ErrFirstNameTooLong   = errors.New("first name too long, max character is 30")
	ErrLastNameIsEmpty    = errors.New("last name can not be blank")
	ErrLastNameTooLong    = errors.New("last name too long, max character is 30")
	ErrCannotRegister     = errors.New("cannot register")
)
