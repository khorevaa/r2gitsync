package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/khorevaa/r2gitsync/internal/services/db/ent/mixin"
)

// Project holds the schema definition for the Project entity.
type Project struct {
	ent.Schema
}

// Fields of the Project.
func (Project) Fields() []ent.Field {
	return []ent.Field{
		field.String("code").MinLen(5).Unique(),
		field.String("name"),
		field.String("Description"),
		field.Enum("type").Values("configuration", "extension"),
		// field("parent").Optional(),
	}
}

// Edges of the Project.
func (Project) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("storages", Storage.Type),
		edge.To("master_storage", Storage.Type).Unique().Comment("хранилище master ветки"),
		edge.To("develop_storage", Storage.Type).Unique().Comment("хранилище develop ветки"),
	}
}

func (Project) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.ModelMixin{},
	}
}
