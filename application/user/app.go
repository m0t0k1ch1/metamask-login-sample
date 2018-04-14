package user

import (
	"context"

	"github.com/m0t0k1ch1/metamask-login-sample/config"
	"github.com/m0t0k1ch1/metamask-login-sample/domain"
	"github.com/m0t0k1ch1/metamask-login-sample/domain/model"
	"github.com/m0t0k1ch1/metamask-login-sample/domain/repository"
)

type Application struct {
	userRepo repository.User
}

func NewApplication(conf *config.AppConfig, container *domain.Container) *Application {
	return &Application{
		userRepo: container.NewUserRepository(),
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

func (app *Application) getUser(ctx context.Context, address model.Address) (*model.User, error) {
	return app.userRepo.Get(ctx, address)
}
