package handlers

import (
	"database/sql"
	"fmt"
	"go-http-cleanArch/internal/entities"
	"strconv"

	"github.com/gin-gonic/gin"
)

type service interface {
	AddUser(user *entities.User) (err error)
	GetAllUsers() (users *[]entities.User, err error)
	GetUserByID(id int) (user *entities.User, err error)
}

type handler struct {
	ser service
}

func InitHandler(ser service) handler {
	return handler{
		ser: ser,
	}
}

func (h *handler) AddUser(c *gin.Context) {
	user := entities.User{}

	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(400, gin.H{"error": "Неверный JSON"})
		return
	}

	err = h.ser.AddUser(&user)
	if err != nil {
		c.JSON(500, gin.H{"error": "Bad request"})
		return
	}

	c.JSON(200, gin.H{
		"message": fmt.Sprintf("Имя: %s. Возраст: %d. ID: %d", user.Name, user.Age, user.Id),
	})

}

func (h *handler) GetAllUsers(c *gin.Context) {
	users, err := h.ser.GetAllUsers()
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}

	c.IndentedJSON(200, users)
}

func (h *handler) GetUserByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(400, gin.H{"error": "Неверный ID (не число)"})
		return
	}

	user, err := h.ser.GetUserByID(id)
	if err == sql.ErrNoRows {
		c.JSON(404, gin.H{"error": "Пользователь не найден"})
		return
	} else if err != nil {
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(200, user)
}
