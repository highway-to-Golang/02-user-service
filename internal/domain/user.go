package domain

import (
	"slices"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
}

func NewUser(name, email, role string) User {
	return User{
		ID:        uuid.New().String(),
		Name:      name,
		Email:     email,
		Role:      role,
		CreatedAt: time.Now(),
	}
}

func ValidRoles() []string {
	return []string{"admin", "user", "guest"}
}

func IsValidRole(role string) bool {
	return slices.Contains(ValidRoles(), role)
}
