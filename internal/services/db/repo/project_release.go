package repo

import (
	"time"

	"gorm.io/gorm"
)

type ProjectRelease struct {
	gorm.Model
	ProjectID  uint `gorm:"TYPE:integer REFERENCES projects;uniqueIndex:idx_project_id_version"`
	Project    Project
	Version    string `gorm:"uniqueIndex:idx_project_id_version"`
	ReleasedAt *time.Time
	Assets     []*Asset `gorm:"polymorphic:Owner;"`
}
