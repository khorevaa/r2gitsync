package repo

type Plugin struct {
	UuidModel
	Name        string `gorm:"size:50;uniqueIndex"`
	Description string
}
