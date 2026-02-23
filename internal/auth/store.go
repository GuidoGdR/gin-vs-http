package auth

// store or repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/GuidoGdR/go-speed-test/internal/platform/appErr"
	"github.com/GuidoGdR/go-speed-test/internal/platform/models"
)

type Store interface {
	GetByUsernameActive(ctx context.Context, username string) (*models.User, error)
	Create(ctx context.Context, user *models.User) error
}

type store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) Store {
	return &store{db: db}
}

func (s *store) GetByUsernameActive(ctx context.Context, username string) (*models.User, error) {

	const q = `SELECT id, password, email, first_name, last_name, date_joined FROM users 
	WHERE username=$1 AND is_active=TRUE;`
	var u models.User

	if err := s.db.QueryRowContext(ctx, q, username).Scan(&u.ID, &u.Password, &u.Email, &u.FirstName, &u.LastName, &u.DateJoined); err != nil {

		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("%w: %v", appErr.NotFound, err)
		}
		return nil, err
	}

	u.Username = username
	u.IsActive = true

	return &u, nil
}

func (s *store) Create(ctx context.Context, user *models.User) error {
	const q = `
	INSERT INTO users (username, password, email, first_name, last_name) 
	VALUES ($1, $2, $3, $4, $5)
	RETURNING id, username, email, first_name, last_name, date_joined, is_active;`

	err := s.db.QueryRowContext(ctx, q, user.Username, user.Password, user.Email, user.FirstName, user.LastName).Scan(&user.ID, &user.Username, &user.Email, user.FirstName, user.LastName, &user.DateJoined, &user.IsActive)
	if err != nil {
		return fmt.Errorf("%w: %v", appErr.Internal, err)
	}

	return nil
}
