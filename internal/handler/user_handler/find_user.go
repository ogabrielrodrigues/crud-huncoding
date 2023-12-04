package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/ogabrielrodrigues/crud-huncoding/config/rest"
	"github.com/ogabrielrodrigues/crud-huncoding/config/validation"
)

func (uc *userHandlerInterface) FindUserByID(ctx *gin.Context) {
	_, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		rest_err := rest.NewBadRequestErr("id is not valid", []rest.Cause{
			{
				Field:   "id",
				Message: "param id is not valid",
			},
		})
		ctx.JSON(rest_err.Code, rest_err)
		return
	}
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
}
