package factory

import (
	"github.com/ogabrielrodrigues/crud-huncoding/config/database/postgres"
	handler "github.com/ogabrielrodrigues/crud-huncoding/internal/handler/user_handler"
	"github.com/ogabrielrodrigues/crud-huncoding/internal/model/repository"
	"github.com/ogabrielrodrigues/crud-huncoding/internal/model/service"
)

func NewUserFactory() handler.UserHandlerInterface {
	user_repository := repository.NewUserRepository(postgres.GetDB())
	user_service := service.NewUserDomainService(user_repository)
	return handler.NewUserHandlerInterface(user_service)
}
