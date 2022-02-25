package service

import (
	"myToDoApp/entities"
	"myToDoApp/internal/repository"
	"myToDoApp/internal/service/user"
)

type Authorization interface {
	CreateUser(user user.User) (int, error)
	GenerateToken(nickName, password string) (string, error)
	ParseToken(token string) (int, error)
}

type TodoList interface {
	Create(userId int, list entities.TodoList) (int, error)
	GetAll(userId int) ([]entities.TodoList, error)
	GetById(userId, listId int) (entities.TodoList, error)
	Delete(userId, listId int) error
	Update(userId, listId int, body entities.UpdateList) error
}

type TodoItem interface {
	Create(userId, listId int, item entities.TodoItem) (int, error)
	GetAll(userId, listId int) ([]entities.TodoItem, error)
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		TodoItem:      NewTodoItemService(repos.TodoItem, repos.TodoList),
		TodoList:      NewTodoListService(repos.TodoList),
	}
}
