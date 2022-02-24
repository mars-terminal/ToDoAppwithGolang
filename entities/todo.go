package entities

import (
	"errors"
)

type TodoList struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title" binding:"required"`
	Description string `json:"description" db:"description"`
}

type UserList struct {
	Id     int
	UserId string
	ListId string
}

type TodoItem struct {
	Id          int    `json:"id"`
	Tittle      string `json:"tittle"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

type ListItem struct {
	Id     int
	ListId int
	ItemId int
}

type UpdateList struct {
	Title       *string `json:"title"`
	Description *string `json:"description" `
}

func (i UpdateList) Validate() error {
	if i.Title == nil && i.Description == nil {
		return errors.New("value is empty")
	}
	return nil
}
