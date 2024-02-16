// Code generated by ent, DO NOT EDIT.

package ent

import (
	"Real-Time-Chat/ent/chat"
	"Real-Time-Chat/ent/predicate"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ChatUpdate is the builder for updating Chat entities.
type ChatUpdate struct {
	config
	hooks    []Hook
	mutation *ChatMutation
}

// Where appends a list predicates to the ChatUpdate builder.
func (cu *ChatUpdate) Where(ps ...predicate.Chat) *ChatUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetSenderID sets the "sender_id" field.
func (cu *ChatUpdate) SetSenderID(i int) *ChatUpdate {
	cu.mutation.ResetSenderID()
	cu.mutation.SetSenderID(i)
	return cu
}

// SetNillableSenderID sets the "sender_id" field if the given value is not nil.
func (cu *ChatUpdate) SetNillableSenderID(i *int) *ChatUpdate {
	if i != nil {
		cu.SetSenderID(*i)
	}
	return cu
}

// AddSenderID adds i to the "sender_id" field.
func (cu *ChatUpdate) AddSenderID(i int) *ChatUpdate {
	cu.mutation.AddSenderID(i)
	return cu
}

// SetReceiverID sets the "receiver_id" field.
func (cu *ChatUpdate) SetReceiverID(i int) *ChatUpdate {
	cu.mutation.ResetReceiverID()
	cu.mutation.SetReceiverID(i)
	return cu
}

// SetNillableReceiverID sets the "receiver_id" field if the given value is not nil.
func (cu *ChatUpdate) SetNillableReceiverID(i *int) *ChatUpdate {
	if i != nil {
		cu.SetReceiverID(*i)
	}
	return cu
}

// AddReceiverID adds i to the "receiver_id" field.
func (cu *ChatUpdate) AddReceiverID(i int) *ChatUpdate {
	cu.mutation.AddReceiverID(i)
	return cu
}

// SetMessage sets the "message" field.
func (cu *ChatUpdate) SetMessage(s string) *ChatUpdate {
	cu.mutation.SetMessage(s)
	return cu
}

// SetNillableMessage sets the "message" field if the given value is not nil.
func (cu *ChatUpdate) SetNillableMessage(s *string) *ChatUpdate {
	if s != nil {
		cu.SetMessage(*s)
	}
	return cu
}

// SetSentAt sets the "sent_at" field.
func (cu *ChatUpdate) SetSentAt(t time.Time) *ChatUpdate {
	cu.mutation.SetSentAt(t)
	return cu
}

// SetNillableSentAt sets the "sent_at" field if the given value is not nil.
func (cu *ChatUpdate) SetNillableSentAt(t *time.Time) *ChatUpdate {
	if t != nil {
		cu.SetSentAt(*t)
	}
	return cu
}

// Mutation returns the ChatMutation object of the builder.
func (cu *ChatUpdate) Mutation() *ChatMutation {
	return cu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *ChatUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, cu.sqlSave, cu.mutation, cu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cu *ChatUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *ChatUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *ChatUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cu *ChatUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(chat.Table, chat.Columns, sqlgraph.NewFieldSpec(chat.FieldID, field.TypeInt))
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.SenderID(); ok {
		_spec.SetField(chat.FieldSenderID, field.TypeInt, value)
	}
	if value, ok := cu.mutation.AddedSenderID(); ok {
		_spec.AddField(chat.FieldSenderID, field.TypeInt, value)
	}
	if value, ok := cu.mutation.ReceiverID(); ok {
		_spec.SetField(chat.FieldReceiverID, field.TypeInt, value)
	}
	if value, ok := cu.mutation.AddedReceiverID(); ok {
		_spec.AddField(chat.FieldReceiverID, field.TypeInt, value)
	}
	if value, ok := cu.mutation.Message(); ok {
		_spec.SetField(chat.FieldMessage, field.TypeString, value)
	}
	if value, ok := cu.mutation.SentAt(); ok {
		_spec.SetField(chat.FieldSentAt, field.TypeTime, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{chat.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cu.mutation.done = true
	return n, nil
}

// ChatUpdateOne is the builder for updating a single Chat entity.
type ChatUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ChatMutation
}

// SetSenderID sets the "sender_id" field.
func (cuo *ChatUpdateOne) SetSenderID(i int) *ChatUpdateOne {
	cuo.mutation.ResetSenderID()
	cuo.mutation.SetSenderID(i)
	return cuo
}

// SetNillableSenderID sets the "sender_id" field if the given value is not nil.
func (cuo *ChatUpdateOne) SetNillableSenderID(i *int) *ChatUpdateOne {
	if i != nil {
		cuo.SetSenderID(*i)
	}
	return cuo
}

// AddSenderID adds i to the "sender_id" field.
func (cuo *ChatUpdateOne) AddSenderID(i int) *ChatUpdateOne {
	cuo.mutation.AddSenderID(i)
	return cuo
}

// SetReceiverID sets the "receiver_id" field.
func (cuo *ChatUpdateOne) SetReceiverID(i int) *ChatUpdateOne {
	cuo.mutation.ResetReceiverID()
	cuo.mutation.SetReceiverID(i)
	return cuo
}

// SetNillableReceiverID sets the "receiver_id" field if the given value is not nil.
func (cuo *ChatUpdateOne) SetNillableReceiverID(i *int) *ChatUpdateOne {
	if i != nil {
		cuo.SetReceiverID(*i)
	}
	return cuo
}

// AddReceiverID adds i to the "receiver_id" field.
func (cuo *ChatUpdateOne) AddReceiverID(i int) *ChatUpdateOne {
	cuo.mutation.AddReceiverID(i)
	return cuo
}

// SetMessage sets the "message" field.
func (cuo *ChatUpdateOne) SetMessage(s string) *ChatUpdateOne {
	cuo.mutation.SetMessage(s)
	return cuo
}

// SetNillableMessage sets the "message" field if the given value is not nil.
func (cuo *ChatUpdateOne) SetNillableMessage(s *string) *ChatUpdateOne {
	if s != nil {
		cuo.SetMessage(*s)
	}
	return cuo
}

// SetSentAt sets the "sent_at" field.
func (cuo *ChatUpdateOne) SetSentAt(t time.Time) *ChatUpdateOne {
	cuo.mutation.SetSentAt(t)
	return cuo
}

// SetNillableSentAt sets the "sent_at" field if the given value is not nil.
func (cuo *ChatUpdateOne) SetNillableSentAt(t *time.Time) *ChatUpdateOne {
	if t != nil {
		cuo.SetSentAt(*t)
	}
	return cuo
}

// Mutation returns the ChatMutation object of the builder.
func (cuo *ChatUpdateOne) Mutation() *ChatMutation {
	return cuo.mutation
}

// Where appends a list predicates to the ChatUpdate builder.
func (cuo *ChatUpdateOne) Where(ps ...predicate.Chat) *ChatUpdateOne {
	cuo.mutation.Where(ps...)
	return cuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *ChatUpdateOne) Select(field string, fields ...string) *ChatUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Chat entity.
func (cuo *ChatUpdateOne) Save(ctx context.Context) (*Chat, error) {
	return withHooks(ctx, cuo.sqlSave, cuo.mutation, cuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *ChatUpdateOne) SaveX(ctx context.Context) *Chat {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *ChatUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *ChatUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cuo *ChatUpdateOne) sqlSave(ctx context.Context) (_node *Chat, err error) {
	_spec := sqlgraph.NewUpdateSpec(chat.Table, chat.Columns, sqlgraph.NewFieldSpec(chat.FieldID, field.TypeInt))
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Chat.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, chat.FieldID)
		for _, f := range fields {
			if !chat.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != chat.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.SenderID(); ok {
		_spec.SetField(chat.FieldSenderID, field.TypeInt, value)
	}
	if value, ok := cuo.mutation.AddedSenderID(); ok {
		_spec.AddField(chat.FieldSenderID, field.TypeInt, value)
	}
	if value, ok := cuo.mutation.ReceiverID(); ok {
		_spec.SetField(chat.FieldReceiverID, field.TypeInt, value)
	}
	if value, ok := cuo.mutation.AddedReceiverID(); ok {
		_spec.AddField(chat.FieldReceiverID, field.TypeInt, value)
	}
	if value, ok := cuo.mutation.Message(); ok {
		_spec.SetField(chat.FieldMessage, field.TypeString, value)
	}
	if value, ok := cuo.mutation.SentAt(); ok {
		_spec.SetField(chat.FieldSentAt, field.TypeTime, value)
	}
	_node = &Chat{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{chat.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cuo.mutation.done = true
	return _node, nil
}