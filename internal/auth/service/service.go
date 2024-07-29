package service

import (
	"alexdenkk/auth-api/internal/auth"
)

// Service - структура слоя сервиса
type Service struct {
	SignKey    []byte
	Repository auth.Repository
}

// New - функция для создания экземпляра структуры Service
func New(repo auth.Repository, key []byte) *Service {
	return &Service{
		Repository: repo,
		SignKey:    key,
	}
}
