package application

import "time"

type Config struct {
	Auth *AuthConfig `json:"auth"`
}

type AuthConfig struct {
	Secret                string `json:"secret"`
	ExpiryDurationSeconds int    `json:"expiry_duration_seconds"`
}

func (conf *AuthConfig) ExpiryDuration() time.Duration {
	return time.Duration(conf.ExpiryDurationSeconds) * time.Second
}
