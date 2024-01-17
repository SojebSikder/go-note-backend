package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Note holds the schema definition for the Note entity.
type Note struct {
	ent.Schema
}

// Fields of the Note.
func (Note) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").Nillable(),
		field.String("content").Nillable(),
		field.Int("is_pinned").Default(0),
		field.Time("created_at").
			Default(time.Now()).Immutable(),
		field.Time("updated_at").
			Default(time.Now()),
	}
}

// Edges of the Note.
func (Note) Edges() []ent.Edge {
	// return []ent.Edge{
	//     edge.To("cars", Car.Type),
	// }
	return nil
}
