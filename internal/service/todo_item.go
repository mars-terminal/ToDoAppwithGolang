package service

import (
	"myToDoApp/entities"
	"myToDoApp/internal/repository"
)

type TodoItemService struct {
	repo     repository.TodoItem
	listRepo repository.TodoList
}

func NewTodoItemService(repo repository.TodoItem, listRepo repository.TodoList) *TodoItemService {
	return &TodoItemService{repo: repo, listRepo: listRepo}
}

func (t *TodoItemService) Create(userId, listId int, item entities.TodoItem) (int, error) {
	_, err := t.listRepo.GetById(userId, listId)
	if err != nil {
		return 0, err
	}

	return t.repo.Create(listId, item)
}

func (t *TodoItemService) GetAll(userId, listId int) ([]entities.TodoItem, error) {
	return t.repo.GetAll(userId, listId)
}
