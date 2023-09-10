package service

import (
	"github.com/SamsonAirapetyan/todo-app"
	"github.com/SamsonAirapetyan/todo-app/pkg/repository"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
	//CreateUser(user todo.User) error
}

type TodoList interface {
}

type TodoItem interface {
}

type Service struct {
	Authorization
	TodoItem
	TodoList
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
