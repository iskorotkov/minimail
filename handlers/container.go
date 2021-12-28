package handlers

import "gorm.io/gorm"

// Container will hold all dependencies for your application.
type Container struct {
	db *gorm.DB
}

// NewContainer returns an empty or an initialized container for your handlers.
func NewContainer(db *gorm.DB) (Container, error) {
	return Container{db: db}, nil
}
