package models

import (
	"github.com/dgrijalva/jwt-go"
)

// Credentials struct para armazenar o cpf e secret no corpo do request
type Credentials struct {
	CPF    string `json:"cpf" validate:"required"`
	Secret string `json:"secret" validate:"required"`
}

// Claims struct que será criptografado em um token JWT
type Claims struct {
	CPF string `json:"cpf" validate:"required"`
	jwt.StandardClaims
}
