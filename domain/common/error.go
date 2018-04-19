package common

import (
	"fmt"
)

var (
	ErrCodeUnexpected = 1000

	ErrInvalidAddressHex      = NewError(2001, "Invalid address hex")
	ErrInvalidSignatureLength = NewError(2002, "Invalid signature length")
	ErrInvalidSignatureHex    = NewError(2003, "Invalid signature hex")
	ErrInvalidSignature       = NewError(2004, "Invalid signature")

	ErrTooShortUserName = NewError(3001,
		fmt.Sprintf("Too short user name (min: %d)", UserNameLengthMin),
	)
	ErrTooLongUserName = NewError(3002,
		fmt.Sprintf("Too long user name (max: %d)", UserNameLengthMax),
	)
	ErrUserNotFound      = NewError(3003, "User not found")
	ErrUserAlreadyExists = NewError(3004, "User already exists")
	ErrUserBroken        = NewError(3005, "User broken")
)

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (err *Error) Error() string {
	return fmt.Sprintf("%s [%d]", err.Message, err.Code)
}

func NewError(code int, msg string) *Error {
	return &Error{
		Code:    code,
		Message: msg,
	}
}

func NewUnexpectedError() *Error {
	return NewError(ErrCodeUnexpected, "Internal server error")
}
