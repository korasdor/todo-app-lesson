package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/korasdor/todo-app/models"
)

type AuthMysql struct {
	db *sqlx.DB
}

func NewAuthMysql(db *sqlx.DB) *AuthMysql {
	return &AuthMysql{db: db}
}

func (r *AuthMysql) CreateUser(user models.User) (int, error) {
	query := fmt.Sprintf(`INSERT INTO %s (name, username, password_hash) values (?, ?, ?)`, usersTable)

	res, err := r.db.Exec(query, user.Name, user.Username, user.Password)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (r *AuthMysql) GetUser(username string, password string) (models.User, error) {
	var user models.User
	query := fmt.Sprintf(`SELECT id FROM %s WHERE username="%s" AND password_hash="%s"`, usersTable, username, password)

	err := r.db.Get(&user, query)

	return user, err
}
