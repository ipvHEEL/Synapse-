package handlers

import (
	"auth_service/database"
	users "auth_service/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authorise(c *gin.Context) {
	var userdata users.User
	database.DB.Find(&userdata)
	c.JSON(http.StatusOK, userdata)
	fmt.Println("Ok")
}
