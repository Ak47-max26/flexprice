package mixin

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// EnvironmentMixin implements the ent.Mixin for sharing environment_id field with package schemas.
type EnvironmentMixin struct {
	mixin.Schema
}

// Fields of the EnvironmentMixin.
func (EnvironmentMixin) Fields() []ent.Field {
	return []ent.Field{
		field.String("environment_id").
			SchemaType(map[string]string{
				"postgres": "varchar(50)",
			}).
			Optional().
			Immutable().
			Default(""),
	}
}

// Hooks of the EnvironmentMixin.
func (EnvironmentMixin) Hooks() []ent.Hook {
	return []ent.Hook{
		// Add hooks if needed
	}
}
