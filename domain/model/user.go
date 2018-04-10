package model

import "github.com/ethereum/go-ethereum/common"

type User struct {
	Name      string
	Address   common.Address
	Challenge string
}

func NewUser(address common.Address) *User {
	return &User{
		Address: address,
	}
}

// TODO
func (u *User) UpdateChallenge() {}
