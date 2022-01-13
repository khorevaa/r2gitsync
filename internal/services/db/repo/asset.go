package repo

import "gorm.io/gorm"

type Asset struct {
	gorm.Model

	OwnerID   int
	OwnerType string

	Filename string
	Size     uint
	MD5      []byte
}
