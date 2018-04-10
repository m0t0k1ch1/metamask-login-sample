package domain

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

func (u *User) AddressHex() string {
	return u.Address.Hex()
}

func (u *User) UpdateChallenge() {
	// TODO: generate new challenge
	u.Challenge = u.AddressHex()
}
