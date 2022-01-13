package repo

import (
	"gorm.io/gorm"
)

type Plugin struct {
	gorm.Model
	Name        string `gorm:"size:50;uniqueIndex"`
	Description string
}
