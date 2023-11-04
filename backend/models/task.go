package models

import (
	"github.com/jinzhu/gorm"
)

type Task struct {
	gorm.model

	Title         string
	Description   string
	Status        string
}
