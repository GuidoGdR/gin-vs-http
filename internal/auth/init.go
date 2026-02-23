package auth

import (
	"database/sql"

	"github.com/GuidoGdR/go-speed-test/pkg/token"
	"github.com/go-playground/validator/v10"
)

type authHTTP struct {
	Store   Store
	Service *Service
	Handler *hTTPHandler
}

func InitAuthHTTP(db *sql.DB, jwtManager *token.JWTManager, validate *validator.Validate) *authHTTP {
	store := NewStore(db)
	service := NewService(store, jwtManager)
	handler := NewHTTPHandler(service, validate)

	return &authHTTP{
		Store:   store,
		Service: service,
		Handler: handler,
	}
}

type authAdapter struct {
	Store   Store
	Service *Service
	Handler *AdapterHandler
}

func InitAuthAdapter(db *sql.DB, jwtManager *token.JWTManager, validate *validator.Validate) *authAdapter {
	store := NewStore(db)
	service := NewService(store, jwtManager)
	handler := NewAdapterHandler(service, validate)

	return &authAdapter{
		Store:   store,
		Service: service,
		Handler: handler,
	}
}

type authGin struct {
	Store   Store
	Service *Service
	Handler *ginHandler
}

func InitAuthGin(db *sql.DB, jwtManager *token.JWTManager) *authGin {
	store := NewStore(db)
	service := NewService(store, jwtManager)
	handler := NewGinHandler(service)

	return &authGin{
		Store:   store,
		Service: service,
		Handler: handler,
	}
}
