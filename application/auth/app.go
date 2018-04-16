package auth

import (
	"context"

	"github.com/m0t0k1ch1/metamask-login-sample/application"
	"github.com/m0t0k1ch1/metamask-login-sample/domain/auth"
	"github.com/m0t0k1ch1/metamask-login-sample/domain/common"
	"github.com/m0t0k1ch1/metamask-login-sample/domain/user"
)

type Application struct {
	*application.Core
	authService auth.Service
	userRepo    user.Repository
}

func NewApplication(core *application.Core) *Application {
	return &Application{
		Core:        core,
		authService: core.Container.AuthService,
		userRepo:    core.Container.UserRepo,
	}
}

func (app *Application) Challenge(ctx context.Context, in *ChallengeInput) (*ChallengeOutput, error) {
	if err := in.Validate(); err != nil {
		return nil, err
	}

	address := in.Address()

	u, err := app.userRepo.Get(ctx, address)
	switch err {
	case nil:
		if err := app.authService.SetUpChallenge(u); err != nil {
			return nil, err
		}
		if err := app.userRepo.Update(ctx, u); err != nil {
			return nil, err
		}
	case common.ErrUserNotFound:
		u = user.NewUser("", address)
		if err := app.authService.SetUpChallenge(u); err != nil {
			return nil, err
		}
		if err := app.userRepo.Add(ctx, u); err != nil {
			return nil, err
		}
	default:
		return nil, err
	}

	return NewChallengeOutput(u.Challenge), nil
}

func (app *Application) Authorize(ctx context.Context, in *AuthorizeInput) (*AuthorizeOutput, error) {
	if err := in.Validate(); err != nil {
		return nil, err
	}

	address := in.Address()
	sig := in.Signature()

	u, err := app.userRepo.Get(ctx, address)
	if err != nil {
		return nil, err
	}

	if err := app.authService.VerifyResponse(u, sig.Bytes()); err != nil {
		return nil, err
	}
	tokenBytes, err := app.authService.IssueToken(u)
	if err != nil {
		return nil, err
	}

	return NewAuthorizeOutput(string(tokenBytes)), nil
}
