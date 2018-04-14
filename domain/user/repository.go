package user

import (
	"context"

	"github.com/m0t0k1ch1/metamask-login-sample/domain"
)

var NewRepository func() Repository

type Repository interface {
	Add(ctx context.Context, user *domain.User) error
	Get(ctx context.Context, address domain.Address) (*domain.User, error)
	Update(ctx context.Context, user *domain.User) error
	Delete(ctx context.Context, user *domain.User) error
}
