package user

import (
	"context"

	"github.com/m0t0k1ch1/metamask-login-sample/application"
	"github.com/m0t0k1ch1/metamask-login-sample/domain/common"
	"github.com/m0t0k1ch1/metamask-login-sample/domain/user"
)

type Application struct {
	*application.Base
	userRepo user.Repository
}

func NewApplication(core *application.Core) *Application {
	base := application.NewBase(core)

	return &Application{
		Base:     base,
		userRepo: base.Container().NewUserRepository(),
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

func (app *Application) getUser(ctx context.Context, address common.Address) (*user.User, error) {
	return app.userRepo.Get(ctx, address)
}
