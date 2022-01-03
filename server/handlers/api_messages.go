package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/iskorotkov/minimail/models"
	"github.com/labstack/echo/v4"
)

// AddMessage - Создаёт новое сообщение
func (c *Container) AddMessage(ctx echo.Context) error {
	var dto models.NewMessageDto
	if err := ctx.Bind(&dto); err != nil {
		return echo.ErrBadRequest.SetInternal(err)
	}

	message, err := c.service.Add(dto)
	if err != nil {
		var httpErr *echo.HTTPError
		if errors.As(err, &httpErr) {
			return ctx.JSON(http.StatusUnprocessableEntity, models.Error{Message: httpErr.Internal.Error()})
		}

		return err
	}

	return ctx.JSON(http.StatusCreated, message)
}

// ClapMessage - Увеличивает количество хлопков сообщения на 1
func (c *Container) ClapMessage(ctx echo.Context) error {
	id, err := strconv.ParseUint(ctx.Param("messageId"), 10, 32)
	if err != nil {
		return echo.ErrBadRequest.SetInternal(err)
	}

	message, err := c.service.Clap(uint(id))
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusCreated, models.ClapsCount{
		Count: message.Claps,
	})
}

// GetMessage - Возвращает сообщение по ID
func (c *Container) GetMessage(ctx echo.Context) error {
	id, err := strconv.ParseUint(ctx.Param("messageId"), 10, 32)
	if err != nil {
		return echo.ErrBadRequest.SetInternal(err)
	}

	message, err := c.service.ByID(uint(id))
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, message)
}

// GetMessages - Возвращает список всех сообщений
func (c *Container) GetMessages(ctx echo.Context) error {
	messages, err := c.service.All()
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, messages)
}
