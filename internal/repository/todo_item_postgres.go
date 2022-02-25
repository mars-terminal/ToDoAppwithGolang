package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	"myToDoApp/entities"
)

type TodoItemPostgres struct {
	db *sqlx.DB
}

func NewTodoItemPostgres(db *sqlx.DB) *TodoItemPostgres {
	return &TodoItemPostgres{db: db}
}

func (t *TodoItemPostgres) Create(listId int, item entities.TodoItem) (int, error) {
	tx, err := t.db.Begin()
	if err != nil {
		return 0, err
	}

	var itemId int
	createItemQuery := fmt.Sprintf(`INSERT INTO %s(title, description) VALUES ($1, $2) RETURNING id`,
		todoItemsTable)
	row := tx.QueryRow(createItemQuery, item.Title, item.Description)
	err = row.Scan(&itemId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	createListItemQuery := fmt.Sprintf(`INSERT INTO %s (todo_id, item_id) VALUES ($1, $2)`,
		listItemsTable)
	_, err = tx.Exec(createListItemQuery, listId, itemId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return itemId, tx.Commit()
}

func (t *TodoItemPostgres) GetAll(userId, listId int) ([]entities.TodoItem, error) {
	var items []entities.TodoItem
	query := fmt.Sprintf(`SELECT tl.id, tl.title, tl.description, tl.done FROM %s ti INNER JOIN %s li on li.item_id= ti.id
	INNER JOIN %s ul on ul.todo_id = li.todo_id WHERE tl.todo_id = $1 AND ul.user_id =$2`,
		todoItemsTable, todoListsTable, usersListsTable)
	if err := t.db.Select(&items, query, listId, userId); err != nil {
		return nil, err
	}

	return items, nil
}
