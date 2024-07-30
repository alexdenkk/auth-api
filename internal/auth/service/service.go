package service

import (
	"alexdenkk/auth-api/internal/auth/repository"
)

// Service - service layer struct
type Service struct {
	JwtSignKey []byte
	Repository *repository.Repository
}

// New - function for creating Service instance
func New(repo *repository.Repository, key []byte) *Service {
	return &Service{
		Repository: repo,
		JwtSignKey: key,
	}
}
