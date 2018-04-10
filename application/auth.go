package application

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/m0t0k1ch1/metamask-login-sample/application/input"
	"github.com/m0t0k1ch1/metamask-login-sample/application/output"
	"github.com/m0t0k1ch1/metamask-login-sample/domain"
	"github.com/m0t0k1ch1/metamask-login-sample/domain/repository"
	"github.com/m0t0k1ch1/metamask-login-sample/infrastructure/crypto"
)

type AuthApplication struct {
	userRepo repository.UserRepository
}

func NewAuthApplication(repo repository.UserRepository) *AuthApplication {
	return &AuthApplication{
		userRepo: repo,
	}
}

func (app *AuthApplication) Challenge(in *input.AuthChallengeInput) (
	*output.AuthChallengeOutput,
	error,
) {
	if err := in.Validate(); err != nil {
		return nil, err
	}

	address := in.Address()

	user, err := app.getUser(address)
	switch err {
	case nil:
		user.UpdateChallenge()
	case domain.ErrUserNotFound:
		user, err = app.createUser(address)
		if err != nil {
			return nil, err
		}
	default:
		return nil, err
	}

	out := output.NewAuthChallengeOutput(user.Challenge)

	return out, nil
}

func (app *AuthApplication) Authorize(in *input.AuthAuthorizeInput) (
	*output.AuthAuthorizeOutput,
	error,
) {
	if err := in.Validate(); err != nil {
		return nil, err
	}

	address := in.Address()
	sigBytes := in.SigBytes()

	user, err := app.getUser(address)
	if err != nil {
		return nil, err
	}

	// TODO: refactoring
	hashBytes := domain.NewMyTypedData(user.Challenge).SignatureHashBytes()
	pubkey, err := crypto.RecoverTypedSignature(hashBytes, sigBytes)
	if err != nil {
		return nil, err
	}
	if crypto.PubkeyToAddressHex(pubkey) != address.Hex() {
		return nil, domain.ErrInvalidSignature
	}

	// TODO: create JWT
	token := "success"

	out := output.NewAuthAuthorizeOutput(token)

	return out, nil
}

func (app *AuthApplication) createUser(address common.Address) (*domain.User, error) {
	user := domain.NewUser(address)
	user.UpdateChallenge()

	if err := app.userRepo.Add(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (app *AuthApplication) getUser(address common.Address) (*domain.User, error) {
	return app.userRepo.Get(address)
}
