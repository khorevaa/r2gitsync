package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/khorevaa/r2gitsync/internal/services/db/ent/mixin"
)

// Asset holds the schema definition for the Asset entity.
type Asset struct {
	ent.Schema
}

// Fields of the Asset.
func (Asset) Fields() []ent.Field {
	return []ent.Field{
		field.String("file_name"),
	}
}
func (Asset) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.ModelMixin{},
	}
}

// Edges of the Asset.
func (Asset) Edges() []ent.Edge {
	return nil
}
