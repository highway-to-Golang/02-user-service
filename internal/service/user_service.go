package service

import (
	"fmt"

	"github.com/highway-to-Golang/02-user-service/internal/domain"
	"github.com/highway-to-Golang/02-user-service/internal/errors"
)

type UserRepository interface {
	Save(user domain.User) error

	FindByID(id string) (domain.User, error)

	FindAll() []domain.User

	DeleteByID(id string) error
}

type UserService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) CreateUser(name, email, role string) (domain.User, error) {
	if !domain.IsValidRole(role) {
		return domain.User{}, fmt.Errorf("%w: %s. Valid roles are: %v", errors.ErrInvalidRole, role, domain.ValidRoles())
	}

	user, err := domain.NewUser(name, email, role)

	if err != nil {
		return domain.User{}, fmt.Errorf("%w: %w", errors.ErrFailedToSave, err)
	}

	err = s.repo.Save(user);
	if  err != nil {
		return domain.User{}, fmt.Errorf("%w: %w", errors.ErrFailedToSave, err)
	}

	return user, nil
}

func (s *UserService) GetUser(id string) (domain.User, error) {
	if id == "" {
		return domain.User{}, errors.ErrUserIDEmpty
	}

	user, err := s.repo.FindByID(id)
	if err != nil {
		return domain.User{}, fmt.Errorf("%w: %w", errors.ErrFailedToGet, err)
	}

	return user, nil
}

func (s *UserService) ListUsers() []domain.User {
	return s.repo.FindAll()
}

func (s *UserService) RemoveUser(id string) error {
	if id == "" {
		return errors.ErrUserIDEmpty
	}

	if err := s.repo.DeleteByID(id); err != nil {
		return fmt.Errorf("%w: %w", errors.ErrFailedToRemove, err)
	}

	return nil
}
