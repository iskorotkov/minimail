package main

import (
	"github.com/iskorotkov/minimail/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	//todo: handle the error!
	c, _ := handlers.NewContainer()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

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
