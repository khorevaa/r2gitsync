package dto

type Asset struct {
	Uuid string

	OwnerID   string
	OwnerType string

	Filename string
	Size     uint
}

type Assets []*Asset
