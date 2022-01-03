package handlers

import (
	"github.com/iskorotkov/minimail/services"
)

// Container will hold all dependencies for your application.
type Container struct {
	service services.Messages
}

// NewContainer returns an empty or an initialized container for your handlers.
func NewContainer(service services.Messages) (Container, error) {
	return Container{service: service}, nil
}
