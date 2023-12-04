package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/ogabrielrodrigues/crud-huncoding/internal/model/service"
)

type UserHandlerInterface interface {
	CreateUser(ctx *gin.Context)
	FindUserByID(ctx *gin.Context)
	FindUserByEmail(ctx *gin.Context)
	UpdateUser(ctx *gin.Context)
	DeleteUser(ctx *gin.Context)
}

type userHandlerInterface struct {
	service service.UserDomainService
}

func NewUserHandlerInterface(service_interface service.UserDomainService) UserHandlerInterface {
	return &userHandlerInterface{
		service: service_interface,
	}
}
