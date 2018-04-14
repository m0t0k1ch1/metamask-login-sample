package auth

import (
	"context"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/m0t0k1ch1/metamask-login-sample/domain"
	"github.com/m0t0k1ch1/metamask-login-sample/domain/user"
)

const (
	secret = "XMa14bGR68BbHzo7IzrmIVzhn83hYi8u"
)

type Application struct {
	secret   string
	userRepo user.Repository
}

func NewApplication() *Application {
	return &Application{
		secret:   secret,
		userRepo: user.NewRepository(),
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
		user.UpdateToken()
	case domain.ErrUserNotFound:
		user, err = app.createUser(ctx, address)
		if err != nil {
			return nil, err
		}
	default:
		return nil, err
	}

	out := NewChallengeOutput(user.Token())

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

func (app *Application) createUser(ctx context.Context, address domain.Address) (*domain.User, error) {
	user := domain.NewUser(address)
	user.UpdateToken()

	if err := app.userRepo.Add(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}

func (app *Application) getUser(ctx context.Context, address domain.Address) (*domain.User, error) {
	return app.userRepo.Get(ctx, address)
}

func (app *Application) newSignedToken(address domain.Address) (string, error) {
	return jwt.NewWithClaims(
		jwt.SigningMethodHS256, domain.NewAuthClaims(address),
	).SignedString([]byte(app.secret))
}
