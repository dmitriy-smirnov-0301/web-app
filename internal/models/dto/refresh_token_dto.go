package dto

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required,gt=0"`
}

type RefreshTokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
