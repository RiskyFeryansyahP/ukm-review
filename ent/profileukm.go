// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/confus1on/UKM/ent/profile"
	"github.com/confus1on/UKM/ent/profileukm"
	"github.com/confus1on/UKM/ent/roleukm"
	"github.com/confus1on/UKM/ent/ukm"
	"github.com/facebookincubator/ent/dialect/sql"
)

// ProfileUKM is the model entity for the ProfileUKM schema.
type ProfileUKM struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Reason holds the value of the "reason" field.
	Reason string `json:"reason,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ProfileUKMQuery when eager-loading is set.
	Edges                  ProfileUKMEdges `json:"edges"`
	profile_ukms           *int
	role_ukm_profile_roles *int
	ukm_profiles           *int
}

// ProfileUKMEdges holds the relations/edges for other nodes in the graph.
type ProfileUKMEdges struct {
	// OwnerProfile holds the value of the owner_profile edge.
	OwnerProfile *Profile
	// OwnerUkm holds the value of the owner_ukm edge.
	OwnerUkm *Ukm
	// OwnerRole holds the value of the owner_role edge.
	OwnerRole *RoleUKM
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// OwnerProfileOrErr returns the OwnerProfile value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ProfileUKMEdges) OwnerProfileOrErr() (*Profile, error) {
	if e.loadedTypes[0] {
		if e.OwnerProfile == nil {
			// The edge owner_profile was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: profile.Label}
		}
		return e.OwnerProfile, nil
	}
	return nil, &NotLoadedError{edge: "owner_profile"}
}

// OwnerUkmOrErr returns the OwnerUkm value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ProfileUKMEdges) OwnerUkmOrErr() (*Ukm, error) {
	if e.loadedTypes[1] {
		if e.OwnerUkm == nil {
			// The edge owner_ukm was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: ukm.Label}
		}
		return e.OwnerUkm, nil
	}
	return nil, &NotLoadedError{edge: "owner_ukm"}
}

// OwnerRoleOrErr returns the OwnerRole value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ProfileUKMEdges) OwnerRoleOrErr() (*RoleUKM, error) {
	if e.loadedTypes[2] {
		if e.OwnerRole == nil {
			// The edge owner_role was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: roleukm.Label}
		}
		return e.OwnerRole, nil
	}
	return nil, &NotLoadedError{edge: "owner_role"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ProfileUKM) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // reason
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*ProfileUKM) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // profile_ukms
		&sql.NullInt64{}, // role_ukm_profile_roles
		&sql.NullInt64{}, // ukm_profiles
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ProfileUKM fields.
func (pu *ProfileUKM) assignValues(values ...interface{}) error {
	if m, n := len(values), len(profileukm.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	pu.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field reason", values[0])
	} else if value.Valid {
		pu.Reason = value.String
	}
	values = values[1:]
	if len(values) == len(profileukm.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field profile_ukms", value)
		} else if value.Valid {
			pu.profile_ukms = new(int)
			*pu.profile_ukms = int(value.Int64)
		}
		if value, ok := values[1].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field role_ukm_profile_roles", value)
		} else if value.Valid {
			pu.role_ukm_profile_roles = new(int)
			*pu.role_ukm_profile_roles = int(value.Int64)
		}
		if value, ok := values[2].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field ukm_profiles", value)
		} else if value.Valid {
			pu.ukm_profiles = new(int)
			*pu.ukm_profiles = int(value.Int64)
		}
	}
	return nil
}

// QueryOwnerProfile queries the owner_profile edge of the ProfileUKM.
func (pu *ProfileUKM) QueryOwnerProfile() *ProfileQuery {
	return (&ProfileUKMClient{config: pu.config}).QueryOwnerProfile(pu)
}

// QueryOwnerUkm queries the owner_ukm edge of the ProfileUKM.
func (pu *ProfileUKM) QueryOwnerUkm() *UkmQuery {
	return (&ProfileUKMClient{config: pu.config}).QueryOwnerUkm(pu)
}

// QueryOwnerRole queries the owner_role edge of the ProfileUKM.
func (pu *ProfileUKM) QueryOwnerRole() *RoleUKMQuery {
	return (&ProfileUKMClient{config: pu.config}).QueryOwnerRole(pu)
}

// Update returns a builder for updating this ProfileUKM.
// Note that, you need to call ProfileUKM.Unwrap() before calling this method, if this ProfileUKM
// was returned from a transaction, and the transaction was committed or rolled back.
func (pu *ProfileUKM) Update() *ProfileUKMUpdateOne {
	return (&ProfileUKMClient{config: pu.config}).UpdateOne(pu)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (pu *ProfileUKM) Unwrap() *ProfileUKM {
	tx, ok := pu.config.driver.(*txDriver)
	if !ok {
		panic("ent: ProfileUKM is not a transactional entity")
	}
	pu.config.driver = tx.drv
	return pu
}

// String implements the fmt.Stringer.
func (pu *ProfileUKM) String() string {
	var builder strings.Builder
	builder.WriteString("ProfileUKM(")
	builder.WriteString(fmt.Sprintf("id=%v", pu.ID))
	builder.WriteString(", reason=")
	builder.WriteString(pu.Reason)
	builder.WriteByte(')')
	return builder.String()
}

// ProfileUKMs is a parsable slice of ProfileUKM.
type ProfileUKMs []*ProfileUKM

func (pu ProfileUKMs) config(cfg config) {
	for _i := range pu {
		pu[_i].config = cfg
	}
}
