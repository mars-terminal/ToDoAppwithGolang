package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"

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
	query := fmt.Sprintf(`SELECT tl.id, tl.tittle, tl.description FROM %s tl 
					INNER JOIN %s ul on tl.id = ul.todo_id WHERE ul.user_id = $1 AND ul.todo_id = $2 `,
		todoListsTable, usersListsTable)
	err := r.db.Get(&list, query, userId, listId)
	return list, err
}
