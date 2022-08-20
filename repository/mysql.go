package repository

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

const (
	usersTable      = "users"
	todoListsTable  = "todo_lists"
	usersListsTable = "users_lists"
	todoItemsTable  = "todo_items"
	listsItemsTable = "lists_items"
)

func NewMysqlDB(dataSourceName string) (*sqlx.DB, error) {
	db, err := sqlx.Connect("mysql", dataSourceName)

	if err != nil {
		return nil, err
	}

	return db, err
}
