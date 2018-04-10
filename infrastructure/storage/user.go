package storage

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/m0t0k1ch1/metamask-login-sample/domain/model"
)

// TODO
type UserStorage struct{}

// TODO
func NewUserStorage() *UserStorage {
	return &UserStorage{}
}

// TODO
func (storage *UserStorage) Add(*model.User) error {
	return nil
}

// TODO
func (storage *UserStorage) Get(common.Address) (*model.User, error) {
	return nil, nil
}

// TODO
func (storage *UserStorage) Update(*model.User) error {
	return nil
}
