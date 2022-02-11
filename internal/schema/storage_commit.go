package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/khorevaa/r2gitsync/internal/schema/mixin"
)

// StorageCommit holds the schema definition for the StorageCommit entity.
type StorageCommit struct {
	ent.Schema
}

// Fields of the Storage.
func (StorageCommit) Fields() []ent.Field {

	return []ent.Field{
		field.Uint("number").Positive(),
		field.String("configuration_version"),
		field.String("author"),
		field.String("description"),
		field.String("tag"),
		field.String("tag_description"),
		field.Time("commit_at"),
		// field("parent").Optional(),
	}
}

// Edges of the Storage.
func (StorageCommit) Edges() []ent.Edge {

	return []ent.Edge{
		edge.To("storage", Storage.Type).Required().Unique(),
	}
}

func (StorageCommit) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.ModelMixin{},
	}
}

func (StorageCommit) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("number").Edges("storage").
			Unique(),
		index.Fields("commit_at"),
	}
}
