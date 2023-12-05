package service

import (
	"github.com/ogabrielrodrigues/crud-huncoding/config/rest"
	"github.com/ogabrielrodrigues/crud-huncoding/internal/model"
)

func (us *userDomainService) CreateUser(user_domain model.UserDomainInterface) (model.UserDomainInterface, *rest.RestErr) {
	user_domain.EncryptPassword()

	return us.repository.CreateUser(user_domain)
}
