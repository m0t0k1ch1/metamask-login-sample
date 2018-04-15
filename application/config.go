package application

import "time"

type Config struct {
	Auth *AuthConfig `json:"auth"`
}

type AuthConfig struct {
	Secret                     string `json:"secret"`
	TokenExpiryDurationSeconds int    `json:"token_expiry_duration_seconds"`
}

func (conf *AuthConfig) TokenExpiryDuration() time.Duration {
	return time.Duration(conf.TokenExpiryDurationSeconds) * time.Second
}
