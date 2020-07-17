// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/confus1on/UKM/ent/announcement"
	"github.com/confus1on/UKM/ent/predicate"
	"github.com/confus1on/UKM/ent/ukm"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// AnnouncementUpdate is the builder for updating Announcement entities.
type AnnouncementUpdate struct {
	config
	hooks      []Hook
	mutation   *AnnouncementMutation
	predicates []predicate.Announcement
}

// Where adds a new predicate for the builder.
func (au *AnnouncementUpdate) Where(ps ...predicate.Announcement) *AnnouncementUpdate {
	au.predicates = append(au.predicates, ps...)
	return au
}

// SetTitle sets the title field.
func (au *AnnouncementUpdate) SetTitle(s string) *AnnouncementUpdate {
	au.mutation.SetTitle(s)
	return au
}

// SetDescription sets the description field.
func (au *AnnouncementUpdate) SetDescription(s string) *AnnouncementUpdate {
	au.mutation.SetDescription(s)
	return au
}

// SetTime sets the time field.
func (au *AnnouncementUpdate) SetTime(s string) *AnnouncementUpdate {
	au.mutation.SetTime(s)
	return au
}

// SetOwnerAnnouncementID sets the owner_announcement edge to Ukm by id.
func (au *AnnouncementUpdate) SetOwnerAnnouncementID(id int) *AnnouncementUpdate {
	au.mutation.SetOwnerAnnouncementID(id)
	return au
}

// SetNillableOwnerAnnouncementID sets the owner_announcement edge to Ukm by id if the given value is not nil.
func (au *AnnouncementUpdate) SetNillableOwnerAnnouncementID(id *int) *AnnouncementUpdate {
	if id != nil {
		au = au.SetOwnerAnnouncementID(*id)
	}
	return au
}

// SetOwnerAnnouncement sets the owner_announcement edge to Ukm.
func (au *AnnouncementUpdate) SetOwnerAnnouncement(u *Ukm) *AnnouncementUpdate {
	return au.SetOwnerAnnouncementID(u.ID)
}

// ClearOwnerAnnouncement clears the owner_announcement edge to Ukm.
func (au *AnnouncementUpdate) ClearOwnerAnnouncement() *AnnouncementUpdate {
	au.mutation.ClearOwnerAnnouncement()
	return au
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (au *AnnouncementUpdate) Save(ctx context.Context) (int, error) {
	if v, ok := au.mutation.Title(); ok {
		if err := announcement.TitleValidator(v); err != nil {
			return 0, fmt.Errorf("ent: validator failed for field \"title\": %v", err)
		}
	}
	if v, ok := au.mutation.Description(); ok {
		if err := announcement.DescriptionValidator(v); err != nil {
			return 0, fmt.Errorf("ent: validator failed for field \"description\": %v", err)
		}
	}
	if v, ok := au.mutation.Time(); ok {
		if err := announcement.TimeValidator(v); err != nil {
			return 0, fmt.Errorf("ent: validator failed for field \"time\": %v", err)
		}
	}

	var (
		err      error
		affected int
	)
	if len(au.hooks) == 0 {
		affected, err = au.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AnnouncementMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			au.mutation = mutation
			affected, err = au.sqlSave(ctx)
			return affected, err
		})
		for i := len(au.hooks) - 1; i >= 0; i-- {
			mut = au.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, au.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (au *AnnouncementUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *AnnouncementUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *AnnouncementUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

func (au *AnnouncementUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   announcement.Table,
			Columns: announcement.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: announcement.FieldID,
			},
		},
	}
	if ps := au.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := au.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: announcement.FieldTitle,
		})
	}
	if value, ok := au.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: announcement.FieldDescription,
		})
	}
	if value, ok := au.mutation.Time(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: announcement.FieldTime,
		})
	}
	if au.mutation.OwnerAnnouncementCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   announcement.OwnerAnnouncementTable,
			Columns: []string{announcement.OwnerAnnouncementColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: ukm.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.OwnerAnnouncementIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   announcement.OwnerAnnouncementTable,
			Columns: []string{announcement.OwnerAnnouncementColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: ukm.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, au.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{announcement.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// AnnouncementUpdateOne is the builder for updating a single Announcement entity.
type AnnouncementUpdateOne struct {
	config
	hooks    []Hook
	mutation *AnnouncementMutation
}

// SetTitle sets the title field.
func (auo *AnnouncementUpdateOne) SetTitle(s string) *AnnouncementUpdateOne {
	auo.mutation.SetTitle(s)
	return auo
}

// SetDescription sets the description field.
func (auo *AnnouncementUpdateOne) SetDescription(s string) *AnnouncementUpdateOne {
	auo.mutation.SetDescription(s)
	return auo
}

// SetTime sets the time field.
func (auo *AnnouncementUpdateOne) SetTime(s string) *AnnouncementUpdateOne {
	auo.mutation.SetTime(s)
	return auo
}

// SetOwnerAnnouncementID sets the owner_announcement edge to Ukm by id.
func (auo *AnnouncementUpdateOne) SetOwnerAnnouncementID(id int) *AnnouncementUpdateOne {
	auo.mutation.SetOwnerAnnouncementID(id)
	return auo
}

// SetNillableOwnerAnnouncementID sets the owner_announcement edge to Ukm by id if the given value is not nil.
func (auo *AnnouncementUpdateOne) SetNillableOwnerAnnouncementID(id *int) *AnnouncementUpdateOne {
	if id != nil {
		auo = auo.SetOwnerAnnouncementID(*id)
	}
	return auo
}

// SetOwnerAnnouncement sets the owner_announcement edge to Ukm.
func (auo *AnnouncementUpdateOne) SetOwnerAnnouncement(u *Ukm) *AnnouncementUpdateOne {
	return auo.SetOwnerAnnouncementID(u.ID)
}

// ClearOwnerAnnouncement clears the owner_announcement edge to Ukm.
func (auo *AnnouncementUpdateOne) ClearOwnerAnnouncement() *AnnouncementUpdateOne {
	auo.mutation.ClearOwnerAnnouncement()
	return auo
}

// Save executes the query and returns the updated entity.
func (auo *AnnouncementUpdateOne) Save(ctx context.Context) (*Announcement, error) {
	if v, ok := auo.mutation.Title(); ok {
		if err := announcement.TitleValidator(v); err != nil {
			return nil, fmt.Errorf("ent: validator failed for field \"title\": %v", err)
		}
	}
	if v, ok := auo.mutation.Description(); ok {
		if err := announcement.DescriptionValidator(v); err != nil {
			return nil, fmt.Errorf("ent: validator failed for field \"description\": %v", err)
		}
	}
	if v, ok := auo.mutation.Time(); ok {
		if err := announcement.TimeValidator(v); err != nil {
			return nil, fmt.Errorf("ent: validator failed for field \"time\": %v", err)
		}
	}

	var (
		err  error
		node *Announcement
	)
	if len(auo.hooks) == 0 {
		node, err = auo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AnnouncementMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			auo.mutation = mutation
			node, err = auo.sqlSave(ctx)
			return node, err
		})
		for i := len(auo.hooks) - 1; i >= 0; i-- {
			mut = auo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, auo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (auo *AnnouncementUpdateOne) SaveX(ctx context.Context) *Announcement {
	a, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return a
}

// Exec executes the query on the entity.
func (auo *AnnouncementUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *AnnouncementUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (auo *AnnouncementUpdateOne) sqlSave(ctx context.Context) (a *Announcement, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   announcement.Table,
			Columns: announcement.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: announcement.FieldID,
			},
		},
	}
	id, ok := auo.mutation.ID()
	if !ok {
		return nil, fmt.Errorf("missing Announcement.ID for update")
	}
	_spec.Node.ID.Value = id
	if value, ok := auo.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: announcement.FieldTitle,
		})
	}
	if value, ok := auo.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: announcement.FieldDescription,
		})
	}
	if value, ok := auo.mutation.Time(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: announcement.FieldTime,
		})
	}
	if auo.mutation.OwnerAnnouncementCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   announcement.OwnerAnnouncementTable,
			Columns: []string{announcement.OwnerAnnouncementColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: ukm.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.OwnerAnnouncementIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   announcement.OwnerAnnouncementTable,
			Columns: []string{announcement.OwnerAnnouncementColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: ukm.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	a = &Announcement{config: auo.config}
	_spec.Assign = a.assignValues
	_spec.ScanValues = a.scanValues()
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{announcement.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return a, nil
}
