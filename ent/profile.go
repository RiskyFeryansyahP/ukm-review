// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"github.com/confus1on/UKM/ent/profile"
	"github.com/facebookincubator/ent/dialect/sql"
)

// Profile is the model entity for the Profile schema.
type Profile struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// FirstName holds the value of the "firstName" field.
	FirstName string `json:"firstName,omitempty"`
	// LastName holds the value of the "lastName" field.
	LastName string `json:"lastName,omitempty"`
	// Jk holds the value of the "jk" field.
	Jk profile.Jk `json:"jk,omitempty"`
	// Address holds the value of the "address" field.
	Address string `json:"address,omitempty"`
	// BirthDate holds the value of the "birth_date" field.
	BirthDate string `json:"birth_date,omitempty"`
	// YearGeneration holds the value of the "year_generation" field.
	YearGeneration string `json:"year_generation,omitempty"`
	// Phone holds the value of the "phone" field.
	Phone string `json:"phone,omitempty"`
	// Status holds the value of the "status" field.
	Status bool `json:"status,omitempty"`
	// ImgProfile holds the value of the "img_profile" field.
	ImgProfile string `json:"img_profile,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ProfileQuery when eager-loading is set.
	Edges ProfileEdges `json:"edges"`
}

// ProfileEdges holds the relations/edges for other nodes in the graph.
type ProfileEdges struct {
	// Owner holds the value of the owner edge.
	Owner []*User
	// Ukms holds the value of the ukms edge.
	Ukms []*ProfileUKM
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// OwnerOrErr returns the Owner value or an error if the edge
// was not loaded in eager-loading.
func (e ProfileEdges) OwnerOrErr() ([]*User, error) {
	if e.loadedTypes[0] {
		return e.Owner, nil
	}
	return nil, &NotLoadedError{edge: "owner"}
}

// UkmsOrErr returns the Ukms value or an error if the edge
// was not loaded in eager-loading.
func (e ProfileEdges) UkmsOrErr() ([]*ProfileUKM, error) {
	if e.loadedTypes[1] {
		return e.Ukms, nil
	}
	return nil, &NotLoadedError{edge: "ukms"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Profile) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // firstName
		&sql.NullString{}, // lastName
		&sql.NullString{}, // jk
		&sql.NullString{}, // address
		&sql.NullString{}, // birth_date
		&sql.NullString{}, // year_generation
		&sql.NullString{}, // phone
		&sql.NullBool{},   // status
		&sql.NullString{}, // img_profile
		&sql.NullTime{},   // created_at
		&sql.NullTime{},   // updated_at
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Profile fields.
func (pr *Profile) assignValues(values ...interface{}) error {
	if m, n := len(values), len(profile.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	pr.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field firstName", values[0])
	} else if value.Valid {
		pr.FirstName = value.String
	}
	if value, ok := values[1].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field lastName", values[1])
	} else if value.Valid {
		pr.LastName = value.String
	}
	if value, ok := values[2].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field jk", values[2])
	} else if value.Valid {
		pr.Jk = profile.Jk(value.String)
	}
	if value, ok := values[3].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field address", values[3])
	} else if value.Valid {
		pr.Address = value.String
	}
	if value, ok := values[4].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field birth_date", values[4])
	} else if value.Valid {
		pr.BirthDate = value.String
	}
	if value, ok := values[5].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field year_generation", values[5])
	} else if value.Valid {
		pr.YearGeneration = value.String
	}
	if value, ok := values[6].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field phone", values[6])
	} else if value.Valid {
		pr.Phone = value.String
	}
	if value, ok := values[7].(*sql.NullBool); !ok {
		return fmt.Errorf("unexpected type %T for field status", values[7])
	} else if value.Valid {
		pr.Status = value.Bool
	}
	if value, ok := values[8].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field img_profile", values[8])
	} else if value.Valid {
		pr.ImgProfile = value.String
	}
	if value, ok := values[9].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field created_at", values[9])
	} else if value.Valid {
		pr.CreatedAt = value.Time
	}
	if value, ok := values[10].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field updated_at", values[10])
	} else if value.Valid {
		pr.UpdatedAt = value.Time
	}
	return nil
}

// QueryOwner queries the owner edge of the Profile.
func (pr *Profile) QueryOwner() *UserQuery {
	return (&ProfileClient{config: pr.config}).QueryOwner(pr)
}

// QueryUkms queries the ukms edge of the Profile.
func (pr *Profile) QueryUkms() *ProfileUKMQuery {
	return (&ProfileClient{config: pr.config}).QueryUkms(pr)
}

// Update returns a builder for updating this Profile.
// Note that, you need to call Profile.Unwrap() before calling this method, if this Profile
// was returned from a transaction, and the transaction was committed or rolled back.
func (pr *Profile) Update() *ProfileUpdateOne {
	return (&ProfileClient{config: pr.config}).UpdateOne(pr)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (pr *Profile) Unwrap() *Profile {
	tx, ok := pr.config.driver.(*txDriver)
	if !ok {
		panic("ent: Profile is not a transactional entity")
	}
	pr.config.driver = tx.drv
	return pr
}

// String implements the fmt.Stringer.
func (pr *Profile) String() string {
	var builder strings.Builder
	builder.WriteString("Profile(")
	builder.WriteString(fmt.Sprintf("id=%v", pr.ID))
	builder.WriteString(", firstName=")
	builder.WriteString(pr.FirstName)
	builder.WriteString(", lastName=")
	builder.WriteString(pr.LastName)
	builder.WriteString(", jk=")
	builder.WriteString(fmt.Sprintf("%v", pr.Jk))
	builder.WriteString(", address=")
	builder.WriteString(pr.Address)
	builder.WriteString(", birth_date=")
	builder.WriteString(pr.BirthDate)
	builder.WriteString(", year_generation=")
	builder.WriteString(pr.YearGeneration)
	builder.WriteString(", phone=")
	builder.WriteString(pr.Phone)
	builder.WriteString(", status=")
	builder.WriteString(fmt.Sprintf("%v", pr.Status))
	builder.WriteString(", img_profile=")
	builder.WriteString(pr.ImgProfile)
	builder.WriteString(", created_at=")
	builder.WriteString(pr.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(pr.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Profiles is a parsable slice of Profile.
type Profiles []*Profile

func (pr Profiles) config(cfg config) {
	for _i := range pr {
		pr[_i].config = cfg
	}
}