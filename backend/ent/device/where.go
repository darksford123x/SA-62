// Code generated by entc, DO NOT EDIT.

package device

import (
	"github.com/darksford123x/repairinvoice-record/ent/predicate"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
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
func IDGT(id int) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Dname applies equality check predicate on the "Dname" field. It's identical to DnameEQ.
func Dname(v string) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDname), v))
	})
}

// DnameEQ applies the EQ predicate on the "Dname" field.
func DnameEQ(v string) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDname), v))
	})
}

// DnameNEQ applies the NEQ predicate on the "Dname" field.
func DnameNEQ(v string) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDname), v))
	})
}

// DnameIn applies the In predicate on the "Dname" field.
func DnameIn(vs ...string) predicate.Device {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Device(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDname), v...))
	})
}

// DnameNotIn applies the NotIn predicate on the "Dname" field.
func DnameNotIn(vs ...string) predicate.Device {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Device(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDname), v...))
	})
}

// DnameGT applies the GT predicate on the "Dname" field.
func DnameGT(v string) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDname), v))
	})
}

// DnameGTE applies the GTE predicate on the "Dname" field.
func DnameGTE(v string) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDname), v))
	})
}

// DnameLT applies the LT predicate on the "Dname" field.
func DnameLT(v string) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDname), v))
	})
}

// DnameLTE applies the LTE predicate on the "Dname" field.
func DnameLTE(v string) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDname), v))
	})
}

// DnameContains applies the Contains predicate on the "Dname" field.
func DnameContains(v string) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldDname), v))
	})
}

// DnameHasPrefix applies the HasPrefix predicate on the "Dname" field.
func DnameHasPrefix(v string) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldDname), v))
	})
}

// DnameHasSuffix applies the HasSuffix predicate on the "Dname" field.
func DnameHasSuffix(v string) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldDname), v))
	})
}

// DnameEqualFold applies the EqualFold predicate on the "Dname" field.
func DnameEqualFold(v string) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldDname), v))
	})
}

// DnameContainsFold applies the ContainsFold predicate on the "Dname" field.
func DnameContainsFold(v string) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldDname), v))
	})
}

// HasRepairInformation applies the HasEdge predicate on the "repair_information" edge.
func HasRepairInformation() predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(RepairInformationTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, RepairInformationTable, RepairInformationColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRepairInformationWith applies the HasEdge predicate on the "repair_information" edge with a given conditions (other predicates).
func HasRepairInformationWith(preds ...predicate.RepairInvoice) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(RepairInformationInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, RepairInformationTable, RepairInformationColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Device) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Device) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
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
func Not(p predicate.Device) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		p(s.Not())
	})
}
