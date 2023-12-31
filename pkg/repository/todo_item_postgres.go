package repository

import (
	"fmt"
	"github.com/SamsonAirapetyan/todo-app"
	"github.com/jmoiron/sqlx"
	"strings"
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

func (r *TodoItemPostgres) GetItemId(userid, itemid int) (todo.TodoItem, error) {
	var item todo.TodoItem
	query := fmt.Sprintf(`SELECT ti.id, ti.title, ti.description, ti.done FROM %s ti INNER JOIN %s li on li.item_id = ti.id INNER JOIN %s ul on ul.list_id = li.list_id WHERE ti.id = $1 AND ul.user_id = $2`, todoItemsTable, listItemsTable, usersListsTable)
	err := r.db.Get(&item, query, itemid, userid)
	return item, err
}

func (r *TodoItemPostgres) UpdateItem(userid, itemid int, input todo.UpdateItemInput) error {
	setValue := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Title != nil {
		setValue = append(setValue, fmt.Sprintf("title = $%d", argId))
		args = append(args, *input.Title)
		argId++
	}

	if input.Description != nil {
		setValue = append(setValue, fmt.Sprintf("description = $%d", argId))
		args = append(args, *input.Description)
		argId++
	}
	if input.Done != nil {
		setValue = append(setValue, fmt.Sprintf("done = $%d", argId))
		args = append(args, *input.Done)
		argId++
	}
	setQuery := strings.Join(setValue, ", ")
	query := fmt.Sprintf("UPDATE %s ti SET %s FROM %s li, %s ul WHERE ti.id = li.item_id AND li.list_id = ul.list_id AND ul.user_id = $%d AND ti.id = $%d", todoItemsTable, setQuery, listItemsTable, usersListsTable, argId, argId+1)

	args = append(args, userid, itemid)

	_, err := r.db.Exec(query, args...)
	return err
}

func (r *TodoItemPostgres) DeleteIdItem(userid, itemid int) error {
	query := fmt.Sprintf(`DELETE FROM %s ti USING %s li, %s ul WHERE ti.id = li.item_id AND li.list_id = ul.list_id AND ul.user_id = $1 AND ti.id = $2`, todoItemsTable, listItemsTable, usersListsTable)
	_, err := r.db.Exec(query, userid, itemid)

	return err
}
