package schema

import (
	"entgo.io/contrib/entproto"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// UserRelation holds the schema definition for the UserRelations entity.
type UserRelation struct {
	ent.Schema
}

func (UserRelation) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
		entproto.Service(),
	}
}

// Fields of the UserRelations.
func (UserRelation) Fields() []ent.Field {
	return []ent.Field{
		field.Int("user_id_1").
			Annotations(
				entproto.Field(2),
			),
		field.Int("user_id_2").
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

func (UserRelation) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user1", User.Type).
			Ref("user_relations_1").
			Annotations(entproto.Field(5)).
			Unique(),
		edge.From("user2", User.Type).
			Ref("user_relations_2").
			Annotations(entproto.Field(6)).
			Unique(),
	}
}
