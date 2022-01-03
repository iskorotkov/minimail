package models

type NewMessageDto struct {

	// Имя автора сообщения
	Author string `json:"author" form:"sender" valid:"maxstringlength(30)~Автор не должен быть длиннее 30 символов,required~Автор не должен быть пустым"`

	// Текст сообщения
	Message string `json:"message" form:"message" valid:"maxstringlength(1000)~Текст сообщения не должен быть длиннее 1000 символов,required~Текст сообщения не должен быть пустым"`
}
