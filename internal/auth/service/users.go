package service

import (
	"alexdenkk/auth-api/internal/auth/entity"
	"alexdenkk/auth-api/pkg/hash"
	"alexdenkk/auth-api/pkg/token"
	"context"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Login - user login function
func (s *Service) Login(
	ctx context.Context, request entity.LoginRequest,
) entity.LoginResponse {

	// getting user by login
	response := s.Repository.GetUserByLogin(ctx, request)

	if response.Err != nil {
		response.Err = errors.New("user not found")
		return response
	}

	// checking password
	if response.User.Password != hash.Hash(request.Password) {
		response.Err = errors.New("incorrect password")
		return response
	}

	// generating token
	claims := token.Claims{
		ID:    response.User.ID,
		Login: response.User.Login,

		RegisteredClaims: &jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 1000)),
		},
	}

	tkn, err := token.GenerateJWT(claims, s.JwtSignKey)

	if err != nil {
		response.Err = err
		return response
	}

	response.Token = tkn
	return response
}
