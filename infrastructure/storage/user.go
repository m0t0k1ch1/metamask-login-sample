package storage

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/m0t0k1ch1/metamask-login-sample/domain/model"
	"github.com/m0t0k1ch1/metamask-login-sample/library/kvs"
)

type UserStorage struct{}

func NewUserStorage() *UserStorage {
	return &UserStorage{}
}

func (storage *UserStorage) Add(u *model.User) error {
	if _, ok := kvs.Get(u.AddressHex()); ok {
		return model.ErrUserAlreadyExists
	}

	kvs.Set(u.AddressHex(), u)

	return nil
}

func (storage *UserStorage) Get(address common.Address) (*model.User, error) {
	data, ok := kvs.Get(address.Hex())
	if !ok {
		return nil, model.ErrUserNotFound
	}

	u, ok := data.(*model.User)
	if !ok {
		return nil, model.ErrUserBroken
	}

	return u, nil
}

func (storage *UserStorage) Update(u *model.User) error {
	if _, ok := kvs.Get(u.AddressHex()); !ok {
		return model.ErrUserNotFound
	}

	kvs.Set(u.AddressHex(), u)

	return nil
}
