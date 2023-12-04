package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ogabrielrodrigues/crud-huncoding/config/validation"
	"github.com/ogabrielrodrigues/crud-huncoding/internal/handler/model/request"
	"github.com/ogabrielrodrigues/crud-huncoding/internal/model"
	"github.com/ogabrielrodrigues/crud-huncoding/internal/view"
)

var (
	UserDomainInterface model.UserDomainInterface
)

func (uc *userHandlerInterface) CreateUser(ctx *gin.Context) {
	var user_request = request.UserRequest{}

	if err := ctx.BindJSON(&user_request); err != nil {
		rest_err := validation.ValidateErr(err)
		ctx.JSON(rest_err.Code, rest_err)
		return
	}

	domain := model.NewUserDomain(
		user_request.Name,
		user_request.Email,
		user_request.Password,
		user_request.Age,
	)

	if err := uc.service.CreateUser(domain); err != nil {
		ctx.JSON(err.Code, err)
		return
	}

	ctx.JSON(http.StatusCreated, view.ConvertDomainToResponse(domain))
}
