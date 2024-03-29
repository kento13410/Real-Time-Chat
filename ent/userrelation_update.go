// Code generated by ent, DO NOT EDIT.

package ent

import (
	"Real-Time-Chat/ent/predicate"
	"Real-Time-Chat/ent/user"
	"Real-Time-Chat/ent/userrelation"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserRelationUpdate is the builder for updating UserRelation entities.
type UserRelationUpdate struct {
	config
	hooks    []Hook
	mutation *UserRelationMutation
}

// Where appends a list predicates to the UserRelationUpdate builder.
func (uru *UserRelationUpdate) Where(ps ...predicate.UserRelation) *UserRelationUpdate {
	uru.mutation.Where(ps...)
	return uru
}

// SetUser1ID sets the "user1" edge to the User entity by ID.
func (uru *UserRelationUpdate) SetUser1ID(id int) *UserRelationUpdate {
	uru.mutation.SetUser1ID(id)
	return uru
}

// SetNillableUser1ID sets the "user1" edge to the User entity by ID if the given value is not nil.
func (uru *UserRelationUpdate) SetNillableUser1ID(id *int) *UserRelationUpdate {
	if id != nil {
		uru = uru.SetUser1ID(*id)
	}
	return uru
}

// SetUser1 sets the "user1" edge to the User entity.
func (uru *UserRelationUpdate) SetUser1(u *User) *UserRelationUpdate {
	return uru.SetUser1ID(u.ID)
}

// SetUser2ID sets the "user2" edge to the User entity by ID.
func (uru *UserRelationUpdate) SetUser2ID(id int) *UserRelationUpdate {
	uru.mutation.SetUser2ID(id)
	return uru
}

// SetNillableUser2ID sets the "user2" edge to the User entity by ID if the given value is not nil.
func (uru *UserRelationUpdate) SetNillableUser2ID(id *int) *UserRelationUpdate {
	if id != nil {
		uru = uru.SetUser2ID(*id)
	}
	return uru
}

// SetUser2 sets the "user2" edge to the User entity.
func (uru *UserRelationUpdate) SetUser2(u *User) *UserRelationUpdate {
	return uru.SetUser2ID(u.ID)
}

// Mutation returns the UserRelationMutation object of the builder.
func (uru *UserRelationUpdate) Mutation() *UserRelationMutation {
	return uru.mutation
}

// ClearUser1 clears the "user1" edge to the User entity.
func (uru *UserRelationUpdate) ClearUser1() *UserRelationUpdate {
	uru.mutation.ClearUser1()
	return uru
}

// ClearUser2 clears the "user2" edge to the User entity.
func (uru *UserRelationUpdate) ClearUser2() *UserRelationUpdate {
	uru.mutation.ClearUser2()
	return uru
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uru *UserRelationUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, uru.sqlSave, uru.mutation, uru.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uru *UserRelationUpdate) SaveX(ctx context.Context) int {
	affected, err := uru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uru *UserRelationUpdate) Exec(ctx context.Context) error {
	_, err := uru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uru *UserRelationUpdate) ExecX(ctx context.Context) {
	if err := uru.Exec(ctx); err != nil {
		panic(err)
	}
}

func (uru *UserRelationUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(userrelation.Table, userrelation.Columns, sqlgraph.NewFieldSpec(userrelation.FieldID, field.TypeInt))
	if ps := uru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if uru.mutation.User1Cleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uru.mutation.User1IDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uru.mutation.User2Cleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uru.mutation.User2IDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{userrelation.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	uru.mutation.done = true
	return n, nil
}

// UserRelationUpdateOne is the builder for updating a single UserRelation entity.
type UserRelationUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserRelationMutation
}

// SetUser1ID sets the "user1" edge to the User entity by ID.
func (uruo *UserRelationUpdateOne) SetUser1ID(id int) *UserRelationUpdateOne {
	uruo.mutation.SetUser1ID(id)
	return uruo
}

// SetNillableUser1ID sets the "user1" edge to the User entity by ID if the given value is not nil.
func (uruo *UserRelationUpdateOne) SetNillableUser1ID(id *int) *UserRelationUpdateOne {
	if id != nil {
		uruo = uruo.SetUser1ID(*id)
	}
	return uruo
}

// SetUser1 sets the "user1" edge to the User entity.
func (uruo *UserRelationUpdateOne) SetUser1(u *User) *UserRelationUpdateOne {
	return uruo.SetUser1ID(u.ID)
}

// SetUser2ID sets the "user2" edge to the User entity by ID.
func (uruo *UserRelationUpdateOne) SetUser2ID(id int) *UserRelationUpdateOne {
	uruo.mutation.SetUser2ID(id)
	return uruo
}

// SetNillableUser2ID sets the "user2" edge to the User entity by ID if the given value is not nil.
func (uruo *UserRelationUpdateOne) SetNillableUser2ID(id *int) *UserRelationUpdateOne {
	if id != nil {
		uruo = uruo.SetUser2ID(*id)
	}
	return uruo
}

// SetUser2 sets the "user2" edge to the User entity.
func (uruo *UserRelationUpdateOne) SetUser2(u *User) *UserRelationUpdateOne {
	return uruo.SetUser2ID(u.ID)
}

// Mutation returns the UserRelationMutation object of the builder.
func (uruo *UserRelationUpdateOne) Mutation() *UserRelationMutation {
	return uruo.mutation
}

// ClearUser1 clears the "user1" edge to the User entity.
func (uruo *UserRelationUpdateOne) ClearUser1() *UserRelationUpdateOne {
	uruo.mutation.ClearUser1()
	return uruo
}

// ClearUser2 clears the "user2" edge to the User entity.
func (uruo *UserRelationUpdateOne) ClearUser2() *UserRelationUpdateOne {
	uruo.mutation.ClearUser2()
	return uruo
}

// Where appends a list predicates to the UserRelationUpdate builder.
func (uruo *UserRelationUpdateOne) Where(ps ...predicate.UserRelation) *UserRelationUpdateOne {
	uruo.mutation.Where(ps...)
	return uruo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uruo *UserRelationUpdateOne) Select(field string, fields ...string) *UserRelationUpdateOne {
	uruo.fields = append([]string{field}, fields...)
	return uruo
}

// Save executes the query and returns the updated UserRelation entity.
func (uruo *UserRelationUpdateOne) Save(ctx context.Context) (*UserRelation, error) {
	return withHooks(ctx, uruo.sqlSave, uruo.mutation, uruo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uruo *UserRelationUpdateOne) SaveX(ctx context.Context) *UserRelation {
	node, err := uruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uruo *UserRelationUpdateOne) Exec(ctx context.Context) error {
	_, err := uruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uruo *UserRelationUpdateOne) ExecX(ctx context.Context) {
	if err := uruo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (uruo *UserRelationUpdateOne) sqlSave(ctx context.Context) (_node *UserRelation, err error) {
	_spec := sqlgraph.NewUpdateSpec(userrelation.Table, userrelation.Columns, sqlgraph.NewFieldSpec(userrelation.FieldID, field.TypeInt))
	id, ok := uruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "UserRelation.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, userrelation.FieldID)
		for _, f := range fields {
			if !userrelation.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != userrelation.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if uruo.mutation.User1Cleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uruo.mutation.User1IDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uruo.mutation.User2Cleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uruo.mutation.User2IDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &UserRelation{config: uruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{userrelation.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	uruo.mutation.done = true
	return _node, nil
}
