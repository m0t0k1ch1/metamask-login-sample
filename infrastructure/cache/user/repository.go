package user

import (
	"context"

	"github.com/m0t0k1ch1/metamask-login-sample/domain/common"
	"github.com/m0t0k1ch1/metamask-login-sample/domain/user"
	"github.com/m0t0k1ch1/metamask-login-sample/library/kvs"
)

type repository struct{}

func NewRepository() user.Repository {
	return &repository{}
}

func (repo *repository) Add(ctx context.Context, u *user.User) error {
	if _, ok := kvs.Get(u.Address.Hex()); ok {
		return common.ErrUserAlreadyExists
	}

	kvs.Set(u.Address.Hex(), u)

	return nil
}

func (repo *repository) Get(ctx context.Context, address common.Address) (*user.User, error) {
	data, ok := kvs.Get(address.Hex())
	if !ok {
		return nil, common.ErrUserNotFound
	}

	u, ok := data.(*user.User)
	if !ok {
		return nil, common.ErrUserBroken
	}

	return u, nil
}

func (repo *repository) Update(ctx context.Context, u *user.User) error {
	if _, ok := kvs.Get(u.Address.Hex()); !ok {
		return common.ErrUserNotFound
	}

	kvs.Set(u.Address.Hex(), u)

	return nil
}

func (repo *repository) Delete(ctx context.Context, u *user.User) error {
	if _, ok := kvs.Delete(u.Address.Hex()); !ok {
		return common.ErrUserNotFound
	}

	return nil
}
