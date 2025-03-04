package domain

import "time"

type Token struct {
	ID           int
	UserID       uint32
	UserName     string
	TokenType    int
	AccessToken  string
	RefreshToken string
	ExpiresAt    time.Time
	CreatedAt    time.Time
	RevokedAt    time.Time
	IsValid      bool
}
