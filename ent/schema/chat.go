package schema

import (
	"entgo.io/contrib/entproto"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// Chat holds the schema definition for the Chat entity.
type Chat struct {
	ent.Schema
}

func (Chat) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
		entproto.Service(),
	}
}

// Fields of the Chat.
func (Chat) Fields() []ent.Field {
	return []ent.Field{
		field.String("message").
			Annotations(
				entproto.Field(2),
			),
		field.Time("sent_at").Immutable().
			Default(func() time.Time {
				return time.Now()
			}).
			Annotations(
				entproto.Field(3),
			),
	}
}

func (Chat) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("sent_user", User.Type).
			Ref("sent_messages").
			Annotations(entproto.Field(4)).
			Unique(),
		edge.From("received_user", User.Type).
			Ref("received_messages").
			Annotations(entproto.Field(5)).
			Unique(),
	}
}
