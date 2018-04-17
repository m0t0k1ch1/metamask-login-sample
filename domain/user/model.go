package user

import (
	"unicode/utf8"

	"github.com/m0t0k1ch1/metamask-login-sample/domain/common"
)

type User struct {
	Name      string
	Address   common.Address
	Challenge string
}

func NewUser(name string, address common.Address) *User {
	return &User{
		Name:    name,
		Address: address,
	}
}

func ValidateUserName(name string) error {
	l := utf8.RuneCountInString(name)
	if l <= common.UserNameLengthMin {
		return common.ErrTooShortUserName
	}
	if l >= common.UserNameLengthMax {
		return common.ErrTooLongUserName
	}
	return nil
}
