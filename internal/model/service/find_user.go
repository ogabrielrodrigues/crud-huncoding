package service

import (
	"github.com/ogabrielrodrigues/crud-huncoding/config/rest"
	"github.com/ogabrielrodrigues/crud-huncoding/internal/model"
)

func (ud *userDomainService) FindUserByID(id string) (model.UserDomainInterface, *rest.RestErr) {
	return ud.repository.FindUserByID(id)
}

func (ud *userDomainService) FindUserByEmail(email string) (model.UserDomainInterface, *rest.RestErr) {
	return ud.repository.FindUserByEmail(email)
}
