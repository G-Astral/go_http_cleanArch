package services

import (
	"go-http-cleanArch/internal/entities"
	"strings"
)

type repo interface {
	// AddUser(user *entities.User) (id int, err error)
	AddUser(user *entities.User) (err error)
}

type service struct {
	repo repo
}

func InitService(repo repo) service {
	return service{
		repo: repo,
	}
}

// func (s *service) AddUser(user *entities.User) (id int, err error) {
func (s *service) AddUser(user *entities.User) (err error) {
	//БИЗНЕС ЛОГИКА
	user.Name = strings.ToUpper(user.Name)

	// //repo
	// id, err = s.repo.AddUser(user)
	// if err != nil {
	// 	return 0, err
	// }

	// return id, err

	err = s.repo.AddUser(user)

	return err
}
