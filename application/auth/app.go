package auth

import (
	"context"

	"github.com/m0t0k1ch1/metamask-login-sample/config"
	"github.com/m0t0k1ch1/metamask-login-sample/domain"
	"github.com/m0t0k1ch1/metamask-login-sample/domain/model"
	"github.com/m0t0k1ch1/metamask-login-sample/domain/repository"
)

type Application struct {
	secret   string
	userRepo repository.User
}

func NewApplication(conf *config.AppConfig, container *domain.Container) *Application {
	return &Application{
		secret:   conf.Secret,
		userRepo: container.NewUserRepository(),
	}
}

func (app *Application) Challenge(ctx context.Context, in *ChallengeInput) (*ChallengeOutput, error) {
	if err := in.Validate(); err != nil {
		return nil, err
	}

	address := in.Address()

	user, err := app.getUser(ctx, address)
	switch err {
	case nil:
		user.UpdateChallenge()
	case domain.ErrUserNotFound:
		user, err = app.createUser(ctx, address)
		if err != nil {
			return nil, err
		}
	default:
		return nil, err
	}

	out := NewChallengeOutput(user.Challenge())

	return out, nil
}

func (app *Application) Authorize(ctx context.Context, in *AuthorizeInput) (*AuthorizeOutput, error) {
	if err := in.Validate(); err != nil {
		return nil, err
	}

	address := in.Address()
	sig := in.Signature()

	user, err := app.getUser(ctx, address)
	if err != nil {
		return nil, err
	}

	pubkey, err := user.AuthTypedData().RecoverPubkey(sig)
	if err != nil {
		return nil, err
	}
	if pubkey.Address().Hex() != address.Hex() {
		return nil, domain.ErrInvalidSignature
	}

	token, err := app.newSignedToken(address)
	if err != nil {
		return nil, err
	}

	out := NewAuthorizeOutput(token)

	return out, nil
}

func (app *Application) createUser(ctx context.Context, address model.Address) (*model.User, error) {
	user := model.NewUser(address)
	user.UpdateChallenge()

	if err := app.userRepo.Add(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}

func (app *Application) getUser(ctx context.Context, address model.Address) (*model.User, error) {
	return app.userRepo.Get(ctx, address)
}

func (app *Application) newSignedToken(address model.Address) (string, error) {
	return model.NewAuthToken(address).SignedString(app.secret)
}
