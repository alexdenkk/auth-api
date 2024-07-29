package token

import (
	"github.com/golang-jwt/jwt/v5"
)

// GenerateJWT - generate JWT token
func GenerateJWT(claims Claims, key []byte) (string, error) {
	// создание токена
	t := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	// преобразовение токена в строку
	tokenString, err := t.SignedString(key)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ParseJWT - decode JWT token
func ParseJWT(token string, key []byte) (*Claims, error) {

	// парсинг токена
	t, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})

	// проверка на валидность
	if claims, ok := t.Claims.(*Claims); ok && t.Valid {
		return claims, nil
	}

	return &Claims{}, err
}
