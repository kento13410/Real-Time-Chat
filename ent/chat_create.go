// Code generated by ent, DO NOT EDIT.

package ent

import (
	"Real-Time-Chat/ent/chat"
	"Real-Time-Chat/ent/user"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ChatCreate is the builder for creating a Chat entity.
type ChatCreate struct {
	config
	mutation *ChatMutation
	hooks    []Hook
}

// SetMessage sets the "message" field.
func (cc *ChatCreate) SetMessage(s string) *ChatCreate {
	cc.mutation.SetMessage(s)
	return cc
}

// SetSentAt sets the "sent_at" field.
func (cc *ChatCreate) SetSentAt(t time.Time) *ChatCreate {
	cc.mutation.SetSentAt(t)
	return cc
}

// SetNillableSentAt sets the "sent_at" field if the given value is not nil.
func (cc *ChatCreate) SetNillableSentAt(t *time.Time) *ChatCreate {
	if t != nil {
		cc.SetSentAt(*t)
	}
	return cc
}

// SetSentUserID sets the "sent_user" edge to the User entity by ID.
func (cc *ChatCreate) SetSentUserID(id int) *ChatCreate {
	cc.mutation.SetSentUserID(id)
	return cc
}

// SetNillableSentUserID sets the "sent_user" edge to the User entity by ID if the given value is not nil.
func (cc *ChatCreate) SetNillableSentUserID(id *int) *ChatCreate {
	if id != nil {
		cc = cc.SetSentUserID(*id)
	}
	return cc
}

// SetSentUser sets the "sent_user" edge to the User entity.
func (cc *ChatCreate) SetSentUser(u *User) *ChatCreate {
	return cc.SetSentUserID(u.ID)
}

// SetReceivedUserID sets the "received_user" edge to the User entity by ID.
func (cc *ChatCreate) SetReceivedUserID(id int) *ChatCreate {
	cc.mutation.SetReceivedUserID(id)
	return cc
}

// SetNillableReceivedUserID sets the "received_user" edge to the User entity by ID if the given value is not nil.
func (cc *ChatCreate) SetNillableReceivedUserID(id *int) *ChatCreate {
	if id != nil {
		cc = cc.SetReceivedUserID(*id)
	}
	return cc
}

// SetReceivedUser sets the "received_user" edge to the User entity.
func (cc *ChatCreate) SetReceivedUser(u *User) *ChatCreate {
	return cc.SetReceivedUserID(u.ID)
}

// Mutation returns the ChatMutation object of the builder.
func (cc *ChatCreate) Mutation() *ChatMutation {
	return cc.mutation
}

// Save creates the Chat in the database.
func (cc *ChatCreate) Save(ctx context.Context) (*Chat, error) {
	cc.defaults()
	return withHooks(ctx, cc.sqlSave, cc.mutation, cc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (cc *ChatCreate) SaveX(ctx context.Context) *Chat {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cc *ChatCreate) Exec(ctx context.Context) error {
	_, err := cc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cc *ChatCreate) ExecX(ctx context.Context) {
	if err := cc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cc *ChatCreate) defaults() {
	if _, ok := cc.mutation.SentAt(); !ok {
		v := chat.DefaultSentAt()
		cc.mutation.SetSentAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cc *ChatCreate) check() error {
	if _, ok := cc.mutation.Message(); !ok {
		return &ValidationError{Name: "message", err: errors.New(`ent: missing required field "Chat.message"`)}
	}
	if _, ok := cc.mutation.SentAt(); !ok {
		return &ValidationError{Name: "sent_at", err: errors.New(`ent: missing required field "Chat.sent_at"`)}
	}
	return nil
}

func (cc *ChatCreate) sqlSave(ctx context.Context) (*Chat, error) {
	if err := cc.check(); err != nil {
		return nil, err
	}
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	cc.mutation.id = &_node.ID
	cc.mutation.done = true
	return _node, nil
}

func (cc *ChatCreate) createSpec() (*Chat, *sqlgraph.CreateSpec) {
	var (
		_node = &Chat{config: cc.config}
		_spec = sqlgraph.NewCreateSpec(chat.Table, sqlgraph.NewFieldSpec(chat.FieldID, field.TypeInt))
	)
	if value, ok := cc.mutation.Message(); ok {
		_spec.SetField(chat.FieldMessage, field.TypeString, value)
		_node.Message = value
	}
	if value, ok := cc.mutation.SentAt(); ok {
		_spec.SetField(chat.FieldSentAt, field.TypeTime, value)
		_node.SentAt = value
	}
	if nodes := cc.mutation.SentUserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   chat.SentUserTable,
			Columns: []string{chat.SentUserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_sent_messages = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cc.mutation.ReceivedUserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   chat.ReceivedUserTable,
			Columns: []string{chat.ReceivedUserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_received_messages = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ChatCreateBulk is the builder for creating many Chat entities in bulk.
type ChatCreateBulk struct {
	config
	err      error
	builders []*ChatCreate
}

// Save creates the Chat entities in the database.
func (ccb *ChatCreateBulk) Save(ctx context.Context) ([]*Chat, error) {
	if ccb.err != nil {
		return nil, ccb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Chat, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ChatMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ccb *ChatCreateBulk) SaveX(ctx context.Context) []*Chat {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccb *ChatCreateBulk) Exec(ctx context.Context) error {
	_, err := ccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccb *ChatCreateBulk) ExecX(ctx context.Context) {
	if err := ccb.Exec(ctx); err != nil {
		panic(err)
	}
}
