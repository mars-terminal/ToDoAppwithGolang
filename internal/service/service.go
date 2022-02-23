package service

import (
	"myToDoApp/internal/repository"
	"myToDoApp/internal/service/user"
)

type Authorization interface {
	CreateUser(user user.User) (int, error)
	GenerateToken(nickName, password string) (string, error)
	ParseToken(token string) (int, error)
}

type TodoList interface {
}

type TodoItem interface {
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization)}
}
