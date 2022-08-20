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
	query := fmt.Sprintf(`INSERT INTO %s (name, username, password_hash) values ("%s", "%s", "%s")`, usersTable, user.Name, user.Username, user.Password)

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return 0, err
	}

	res, err := stmt.Exec()
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (r *AuthMysql) CreateUser(user models.User) (int, error) {

}
