package services

import (
	"go-http-cleanArch/internal/entities"
	"strings"
)

type repo interface {
	AddUser(user *entities.User) (err error)
	GetAllUsers() (users *[]entities.User, err error)
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
	//БИЗНЕС ЛОГИКА
	user.Name = strings.ToUpper(user.Name)

	err = s.repo.AddUser(user)

	return err
}

func (s *service) GetAllUsers() (users *[]entities.User, err error) {
	users, err = s.repo.GetAllUsers()

	// КАКАЯ-ТО БИЗНЕС ЛОГИКА
	for _, v := range *users {
		v.Name = strings.ToLower(v.Name)
	}

	return users, err
}
