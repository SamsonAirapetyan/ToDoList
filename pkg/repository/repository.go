package repository

import (
	"github.com/SamsonAirapetyan/todo-app"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GetUser(username, password string) (todo.User, error)
}

type TodoList interface {
	Create(userId int, list todo.TodoList) (int, error)
	GetAllLists(userId int) ([]todo.TodoList, error)
	GetIdList(userid int, listId int) (todo.TodoList, error)
	DeleteIdList(userid, listid int) error
	Update(userid int, listId int, input todo.UpdateListInput) error
}

type TodoItem interface {
	Create(listid int, input todo.TodoItem) (int, error)
	GetAll(userid, listid int) ([]todo.TodoItem, error)
	GetItemId(userid, itemid int) (todo.TodoItem, error)
	UpdateItem(userid, itemid int, input todo.UpdateItemInput) error
	DeleteIdItem(userid, itemid int) error
}

type Repository struct {
	Authorization
	TodoItem
	TodoList
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		TodoList:      NewTodoListPostgres(db),
		TodoItem:      NewTodoItemPostgres(db),
	}
}
