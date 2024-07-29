package service

import (
	"context"

	"alexdenkk/auth-api/internal/auth"
	"alexdenkk/auth-api/pkg/hash"
	"alexdenkk/auth-api/pkg/token"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Login - логин пользователя
func (s *Service) Login(ctx context.Context, login, password string) (string, error) {
	// получение пользователя по логину
	user, err := s.Repository.GetUserByLogin(ctx, login)

	if err != nil {
		return "", auth.ErrUserNotFound
	}

	// проверка пароля
	if user.Password != hash.Hash(password) {
		return "", auth.ErrIncorrectPassword
	}

	// генерация токена
	claims := token.Claims{
		ID:    user.ID,
		Login: user.Login,

		RegisteredClaims: &jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 1000)),
		},
	}

	tkn, err := token.GenerateJWT(claims, s.SignKey)

	if err != nil {
		return "", err
	}

	return tkn, nil
}
