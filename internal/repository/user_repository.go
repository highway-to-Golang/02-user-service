package repository

import "user-service/internal/domain"

type UserRepository interface {
	Save(user domain.User) error

	FindByID(id string) (domain.User, error)

	FindAll() []domain.User

	DeleteByID(id string) error
}
