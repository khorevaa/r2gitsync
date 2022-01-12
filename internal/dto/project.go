package dto

type Project struct {
	Id          uint
	Code        string
	Name        string
	Description string
}

type Projects []*Project
