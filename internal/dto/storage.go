package dto

import (
	"github.com/google/uuid"
	"github.com/khorevaa/r2gitsync/internal/services/datastore/ent"
	"github.com/khorevaa/r2gitsync/internal/services/datastore/ent/storage"
)

type StorageType uint

type Storage struct {
	ID               uuid.UUID
	ConnectionString string       `json:"connection_string,omitempty"`
	Develop          bool         `json:"develop,omitempty"`
	Extension        *string      `json:"extension,omitempty"`
	Type             storage.Type `json:"type,omitempty"`
	Project          *Project     `json:"project,omitempty"`
	Parent           *Storage     `json:"parent,omitempty"`
}

type Storages []*Storage

func (p *Storage) FromEnt(edm *ent.Storage) *Storage {
	if edm == nil {
		return nil
	}

	return &Storage{
		ID:               edm.ID,
		ConnectionString: edm.ConnectionString,
		Develop:          edm.Develop,
		Extension:        edm.Extension,
		Type:             edm.Type,
		Project:          (&Project{}).FromEnt(edm.Edges.Project),
		Parent:           (&Storage{}).FromEnt(edm.Edges.Parent),
	}
}

func (p Storages) FromEnt(edm ent.Storages) Storages {
	if len(edm) == 0 {
		return nil
	}
	return nil
}
