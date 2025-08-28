package main

import (
	"auth_service/database"
	"auth_service/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()

	server := gin.Default()

	server.POST("/login", handlers.Authorise)
	server.Run(":8081")
}
