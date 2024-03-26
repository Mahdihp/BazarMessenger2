package user

import (
	"BazarMessenger/dto"
	"BazarMessenger/repository/postgres"
)

type UserRepository interface {
	Insert() (dto.UserInfo, error)
	//Update() (dto.UserInfo, error)

	//IsMobileUnique() (bool, error)
	//IsUsernameUnique() (bool, error)

	//GetByMobile() (dto.UserInfo, error)
	//GetByUsername() (dto.UserInfo, error)
}
type UserRepositoryImpl struct {
	conn *postgres.PostgresqlDB
}

func New(conn *postgres.PostgresqlDB) *UserRepositoryImpl {
	return &UserRepositoryImpl{
		conn: conn,
	}
}

func (this *UserRepositoryImpl) Insert() (dto.UserInfo, error) {
	return dto.UserInfo{Mobile: "0900000000000000"}, nil
}
