package auth

import (
	"context"

	"github.com/m0t0k1ch1/metamask-login-sample/domain"
	"github.com/m0t0k1ch1/metamask-login-sample/domain/user"
	"github.com/m0t0k1ch1/metamask-login-sample/library/crypto"
)

type Application struct {
	userRepo user.Repository
}

func NewApplication() *Application {
	return &Application{
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

	// TODO: refactoring
	hashBytes := domain.NewMyTypedData(user.Token()).SignatureHashBytes()
	pubkey, err := crypto.RecoverTypedSignature(hashBytes, sig.Bytes())
	if err != nil {
		return nil, err
	}
	if crypto.PubkeyToAddressHex(pubkey) != address.Hex() {
		return nil, domain.ErrInvalidSignature
	}

	// TODO: create JWT
	token := "success"

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
