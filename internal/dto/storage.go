package dto

type Storage struct {
	Id               string
	ConnectionString string
	Type             StorageType
	Develop          bool
	ParentId         *uint

	ProjectId   *uint
	ProjectName string

	RemoteUrl  string
	BranchName string
}

type Storages []*Storage
