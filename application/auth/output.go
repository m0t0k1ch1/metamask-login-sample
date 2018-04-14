package auth

import "github.com/m0t0k1ch1/metamask-login-sample/domain/auth"

type ChallengeOutput struct {
	Challenge string `json:"challenge"`
}

func NewChallengeOutput(challenge *auth.Challenge) *ChallengeOutput {
	return &ChallengeOutput{
		Challenge: challenge.Value,
	}
}

type AuthorizeOutput struct {
	Token string `json:"token"`
}

func NewAuthorizeOutput(token string) *AuthorizeOutput {
	return &AuthorizeOutput{
		Token: token,
	}
}
