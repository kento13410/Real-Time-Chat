package schema

import (
	"entgo.io/contrib/entproto"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// UserRelations holds the schema definition for the UserRelations entity.
type UserRelations struct {
	ent.Schema
}

func (UserRelations) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
		entproto.Service(),
	}
}

// Fields of the UserRelations.
func (UserRelations) Fields() []ent.Field {
	return []ent.Field{
		field.Int("user_id_1").
			Annotations(
				entproto.Field(2),
			),
		field.Int("user_id_2").
			Annotations(
				entproto.Field(3),
			),
	}
}
