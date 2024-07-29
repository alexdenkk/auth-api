package auth

import (
	"context"
)

type Service interface {
	Login(context.Context, string, string) (string, error)
}
