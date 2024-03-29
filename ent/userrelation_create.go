// Code generated by ent, DO NOT EDIT.

package ent

import (
	"Real-Time-Chat/ent/user"
	"Real-Time-Chat/ent/userrelation"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserRelationCreate is the builder for creating a UserRelation entity.
type UserRelationCreate struct {
	config
	mutation *UserRelationMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (urc *UserRelationCreate) SetCreatedAt(t time.Time) *UserRelationCreate {
	urc.mutation.SetCreatedAt(t)
	return urc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (urc *UserRelationCreate) SetNillableCreatedAt(t *time.Time) *UserRelationCreate {
	if t != nil {
		urc.SetCreatedAt(*t)
	}
	return urc
}

// SetUser1ID sets the "user1" edge to the User entity by ID.
func (urc *UserRelationCreate) SetUser1ID(id int) *UserRelationCreate {
	urc.mutation.SetUser1ID(id)
	return urc
}

// SetNillableUser1ID sets the "user1" edge to the User entity by ID if the given value is not nil.
func (urc *UserRelationCreate) SetNillableUser1ID(id *int) *UserRelationCreate {
	if id != nil {
		urc = urc.SetUser1ID(*id)
	}
	return urc
}

// SetUser1 sets the "user1" edge to the User entity.
func (urc *UserRelationCreate) SetUser1(u *User) *UserRelationCreate {
	return urc.SetUser1ID(u.ID)
}

// SetUser2ID sets the "user2" edge to the User entity by ID.
func (urc *UserRelationCreate) SetUser2ID(id int) *UserRelationCreate {
	urc.mutation.SetUser2ID(id)
	return urc
}

// SetNillableUser2ID sets the "user2" edge to the User entity by ID if the given value is not nil.
func (urc *UserRelationCreate) SetNillableUser2ID(id *int) *UserRelationCreate {
	if id != nil {
		urc = urc.SetUser2ID(*id)
	}
	return urc
}

// SetUser2 sets the "user2" edge to the User entity.
func (urc *UserRelationCreate) SetUser2(u *User) *UserRelationCreate {
	return urc.SetUser2ID(u.ID)
}

// Mutation returns the UserRelationMutation object of the builder.
func (urc *UserRelationCreate) Mutation() *UserRelationMutation {
	return urc.mutation
}

// Save creates the UserRelation in the database.
func (urc *UserRelationCreate) Save(ctx context.Context) (*UserRelation, error) {
	urc.defaults()
	return withHooks(ctx, urc.sqlSave, urc.mutation, urc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (urc *UserRelationCreate) SaveX(ctx context.Context) *UserRelation {
	v, err := urc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (urc *UserRelationCreate) Exec(ctx context.Context) error {
	_, err := urc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (urc *UserRelationCreate) ExecX(ctx context.Context) {
	if err := urc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (urc *UserRelationCreate) defaults() {
	if _, ok := urc.mutation.CreatedAt(); !ok {
		v := userrelation.DefaultCreatedAt()
		urc.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (urc *UserRelationCreate) check() error {
	if _, ok := urc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "UserRelation.created_at"`)}
	}
	return nil
}

func (urc *UserRelationCreate) sqlSave(ctx context.Context) (*UserRelation, error) {
	if err := urc.check(); err != nil {
		return nil, err
	}
	_node, _spec := urc.createSpec()
	if err := sqlgraph.CreateNode(ctx, urc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	urc.mutation.id = &_node.ID
	urc.mutation.done = true
	return _node, nil
}

func (urc *UserRelationCreate) createSpec() (*UserRelation, *sqlgraph.CreateSpec) {
	var (
		_node = &UserRelation{config: urc.config}
		_spec = sqlgraph.NewCreateSpec(userrelation.Table, sqlgraph.NewFieldSpec(userrelation.FieldID, field.TypeInt))
	)
	if value, ok := urc.mutation.CreatedAt(); ok {
		_spec.SetField(userrelation.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if nodes := urc.mutation.User1IDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   userrelation.User1Table,
			Columns: []string{userrelation.User1Column},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_user_relations_1 = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := urc.mutation.User2IDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   userrelation.User2Table,
			Columns: []string{userrelation.User2Column},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_user_relations_2 = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// UserRelationCreateBulk is the builder for creating many UserRelation entities in bulk.
type UserRelationCreateBulk struct {
	config
	err      error
	builders []*UserRelationCreate
}

// Save creates the UserRelation entities in the database.
func (urcb *UserRelationCreateBulk) Save(ctx context.Context) ([]*UserRelation, error) {
	if urcb.err != nil {
		return nil, urcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(urcb.builders))
	nodes := make([]*UserRelation, len(urcb.builders))
	mutators := make([]Mutator, len(urcb.builders))
	for i := range urcb.builders {
		func(i int, root context.Context) {
			builder := urcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserRelationMutation)
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
					_, err = mutators[i+1].Mutate(root, urcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, urcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, urcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (urcb *UserRelationCreateBulk) SaveX(ctx context.Context) []*UserRelation {
	v, err := urcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (urcb *UserRelationCreateBulk) Exec(ctx context.Context) error {
	_, err := urcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (urcb *UserRelationCreateBulk) ExecX(ctx context.Context) {
	if err := urcb.Exec(ctx); err != nil {
		panic(err)
	}
}
