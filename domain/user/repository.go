package user

import (
	"context"

	"github.com/m0t0k1ch1/metamask-login-sample/domain/common"
)

type Repository interface {
	Add(ctx context.Context, user *User) error
	Get(ctx context.Context, address common.Address) (*User, error)
	Update(ctx context.Context, user *User) error
	Delete(ctx context.Context, user *User) error
}
