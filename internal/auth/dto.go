package auth

import "github.com/GuidoGdR/go-speed-test/internal/platform/models"

// login
type loginRequest struct {
	Username string `json:"username" binding:"required,min=4,max=20" validate:"required,min=4,max=20"`
	Password string `json:"password" binding:"required,min=8,max=255" validate:"required,min=8,max=255"`
}

type loginResult struct {
	AccessTkn  string
	RefreshTkn string
	TokenType  string
	User       *models.User
}

type loginResponse struct {
	Access    string       `json:"access"`
	Refresh   string       `json:"refresh"`
	TokenType string       `json:"token_type"`
	User      *models.User `json:"user"`
}

// Refresh
type refreshRequest struct {
	Refresh string `json:"refresh" binding:"required,min=4" validate:"required,min=4"`
}

type refreshResult struct {
	AccessTkn  string
	RefreshTkn string
}

type refreshResponse struct {
	Access  string `json:"access"`
	Refresh string `json:"refresh"`
}

// Register
type registerRequest struct {
	Username string `json:"username" binding:"required,min=4,max=100" validate:"required,min=4,max=100"`
	Password string `json:"password" binding:"required,min=8,max=255" validate:"required,min=8,max=255"`
	Email    string `json:"email" binding:"required,email,max=255" validate:"required,email,max=255"`
}
