package common

import (
	"fmt"
)

var (
	ErrCodeUnexpected = 1000

	ErrInvalidAddressHex = NewError(
		2001,
		"Address is not hex",
	)
	ErrInvalidSignatureLength = NewError(
		2002,
		fmt.Sprintf("Signature must be %d bytes", SignatureLength),
	)
	ErrInvalidSignatureHex = NewError(
		2003,
		"Signature is not hex",
	)
	ErrInvalidSignature = NewError(
		2004,
		"Signature is invalid",
	)

	ErrTooShortUserName = NewError(
		3001,
		fmt.Sprintf("User name must be %d characters or more", UserNameLengthMin),
	)
	ErrTooLongUserName = NewError(
		3002,
		fmt.Sprintf("User name must be %d characters or less", UserNameLengthMax),
	)
	ErrUserNotFound = NewError(
		3003,
		"User is not found",
	)
	ErrUserAlreadyExists = NewError(
		3004,
		"User already exists",
	)
	ErrUserBroken = NewError(
		3005,
		"User was broken",
	)
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
