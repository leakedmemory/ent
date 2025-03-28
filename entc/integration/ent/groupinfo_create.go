// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entc/integration/ent/group"
	"entgo.io/ent/entc/integration/ent/groupinfo"
	"entgo.io/ent/schema/field"
)

// GroupInfoCreate is the builder for creating a GroupInfo entity.
type GroupInfoCreate struct {
	config
	mutation *GroupInfoMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetDesc sets the "desc" field.
func (_c *GroupInfoCreate) SetDesc(v string) *GroupInfoCreate {
	_c.mutation.SetDesc(v)
	return _c
}

// SetMaxUsers sets the "max_users" field.
func (_c *GroupInfoCreate) SetMaxUsers(v int) *GroupInfoCreate {
	_c.mutation.SetMaxUsers(v)
	return _c
}

// SetNillableMaxUsers sets the "max_users" field if the given value is not nil.
func (_c *GroupInfoCreate) SetNillableMaxUsers(v *int) *GroupInfoCreate {
	if v != nil {
		_c.SetMaxUsers(*v)
	}
	return _c
}

// AddGroupIDs adds the "groups" edge to the Group entity by IDs.
func (_c *GroupInfoCreate) AddGroupIDs(ids ...int) *GroupInfoCreate {
	_c.mutation.AddGroupIDs(ids...)
	return _c
}

// AddGroups adds the "groups" edges to the Group entity.
func (_c *GroupInfoCreate) AddGroups(v ...*Group) *GroupInfoCreate {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return _c.AddGroupIDs(ids...)
}

// Mutation returns the GroupInfoMutation object of the builder.
func (_c *GroupInfoCreate) Mutation() *GroupInfoMutation {
	return _c.mutation
}

// Save creates the GroupInfo in the database.
func (_c *GroupInfoCreate) Save(ctx context.Context) (*GroupInfo, error) {
	_c.defaults()
	return withHooks(ctx, _c.sqlSave, _c.mutation, _c.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (_c *GroupInfoCreate) SaveX(ctx context.Context) *GroupInfo {
	v, err := _c.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (_c *GroupInfoCreate) Exec(ctx context.Context) error {
	_, err := _c.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (_c *GroupInfoCreate) ExecX(ctx context.Context) {
	if err := _c.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (_c *GroupInfoCreate) defaults() {
	if _, ok := _c.mutation.MaxUsers(); !ok {
		v := groupinfo.DefaultMaxUsers
		_c.mutation.SetMaxUsers(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (_c *GroupInfoCreate) check() error {
	if _, ok := _c.mutation.Desc(); !ok {
		return &ValidationError{Name: "desc", err: errors.New(`ent: missing required field "GroupInfo.desc"`)}
	}
	if _, ok := _c.mutation.MaxUsers(); !ok {
		return &ValidationError{Name: "max_users", err: errors.New(`ent: missing required field "GroupInfo.max_users"`)}
	}
	return nil
}

func (_c *GroupInfoCreate) sqlSave(ctx context.Context) (*GroupInfo, error) {
	if err := _c.check(); err != nil {
		return nil, err
	}
	_node, _spec := _c.createSpec()
	if err := sqlgraph.CreateNode(ctx, _c.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	_c.mutation.id = &_node.ID
	_c.mutation.done = true
	return _node, nil
}

func (_c *GroupInfoCreate) createSpec() (*GroupInfo, *sqlgraph.CreateSpec) {
	var (
		_node = &GroupInfo{config: _c.config}
		_spec = sqlgraph.NewCreateSpec(groupinfo.Table, sqlgraph.NewFieldSpec(groupinfo.FieldID, field.TypeInt))
	)
	_spec.OnConflict = _c.conflict
	if value, ok := _c.mutation.Desc(); ok {
		_spec.SetField(groupinfo.FieldDesc, field.TypeString, value)
		_node.Desc = value
	}
	if value, ok := _c.mutation.MaxUsers(); ok {
		_spec.SetField(groupinfo.FieldMaxUsers, field.TypeInt, value)
		_node.MaxUsers = value
	}
	if nodes := _c.mutation.GroupsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   groupinfo.GroupsTable,
			Columns: []string{groupinfo.GroupsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.GroupInfo.Create().
//		SetDesc(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.GroupInfoUpsert) {
//			SetDesc(v+v).
//		}).
//		Exec(ctx)
func (_c *GroupInfoCreate) OnConflict(opts ...sql.ConflictOption) *GroupInfoUpsertOne {
	_c.conflict = opts
	return &GroupInfoUpsertOne{
		create: _c,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.GroupInfo.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (_c *GroupInfoCreate) OnConflictColumns(columns ...string) *GroupInfoUpsertOne {
	_c.conflict = append(_c.conflict, sql.ConflictColumns(columns...))
	return &GroupInfoUpsertOne{
		create: _c,
	}
}

type (
	// GroupInfoUpsertOne is the builder for "upsert"-ing
	//  one GroupInfo node.
	GroupInfoUpsertOne struct {
		create *GroupInfoCreate
	}

	// GroupInfoUpsert is the "OnConflict" setter.
	GroupInfoUpsert struct {
		*sql.UpdateSet
	}
)

// SetDesc sets the "desc" field.
func (u *GroupInfoUpsert) SetDesc(v string) *GroupInfoUpsert {
	u.Set(groupinfo.FieldDesc, v)
	return u
}

// UpdateDesc sets the "desc" field to the value that was provided on create.
func (u *GroupInfoUpsert) UpdateDesc() *GroupInfoUpsert {
	u.SetExcluded(groupinfo.FieldDesc)
	return u
}

// SetMaxUsers sets the "max_users" field.
func (u *GroupInfoUpsert) SetMaxUsers(v int) *GroupInfoUpsert {
	u.Set(groupinfo.FieldMaxUsers, v)
	return u
}

// UpdateMaxUsers sets the "max_users" field to the value that was provided on create.
func (u *GroupInfoUpsert) UpdateMaxUsers() *GroupInfoUpsert {
	u.SetExcluded(groupinfo.FieldMaxUsers)
	return u
}

// AddMaxUsers adds v to the "max_users" field.
func (u *GroupInfoUpsert) AddMaxUsers(v int) *GroupInfoUpsert {
	u.Add(groupinfo.FieldMaxUsers, v)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.GroupInfo.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *GroupInfoUpsertOne) UpdateNewValues() *GroupInfoUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.GroupInfo.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *GroupInfoUpsertOne) Ignore() *GroupInfoUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *GroupInfoUpsertOne) DoNothing() *GroupInfoUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the GroupInfoCreate.OnConflict
// documentation for more info.
func (u *GroupInfoUpsertOne) Update(set func(*GroupInfoUpsert)) *GroupInfoUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&GroupInfoUpsert{UpdateSet: update})
	}))
	return u
}

// SetDesc sets the "desc" field.
func (u *GroupInfoUpsertOne) SetDesc(v string) *GroupInfoUpsertOne {
	return u.Update(func(s *GroupInfoUpsert) {
		s.SetDesc(v)
	})
}

// UpdateDesc sets the "desc" field to the value that was provided on create.
func (u *GroupInfoUpsertOne) UpdateDesc() *GroupInfoUpsertOne {
	return u.Update(func(s *GroupInfoUpsert) {
		s.UpdateDesc()
	})
}

// SetMaxUsers sets the "max_users" field.
func (u *GroupInfoUpsertOne) SetMaxUsers(v int) *GroupInfoUpsertOne {
	return u.Update(func(s *GroupInfoUpsert) {
		s.SetMaxUsers(v)
	})
}

// AddMaxUsers adds v to the "max_users" field.
func (u *GroupInfoUpsertOne) AddMaxUsers(v int) *GroupInfoUpsertOne {
	return u.Update(func(s *GroupInfoUpsert) {
		s.AddMaxUsers(v)
	})
}

// UpdateMaxUsers sets the "max_users" field to the value that was provided on create.
func (u *GroupInfoUpsertOne) UpdateMaxUsers() *GroupInfoUpsertOne {
	return u.Update(func(s *GroupInfoUpsert) {
		s.UpdateMaxUsers()
	})
}

// Exec executes the query.
func (u *GroupInfoUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for GroupInfoCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *GroupInfoUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *GroupInfoUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *GroupInfoUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// GroupInfoCreateBulk is the builder for creating many GroupInfo entities in bulk.
type GroupInfoCreateBulk struct {
	config
	err      error
	builders []*GroupInfoCreate
	conflict []sql.ConflictOption
}

// Save creates the GroupInfo entities in the database.
func (_c *GroupInfoCreateBulk) Save(ctx context.Context) ([]*GroupInfo, error) {
	if _c.err != nil {
		return nil, _c.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(_c.builders))
	nodes := make([]*GroupInfo, len(_c.builders))
	mutators := make([]Mutator, len(_c.builders))
	for i := range _c.builders {
		func(i int, root context.Context) {
			builder := _c.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*GroupInfoMutation)
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
					_, err = mutators[i+1].Mutate(root, _c.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = _c.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, _c.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, _c.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (_c *GroupInfoCreateBulk) SaveX(ctx context.Context) []*GroupInfo {
	v, err := _c.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (_c *GroupInfoCreateBulk) Exec(ctx context.Context) error {
	_, err := _c.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (_c *GroupInfoCreateBulk) ExecX(ctx context.Context) {
	if err := _c.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.GroupInfo.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.GroupInfoUpsert) {
//			SetDesc(v+v).
//		}).
//		Exec(ctx)
func (_c *GroupInfoCreateBulk) OnConflict(opts ...sql.ConflictOption) *GroupInfoUpsertBulk {
	_c.conflict = opts
	return &GroupInfoUpsertBulk{
		create: _c,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.GroupInfo.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (_c *GroupInfoCreateBulk) OnConflictColumns(columns ...string) *GroupInfoUpsertBulk {
	_c.conflict = append(_c.conflict, sql.ConflictColumns(columns...))
	return &GroupInfoUpsertBulk{
		create: _c,
	}
}

// GroupInfoUpsertBulk is the builder for "upsert"-ing
// a bulk of GroupInfo nodes.
type GroupInfoUpsertBulk struct {
	create *GroupInfoCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.GroupInfo.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *GroupInfoUpsertBulk) UpdateNewValues() *GroupInfoUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.GroupInfo.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *GroupInfoUpsertBulk) Ignore() *GroupInfoUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *GroupInfoUpsertBulk) DoNothing() *GroupInfoUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the GroupInfoCreateBulk.OnConflict
// documentation for more info.
func (u *GroupInfoUpsertBulk) Update(set func(*GroupInfoUpsert)) *GroupInfoUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&GroupInfoUpsert{UpdateSet: update})
	}))
	return u
}

// SetDesc sets the "desc" field.
func (u *GroupInfoUpsertBulk) SetDesc(v string) *GroupInfoUpsertBulk {
	return u.Update(func(s *GroupInfoUpsert) {
		s.SetDesc(v)
	})
}

// UpdateDesc sets the "desc" field to the value that was provided on create.
func (u *GroupInfoUpsertBulk) UpdateDesc() *GroupInfoUpsertBulk {
	return u.Update(func(s *GroupInfoUpsert) {
		s.UpdateDesc()
	})
}

// SetMaxUsers sets the "max_users" field.
func (u *GroupInfoUpsertBulk) SetMaxUsers(v int) *GroupInfoUpsertBulk {
	return u.Update(func(s *GroupInfoUpsert) {
		s.SetMaxUsers(v)
	})
}

// AddMaxUsers adds v to the "max_users" field.
func (u *GroupInfoUpsertBulk) AddMaxUsers(v int) *GroupInfoUpsertBulk {
	return u.Update(func(s *GroupInfoUpsert) {
		s.AddMaxUsers(v)
	})
}

// UpdateMaxUsers sets the "max_users" field to the value that was provided on create.
func (u *GroupInfoUpsertBulk) UpdateMaxUsers() *GroupInfoUpsertBulk {
	return u.Update(func(s *GroupInfoUpsert) {
		s.UpdateMaxUsers()
	})
}

// Exec executes the query.
func (u *GroupInfoUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the GroupInfoCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for GroupInfoCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *GroupInfoUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
