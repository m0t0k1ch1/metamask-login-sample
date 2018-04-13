package storage

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/m0t0k1ch1/metamask-login-sample/domain"
	"github.com/m0t0k1ch1/metamask-login-sample/library/kvs"
)

type UserStorage struct{}

func NewUserStorage() *UserStorage {
	return &UserStorage{}
}

func (storage *UserStorage) Add(user *domain.User) error {
	if _, ok := kvs.Get(user.AddressHex()); ok {
		return domain.ErrUserAlreadyExists
	}

	kvs.Set(user.AddressHex(), user)

	return nil
}

func (storage *UserStorage) Get(address common.Address) (*domain.User, error) {
	data, ok := kvs.Get(address.Hex())
	if !ok {
		return nil, domain.ErrUserNotFound
	}

	user, ok := data.(*domain.User)
	if !ok {
		return nil, domain.ErrUserBroken
	}

	return user, nil
}

func (storage *UserStorage) Update(user *domain.User) error {
	if _, ok := kvs.Get(user.AddressHex()); !ok {
		return domain.ErrUserNotFound
	}

	kvs.Set(user.AddressHex(), user)

	return nil
}

func (storage *UserStorage) Delete(user *domain.User) error {
	if _, ok := kvs.Delete(user.AddressHex()); !ok {
		return domain.ErrUserNotFound
	}

	return nil
}
