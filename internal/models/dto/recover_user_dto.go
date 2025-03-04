package dto

type RecoverUserRequest struct {
	UserName    string `json:"user_name" binding:"required,min=5"`
	SecretWord  string `json:"secret_word" binding:"required,min=5"`
	PasswordNew string `json:"password_new" binding:"required,min=5"`
}

type RecoverUserResponse struct {
}
