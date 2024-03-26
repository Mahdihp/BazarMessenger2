package user

import (
	"BazarMessenger/dto"
	"BazarMessenger/repository/postgres/user"
)

type UserService struct {
	repo user.UserRepository
}

func New(repo user.UserRepository) UserService {
	return UserService{repo: repo}
}

func (this *UserService) Register() (dto.UserInfo, error) {
	return this.repo.Insert()
}
