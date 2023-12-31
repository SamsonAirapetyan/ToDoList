package service

import (
	"github.com/SamsonAirapetyan/todo-app"
	"github.com/SamsonAirapetyan/todo-app/pkg/repository"
)

type ToDoItemServise struct {
	repo     repository.TodoItem
	listRepo repository.TodoList
}

func NewToDoItemServise(repo repository.TodoItem, listRepo repository.TodoList) *ToDoItemServise {
	return &ToDoItemServise{repo: repo, listRepo: listRepo}
}

func (s *ToDoItemServise) Create(userid, listid int, input todo.TodoItem) (int, error) {
	_, err := s.listRepo.GetIdList(userid, listid)
	if err != nil {
		//list does not exist or does not belong user
		return 0, err
	}

	return s.repo.Create(listid, input)
}

func (s *ToDoItemServise) GetAll(userid, listid int) ([]todo.TodoItem, error) {
	return s.repo.GetAll(userid, listid)
}

func (s *ToDoItemServise) GetItemId(userid, itemid int) (todo.TodoItem, error) {
	return s.repo.GetItemId(userid, itemid)
}

func (s *ToDoItemServise) UpdateItem(userid, itemid int, input todo.UpdateItemInput) error {
	if err := input.ValidItem(); err != nil {
		return err
	}
	return s.repo.UpdateItem(userid, itemid, input)
}

func (s *ToDoItemServise) DeleteIdItem(userid, itemid int) error {
	return s.repo.DeleteIdItem(userid, itemid)
}
