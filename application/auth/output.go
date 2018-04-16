package auth

type ChallengeOutput struct {
	Challenge string `json:"challenge"`
}

func NewChallengeOutput(challenge string) *ChallengeOutput {
	return &ChallengeOutput{
		Challenge: challenge,
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
