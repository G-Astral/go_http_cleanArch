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
	DelUserById(id int) (rowsAffected int64, err error)
	UpdUserById(user *entities.User, id int) (err error)
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
		c.Set("logMessage", "Ошибка при добавлении пользователя: неверный запрос")
		c.JSON(400, gin.H{"error": "Неверный JSON"})
		return
	}

	err = h.ser.AddUser(&user)
	if err != nil {
		c.Set("logMessage", "Ошибка при добавлении пользователя")
		c.JSON(500, gin.H{"error": "Bad request"})
		return
	}

	c.Set("logMessage", fmt.Sprintf("Добавлен новый пользователь. Имя: %s. Возраст: %d. ID: %d", user.Name, user.Age, user.Id))
	c.JSON(200, gin.H{
		"message": fmt.Sprintf("Имя: %s. Возраст: %d. ID: %d", user.Name, user.Age, user.Id),
	})

}

func (h *handler) GetAllUsers(c *gin.Context) {
	users, err := h.ser.GetAllUsers()
	if err != nil {
		c.Set("logMessage", "Ошибка сервера при получении всех пользователей")
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}

	c.Set("logMessage", "Запрошены все пользователи")
	c.IndentedJSON(200, users)
}

func (h *handler) GetUserByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.Set("logMessage", "Неверный ID (не число)")
		c.JSON(400, gin.H{"error": "Неверный ID (не число)"})
		return
	}

	user, err := h.ser.GetUserByID(id)
	if err == sql.ErrNoRows {
		c.Set("logMessage", "Пользователь не найден")
		c.JSON(404, gin.H{"error": "Пользователь не найден"})
		return
	} else if err != nil {
		c.Set("logMessage", "Ошибка сервера при получении пользователя")
		c.JSON(500, gin.H{"error": "Ошибка сервера при получении пользователя"})
		return
	}

	c.Set("logMessage", fmt.Sprintf("Запрошен пользователь с ID: %d", id))
	c.JSON(200, user)
}

func (h *handler) DelUserById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.Set("logMessage", "Неверный ID (не число)")
		c.JSON(400, gin.H{"error": "Неверный ID (не число)"})
		return
	}

	rowsAffected, err := h.ser.DelUserById(id)
	if err != nil {
		c.Set("logMessage", "Ошибка сервера при удалении пользователя")
		c.JSON(500, gin.H{"error": "Ошибка сервера при удалении пользователя"})
		return
	}

	if rowsAffected > 0 {
		c.Set("logMessage", fmt.Sprintf("Удален пользователь с ID: %d", id))
		c.JSON(200, gin.H{"message": "Пользователь удален"})
	} else {
		c.Set("logMessage", "Пользователь не найден")
		c.JSON(404, gin.H{"error": "Пользователь не найден"})
	}
}

func (h *handler) UpdUserById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.Set("logMessage", "Неверный ID (не число)")
		c.JSON(400, gin.H{"error": "Неверный ID (не число)"})
		return
	}

	user := entities.User{}

	err = c.BindJSON(&user)
	if err != nil {
		c.Set("logMessage", "Ошибка при изменении пользователя: неверный запрос")
		c.JSON(400, gin.H{"error": "Неверный JSON"})
		return
	}

	err = h.ser.UpdUserById(&user, id)
	if err == sql.ErrNoRows {
		c.Set("logMessage", "Пользователь не найден")
		c.JSON(404, gin.H{"error": "Пользователь не найден"})
		return
	} else if err != nil {
		c.Set("logMessage", "Ошибка сервера при изменении пользователя")
		c.JSON(500, gin.H{"error": "Ошибка сервера при изменении пользователя"})
		return
	}

	c.Set("logMessage", fmt.Sprintf("Изменен пользователь с ID: %d", id))
	c.JSON(200, &user)
}
