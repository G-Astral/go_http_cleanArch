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

func (r *repo) GetUserByID(id int) (user *entities.User, err error) {
	user = &entities.User{}
	query := "SELECT * FROM users WHERE id = $1"

	err = r.db.QueryRow(query, id).Scan(&user.Id, &user.Name, &user.Age)
	if err != nil {
		return nil, err
	}

	return user, err
}

func (r *repo) DelUserById(id int) (rowsAffected int64, err error) {
	query := "DELETE FROM users WHERE id = $1"
	res, err := r.db.Exec(query, id)
	if err != nil {
		return rowsAffected, err
	}

	rowsAffected, err = res.RowsAffected()
	if err != nil {
		return rowsAffected, err
	}

	return rowsAffected, nil
}

func (r *repo) UpdUserById(user *entities.User, id int) (err error) {
	query := "UPDATE users SET name = $1, age = $2 WHERE id = $3"

	res, err := r.db.Exec(query, user.Name, user.Age, id)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected > 0 {
		user.Id = id
		return nil
	} else {
		return sql.ErrNoRows
	}
}
