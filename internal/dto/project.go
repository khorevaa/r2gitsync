package dto

import (
	"github.com/google/uuid"
	"github.com/khorevaa/r2gitsync/internal/services/datastore/ent"
)

type ProjectType string

// Type values.
const (
	ConfigurationProjectType ProjectType = "configuration"
	ExtensionProjectType     ProjectType = "extension"
)

type Project struct {
	ID             uuid.UUID
	Code           string
	Name           string
	Description    string
	Type           ProjectType
	MasterStorage  *Storage
	DevelopStorage *Storage

	Storages []*Storage
}

type Projects []*Project

func (p *Project) FromEnt(edm *ent.Project) *Project {
	if edm == nil {
		return nil
	}

	return &Project{
		ID:             edm.ID,
		Code:           edm.Code,
		Name:           edm.Name,
		Description:    edm.Description,
		Type:           ProjectType(edm.Type),
		MasterStorage:  (&Storage{}).FromEnt(edm.Edges.MasterStorage),
		DevelopStorage: (&Storage{}).FromEnt(edm.Edges.DevelopStorage),
		Storages:       (&Storages{}).FromEnt(edm.Edges.Storages),
	}
}

func (p Projects) FromEnt(edm ent.Projects) Projects {
	if edm == nil {
		return nil
	}

	return nil
}
