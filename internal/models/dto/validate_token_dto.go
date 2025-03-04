package dto

import "time"

type ValidateTokenRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required,gt=0"`
}

type ValidateTokenResponse struct {
	UserName  string    `json:"user_name"`
	ExpiresAt time.Time `json:"expires_at"`
	CreatedAt time.Time `json:"created_at"`
	IsValid   bool      `json:"is_valid"`
}
