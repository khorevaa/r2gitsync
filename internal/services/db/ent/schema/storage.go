package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/khorevaa/r2gitsync/internal/services/db/ent/mixin"
)

// Storage holds the schema definition for the Storage entity.
type Storage struct {
	ent.Schema
}

// Fields of the Storage.
func (Storage) Fields() []ent.Field {
	// ConnectionString string
	// Type             dto.StorageType
	// Develop          bool
	// Extension        *string
	// ParentUuid       *uint
	// Parent           *Storage
	return []ent.Field{
		field.String("connection_string"),
		field.Bool("develop"),
		field.String("extension").Optional().Nillable(),
		field.Enum("type").Values("http", "file"),
		// field("parent").Optional(),
	}
}

// Edges of the Storage.
func (Storage) Edges() []ent.Edge {

	return []ent.Edge{
		edge.From("project", Project.Type).
			Ref("storages").Unique().
			Required(),
		edge.To("parent", Storage.Type).
			Unique(),
	}
}

func (Storage) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.ModelMixin{},
	}
}

func (Storage) Indexes() []ent.Index {
	return []ent.Index{
		// index.Fields("name").
		// 	Unique(),
	}
}
