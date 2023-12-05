package repository

import (
	"github.com/ogabrielrodrigues/crud-huncoding/config/database"
	"github.com/ogabrielrodrigues/crud-huncoding/config/rest"
	"github.com/ogabrielrodrigues/crud-huncoding/internal/model"
)

type userRepository struct {
	database database.DB
}

type UserRepository interface {
	CreateUser(user_domain model.UserDomainInterface) (model.UserDomainInterface, *rest.RestErr)
	FindUserByID(id string) (model.UserDomainInterface, *rest.RestErr)
	FindUserByEmail(id string) (model.UserDomainInterface, *rest.RestErr)
}

func NewUserRepository(database database.DB) UserRepository {
	return &userRepository{database}
}
