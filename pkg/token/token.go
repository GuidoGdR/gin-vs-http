package token

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTManager struct {
	secret     string
	secretByte []byte
}

func NewJWTManager(secret string) *JWTManager {
	return &JWTManager{secret: secret, secretByte: []byte(secret)}
}

func (m *JWTManager) ValidateAccessToken(tokenString string) (*jwtClaims, error) {
	result, err := m.validateToken(tokenString)
	if err != nil {
		return nil, err
	}

	if result.TokenType != "accesss" {
		return nil, fmt.Errorf("%w: Token type: %v", Errors.InvalidToken, result.TokenType)
	}

	return m.validateToken(tokenString)
}

func (m *JWTManager) ValidateRefreshToken(tokenString string) (*jwtClaims, error) {
	result, err := m.validateToken(tokenString)
	if err != nil {
		return nil, err
	}

	if result.TokenType != "refresh" {
		return nil, fmt.Errorf("%w: Obtined TokenType: %v", Errors.InvalidToken, result.TokenType)
	}

	return m.validateToken(tokenString)
}

func (m *JWTManager) validateToken(tokenString string) (*jwtClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signature method: %v", token.Header["alg"])
		}
		return m.secretByte, nil
	})
	if err != nil {
		return nil, fmt.Errorf("%w: %v", Errors.InvalidToken, err)
	}

	if claims, ok := token.Claims.(*jwtClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("%w: Invalid claims", Errors.InvalidToken)

}

func (m *JWTManager) NewAccessJWT(usr_id string) (string, error) {
	tkn, err := m.newJWT(usr_id, "access", time.Minute*15)
	if err != nil {
		err = fmt.Errorf("%w: Token type: access", err)
	}
	return tkn, err
}

func (m *JWTManager) NewRefreshJWT(usr_id string) (string, error) {
	tkn, err := m.newJWT(usr_id, "refresh", time.Hour*24)
	if err != nil {
		err = fmt.Errorf("%w: Token type: refresh", err)
	}
	return tkn, err
}

func (m *JWTManager) newJWT(usr_id string, ttype string, exp time.Duration) (string, error) {
	claims := jwtClaims{
		TokenType: ttype,

		UserID: usr_id,

		Exp: time.Now().Add(exp).Unix(),
		Iat: time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(m.secret)
	if err != nil {
		return "", fmt.Errorf("%w: %v", Errors.MakingToken, err)
	}

	return tokenString, nil
}
