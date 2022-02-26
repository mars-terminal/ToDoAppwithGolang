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
	query := fmt.Sprintf(`SELECT ti.id, ti.title, ti.description, ti.done FROM %s ti INNER JOIN %s li on li.item_id
	= ti.id INNER JOIN %s ul on ul.todo_id = li.todo_id WHERE li.todo_id=$1 AND ul.user_id=$2`,
		todoItemsTable, listItemsTable, usersListsTable)
	if err := t.db.Select(&items, query, listId, userId); err != nil {
		return nil, err
	}

	return items, nil
}

func (t *TodoItemPostgres) GetById(userId, itemId int) (entities.TodoItem, error) {
	var item entities.TodoItem
	query := fmt.Sprintf(`SELECT ti.id, ti.title, ti.description, ti.done FROM %s ti INNER JOIN %s li on li.item_id
	=ti.id INNER JOIN %s ul on ul.todo_id = li.todo_id WHERE ti.id=$1 AND ul.user_id=$2`,
		todoItemsTable, listItemsTable, usersListsTable)
	if err := t.db.Get(&item, query, itemId, userId); err != nil {
		return item, err
	}

	return item, nil
}

func (t *TodoItemPostgres) Delete(userId, itemId int) error {
	query := fmt.Sprintf(`DELETE FROM %s ti USING %s li, %s ul
	WHERE ti.id = li.item_id AND li.todo_id = ul.todo_id AND ul.user_id=$1 AND ti.id=$2`,
		todoItemsTable, listItemsTable, usersListsTable)

	_, err := t.db.Exec(query, userId, itemId)
	return err
}
