package dto

type SignupUserRequest struct {
	UserName   string `json:"user_name" binding:"required,min=5"`
	Email      string `json:"email" binding:"required"`
	Password   string `json:"password" binding:"required,min=5"`
	SecretWord string `json:"secret_word" binding:"required,min=5"`
}

type SignupUserResponse struct {
	ID        int    `json:"id"`
	UserName  string `json:"user_name"`
	CreatedAt string `json:"created_at"`
}
