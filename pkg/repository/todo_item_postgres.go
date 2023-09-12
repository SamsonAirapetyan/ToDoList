package repository

import (
	"fmt"
	"github.com/SamsonAirapetyan/todo-app"
	"github.com/jmoiron/sqlx"
)

type TodoItemPostgres struct {
	db *sqlx.DB
}

func NewTodoItemPostgres(db *sqlx.DB) *TodoItemPostgres {
	return &TodoItemPostgres{db: db}
}

func (r *TodoItemPostgres) Create(listid int, input todo.TodoItem) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var itemid int
	createItemQuery := fmt.Sprintf("INSERT INTO %s (title, description) values ($1, $2) RETURNING id", todoItemsTable)
	row := tx.QueryRow(createItemQuery, input.Title, input.Description)
	err = row.Scan(&itemid)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	createListQuery := fmt.Sprintf("INSERT INTO %s (list_id, item_id) values ($1, $2)", listItemsTable)
	_, err = tx.Exec(createListQuery, listid, itemid)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	return itemid, tx.Commit()
}

func (r *TodoItemPostgres) GetAll(userid, listid int) ([]todo.TodoItem, error) {
	var items []todo.TodoItem
	query := fmt.Sprintf("SELECT ti.id, ti.title, ti.description, ti.done FROM %s ti INNER JOIN %s li on li.item_id = ti.id INNER JOIN %s ul on ul.list_id = li.list_id WHERE li.list_id = $1 AND ul.user_id = $2", todoItemsTable, listItemsTable, usersListsTable)
	if err := r.db.Select(&items, query, listid, userid); err != nil {
		return nil, err
	}
	return items, nil
}
