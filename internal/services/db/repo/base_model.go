package repo

import "time"
import "gorm.io/gorm"

type UuidModel struct {
	Uuid      string `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
