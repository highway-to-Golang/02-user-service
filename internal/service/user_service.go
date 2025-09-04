package service

import (
	"fmt"

	"user-service/internal/domain"
	"user-service/internal/repository"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) CreateUser(name, email, role string) (domain.User, error) {
	if !domain.IsValidRole(role) {
		return domain.User{}, fmt.Errorf("invalid role: %s. Valid roles are: %v", role, domain.ValidRoles())
	}

	user := domain.NewUser(name, email, role)

	if err := s.repo.Save(user); err != nil {
		return domain.User{}, fmt.Errorf("failed to save user: %w", err)
	}

	return user, nil
}

func (s *UserService) GetUser(id string) (domain.User, error) {
	if id == "" {
		return domain.User{}, fmt.Errorf("user ID cannot be empty")
	}

	user, err := s.repo.FindByID(id)
	if err != nil {
		return domain.User{}, fmt.Errorf("failed to get user: %w", err)
	}

	return user, nil
}

func (s *UserService) ListUsers() []domain.User {
	return s.repo.FindAll()
}

func (s *UserService) RemoveUser(id string) error {
	if id == "" {
		return fmt.Errorf("user ID cannot be empty")
	}

	if err := s.repo.DeleteByID(id); err != nil {
		return fmt.Errorf("failed to remove user: %w", err)
	}

	return nil
}
