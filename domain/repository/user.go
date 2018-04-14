package repository

import (
	"context"

	"github.com/m0t0k1ch1/metamask-login-sample/domain/model"
)

type User interface {
	Add(ctx context.Context, user *model.User) error
	Get(ctx context.Context, address model.Address) (*model.User, error)
	Update(ctx context.Context, user *model.User) error
	Delete(ctx context.Context, user *model.User) error
}
