package repository

import (
	"database/sql"
	"go-http-cleanArch/internal/entities"
)

type repo struct {
	db *sql.DB
}

func InitRepo(db *sql.DB) repo {
	return repo{
		db: db,
	}
}

// func(r *repo) AddUser(user *entities.User) (id int, err error) {
func (r *repo) AddUser(user *entities.User) (err error) {
	query := "INSERT INTO users (name, age) VALUES ($1, $2) RETURNING id"
	err = r.db.QueryRow(query, user.Name, user.Age).Scan(&user.Id)
	if err != nil {

		return err
	}

	return nil
}
