package cache

import (
	"context"

	"github.com/m0t0k1ch1/metamask-login-sample/domain"
	"github.com/m0t0k1ch1/metamask-login-sample/domain/model"
	"github.com/m0t0k1ch1/metamask-login-sample/domain/repository"
	"github.com/m0t0k1ch1/metamask-login-sample/library/kvs"
)

type user struct{}

func NewUser() repository.User {
	return &user{}
}

func (repo *user) Add(ctx context.Context, user *model.User) error {
	if _, ok := kvs.Get(user.Address.Hex()); ok {
		return domain.ErrUserAlreadyExists
	}

	kvs.Set(user.Address.Hex(), user)

	return nil
}

func (repo *user) Get(ctx context.Context, address model.Address) (*model.User, error) {
	data, ok := kvs.Get(address.Hex())
	if !ok {
		return nil, domain.ErrUserNotFound
	}

	user, ok := data.(*model.User)
	if !ok {
		return nil, domain.ErrUserBroken
	}

	return user, nil
}

func (repo *user) Update(ctx context.Context, user *model.User) error {
	if _, ok := kvs.Get(user.Address.Hex()); !ok {
		return domain.ErrUserNotFound
	}

	kvs.Set(user.Address.Hex(), user)

	return nil
}

func (repo *user) Delete(ctx context.Context, user *model.User) error {
	if _, ok := kvs.Delete(user.Address.Hex()); !ok {
		return domain.ErrUserNotFound
	}

	return nil
}
