package auth

import (
	"context"

	"github.com/m0t0k1ch1/metamask-login-sample/application"
	"github.com/m0t0k1ch1/metamask-login-sample/domain"
)

type Application interface {
	Challenge(ctx context.Context, in *ChallengeInput) (*ChallengeOutput, error)
	Authorize(ctx context.Context, in *AuthorizeInput) (*AuthorizeOutput, error)
}

type applicationImpl struct {
	*application.Core
}

func NewApplication(core *application.Core) Application {
	return &applicationImpl{core}
}

func (app *applicationImpl) Challenge(ctx context.Context, in *ChallengeInput) (*ChallengeOutput, error) {
	if err := in.Validate(); err != nil {
		return nil, err
	}

	address := in.Address()

	u, err := app.Repositories.User.Get(ctx, address)
	switch err {
	case nil:
		if err := app.Services.Auth.SetUpChallenge(u); err != nil {
			return nil, err
		}
		if err := app.Repositories.User.Update(ctx, u); err != nil {
			return nil, err
		}
	case domain.ErrUserNotFound:
		u = domain.NewUser("", address)
		if err := app.Services.Auth.SetUpChallenge(u); err != nil {
			return nil, err
		}
		if err := app.Repositories.User.Add(ctx, u); err != nil {
			return nil, err
		}
	default:
		return nil, err
	}

	return NewChallengeOutput(u.Challenge), nil
}

func (app *applicationImpl) Authorize(ctx context.Context, in *AuthorizeInput) (*AuthorizeOutput, error) {
	if err := in.Validate(); err != nil {
		return nil, err
	}

	address := in.Address()
	sig := in.Signature()

	u, err := app.Repositories.User.Get(ctx, address)
	if err != nil {
		return nil, err
	}

	if err := app.Services.Auth.VerifyResponse(u, sig.Bytes()); err != nil {
		return nil, err
	}
	tokenBytes, err := app.Services.Auth.IssueToken(u)
	if err != nil {
		return nil, err
	}

	return NewAuthorizeOutput(string(tokenBytes)), nil
}
