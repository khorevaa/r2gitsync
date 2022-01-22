package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/khorevaa/r2gitsync/internal/services/db/ent/mixin"
)

// StoragePluginProperty holds the schema definition for the StoragePluginProperty entity.
type StoragePluginProperty struct {
	ent.Schema
}

// Fields of the StoragePluginProperty.
func (StoragePluginProperty) Fields() []ent.Field {

	// PluginPropertyUuid string `gorm:"TYPE:uuid REFERENCES plugin_properties;index;uniqueIndex:idx_plugin_property_uuid_name"`
	// Name               string `gorm:"uniqueIndex:idx_plugin_property_uuid_name"`
	// Value              *string
	return []ent.Field{
		field.String("name"),
		field.String("value"),
	}
}

// Edges of the StoragePluginProperty.
func (StoragePluginProperty) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("plugin", StoragePlugin.Type).
			Ref("properties").
			Required().Unique(),
	}
}

func (StoragePluginProperty) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.ModelMixin{},
	}
}

func (StoragePluginProperty) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("name").Edges("plugin").
			Unique(),
	}
}
