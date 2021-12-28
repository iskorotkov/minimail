package models

type NewMessageDto struct {

	// Имя автора сообщения
	Author string `json:"author"`

	// Текст сообщения
	Message string `json:"message"`
}
