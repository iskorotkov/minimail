package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/iskorotkov/minimail/models"
	"github.com/labstack/echo/v4"
)

// AddMessage - Создаёт новое сообщение
func (c *Container) AddMessagePage(ctx echo.Context) error {
	var dto models.NewMessageDto
	if err := ctx.Bind(&dto); err != nil {
		return echo.ErrBadRequest.SetInternal(err)
	}

	_, err := c.service.Add(dto)
	if err != nil {
		messages, err2 := c.service.All()
		if err2 != nil {
			return c.renderMessagesTemplate(ctx, nil, err2)
		}

		return c.renderMessagesTemplate(ctx, messages, err)
	}

	messages, err := c.service.All()
	if err != nil {
		return c.renderMessagesTemplateAfterSend(ctx, messages, err)
	}

	return c.renderMessagesTemplateAfterSend(ctx, messages, nil)
}

// ClapMessage - Увеличивает количество хлопков сообщения на 1
func (c *Container) ClapMessagePage(ctx echo.Context) error {
	id, err := strconv.ParseUint(ctx.Param("messageId"), 10, 32)
	if err != nil {
		ctx.Logger().Error(err)
		return c.redirectToIndex(ctx)
	}

	if _, err := c.service.Clap(uint(id)); err != nil {
		ctx.Logger().Error(err)
		return c.redirectToIndex(ctx)
	}

	if ctx.QueryParam("page") == "index" {
		return c.redirectToIndex(ctx)
	} else {
		return c.redirectToMessage(ctx, int(id))
	}
}

// GetMessage - Возвращает сообщение по ID
func (c *Container) GetMessagePage(ctx echo.Context) error {
	id, err := strconv.ParseUint(ctx.Param("messageId"), 10, 32)
	if err != nil {
		ctx.Logger().Error(err)
		return c.redirectToIndex(ctx)
	}

	message, err := c.service.ByID(uint(id))
	if err != nil {
		ctx.Logger().Error(err)
		return c.redirectToIndex(ctx)
	}

	return c.renderMessageTemplate(ctx, message, nil)
}

// GetMessages - Возвращает список всех сообщений
func (c *Container) GetMessagesPage(ctx echo.Context) error {
	messages, err := c.service.All()
	if err != nil {
		ctx.Logger().Error(err)
		return c.renderMessagesTemplate(ctx, messages, err)
	}

	return c.renderMessagesTemplate(ctx, messages, nil)
}

func (c *Container) redirectToIndex(ctx echo.Context) error {
	return ctx.Redirect(http.StatusSeeOther, "/simple/")
}

func (c *Container) redirectToMessage(ctx echo.Context, id int) error {
	return ctx.Redirect(http.StatusSeeOther, fmt.Sprintf("/simple/messages/%d", id))
}

func (c *Container) renderMessagesTemplate(ctx echo.Context, messages []models.Message, err error) error {
	errString := ""

	var httpErr *echo.HTTPError
	if errors.As(err, &httpErr) {
		errString = httpErr.Internal.Error()
	}

	return ctx.Render(http.StatusOK, "index.html", struct {
		Messages []models.Message
		Sent     bool
		Error    string
	}{messages, false, errString})
}

func (c *Container) renderMessagesTemplateAfterSend(ctx echo.Context, messages []models.Message, err error) error {
	errString := ""

	var httpErr *echo.HTTPError
	if errors.As(err, &httpErr) {
		errString = httpErr.Internal.Error()
	}

	return ctx.Render(http.StatusOK, "index.html", struct {
		Messages []models.Message
		Sent     bool
		Error    string
	}{messages, true, errString})
}

func (c *Container) renderMessageTemplate(ctx echo.Context, message models.Message, err error) error {
	errString := ""

	var httpErr *echo.HTTPError
	if errors.As(err, &httpErr) {
		errString = httpErr.Internal.Error()
	}

	return ctx.Render(http.StatusOK, "message.html", struct {
		Message models.Message
		Error   string
	}{message, errString})
}
