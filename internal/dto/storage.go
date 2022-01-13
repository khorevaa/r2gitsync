package dto

type Storage struct {
	Id               string
	ConnectionString string
	Type             StorageType
	Develop          bool
	Extension        *uint
	ParentId         *uint
	Parent           *Storages

	ProjectId *uint
	Project   Project

	RemoteUrl  string
	BranchName string
}

type Storages []*Storage
