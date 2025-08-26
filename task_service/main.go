package main

import (
	"tasksvc/database"
	"tasksvc/handlers"

	"github.com/gin-gonic/gin"
)

func main() {

	database.Connect()

	server := gin.Default()

	server.GET("/tasks", handlers.GetTasks)
	server.POST("/tasks", handlers.CreateTask)
	server.DELETE("/tasks", handlers.DeleteTask)

	server.POST("/teams", handlers.CreateTeam)
	server.GET("/teams", handlers.GetTeams)
	server.Run(":8080")

}
