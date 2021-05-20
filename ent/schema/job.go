package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/edge"

)

// Job holds the schema definition for the Job entity.
type Job struct {
	ent.Schema
}

// Fields of the Job.
func (Job) Fields() []ent.Field {

	return []ent.Field{
		field.String("status").Default("ongoing"),
		field.Time("createdAt").Default(time.Now),
	}
}


// Edges of the Job.
func (Job) Edges() []ent.Edge {
	return []ent.Edge{
        edge.To("visits", Visit.Type),
    }
}

