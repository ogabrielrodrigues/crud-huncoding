package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ogabrielrodrigues/crud-huncoding/config/rest"
	"github.com/ogabrielrodrigues/crud-huncoding/config/validation"
	"github.com/ogabrielrodrigues/crud-huncoding/internal/view"
)

func (uc *userHandlerInterface) FindUserByID(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := validation.Validate.Var(&id, "uuid4"); err != nil {
		rest_err := rest.NewBadRequestErr("id is not valid", []rest.Cause{
			{
				Field:   "id",
				Message: "param id is not valid",
			},
		})
		ctx.JSON(rest_err.Code, rest_err)
		return
	}

	user_domain, err := uc.service.FindUserByID(id)
	if err != nil {
		ctx.JSON(err.Code, err)
		return
	}

	ctx.JSON(http.StatusOK, view.ConvertDomainToResponse(user_domain))
}

func (uc *userHandlerInterface) FindUserByEmail(ctx *gin.Context) {
	email := ctx.Param("email")
	if err := validation.Validate.Var(&email, "email"); err != nil {
		rest_err := rest.NewBadRequestErr("email is not valid", []rest.Cause{
			{
				Field:   "email",
				Message: "param email is not valid",
			},
		})
		ctx.JSON(rest_err.Code, rest_err)
		return
	}

	user_domain, err := uc.service.FindUserByEmail(email)
	if err != nil {
		ctx.JSON(err.Code, err)
		return
	}

	ctx.JSON(http.StatusOK, view.ConvertDomainToResponse(user_domain))
}
