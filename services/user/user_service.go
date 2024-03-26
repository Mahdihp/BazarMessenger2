package user

import (
	"BazarMessenger/dto"
	"BazarMessenger/repository/postgres/user"
)

type UserServiceRepositort interface {
	Register() (dto.UserInfo, error)
}
type UserService struct {
	repo user.UserRepository
}

func New(repo user.UserRepository) UserService {
	return UserService{repo: repo}
}

func (this *UserService) Register() (dto.UserInfo, error) {
	return this.repo.Register()
}
