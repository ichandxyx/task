package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Visit holds the schema definition for the Visit entity.
type Visit struct {
	ent.Schema
}

// Fields of the Visit.
func (Visit) Fields() []ent.Field {
	return []ent.Field{
		field.String("storeID"),
		field.Time("visitTime"),
		field.String("error").Nillable().Optional(),
		field.Int("perimeter").Optional(),
		field.Strings("imageURLs"),
	}
}

// Edges of the Visit.
func (Visit) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("job", Job.Type).
			Ref("visits").
			Unique(),
	}
}
