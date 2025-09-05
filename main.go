package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Заданные логин и пароль
const (
	USERNAME = "admin"
	PASSWORD = "secret"
)

// Middleware для проверки Basic Auth
func basicAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Получаем логин и пароль из заголовка Authorization
		username, password, ok := c.Request.BasicAuth()

		if !ok || username != USERNAME || password != PASSWORD {
			c.Header("WWW-Authenticate", `Basic realm="Restricted"`)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized",
			})
			return
		}

		// Авторизация успешна — продолжаем
		c.Next()
	}
}

func main() {
	r := gin.Default()

	// Публичный маршрут
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Привет! Это публичный эндпоинт.",
		})
	})

	// Защищённый маршрут без использования Group
	r.GET("/secure", basicAuthMiddleware(), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Добро пожаловать, вы авторизованы!",
		})
	})

	// Запуск сервера на порту 8080
	r.Run(":8990")
}