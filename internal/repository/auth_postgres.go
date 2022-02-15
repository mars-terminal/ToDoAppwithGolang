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
