package repository

import (
	"github.com/jmoiron/sqlx"

	"myToDoApp/entities"
	"myToDoApp/internal/service/user"
)

type Authorization interface {
	CreateUser(user user.User) (int, error)
	GetUser(nickName, password string) (user.User, error)
}

type TodoList interface {
	Create(userId int, list entities.TodoList) (int, error)
	GetAll(userId int) ([]entities.TodoList, error)
	GetById(userId, listId int) (entities.TodoList, error)
	Delete(userId, listId int) error
}

type TodoItem interface {
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db, "users"),
		TodoList:      NewTodoListRepository(db),
	}
}
