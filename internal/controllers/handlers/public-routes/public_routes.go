package public

import (
	auths "ice-creams-app/internal/services/auth-services"
)

type AuthHandler struct {
	svc auths.AuthorizationService
}

func NewAuthHandler(svc auths.AuthorizationService) *AuthHandler {
	return &AuthHandler{
		svc: svc,
	}
}
