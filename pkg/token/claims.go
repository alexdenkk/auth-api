package token

import (
	"github.com/golang-jwt/jwt/v5"
)

// Claims - JWT token struct
type Claims struct {
	*jwt.RegisteredClaims
	ID                uint   `json:"id"`
	Login             string `json:"login"`
	ConstructTariffID uint8  `json:"construct_tariff_id"`
}
