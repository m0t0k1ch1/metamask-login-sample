package domain

import (
	"fmt"
)

var (
	ErrCodeUnexpected = 1000

	ErrInvalidAddressHex   = NewError(2001, "Invalid address hex")
	ErrInvalidSignatureHex = NewError(2002, "Invalid signature hex")
	ErrInvalidSignature    = NewError(2003, "Invalid signature")

	ErrUserNotFound      = NewError(3001, "User not found")
	ErrUserAlreadyExists = NewError(3002, "User already exists")
	ErrUserBroken        = NewError(3003, "User broken")
)

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (err *Error) Error() string {
	return fmt.Sprintf("%s [%d]", err.Message, err.Code)
}

func NewError(code int, message string) *Error {
	return &Error{
		Code:    code,
		Message: message,
	}
}

func NewUnexpectedError() *Error {
	return NewError(ErrCodeUnexpected, "Internal server error")
}
