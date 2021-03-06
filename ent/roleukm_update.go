// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/confus1on/UKM/ent/predicate"
	"github.com/confus1on/UKM/ent/profileukm"
	"github.com/confus1on/UKM/ent/roleukm"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// RoleUKMUpdate is the builder for updating RoleUKM entities.
type RoleUKMUpdate struct {
	config
	hooks      []Hook
	mutation   *RoleUKMMutation
	predicates []predicate.RoleUKM
}

// Where adds a new predicate for the builder.
func (ruu *RoleUKMUpdate) Where(ps ...predicate.RoleUKM) *RoleUKMUpdate {
	ruu.predicates = append(ruu.predicates, ps...)
	return ruu
}

// SetStatusRole sets the status_role field.
func (ruu *RoleUKMUpdate) SetStatusRole(s string) *RoleUKMUpdate {
	ruu.mutation.SetStatusRole(s)
	return ruu
}

// AddProfileRoleIDs adds the profile_roles edge to ProfileUKM by ids.
func (ruu *RoleUKMUpdate) AddProfileRoleIDs(ids ...int) *RoleUKMUpdate {
	ruu.mutation.AddProfileRoleIDs(ids...)
	return ruu
}

// AddProfileRoles adds the profile_roles edges to ProfileUKM.
func (ruu *RoleUKMUpdate) AddProfileRoles(p ...*ProfileUKM) *RoleUKMUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ruu.AddProfileRoleIDs(ids...)
}

// RemoveProfileRoleIDs removes the profile_roles edge to ProfileUKM by ids.
func (ruu *RoleUKMUpdate) RemoveProfileRoleIDs(ids ...int) *RoleUKMUpdate {
	ruu.mutation.RemoveProfileRoleIDs(ids...)
	return ruu
}

// RemoveProfileRoles removes profile_roles edges to ProfileUKM.
func (ruu *RoleUKMUpdate) RemoveProfileRoles(p ...*ProfileUKM) *RoleUKMUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ruu.RemoveProfileRoleIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (ruu *RoleUKMUpdate) Save(ctx context.Context) (int, error) {
	if v, ok := ruu.mutation.StatusRole(); ok {
		if err := roleukm.StatusRoleValidator(v); err != nil {
			return 0, fmt.Errorf("ent: validator failed for field \"status_role\": %v", err)
		}
	}

	var (
		err      error
		affected int
	)
	if len(ruu.hooks) == 0 {
		affected, err = ruu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RoleUKMMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ruu.mutation = mutation
			affected, err = ruu.sqlSave(ctx)
			return affected, err
		})
		for i := len(ruu.hooks) - 1; i >= 0; i-- {
			mut = ruu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ruu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ruu *RoleUKMUpdate) SaveX(ctx context.Context) int {
	affected, err := ruu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ruu *RoleUKMUpdate) Exec(ctx context.Context) error {
	_, err := ruu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ruu *RoleUKMUpdate) ExecX(ctx context.Context) {
	if err := ruu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ruu *RoleUKMUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   roleukm.Table,
			Columns: roleukm.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: roleukm.FieldID,
			},
		},
	}
	if ps := ruu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ruu.mutation.StatusRole(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: roleukm.FieldStatusRole,
		})
	}
	if nodes := ruu.mutation.RemovedProfileRolesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   roleukm.ProfileRolesTable,
			Columns: []string{roleukm.ProfileRolesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: profileukm.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruu.mutation.ProfileRolesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   roleukm.ProfileRolesTable,
			Columns: []string{roleukm.ProfileRolesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: profileukm.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ruu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{roleukm.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// RoleUKMUpdateOne is the builder for updating a single RoleUKM entity.
type RoleUKMUpdateOne struct {
	config
	hooks    []Hook
	mutation *RoleUKMMutation
}

// SetStatusRole sets the status_role field.
func (ruuo *RoleUKMUpdateOne) SetStatusRole(s string) *RoleUKMUpdateOne {
	ruuo.mutation.SetStatusRole(s)
	return ruuo
}

// AddProfileRoleIDs adds the profile_roles edge to ProfileUKM by ids.
func (ruuo *RoleUKMUpdateOne) AddProfileRoleIDs(ids ...int) *RoleUKMUpdateOne {
	ruuo.mutation.AddProfileRoleIDs(ids...)
	return ruuo
}

// AddProfileRoles adds the profile_roles edges to ProfileUKM.
func (ruuo *RoleUKMUpdateOne) AddProfileRoles(p ...*ProfileUKM) *RoleUKMUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ruuo.AddProfileRoleIDs(ids...)
}

// RemoveProfileRoleIDs removes the profile_roles edge to ProfileUKM by ids.
func (ruuo *RoleUKMUpdateOne) RemoveProfileRoleIDs(ids ...int) *RoleUKMUpdateOne {
	ruuo.mutation.RemoveProfileRoleIDs(ids...)
	return ruuo
}

// RemoveProfileRoles removes profile_roles edges to ProfileUKM.
func (ruuo *RoleUKMUpdateOne) RemoveProfileRoles(p ...*ProfileUKM) *RoleUKMUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ruuo.RemoveProfileRoleIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (ruuo *RoleUKMUpdateOne) Save(ctx context.Context) (*RoleUKM, error) {
	if v, ok := ruuo.mutation.StatusRole(); ok {
		if err := roleukm.StatusRoleValidator(v); err != nil {
			return nil, fmt.Errorf("ent: validator failed for field \"status_role\": %v", err)
		}
	}

	var (
		err  error
		node *RoleUKM
	)
	if len(ruuo.hooks) == 0 {
		node, err = ruuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RoleUKMMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ruuo.mutation = mutation
			node, err = ruuo.sqlSave(ctx)
			return node, err
		})
		for i := len(ruuo.hooks) - 1; i >= 0; i-- {
			mut = ruuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ruuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ruuo *RoleUKMUpdateOne) SaveX(ctx context.Context) *RoleUKM {
	ru, err := ruuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return ru
}

// Exec executes the query on the entity.
func (ruuo *RoleUKMUpdateOne) Exec(ctx context.Context) error {
	_, err := ruuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ruuo *RoleUKMUpdateOne) ExecX(ctx context.Context) {
	if err := ruuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ruuo *RoleUKMUpdateOne) sqlSave(ctx context.Context) (ru *RoleUKM, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   roleukm.Table,
			Columns: roleukm.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: roleukm.FieldID,
			},
		},
	}
	id, ok := ruuo.mutation.ID()
	if !ok {
		return nil, fmt.Errorf("missing RoleUKM.ID for update")
	}
	_spec.Node.ID.Value = id
	if value, ok := ruuo.mutation.StatusRole(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: roleukm.FieldStatusRole,
		})
	}
	if nodes := ruuo.mutation.RemovedProfileRolesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   roleukm.ProfileRolesTable,
			Columns: []string{roleukm.ProfileRolesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: profileukm.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruuo.mutation.ProfileRolesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   roleukm.ProfileRolesTable,
			Columns: []string{roleukm.ProfileRolesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: profileukm.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	ru = &RoleUKM{config: ruuo.config}
	_spec.Assign = ru.assignValues
	_spec.ScanValues = ru.scanValues()
	if err = sqlgraph.UpdateNode(ctx, ruuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{roleukm.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return ru, nil
}
