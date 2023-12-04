package routes

import (
	"github.com/gin-gonic/gin"
	handler "github.com/ogabrielrodrigues/crud-huncoding/internal/handler/user_handler"
)

func Routes(router *gin.RouterGroup, user_handler handler.UserHandlerInterface) {
	router.GET("/user/:id", user_handler.FindUserByID)
	router.GET("/user_by_email/:email", user_handler.FindUserByEmail)
	router.POST("/user", user_handler.CreateUser)
	router.PUT("/user/:id", user_handler.UpdateUser)
	router.DELETE("/user/:id", user_handler.DeleteUser)
}
