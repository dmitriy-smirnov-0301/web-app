package enums

import (
	"time"
)

const (
	SecretTypePassword   = "password_hash"
	SecretTypeSecretWord = "secret_word_hash"
	TokenTypeAccess      = 0
	TokenTypeRefresh     = 1
	TokenAccessTTL       = 24 * time.Hour
	TokenRefreshTTL      = 5 * 24 * time.Hour
)
