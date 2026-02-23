package auth

import (
	"context"
	"errors"
	"fmt"

	"github.com/GuidoGdR/go-speed-test/internal/platform/appErr"
	"github.com/GuidoGdR/go-speed-test/internal/platform/models"
	"github.com/GuidoGdR/go-speed-test/pkg/token"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	store        Store
	tokenManager *token.JWTManager
}

func NewService(store Store, tokenManager *token.JWTManager) *Service {
	return &Service{store, tokenManager}
}

func (s *Service) Login(ctx context.Context, username string, password string) (*loginResult, error) {

	usr, err := s.store.GetByUsernameActive(ctx, username)
	if err != nil {

		if errors.Is(err, appErr.NotFound) {
			return nil, fmt.Errorf("%w: %v", appErr.Unauthorized, err)
		}
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(usr.Password), []byte(password))
	if err != nil {
		return nil, fmt.Errorf("%w: Incorrect password: %v", appErr.Unauthorized, err)
	}

	access, err := s.tokenManager.NewAccessJWT(usr.ID)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", appErr.Internal, err)
	}

	refresh, err := s.tokenManager.NewRefreshJWT(usr.ID)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", appErr.Internal, err)
	}

	return &loginResult{User: usr, AccessTkn: access, RefreshTkn: refresh, TokenType: "Bearer"}, nil
}

func (s *Service) Refresh(ctx context.Context, refresh string) (*refreshResult, error) {
	claims, err := s.tokenManager.ValidateRefreshToken(refresh)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", appErr.Unauthorized, err)
	}
	access, err := s.tokenManager.NewAccessJWT(claims.UserID)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", appErr.Internal, err)
	}

	return &refreshResult{
		AccessTkn:  access,
		RefreshTkn: refresh,
	}, nil
}

func (s *Service) Register(ctx context.Context, username string, password string, email string) (*models.User, error) {
	usr := &models.User{
		Username: username,
		Password: password,
		Email:    email,
	}

	err := s.store.Create(ctx, usr)
	if err != nil {
		return nil, err
	}

	return usr, nil
}
