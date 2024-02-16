package schema

import (
	"entgo.io/contrib/entproto"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
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
		field.Int("sender_id").
			Annotations(
				entproto.Field(2),
			),
		field.Int("receiver_id").
			Annotations(
				entproto.Field(3),
			),
		field.String("message").
			Annotations(
				entproto.Field(4),
			),
		field.Time("sent_at").
			Annotations(
				entproto.Field(5),
			),
	}
}
