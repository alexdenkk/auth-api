package repository

import (
	"alexdenkk/auth-api/model"
	"context"
)

// GetByLogin - метод для получения пользователя по логину
func (r *Repository) GetUserByLogin(ctx context.Context, login string) (model.User, error) {
	var user model.User
	result := r.DB.Where("login = ?", login).First(&user)
	return user, result.Error
}
