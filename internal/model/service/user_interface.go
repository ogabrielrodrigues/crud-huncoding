package service

import (
	"github.com/ogabrielrodrigues/crud-huncoding/config/rest"
	"github.com/ogabrielrodrigues/crud-huncoding/internal/model"
)

type userDomainService struct{}

type UserDomainService interface {
	CreateUser(model.UserDomainInterface) *rest.RestErr
	FindUser(string) (*model.UserDomainInterface, *rest.RestErr)
	UpdateUser(string, model.UserDomainInterface) *rest.RestErr
	DeleteUser(string) *rest.RestErr
}

func NewUserDomainService() UserDomainService {
	return &userDomainService{}
}
