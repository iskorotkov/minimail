package handlers

import (
	"github.com/iskorotkov/minimail/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

// AddMessage - Создаёт новое сообщение
func (c *Container) AddMessage(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, models.HelloWorld{
		Message: "Hello World",
	})
}

// ClapMessage - Увеличивает количество хлопков сообщения на 1
func (c *Container) ClapMessage(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, models.HelloWorld{
		Message: "Hello World",
	})
}

// GetMessage - Возвращает сообщение по ID
func (c *Container) GetMessage(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, models.HelloWorld{
		Message: "Hello World",
	})
}

// GetMessages - Возвращает список всех сообщений
func (c *Container) GetMessages(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, models.HelloWorld{
		Message: "Hello World",
	})
}
