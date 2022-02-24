package repository

import (
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"

	"myToDoApp/entities"
)

type TodoListRepository struct {
	db *sqlx.DB
}

func NewTodoListRepository(db *sqlx.DB) *TodoListRepository {
	return &TodoListRepository{db: db}
}

func (r *TodoListRepository) Create(userId int, list entities.TodoList) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	createQueryList := fmt.Sprintf(`INSERT INTO %s (title, description) 
									VALUES ($1, $2) RETURNING id`, todoListsTable)
	row := tx.QueryRow(createQueryList, list.Title, list.Description)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	createUserListQuery := fmt.Sprintf(`INSERT INTO %s (user_id, todo_id)
										VALUES ($1, $2)`, usersListsTable)
	_, err = tx.Exec(createUserListQuery, userId, id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	return id, tx.Commit()
}

func (r *TodoListRepository) GetAll(userId int) ([]entities.TodoList, error) {
	var lists []entities.TodoList
	query := fmt.Sprintf(`SELECT tl.id, tl.title, tl.description FROM %s tl 
						INNER JOIN %s ul on tl.id = ul.todo_id WHERE ul.user_id = $1 `,
		todoListsTable, usersListsTable)
	err := r.db.Select(&lists, query, userId)
	return lists, err
}

func (r *TodoListRepository) GetById(userId, listId int) (entities.TodoList, error) {
	var list entities.TodoList
	query := fmt.Sprintf(`SELECT tl.id, tl.title, tl.description FROM %s tl 
					INNER JOIN %s ul on tl.id = ul.todo_id WHERE ul.user_id = $1 AND ul.todo_id = $2 `,
		todoListsTable, usersListsTable)
	err := r.db.Get(&list, query, userId, listId)
	return list, err
}

func (r *TodoListRepository) Delete(userId, listId int) error {
	query := fmt.Sprintf(`DELETE FROM %s tl USING %s ul WHERE 
				tl.id = ul.todo_id AND ul.user_id = $1 AND ul.todo_id = $2`,
		todoListsTable, usersListsTable)

	_, err := r.db.Exec(query, userId, listId)

	return err

}

func (r *TodoListRepository) Update(userId, listId int, body entities.UpdateList) error {

	setValue := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if body.Title != nil {
		setValue = append(setValue, fmt.Sprintf("title = $%d", argId))
		args = append(args, *body.Title)
		argId++
	}

	if body.Description != nil {
		setValue = append(setValue, fmt.Sprintf("description = $%d", argId))
		args = append(args, *body.Description)
		argId++
	}

	setQuery := strings.Join(setValue, ", ")

	query := fmt.Sprintf(`UPDATE %s tl SET %s FROM %s ul WHERE 
			tl.id = ul.todo_id AND ul.todo_id=$%d AND ul.user_id=$%d`,
		todoListsTable, setQuery, usersListsTable, argId, argId+1)

	args = append(args, listId, userId)

	logrus.Debugf("updateQuery: %s", query)
	logrus.Debugf("args: %s", args)

	_, err := r.db.Exec(query, args...)
	return err
}
