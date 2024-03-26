package user

import "BazarMessenger/repository/postgres/user"

type User struct {
	Repo user.UserRepository
}
