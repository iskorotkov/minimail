package services

import (
	"errors"
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/iskorotkov/minimail/models"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Messages struct {
	db *gorm.DB
}

func NewMessages(db *gorm.DB) (Messages, error) {
	return Messages{db: db}, nil
}

func (m Messages) All() ([]models.Message, error) {
	var messages []models.Message
	if err := m.db.Order("claps DESC, id").Find(&messages).Error; err != nil {
		return nil, echo.ErrInternalServerError.SetInternal(err)
	}

	return messages, nil
}

func (m Messages) ByID(id uint) (models.Message, error) {
	var message models.Message
	if err := m.db.First(&message, "id = ?", id).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return models.Message{}, echo.ErrNotFound.SetInternal(err)
	} else if err != nil {
		return models.Message{}, echo.ErrInternalServerError.SetInternal(err)
	}

	return message, nil
}

func (m Messages) Clap(id uint) (models.Message, error) {
	query := m.db.
		Model(&models.Message{}).
		Where("id = ?", id).
		UpdateColumn("claps", gorm.Expr("claps + 1"))
	if err := query.Error; err != nil {
		return models.Message{}, echo.ErrInternalServerError.SetInternal(err)
	}

	if query.RowsAffected == 0 {
		return models.Message{}, echo.ErrNotFound
	}

	var message models.Message
	if err := m.db.First(&message, id).Error; err != nil {
		return models.Message{}, echo.ErrInternalServerError.SetInternal(err)
	}

	return message, nil
}

func (m Messages) Add(dto models.NewMessageDto) (models.Message, error) {
	if ok, err := govalidator.ValidateStruct(dto); !ok {
		return models.Message{}, echo.NewHTTPError(http.StatusUnprocessableEntity).SetInternal(err)
	} else if err != nil {
		return models.Message{}, echo.ErrInternalServerError.SetInternal(err)
	}

	model := models.Message{
		Id:      0,
		Author:  dto.Author,
		Message: dto.Message,
		Claps:   0,
	}
	if err := m.db.Create(&model).Error; err != nil {
		return models.Message{}, echo.ErrInternalServerError.SetInternal(err)
	}

	return model, nil
}
