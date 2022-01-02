package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/asaskevich/govalidator"
	"github.com/iskorotkov/minimail/models"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// AddMessage - Создаёт новое сообщение
func (c *Container) AddMessage(ctx echo.Context) error {
	var dto models.NewMessageDto
	if err := ctx.Bind(&dto); err != nil {
		return echo.ErrBadRequest.SetInternal(err)
	}

	if ok, err := govalidator.ValidateStruct(dto); !ok {
		return ctx.JSON(http.StatusUnprocessableEntity, models.Error{
			Message: err.Error(),
		})
	} else if err != nil {
		return echo.ErrInternalServerError.SetInternal(err)
	}

	model := models.Message{
		Id:      0,
		Author:  dto.Author,
		Message: dto.Message,
		Claps:   0,
	}
	if err := c.db.Create(&model).Error; err != nil {
		return echo.ErrInternalServerError.SetInternal(err)
	}

	return ctx.JSON(http.StatusCreated, model)
}

// ClapMessage - Увеличивает количество хлопков сообщения на 1
func (c *Container) ClapMessage(ctx echo.Context) error {
	id, err := strconv.ParseUint(ctx.Param("messageId"), 10, 32)
	if err != nil {
		return echo.ErrBadRequest.SetInternal(err)
	}

	query := c.db.Model(&models.Message{}).Where("id = ?", id).UpdateColumn("claps", gorm.Expr("claps + 1"))
	if err := query.Error; err != nil {
		return echo.ErrInternalServerError.SetInternal(err)
	}

	if query.RowsAffected == 0 {
		return echo.ErrNotFound
	}

	var message models.Message
	if err := c.db.First(&message, id).Error; err != nil {
		return echo.ErrInternalServerError.SetInternal(err)
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

	var message models.Message
	if err := c.db.First(&message, "id = ?", id).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return echo.ErrNotFound.SetInternal(err)
	} else if err != nil {
		return echo.ErrInternalServerError.SetInternal(err)
	}

	return ctx.JSON(http.StatusOK, message)
}

// GetMessages - Возвращает список всех сообщений
func (c *Container) GetMessages(ctx echo.Context) error {
	var messages []models.Message
	if err := c.db.Order("claps DESC").Find(&messages).Error; err != nil {
		return echo.ErrInternalServerError.SetInternal(err)
	}

	return ctx.JSON(http.StatusOK, messages)
}
