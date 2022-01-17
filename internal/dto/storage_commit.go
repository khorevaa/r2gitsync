package dto

import "time"

type StorageCommit struct {
	Uuid                 string
	StorageUuid          string
	Number               uint
	ConfigurationVersion string
	Author               string
	Description          string
	CommitAt             time.Time
	Tag                  string
	TagDesc              string
}

type StorageCommits []*StorageCommit
