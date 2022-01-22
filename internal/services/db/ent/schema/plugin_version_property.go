package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/khorevaa/r2gitsync/internal/services/db/ent/mixin"
)

// PluginVersionProperty holds the schema definition for the PluginVersionProperty entity.
type PluginVersionProperty struct {
	ent.Schema
}

// Fields of the PluginVersionProperty.
func (PluginVersionProperty) Fields() []ent.Field {

	return []ent.Field{
		field.String("name").MaxLen(50),
		field.String("default").MaxLen(150),
		field.Bool("required"),
		field.Enum("type").Values("bool", "string", "int"),
	}
}

// Edges of the PluginVersionProperty.
func (PluginVersionProperty) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("plugin", Plugin.Type).Required().Unique(),
		edge.To("version", PluginVersion.Type).Required().Unique(),
	}
}
func (PluginVersionProperty) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.ModelMixin{},
	}
}

func (PluginVersionProperty) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("name").Edges("plugin", "version").
			Unique(),
		index.Fields("required"),
	}
}
