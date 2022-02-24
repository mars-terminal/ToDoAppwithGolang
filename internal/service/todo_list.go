package service

import (
	"myToDoApp/entities"
	"myToDoApp/internal/repository"
)

type TodoListService struct {
	repo repository.TodoList
}

func NewTodoListService(repo repository.TodoList) *TodoListService {
	return &TodoListService{repo: repo}
}

func (t *TodoListService) Create(userId int, list entities.TodoList) (int, error) {
	return t.repo.Create(userId, list)
}

func (t *TodoListService) GetAll(userId int) ([]entities.TodoList, error) {
	return t.repo.GetAll(userId)
}

func (t *TodoListService) GetById(userId, listId int) (entities.TodoList, error) {
	return t.repo.GetById(userId, listId)
}

func (t *TodoListService) Delete(userId, listId int) error {
	return t.repo.Delete(userId, listId)
}

func (t *TodoListService) Update(userId, listId int, body entities.UpdateList) error {
	if err := body.Validate(); err != nil {
		return nil
	}
	return t.repo.Update(userId, listId, body)
}
