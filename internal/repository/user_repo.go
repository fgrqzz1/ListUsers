package repository

import (
	"ListUsers/internal/models"
)

type UserRepository interface {
	GetUsers([]models.User, error)
}

type InMemUserRepository struct{}

func (r *InMemUserRepository) GetUsers() ([]models.User, error) {
	users := []models.User{
		{ID: 1, Name: "Алексей", Email: "alex@example.com"},
		{ID: 2, Name: "Мария", Email: "maria@example.com"},
		{ID: 3, Name: "Иван", Email: "ivan@example.com"},
	}

	return users, nil
}
