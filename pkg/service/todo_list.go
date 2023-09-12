package service

import (
	"github.com/SamsonAirapetyan/todo-app"
	"github.com/SamsonAirapetyan/todo-app/pkg/repository"
)

type ToDoListServise struct {
	repo repository.TodoList
}

func NewToDoListServise(repo repository.TodoList) *ToDoListServise {
	return &ToDoListServise{repo: repo}
}

func (s *ToDoListServise) Create(userId int, list todo.TodoList) (int, error) {
	return s.repo.Create(userId, list)
}

func (s *ToDoListServise) GetAllLists(userId int) ([]todo.TodoList, error) {
	return s.repo.GetAllLists(userId)
}

func (s *ToDoListServise) GetIdList(userid int, listId int) (todo.TodoList, error) {
	return s.repo.GetIdList(userid, listId)
}

func (s *ToDoListServise) DeleteIdList(userid, listid int) error {
	return s.repo.DeleteIdList(userid, listid)
}

func (s *ToDoListServise) Update(userid int, listId int, input todo.UpdateListInput) error {
	err := input.ValidList()
	if err != nil {
		return err
	}
	return s.repo.Update(userid, listId, input)
}
