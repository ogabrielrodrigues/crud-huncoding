package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/ogabrielrodrigues/crud-huncoding/config"
	"github.com/ogabrielrodrigues/crud-huncoding/config/logger"
	"github.com/ogabrielrodrigues/crud-huncoding/internal/handler/routes"
	handler "github.com/ogabrielrodrigues/crud-huncoding/internal/handler/user_handler"
	"github.com/ogabrielrodrigues/crud-huncoding/internal/model/service"
)

func main() {
	err := config.Load()
	if err != nil {
		logger.Error(err)
		os.Exit(1)
	}

	user_service := service.NewUserDomainService()
	user_handler := handler.NewUserHandlerInterface(user_service)

	router := gin.Default()
	routes.Routes(&router.RouterGroup, user_handler)

	if err = router.Run(config.GetAPIConfig().Port); err != nil {
		log.Fatal(err)
	}
}
