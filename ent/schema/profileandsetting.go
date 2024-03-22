package schema

import "entgo.io/ent"

// ProfileAndSetting holds the schema definition for the ProfileAndSetting entity.
type ProfileAndSetting struct {
	ent.Schema
}

// Fields of the ProfileAndSetting.
func (ProfileAndSetting) Fields() []ent.Field {
	return nil
}

// Edges of the ProfileAndSetting.
func (ProfileAndSetting) Edges() []ent.Edge {
	return nil
}
