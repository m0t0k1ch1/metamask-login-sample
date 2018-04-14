package domain

import "github.com/m0t0k1ch1/metamask-login-sample/domain/repository"

type Container struct {
	NewUserRepository func() repository.User
}
