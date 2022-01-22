package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/khorevaa/r2gitsync/internal/services/db/ent/mixin"
)

// StoragePlugin holds the schema definition for the StoragePlugin entity.
type StoragePlugin struct {
	ent.Schema
}

// Fields of the StoragePlugin.
func (StoragePlugin) Fields() []ent.Field {
	return []ent.Field{
		field.Bool("disable"),
	}
}

// Edges of the StoragePlugin.
func (StoragePlugin) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("storage", Storage.Type).Required().Unique(),
		edge.To("plugin", PluginVersion.Type).Required().Unique(),
		edge.To("properties", StoragePluginProperty.Type),
	}
}

func (StoragePlugin) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.ModelMixin{},
	}
}

func (StoragePlugin) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields().Edges("storage", "plugin").
			Unique(),
		index.Fields("disable"),
	}
}
