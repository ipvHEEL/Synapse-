package crypto

import (
	users "auth_service/models"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("2a$10$wZ5GGg5/FFHyS68QphdtFepdIBr0eCqMy8.sqyNI3VY95S7svRw4C")

type Claims struct {
	ID    uint   `json:"id"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func GenerateJWT(user users.User) (string, error) {
	expTime := time.Now().Add(24 * time.Hour)

	claims := &Claims{
		ID:    user.ID,
		Email: user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "auth_service",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err // теперь ошибка будет возвращена
	}
	return tokenString, nil
}

func ParseJWT(tokenStr string) (*Claims, error) {
	Claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, Claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, err
	}
	return Claims, nil
}
