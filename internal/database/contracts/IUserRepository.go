package contracts

import (
	"testproject/internal/database/repository"
)

type IUserRepository interface {
	ReadUserByUsername(username string) (*repository.User, error)
}
