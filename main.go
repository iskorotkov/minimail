package main

import (
	"os"

	"github.com/iskorotkov/minimail/handlers"
	"github.com/iskorotkov/minimail/models"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	e := echo.New()

	if err := godotenv.Load(); err != nil {
		e.Logger.Fatal(err)
	}

	db, err := gorm.Open(postgres.Open(os.Getenv("CONN_STRING")))
	if err != nil {
		e.Logger.Fatal(err)
	}

	if err := db.AutoMigrate(models.Message{}); err != nil {
		e.Logger.Fatal(err)
	}

	c, err := handlers.NewContainer(db)
	if err != nil {
		e.Logger.Fatal(err)
	}

	// Middleware
	e.Use(middleware.NonWWWRedirect())
	e.Use(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Use(middleware.Gzip())

	// AddMessage - Создаёт новое сообщение
	e.POST("/api/messages", c.AddMessage)

	// ClapMessage - Увеличивает количество хлопков сообщения на 1
	e.POST("/api/messages/:messageId/claps", c.ClapMessage)

	// GetMessage - Возвращает сообщение по ID
	e.GET("/api/messages/:messageId", c.GetMessage)

	// GetMessages - Возвращает список всех сообщений
	e.GET("/api/messages", c.GetMessages)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
