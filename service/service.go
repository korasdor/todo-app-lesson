package service

import (
	"github.com/korasdor/todo-app/models"
	"github.com/korasdor/todo-app/repository"
)

type Autharization interface {
	CreateUser(user models.User) (int, error)
	GenerateToken(username string, password string) (int, error)
}

type TodoList interface {
}

type TodoItem interface {
}

type Service struct {
	Autharization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Autharization: NewAuthService(repos),
	}
}