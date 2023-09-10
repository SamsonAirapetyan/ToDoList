package repository

import (
	"fmt"
	"github.com/SamsonAirapetyan/todo-app"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user todo.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) values ($1, $2, $3) RETURNING id", usersTable)
	fmt.Println(query)
	row := r.db.QueryRow(query, user.Name, user.Username, user.Password)
	//row := r.db.QueryRow("INSERT INTO users (name, username, password_hash) values ($1, $2, $3) RETURNING id", user.Name, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		fmt.Printf(err.Error())
		return 0, err
	}
	return id, nil
}

//func (r *AuthPostgres) CreateUser(user todo.User) error {
//	query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) values ($1, $2, $3);", usersTable)
//	r.db.QueryRow(query, user.Name, user.Username, user.Password)
//	return nil
//}
