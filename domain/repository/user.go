package repository

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/m0t0k1ch1/metamask-login-sample/domain"
)

type UserRepository interface {
	Add(user *domain.User) error
	Get(address common.Address) (*domain.User, error)
	Update(user *domain.User) error
	Delete(user *domain.User) error
}
