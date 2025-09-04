package repository

import (
	"fmt"

	"user-service/internal/domain"
)

type InMemoryUserRepo struct {
	users map[string]domain.User
}

func NewInMemoryUserRepo() *InMemoryUserRepo {
	return &InMemoryUserRepo{
		users: make(map[string]domain.User),
	}
}

func (r *InMemoryUserRepo) Save(user domain.User) error {
	r.users[user.ID] = user
	return nil
}

func (r *InMemoryUserRepo) FindByID(id string) (domain.User, error) {
	user, ok := r.users[id]
	if !ok {
		return domain.User{}, fmt.Errorf("user with ID %s not found", id)
	}

	return user, nil
}

func (r *InMemoryUserRepo) FindAll() []domain.User {
	users := make([]domain.User, 0, len(r.users))
	for _, user := range r.users {
		users = append(users, user)
	}

	return users
}

func (r *InMemoryUserRepo) DeleteByID(id string) error {
	if _, ok := r.users[id]; !ok {
		return fmt.Errorf("user with ID %s not found", id)
	}

	delete(r.users, id)
	return nil
}
