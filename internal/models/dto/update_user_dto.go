package dto

type UpdateUserRequest struct {
	UserName      string `json:"user_name" binding:"required,min=5"`
	Password      string `json:"password" binding:"required,min=5"`
	EmailNew      string `json:"email_new" binding:"required"`
	PasswordNew   string `json:"password_new" binding:"required,min=5"`
	SecretWordNew string `json:"secret_word_new" binding:"required,min=5"`
}

type UpdateUserResponse struct {
}
