package application

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/m0t0k1ch1/metamask-login-sample/domain/model"
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

func (app *AuthApplication) Challenge(addressHex string) (string, error) {
	address, err := hexToAddress(addressHex)
	if err != nil {
		return "", err
	}

	// TODO: if user exists

	// TODO
	// user, err := app.createUser(address)
	// if err != nil {
	// 	return "", err
	// }

	// return user.Challenge, nil
	return address.Hex(), nil
}

func (app *AuthApplication) Login(addressHex, sigHex string) (string, error) {
	address, err := hexToAddress(addressHex)
	if err != nil {
		return "", err
	}

	// TODO: validate signature format

	// TODO:
	// user, err := app.getUser(address)
	// if err != nil {
	// 	return "", err
	// }

	// data := model.NewMyTypedData(user.Challenge)
	data := model.NewMyTypedData(address.Hex())

	pubkey, err := crypto.RecoverTypedSignature(
		data.SignatureHashBytes(),
		common.FromHex(sigHex),
	)
	if err != nil {
		return "", err
	}

	if crypto.PubkeyToAddressHex(pubkey) != address.Hex() {
		return "", model.ErrInvalidSignature
	}

	// TODO: create JWT
	token := "success"

	return token, nil
}

func (app *AuthApplication) createUser(address common.Address) (*model.User, error) {
	user := model.NewUser(address)
	user.UpdateChallenge()

	if err := app.userRepo.Add(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (app *AuthApplication) getUser(address common.Address) (*model.User, error) {
	return app.userRepo.Get(address)
}
