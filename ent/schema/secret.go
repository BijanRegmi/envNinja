package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Secret holds the schema definition for the Secret entity.
type Secret struct {
	ent.Schema
}

// Fields of the Secret.
func (Secret) Fields() []ent.Field {
	return []ent.Field{
		field.String("key"),
		field.String("value"),
	}
}

// Edges of the Secret.
func (Secret) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("project", Project.Type).Ref("secrets").Unique(),
	}
}

// Indexes of the Secret
func (Secret) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("key").Edges("project").Unique(),
	}
}
