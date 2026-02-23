package token

import "github.com/golang-jwt/jwt/v5"

type jwtClaims struct {
	TokenType string `json:"token_type"`

	UserID string `json:"user_id"`

	Exp int64 `json:"exp"` // expiration time
	Iat int64 `json:"iat"` // emition time

	jwt.RegisteredClaims
}
