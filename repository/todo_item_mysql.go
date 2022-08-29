package repository

import (
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/korasdor/todo-app/models"
	"github.com/sirupsen/logrus"
)

type TodoItemMysql struct {
	db *sqlx.DB
}

func NewTodoItemMysql(db *sqlx.DB) *TodoItemMysql {
	return &TodoItemMysql{db: db}
}

func (r *TodoItemMysql) Create(listId int, item models.TodoItem) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	createItemQuery := fmt.Sprintf("INSERT INTO %s (title, description) VALUES (?, ?)", todoItemsTable)
	res, err := tx.Exec(createItemQuery, item.Title, item.Description)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	itemId, err := res.LastInsertId()
	if err != nil {
		tx.Rollback()
		return 0, nil
	}

	createListItemsQuery := fmt.Sprintf("INSERT INTO %s (list_id, item_id) VALUES (?, ?)", listsItemsTable)
	_, err = tx.Exec(createListItemsQuery, listId, itemId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return int(itemId), tx.Commit()
}

func (r *TodoItemMysql) GetAll(userId int, listId int) ([]models.TodoItem, error) {
	var items []models.TodoItem
	query := fmt.Sprintf(`SELECT ti.id, ti.title, ti.description, ti.done FROM %s ti INNER JOIN %s li on li.item_id = ti.id INNER JOIN %s ul on ul.list_id = li.list_id 
							WHERE li.list_id = ? AND ul.user_id = ?`, todoItemsTable, listsItemsTable, usersListsTable)

	if err := r.db.Select(&items, query, listId, userId); err != nil {
		return nil, err
	}

	return items, nil
}

func (r *TodoItemMysql) GetById(userId int, itemId int) (models.TodoItem, error) {
	var item models.TodoItem
	query := fmt.Sprintf(`SELECT ti.id, ti.title, ti.description, ti.done FROM %s ti INNER JOIN %s li on li.item_id = ti.id INNER JOIN %s ul on ul.list_id = li.list_id 
							WHERE ti.id = ? AND ul.user_id = ?`, todoItemsTable, listsItemsTable, usersListsTable)

	if err := r.db.Get(&item, query, itemId, userId); err != nil {
		return item, err
	}

	return item, nil
}

func (r *TodoItemMysql) Update(userId int, itemId int, input models.UpdateItemInput) error {
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

	if input.Done != nil {
		values = append(values, "done = ?")
		args = append(args, *input.Done)
	}

	setQuery := strings.Join(values, ", ")
	args = append(args, userId, itemId)

	query := fmt.Sprintf(`UPDATE %s ti, %s li, %s ul SET %s WHERE ti.id = li.item_id AND li.list_id = ul.list_id AND ul.user_id = ? AND ti.id = ?`,
		todoItemsTable, listsItemsTable, usersListsTable, setQuery)
	logrus.Println(query)

	_, err := r.db.Exec(query, args...)

	return err
}

func (r *TodoItemMysql) Delete(userId int, itemId int) error {
	query := fmt.Sprintf(`DELETE FROM ti USING %s ti, %s li, %s ul 
		WHERE ti.id = li.item_id AND li.list_id = ul.list_id AND ul.user_id = ? AND ti.id = ?`, todoItemsTable, listsItemsTable, usersListsTable)

	_, err := r.db.Exec(query, userId, itemId)
	return err
}
