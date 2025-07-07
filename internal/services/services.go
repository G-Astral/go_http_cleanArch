package services

import (
	"go-http-cleanArch/internal/entities"
	// "strings"
)

type repo interface {
	AddUser(user *entities.User) (err error)
	GetAllUsers() (users *[]entities.User, err error)
	GetUserByID(id int) (user *entities.User, err error)
	DelUserById(id int) (rowsAffected int64, err error)
}

type service struct {
	repo repo
}

func InitService(repo repo) service {
	return service{
		repo: repo,
	}
}

func (s *service) AddUser(user *entities.User) (err error) {
	// БИЗНЕС ЛОГИКА
	// user.Name = strings.ToUpper(user.Name)

	err = s.repo.AddUser(user)

	return err
}

func (s *service) GetAllUsers() (users *[]entities.User, err error) {
	users, err = s.repo.GetAllUsers()

	// КАКАЯ-ТО БИЗНЕС ЛОГИКА

	return users, err
}

func (s *service) GetUserByID(id int) (user *entities.User, err error) {
	user, err = s.repo.GetUserByID(id)

	// КАКАЯ-ТО БИЗНЕС ЛОГИКА

	return user, err
}

func (s *service) DelUserById(id int) (rowsAffected int64, err error) {
	rowsAffected, err = s.repo.DelUserById(id)

	return rowsAffected, err
}
