package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ogabrielrodrigues/crud-huncoding/config/validation"
	"github.com/ogabrielrodrigues/crud-huncoding/internal/handler/model/request"
)

func CreateUser(ctx *gin.Context) {
	var user_request = request.UserRequest{}

	if err := ctx.BindJSON(&user_request); err != nil {
		rest_err := validation.ValidateUserError(err)
		ctx.JSON(rest_err.Code, rest_err)
		return
	}

	ctx.JSON(http.StatusOK, user_request)
}
