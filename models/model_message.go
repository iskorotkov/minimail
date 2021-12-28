package models

type Message struct {

	// ID сообщения
	Id int32 `json:"id"`

	// Имя автора сообщения
	Author string `json:"author"`

	// Текст сообщения
	Message string `json:"message"`

	// Количество хлопков
	Claps int32 `json:"claps"`
}
