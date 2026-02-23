package models

type User struct {
	ID        string `json:"id"`
	Username  string `json:"username" binding:"min=4,max=100,required" validate:"min=4,max=100,required"`
	Password  string `json:"password" binding:"min=8,max=255,required" validate:"min=8,max=255,required"`
	Email     string `json:"email" binding:"email,max=255,required" validate:"email,max=255,required"`
	FirstName string `json:"first_name" binding:"max=100" validate:"max=100"`
	LastName  string `json:"last_name" binding:"max=100" validate:"max=100"`

	IsActive bool `json:"is_active"`

	DateJoined string `json:"date_joined"`
}
