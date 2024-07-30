package http

import (
	"alexdenkk/auth-api/internal/auth/service"
)

// Delivery - http delivery layer struct
type Delivery struct {
	Service *service.Service
}

// New - function for creating Delivery instance
func New(s *service.Service) *Delivery {
	return &Delivery{
		Service: s,
	}
}
