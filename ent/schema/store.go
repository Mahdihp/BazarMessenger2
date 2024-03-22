package schema

import "entgo.io/ent"

// Store holds the schema definition for the Store entity.
type Store struct {
	ent.Schema
}

// Fields of the Store.
func (Store) Fields() []ent.Field {
	return nil
}

// Edges of the Store.
func (Store) Edges() []ent.Edge {
	return nil
}
