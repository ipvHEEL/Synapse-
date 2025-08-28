package handlers

import (
	"auth_service/crypto"
	"auth_service/database"
	users "auth_service/models"

	"github.com/gin-gonic/gin"
)

func Authorise(c *gin.Context) {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": "Ошибка ввода"})
		return
	}

	var user users.User

	if err := database.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		c.JSON(400, gin.H{"error": "Неверный логин или пароль"})
		return
	}
	if !crypto.CheckPasswordHash(input.Password, user.PasswordHash) {
		c.JSON(400, gin.H{"error": "Неверный логин или пароль"})
		return
	}

	token, err := crypto.GenerateJWT(user)
	if err != nil {
		c.JSON(500, gin.H{"error": "Не удалось создать токен"})
		return
	}

	c.JSON(200, gin.H{
		"message": "Login success",
		"token":   token,
		"user": gin.H{
			"id":    user.ID,
			"email": user.Email,
		},
	})
}
