package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/ogabrielrodrigues/crud-huncoding/config"
	"github.com/ogabrielrodrigues/crud-huncoding/internal/handler/routes"
)

func main() {
	err := config.Load()
	if err != nil {
		panic(err)
	}

	router := gin.Default()
	routes.Routes(&router.RouterGroup)

	if err = router.Run(config.GetAPIConfig().Port); err != nil {
		log.Fatal(err)
	}
}
