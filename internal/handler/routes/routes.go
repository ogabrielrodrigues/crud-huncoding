package routes

import (
	"github.com/gin-gonic/gin"
	handler "github.com/ogabrielrodrigues/crud-huncoding/internal/handler/user_handler"
)

func Routes(router *gin.RouterGroup) {
	router.GET("/user/:id")
	router.POST("/user", handler.CreateUser)
	router.PUT("/user/:id")
	router.DELETE("/user/:id")
}
