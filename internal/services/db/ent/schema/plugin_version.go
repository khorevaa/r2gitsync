package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/khorevaa/r2gitsync/internal/services/db/ent/mixin"
)

// PluginVersion holds the schema definition for the PluginVersion entity.
type PluginVersion struct {
	ent.Schema
}

// Fields of the PluginVersion.
func (PluginVersion) Fields() []ent.Field {
	return []ent.Field{
		field.String("version"),
		field.String("description").MaxLen(150),
		field.Bool("broken"),
	}
}

// Edges of the PluginVersion.
func (PluginVersion) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("plugin", Plugin.Type).Required().Unique(),
	}
}

func (PluginVersion) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.ModelMixin{},
	}
}

func (PluginVersion) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("version").Edges("plugin").
			Unique(),
		index.Fields("broken"),
	}
}
