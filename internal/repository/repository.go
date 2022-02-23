package repository

import (
	"github.com/jmoiron/sqlx"

	"myToDoApp/internal/service/user"
)

type Authorization interface {
	CreateUser(user user.User) (int, error)
	GetUser(nickName, password string) (user.User, error)
}

type TodoList interface {
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
	}
}
