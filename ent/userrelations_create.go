// Code generated by ent, DO NOT EDIT.

package ent

import (
	"Real-Time-Chat/ent/userrelations"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserRelationsCreate is the builder for creating a UserRelations entity.
type UserRelationsCreate struct {
	config
	mutation *UserRelationsMutation
	hooks    []Hook
}

// SetUserID1 sets the "user_id_1" field.
func (urc *UserRelationsCreate) SetUserID1(i int) *UserRelationsCreate {
	urc.mutation.SetUserID1(i)
	return urc
}

// SetUserID2 sets the "user_id_2" field.
func (urc *UserRelationsCreate) SetUserID2(i int) *UserRelationsCreate {
	urc.mutation.SetUserID2(i)
	return urc
}

// SetID sets the "id" field.
func (urc *UserRelationsCreate) SetID(i int) *UserRelationsCreate {
	urc.mutation.SetID(i)
	return urc
}

// Mutation returns the UserRelationsMutation object of the builder.
func (urc *UserRelationsCreate) Mutation() *UserRelationsMutation {
	return urc.mutation
}

// Save creates the UserRelations in the database.
func (urc *UserRelationsCreate) Save(ctx context.Context) (*UserRelations, error) {
	return withHooks(ctx, urc.sqlSave, urc.mutation, urc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (urc *UserRelationsCreate) SaveX(ctx context.Context) *UserRelations {
	v, err := urc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (urc *UserRelationsCreate) Exec(ctx context.Context) error {
	_, err := urc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (urc *UserRelationsCreate) ExecX(ctx context.Context) {
	if err := urc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (urc *UserRelationsCreate) check() error {
	if _, ok := urc.mutation.UserID1(); !ok {
		return &ValidationError{Name: "user_id_1", err: errors.New(`ent: missing required field "UserRelations.user_id_1"`)}
	}
	if _, ok := urc.mutation.UserID2(); !ok {
		return &ValidationError{Name: "user_id_2", err: errors.New(`ent: missing required field "UserRelations.user_id_2"`)}
	}
	return nil
}

func (urc *UserRelationsCreate) sqlSave(ctx context.Context) (*UserRelations, error) {
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
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int(id)
	}
	urc.mutation.id = &_node.ID
	urc.mutation.done = true
	return _node, nil
}

func (urc *UserRelationsCreate) createSpec() (*UserRelations, *sqlgraph.CreateSpec) {
	var (
		_node = &UserRelations{config: urc.config}
		_spec = sqlgraph.NewCreateSpec(userrelations.Table, sqlgraph.NewFieldSpec(userrelations.FieldID, field.TypeInt))
	)
	if id, ok := urc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := urc.mutation.UserID1(); ok {
		_spec.SetField(userrelations.FieldUserID1, field.TypeInt, value)
		_node.UserID1 = value
	}
	if value, ok := urc.mutation.UserID2(); ok {
		_spec.SetField(userrelations.FieldUserID2, field.TypeInt, value)
		_node.UserID2 = value
	}
	return _node, _spec
}

// UserRelationsCreateBulk is the builder for creating many UserRelations entities in bulk.
type UserRelationsCreateBulk struct {
	config
	err      error
	builders []*UserRelationsCreate
}

// Save creates the UserRelations entities in the database.
func (urcb *UserRelationsCreateBulk) Save(ctx context.Context) ([]*UserRelations, error) {
	if urcb.err != nil {
		return nil, urcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(urcb.builders))
	nodes := make([]*UserRelations, len(urcb.builders))
	mutators := make([]Mutator, len(urcb.builders))
	for i := range urcb.builders {
		func(i int, root context.Context) {
			builder := urcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserRelationsMutation)
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
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
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
func (urcb *UserRelationsCreateBulk) SaveX(ctx context.Context) []*UserRelations {
	v, err := urcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (urcb *UserRelationsCreateBulk) Exec(ctx context.Context) error {
	_, err := urcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (urcb *UserRelationsCreateBulk) ExecX(ctx context.Context) {
	if err := urcb.Exec(ctx); err != nil {
		panic(err)
	}
}
