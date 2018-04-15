package user

import (
	"context"

	"github.com/m0t0k1ch1/metamask-login-sample/application"
	"github.com/m0t0k1ch1/metamask-login-sample/domain/user"
)

type Application struct {
	*application.Core
	userRepo user.Repository
}

func NewApplication(core *application.Core) *Application {
	return &Application{
		Core:     core,
		userRepo: core.Container.UserRepo,
	}
}

func (app *Application) GetUser(ctx context.Context, in *GetUserInput) (*GetUserOutput, error) {
	if err := in.Validate(); err != nil {
		return nil, err
	}

	address := in.Address()

	u, err := app.userRepo.Get(ctx, address)
	if err != nil {
		return nil, err
	}

	out := NewGetUserOutput(u)

	return out, nil
}
