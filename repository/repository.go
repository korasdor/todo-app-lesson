package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/korasdor/todo-app/models"
)

type Autharization interface {
	CreateUser(user models.User) (int, error)
	GetUser(username, password string) (models.User, error)
}

type TodoList interface {
}

type TodoItem interface {
}

type Repository struct {
	Autharization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Autharization: NewAuthMysql(db),
	}
}
