package schema

import (
	"entgo.io/contrib/entproto"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the Users entity.
type User struct {
	ent.Schema
}

func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
		entproto.Service(),
	}
}

// Fields of the Users.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("username").Unique().
			Annotations(
				entproto.Field(2),
			),
		field.String("email").Unique().
			Annotations(
				entproto.Field(3),
			),
		field.String("password_hash").
			Annotations(
				entproto.Field(4),
			),
	}
}
