package user

import (
	"context"

	"github.com/m0t0k1ch1/metamask-login-sample/domain"
	"github.com/m0t0k1ch1/metamask-login-sample/domain/user"
)

type Application struct {
	userRepo user.Repository
}

func NewApplication() *Application {
	return &Application{
		userRepo: user.NewRepository(),
	}
}

func (app *Application) GetUser(ctx context.Context, in *GetUserInput) (*GetUserOutput, error) {
	if err := in.Validate(); err != nil {
		return nil, err
	}

	address := in.Address()

	user, err := app.getUser(ctx, address)
	if err != nil {
		return nil, err
	}

	out := NewGetUserOutput(user)

	return out, nil
}

func (app *Application) getUser(ctx context.Context, address domain.Address) (*domain.User, error) {
	return app.userRepo.Get(ctx, address)
}
