package repo

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/m0t0k1ch1/metamask-login-sample/domain/model"
)

type UserRepository interface {
	Add(u *model.User) error
	Get(address common.Address) (*model.User, error)
	Update(u *model.User) error
}
