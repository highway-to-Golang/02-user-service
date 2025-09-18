package repository

import (
	"github.com/highway-to-Golang/02-user-service/internal/domain"
	"github.com/highway-to-Golang/02-user-service/internal/errors"
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
		return domain.User{}, errors.ErrorWithID(errors.ErrUserNotFound, id)
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
		return errors.ErrorWithID(errors.ErrUserNotFound, id)
	}

	delete(r.users, id)
	return nil
}
