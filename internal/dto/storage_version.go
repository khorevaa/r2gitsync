package dto

import "time"

type StorageVersion struct {
	ID                   uint
	ConfigurationVersion string
	Author               string
	Description          string
	CommitAt             time.Time
	Tag                  string
	TagDesc              string
}

type StorageVersions []*StorageVersion
