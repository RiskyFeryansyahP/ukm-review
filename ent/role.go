// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/confus1on/UKM/ent/role"
	"github.com/facebookincubator/ent/dialect/sql"
)

// Role is the model entity for the Role schema.
type Role struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Value holds the value of the "value" field.
	Value string `json:"value,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the RoleQuery when eager-loading is set.
	Edges RoleEdges `json:"edges"`
}

// RoleEdges holds the relations/edges for other nodes in the graph.
type RoleEdges struct {
	// Access holds the value of the access edge.
	Access []*User
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// AccessOrErr returns the Access value or an error if the edge
// was not loaded in eager-loading.
func (e RoleEdges) AccessOrErr() ([]*User, error) {
	if e.loadedTypes[0] {
		return e.Access, nil
	}
	return nil, &NotLoadedError{edge: "access"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Role) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // value
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Role fields.
func (r *Role) assignValues(values ...interface{}) error {
	if m, n := len(values), len(role.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	r.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field value", values[0])
	} else if value.Valid {
		r.Value = value.String
	}
	return nil
}

// QueryAccess queries the access edge of the Role.
func (r *Role) QueryAccess() *UserQuery {
	return (&RoleClient{config: r.config}).QueryAccess(r)
}

// Update returns a builder for updating this Role.
// Note that, you need to call Role.Unwrap() before calling this method, if this Role
// was returned from a transaction, and the transaction was committed or rolled back.
func (r *Role) Update() *RoleUpdateOne {
	return (&RoleClient{config: r.config}).UpdateOne(r)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (r *Role) Unwrap() *Role {
	tx, ok := r.config.driver.(*txDriver)
	if !ok {
		panic("ent: Role is not a transactional entity")
	}
	r.config.driver = tx.drv
	return r
}

// String implements the fmt.Stringer.
func (r *Role) String() string {
	var builder strings.Builder
	builder.WriteString("Role(")
	builder.WriteString(fmt.Sprintf("id=%v", r.ID))
	builder.WriteString(", value=")
	builder.WriteString(r.Value)
	builder.WriteByte(')')
	return builder.String()
}

// Roles is a parsable slice of Role.
type Roles []*Role

func (r Roles) config(cfg config) {
	for _i := range r {
		r[_i].config = cfg
	}
}
