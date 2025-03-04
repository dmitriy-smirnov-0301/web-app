package dto

type LoginUserRequest struct {
	UserName string `json:"user_name" binding:"required,min=5"`
	Password string `json:"password" binding:"required,min=5"`
}

type LoginUserResponse struct {
	UserName     string `json:"user_name"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
