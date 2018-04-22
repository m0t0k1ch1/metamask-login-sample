package user

import (
	"context"

	"github.com/m0t0k1ch1/metamask-login-sample/application"
)

type Application interface {
	GetUser(ctx context.Context, in *GetUserInput) (*GetUserOutput, error)
	UpdateUser(ctx context.Context, in *UpdateUserInput) (*UpdateUserOutput, error)
	DeleteUser(ctx context.Context, in *DeleteUserInput) (*DeleteUserOutput, error)
}

type applicationImpl struct {
	*application.Core
}

func NewApplication(core *application.Core) Application {
	return &applicationImpl{core}
}

func (app *applicationImpl) GetUser(ctx context.Context, in *GetUserInput) (*GetUserOutput, error) {
	if err := in.Validate(); err != nil {
		return nil, err
	}

	address := in.Address()

	u, err := app.Repositories.User.Get(ctx, address)
	if err != nil {
		return nil, err
	}

	return NewGetUserOutput(u), nil
}

func (app *applicationImpl) UpdateUser(ctx context.Context, in *UpdateUserInput) (*UpdateUserOutput, error) {
	if err := in.Validate(); err != nil {
		return nil, err
	}

	address := in.Address()
	name := in.Name

	u, err := app.Repositories.User.Get(ctx, address)
	if err != nil {
		return nil, err
	}

	u.Name = name

	if err := app.Repositories.User.Update(ctx, u); err != nil {
		return nil, err
	}

	return NewUpdateUserOutput(), nil
}

func (app *applicationImpl) DeleteUser(ctx context.Context, in *DeleteUserInput) (*DeleteUserOutput, error) {
	if err := in.Validate(); err != nil {
		return nil, err
	}

	address := in.Address()

	u, err := app.Repositories.User.Get(ctx, address)
	if err != nil {
		return nil, err
	}

	if err := app.Repositories.User.Delete(ctx, u); err != nil {
		return nil, err
	}

	return NewDeleteUserOutput(), nil
}
