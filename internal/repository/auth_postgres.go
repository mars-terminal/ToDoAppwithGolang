package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	"myToDoApp/internal/service/user"
)

type AuthPostgres struct {
	table string

	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB, table string) *AuthPostgres {
	return &AuthPostgres{db: db, table: table}
}

func (r *AuthPostgres) CreateUser(user user.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, nickname, password_hash) values ($1,$2,$3) RETURNING id", r.table)
	row := r.db.QueryRow(query, user.Name, user.NickName, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *AuthPostgres) GetUser(nickName, password string) (user.User, error) {
	var get user.User

	query := fmt.Sprintf(`SELECT id FROM %s WHERE nickname=$1 AND password_hash=$2`, userTable)
	err := r.db.Get(&get, query, nickName, password)
	if err != nil {
		return user.User{}, fmt.Errorf("failed to fetch user: %w", err)
	}

	return get, err
}
