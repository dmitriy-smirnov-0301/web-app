package jwt

import (
	"errors"
	"ice-creams-app/internal/models/domain"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	jwtSecret = []byte("secret_key")
)

func GenerateToken(userID interface{}, duration time.Duration, tokenType int) (string, error) {

	var uid uint32

	switch v := userID.(type) {
	case int:
		uid = uint32(v)
	case int32:
		uid = uint32(v)
	case uint32:
		uid = v
	case int64:
		uid = uint32(v)
	case uint64:
		uid = uint32(v)
	default:
		return "", errors.New("unsupported user ID type")
	}

	claims := jwt.MapClaims{
		"sub":        uid,
		"iat":        time.Now().Unix(),
		"exp":        time.Now().Add(duration).Unix(),
		"token_type": tokenType,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(jwtSecret)

}

func ParseToken(tokenStr string) (*domain.Token, error) {

	token, err := jwt.ParseWithClaims(tokenStr, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return jwtSecret, nil
	})

	if err != nil {
		return nil, errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token claims")
	}

	tknType, ok := claims["token_type"].(float64)
	if !ok {
		return nil, errors.New("token type is missing or invalid")
	}
	tokenType := int(tknType)

	iat, ok := claims["iat"].(float64)
	if !ok || time.Now().Before(time.Unix(int64(iat), 0)) {
		return nil, errors.New("invalid token")
	}
	createdAt := time.Unix(int64(iat), 0)

	exp, ok := claims["exp"].(float64)
	if !ok || time.Now().After(time.Unix(int64(exp), 0)) {
		return nil, errors.New("token has expired")
	}
	expiresAt := time.Unix(int64(exp), 0)

	uid, ok := claims["sub"].(float64)
	if !ok {
		return nil, errors.New("invalid user ID in token")
	}
	userID := uint32(uid)

	parsedToken := &domain.Token{
		UserID:    userID,
		TokenType: tokenType,
		ExpiresAt: expiresAt,
		CreatedAt: createdAt,
		IsValid:   true,
	}

	return parsedToken, nil

}
