package service

import (
	"github.com/ogabrielrodrigues/crud-huncoding/config/rest"
	"github.com/ogabrielrodrigues/crud-huncoding/internal/model"
	"github.com/ogabrielrodrigues/crud-huncoding/internal/model/repository"
)

type userDomainService struct {
	repository repository.UserRepository
}

type UserDomainService interface {
	CreateUser(model.UserDomainInterface) (model.UserDomainInterface, *rest.RestErr)
	FindUserByID(string) (model.UserDomainInterface, *rest.RestErr)
	FindUserByEmail(string) (model.UserDomainInterface, *rest.RestErr)
	UpdateUser(string, model.UserDomainInterface) *rest.RestErr
	DeleteUser(string) *rest.RestErr
}

func NewUserDomainService(repository repository.UserRepository) UserDomainService {
	return &userDomainService{repository}
}
