package service

import (
	"github.com/SamsonAirapetyan/todo-app"
	"github.com/SamsonAirapetyan/todo-app/pkg/repository"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(accesstoken string) (int, error)
}

type TodoList interface {
	Create(userId int, list todo.TodoList) (int, error)
	GetAllLists(userId int) ([]todo.TodoList, error)
	GetIdList(userid int, listId int) (todo.TodoList, error)
	DeleteIdList(userid, listid int) error
	Update(userid int, listId int, input todo.UpdateListInput) error
}

type TodoItem interface {
	Create(userid, listid int, input todo.TodoItem) (int, error)
	GetAll(userid, listid int) ([]todo.TodoItem, error)
}

type Service struct {
	Authorization
	TodoItem
	TodoList
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		TodoList:      NewToDoListServise(repos.TodoList),
		TodoItem:      NewToDoItemServise(repos.TodoItem, repos.TodoList),
	}
}
