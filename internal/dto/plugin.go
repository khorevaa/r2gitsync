package dto

type Plugin struct {
	Uuid        string
	Name        string
	Description string
}

type Plugins []*Plugin
