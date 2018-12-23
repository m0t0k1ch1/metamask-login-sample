package domain

import (
	"unicode/utf8"
)

type User struct {
	Name      string
	Address   Address
	Challenge string
}

func NewUser(name string, address Address) *User {
	return &User{
		Name:    name,
		Address: address,
	}
}

func ValidateUserName(name string) error {
	l := utf8.RuneCountInString(name)
	if l < UserNameLengthMin {
		return ErrTooShortUserName
	}
	if l > UserNameLengthMax {
		return ErrTooLongUserName
	}
	return nil
}
