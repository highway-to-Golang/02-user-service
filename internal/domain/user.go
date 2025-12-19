package domain

import (
	"slices"
	"time"

	"github.com/highway-to-Golang/02-user-service/internal/errors"
	"github.com/google/uuid"
)

type User struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
}

func NewUser(name, email, role string) (User, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return User{}, errors.ErrFailedToBuild
	}
	return User{
		ID:        id.String(),
		Name:      name,
		Email:     email,
		Role:      role,
		CreatedAt: time.Now(),
	}, nil
}

func ValidRoles() []string {
	return []string{"admin", "user", "guest"}
}

func IsValidRole(role string) bool {
	return slices.Contains(ValidRoles(), role)
}
