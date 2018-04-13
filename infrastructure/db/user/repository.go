package user

import (
	"context"

	"github.com/m0t0k1ch1/metamask-login-sample/domain"
	"github.com/m0t0k1ch1/metamask-login-sample/domain/user"
	"github.com/m0t0k1ch1/metamask-login-sample/library/kvs"
)

type repository struct{}

func NewRepository() user.Repository {
	return &repository{}
}

func (repo *repository) Add(ctx context.Context, user *domain.User) error {
	if _, ok := kvs.Get(user.Address.Hex()); ok {
		return domain.ErrUserAlreadyExists
	}

	kvs.Set(user.Address.Hex(), user)

	return nil
}

func (repo *repository) Get(ctx context.Context, address domain.Address) (*domain.User, error) {
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

func (repo *repository) Update(ctx context.Context, user *domain.User) error {
	if _, ok := kvs.Get(user.Address.Hex()); !ok {
		return domain.ErrUserNotFound
	}

	kvs.Set(user.Address.Hex(), user)

	return nil
}

func (repo *repository) Delete(ctx context.Context, user *domain.User) error {
	if _, ok := kvs.Delete(user.Address.Hex()); !ok {
		return domain.ErrUserNotFound
	}

	return nil
}
