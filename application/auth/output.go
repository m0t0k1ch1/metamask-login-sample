package auth

type ChallengeOutput struct {
	Challenge string `json:"challenge"`
}

func NewChallengeOutput(challengeStr string) *ChallengeOutput {
	return &ChallengeOutput{
		Challenge: challengeStr,
	}
}

type AuthorizeOutput struct {
	Token string `json:"token"`
}

func NewAuthorizeOutput(tokenStr string) *AuthorizeOutput {
	return &AuthorizeOutput{
		Token: tokenStr,
	}
}
