package repository

import (
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/korasdor/todo-app/models"
	"github.com/sirupsen/logrus"
)

type TodoListMysql struct {
	db *sqlx.DB
}

func NewTodoListMysql(db *sqlx.DB) *TodoListMysql {
	return &TodoListMysql{db: db}
}

func (r *TodoListMysql) Create(userId int, list models.TodoList) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	createListQuery := fmt.Sprintf("INSERT INTO %s (title, description) VALUES (?, ?)", todoListsTable)

	res, err := tx.Exec(createListQuery, list.Title, list.Description)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	createUsersListQuery := fmt.Sprintf("INSERT INTO %s (user_id, list_id) VALUES (?, ?)", usersListsTable)
	_, err = tx.Exec(createUsersListQuery, userId, id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return int(id), tx.Commit()

}

func (r *TodoListMysql) GetAll(userId int) ([]models.TodoList, error) {
	var lists []models.TodoList

	query := fmt.Sprintf("SELECT tl.id, tl.title, tl.description FROM %s tl INNER JOIN %s ul on tl.id = ul.list_id WHERE ul.user_id = ?", todoListsTable, usersListsTable)
	err := r.db.Select(&lists, query, userId)

	return lists, err
}

func (r *TodoListMysql) GetById(userId int, listId int) (models.TodoList, error) {
	var list models.TodoList

	query := fmt.Sprintf(`SELECT tl.id, tl.title, tl.description FROM %s tl 
							INNER JOIN %s ul on tl.id = ul.list_id WHERE ul.user_id = ? AND ul.list_id = ?`, todoListsTable, usersListsTable)
	err := r.db.Get(&list, query, userId, listId)

	return list, err
}

func (r *TodoListMysql) Delete(userId int, listId int) error {
	query := fmt.Sprintf(`DELETE FROM tl USING %s tl, %s ul WHERE tl.id = ul.list_id AND ul.user_id = ? AND ul.list_id = ?`, todoListsTable, usersListsTable)
	_, err := r.db.Exec(query, userId, listId)

	return err
}

func (r *TodoListMysql) Update(userId, listId int, input models.UpdateListInput) error {
	values := make([]string, 0)
	args := make([]interface{}, 0)

	if input.Title != nil {
		values = append(values, "title = ?")
		args = append(args, *input.Title)
	}

	if input.Descriptione != nil {
		values = append(values, "description = ?")
		args = append(args, *input.Descriptione)
	}

	setQuery := strings.Join(values, ", ")
	args = append(args, listId, userId)

	query := fmt.Sprintf(`UPDATE %s tl, %s ul SET %s WHERE tl.id = ul.list_id AND ul.list_id = ? AND ul.user_id = ?`, todoListsTable, usersListsTable, setQuery)
	logrus.Println(query)

	_, err := r.db.Exec(query, args...)

	return err
}
