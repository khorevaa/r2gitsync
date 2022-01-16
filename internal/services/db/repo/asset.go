package repo

type Asset struct {
	UuidModel

	OwnerID   string
	OwnerType string

	Filename string
	Size     uint
}
