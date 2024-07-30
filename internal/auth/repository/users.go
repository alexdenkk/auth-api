package repository

import (
	"alexdenkk/auth-api/internal/auth/entity"
	"alexdenkk/auth-api/model"
	"context"
)

// GetByLogin - function for getting user by login
func (r *Repository) GetUserByLogin(
	ctx context.Context, request entity.LoginRequest,
) entity.LoginResponse {

	var user model.User

	result := r.DB.Where("login = ?", request.Login).First(&user)

	return entity.LoginResponse{
		User: user,
		Err:  result.Error,
	}
}
