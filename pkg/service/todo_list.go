package service

import (
	"github.com/Kosodaka/todo-app"
	"github.com/Kosodaka/todo-app/pkg/repository"
)

type todoListService struct {
	repo repository.TodoList
}

func newTodoListService(repo repository.TodoList) *todoListService {
	return &todoListService{repo: repo}
}

func (s *todoListService) Create(userId int, list todo.TodoList) (int, error) {
	return s.repo.Create(userId, list)
}

func (s *todoListService) GetAll(userId int) ([]todo.TodoList, error) {
	return s.repo.GetAll(userId)
}

func (s *todoListService) GetById(userId, listId int) (todo.TodoList, error) {
	return s.repo.GetById(userId, listId)
}

func (s *todoListService) Delete(userId, listId int) error {
	return s.repo.Delete(userId, listId)
}

func (s *todoListService) Update(userId, listId int, input todo.UpdateListInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.Update(userId, listId, input)
}
