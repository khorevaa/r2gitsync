package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/khorevaa/r2gitsync/internal/services/db/ent/mixin"
)

// Plugin holds the schema definition for the Plugin entity.
type Plugin struct {
	ent.Schema
}

// Fields of the Plugin.
func (Plugin) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").MaxLen(50).Unique(),
		field.String("description").MaxLen(150),
	}
}

// Edges of the Plugin.
func (Plugin) Edges() []ent.Edge {
	return nil
}
func (Plugin) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.ModelMixin{},
	}
}

func (Plugin) Indexes() []ent.Index {
	return []ent.Index{
		// index.Fields("name").
		// 	Unique(),
	}
}
