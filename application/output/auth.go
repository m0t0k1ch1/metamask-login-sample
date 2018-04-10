package output

type AuthChallengeOutput struct {
	Challenge string `json:"challenge"`
}

func NewAuthChallengeOutput(challenge string) *AuthChallengeOutput {
	return &AuthChallengeOutput{
		Challenge: challenge,
	}
}

type AuthAuthorizeOutput struct {
	Token string `json:"token"`
}

func NewAuthAuthorizeOutput(token string) *AuthAuthorizeOutput {
	return &AuthAuthorizeOutput{
		Token: token,
	}
}
