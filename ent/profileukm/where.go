// Code generated by entc, DO NOT EDIT.

package profileukm

import (
	"github.com/confus1on/UKM/ent/predicate"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.ProfileUKM {
	return predicate.ProfileUKM(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.ProfileUKM {
	return predicate.ProfileUKM(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.ProfileUKM {
	return predicate.ProfileUKM(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.ProfileUKM {
	return predicate.ProfileUKM(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.ProfileUKM {
	return predicate.ProfileUKM(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.ProfileUKM {
	return predicate.ProfileUKM(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.ProfileUKM {
	return predicate.ProfileUKM(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.ProfileUKM {
	return predicate.ProfileUKM(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.ProfileUKM {
	return predicate.ProfileUKM(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Reason applies equality check predicate on the "reason" field. It's identical to ReasonEQ.
func Reason(v string) predicate.ProfileUKM {
	return predicate.ProfileUKM(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldReason), v))
	})
}

// ReasonEQ applies the EQ predicate on the "reason" field.
func ReasonEQ(v string) predicate.ProfileUKM {
	return predicate.ProfileUKM(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldReason), v))
	})
}

// ReasonNEQ applies the NEQ predicate on the "reason" field.
func ReasonNEQ(v string) predicate.ProfileUKM {
	return predicate.ProfileUKM(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldReason), v))
	})
}

// ReasonIn applies the In predicate on the "reason" field.
func ReasonIn(vs ...string) predicate.ProfileUKM {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ProfileUKM(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldReason), v...))
	})
}

// ReasonNotIn applies the NotIn predicate on the "reason" field.
func ReasonNotIn(vs ...string) predicate.ProfileUKM {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ProfileUKM(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldReason), v...))
	})
}

// ReasonGT applies the GT predicate on the "reason" field.
func ReasonGT(v string) predicate.ProfileUKM {
	return predicate.ProfileUKM(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldReason), v))
	})
}

// ReasonGTE applies the GTE predicate on the "reason" field.
func ReasonGTE(v string) predicate.ProfileUKM {
	return predicate.ProfileUKM(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldReason), v))
	})
}

// ReasonLT applies the LT predicate on the "reason" field.
func ReasonLT(v string) predicate.ProfileUKM {
	return predicate.ProfileUKM(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldReason), v))
	})
}

// ReasonLTE applies the LTE predicate on the "reason" field.
func ReasonLTE(v string) predicate.ProfileUKM {
	return predicate.ProfileUKM(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldReason), v))
	})
}

// ReasonContains applies the Contains predicate on the "reason" field.
func ReasonContains(v string) predicate.ProfileUKM {
	return predicate.ProfileUKM(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldReason), v))
	})
}

// ReasonHasPrefix applies the HasPrefix predicate on the "reason" field.
func ReasonHasPrefix(v string) predicate.ProfileUKM {
	return predicate.ProfileUKM(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldReason), v))
	})
}

// ReasonHasSuffix applies the HasSuffix predicate on the "reason" field.
func ReasonHasSuffix(v string) predicate.ProfileUKM {
	return predicate.ProfileUKM(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldReason), v))
	})
}

// ReasonEqualFold applies the EqualFold predicate on the "reason" field.
func ReasonEqualFold(v string) predicate.ProfileUKM {
	return predicate.ProfileUKM(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldReason), v))
	})
}

// ReasonContainsFold applies the ContainsFold predicate on the "reason" field.
func ReasonContainsFold(v string) predicate.ProfileUKM {
	return predicate.ProfileUKM(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldReason), v))
	})
}

// HasOwnerProfile applies the HasEdge predicate on the "owner_profile" edge.
func HasOwnerProfile() predicate.ProfileUKM {
	return predicate.ProfileUKM(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OwnerProfileTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OwnerProfileTable, OwnerProfileColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOwnerProfileWith applies the HasEdge predicate on the "owner_profile" edge with a given conditions (other predicates).
func HasOwnerProfileWith(preds ...predicate.Profile) predicate.ProfileUKM {
	return predicate.ProfileUKM(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OwnerProfileInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OwnerProfileTable, OwnerProfileColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasOwnerUkm applies the HasEdge predicate on the "owner_ukm" edge.
func HasOwnerUkm() predicate.ProfileUKM {
	return predicate.ProfileUKM(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OwnerUkmTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OwnerUkmTable, OwnerUkmColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOwnerUkmWith applies the HasEdge predicate on the "owner_ukm" edge with a given conditions (other predicates).
func HasOwnerUkmWith(preds ...predicate.Ukm) predicate.ProfileUKM {
	return predicate.ProfileUKM(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OwnerUkmInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OwnerUkmTable, OwnerUkmColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasOwnerRole applies the HasEdge predicate on the "owner_role" edge.
func HasOwnerRole() predicate.ProfileUKM {
	return predicate.ProfileUKM(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OwnerRoleTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OwnerRoleTable, OwnerRoleColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOwnerRoleWith applies the HasEdge predicate on the "owner_role" edge with a given conditions (other predicates).
func HasOwnerRoleWith(preds ...predicate.RoleUKM) predicate.ProfileUKM {
	return predicate.ProfileUKM(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OwnerRoleInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OwnerRoleTable, OwnerRoleColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.ProfileUKM) predicate.ProfileUKM {
	return predicate.ProfileUKM(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.ProfileUKM) predicate.ProfileUKM {
	return predicate.ProfileUKM(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.ProfileUKM) predicate.ProfileUKM {
	return predicate.ProfileUKM(func(s *sql.Selector) {
		p(s.Not())
	})
}
