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

func (r *repo) AddUser(user *entities.User) (err error) {
	query := "INSERT INTO users (name, age) VALUES ($1, $2) RETURNING id"
	err = r.db.QueryRow(query, user.Name, user.Age).Scan(&user.Id)
	if err != nil {
		return err
	}

	return nil
}

func (r *repo) GetAllUsers() (users *[]entities.User, err error) {
	query := "SELECT * FROM users"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var usersSlice []entities.User

	for rows.Next() {
		var u entities.User
		if err := rows.Scan(&u.Id, &u.Name, &u.Age); err != nil {
			return nil, err
		}

		usersSlice = append(usersSlice, u)
	}

	users = &usersSlice

	return users, nil
}
