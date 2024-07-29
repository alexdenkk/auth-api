package auth

import (
	"alexdenkk/auth-api/model"
	"context"
)

type Repository interface {
	GetUserByLogin(context.Context, string) (model.User, error)
}
