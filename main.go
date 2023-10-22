package main

import (
	"fmt"
	"log"

	"github.com/Uallessonivo/golangCrud/src/controllers"
	"github.com/Uallessonivo/golangCrud/src/controllers/routes"
	"github.com/Uallessonivo/golangCrud/src/models/services"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	// Init dependencies
	service := services.NewUserDomainService()
	userController := controllers.NewUserControllerInterface(service)

	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup, userController)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
