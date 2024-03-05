package schema

import (
	"entgo.io/contrib/entproto"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
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
		field.String("password_hash").
			Annotations(
				entproto.Field(3),
			),
		field.Time("created_at").Immutable().
			Default(func() time.Time {
				return time.Now()
			}).
			Annotations(
				entproto.Field(4),
			),
	}
}

func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user_relations_1", UserRelation.Type).Annotations(entproto.Field(5)),
		edge.To("user_relations_2", UserRelation.Type).Annotations(entproto.Field(6)),
		edge.To("sent_messages", Chat.Type).Annotations(entproto.Field(7)),
		edge.To("received_messages", Chat.Type).Annotations(entproto.Field(8)),
	}
}
