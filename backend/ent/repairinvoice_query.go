// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"github.com/darksford123x/repairinvoice-record/ent/device"
	"github.com/darksford123x/repairinvoice-record/ent/predicate"
	"github.com/darksford123x/repairinvoice-record/ent/repairinvoice"
	"github.com/darksford123x/repairinvoice-record/ent/statusr"
	"github.com/darksford123x/repairinvoice-record/ent/symptom"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// RepairInvoiceQuery is the builder for querying RepairInvoice entities.
type RepairInvoiceQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	unique     []string
	predicates []predicate.RepairInvoice
	// eager-loading edges.
	withDevice  *DeviceQuery
	withStatus  *StatusRQuery
	withSymptom *SymptomQuery
	withFKs     bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the builder.
func (riq *RepairInvoiceQuery) Where(ps ...predicate.RepairInvoice) *RepairInvoiceQuery {
	riq.predicates = append(riq.predicates, ps...)
	return riq
}

// Limit adds a limit step to the query.
func (riq *RepairInvoiceQuery) Limit(limit int) *RepairInvoiceQuery {
	riq.limit = &limit
	return riq
}

// Offset adds an offset step to the query.
func (riq *RepairInvoiceQuery) Offset(offset int) *RepairInvoiceQuery {
	riq.offset = &offset
	return riq
}

// Order adds an order step to the query.
func (riq *RepairInvoiceQuery) Order(o ...OrderFunc) *RepairInvoiceQuery {
	riq.order = append(riq.order, o...)
	return riq
}

// QueryDevice chains the current query on the device edge.
func (riq *RepairInvoiceQuery) QueryDevice() *DeviceQuery {
	query := &DeviceQuery{config: riq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := riq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(repairinvoice.Table, repairinvoice.FieldID, riq.sqlQuery()),
			sqlgraph.To(device.Table, device.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, repairinvoice.DeviceTable, repairinvoice.DeviceColumn),
		)
		fromU = sqlgraph.SetNeighbors(riq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryStatus chains the current query on the status edge.
func (riq *RepairInvoiceQuery) QueryStatus() *StatusRQuery {
	query := &StatusRQuery{config: riq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := riq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(repairinvoice.Table, repairinvoice.FieldID, riq.sqlQuery()),
			sqlgraph.To(statusr.Table, statusr.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, repairinvoice.StatusTable, repairinvoice.StatusColumn),
		)
		fromU = sqlgraph.SetNeighbors(riq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QuerySymptom chains the current query on the symptom edge.
func (riq *RepairInvoiceQuery) QuerySymptom() *SymptomQuery {
	query := &SymptomQuery{config: riq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := riq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(repairinvoice.Table, repairinvoice.FieldID, riq.sqlQuery()),
			sqlgraph.To(symptom.Table, symptom.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, repairinvoice.SymptomTable, repairinvoice.SymptomColumn),
		)
		fromU = sqlgraph.SetNeighbors(riq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first RepairInvoice entity in the query. Returns *NotFoundError when no repairinvoice was found.
func (riq *RepairInvoiceQuery) First(ctx context.Context) (*RepairInvoice, error) {
	ris, err := riq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(ris) == 0 {
		return nil, &NotFoundError{repairinvoice.Label}
	}
	return ris[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (riq *RepairInvoiceQuery) FirstX(ctx context.Context) *RepairInvoice {
	ri, err := riq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return ri
}

// FirstID returns the first RepairInvoice id in the query. Returns *NotFoundError when no id was found.
func (riq *RepairInvoiceQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = riq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{repairinvoice.Label}
		return
	}
	return ids[0], nil
}

// FirstXID is like FirstID, but panics if an error occurs.
func (riq *RepairInvoiceQuery) FirstXID(ctx context.Context) int {
	id, err := riq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only RepairInvoice entity in the query, returns an error if not exactly one entity was returned.
func (riq *RepairInvoiceQuery) Only(ctx context.Context) (*RepairInvoice, error) {
	ris, err := riq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(ris) {
	case 1:
		return ris[0], nil
	case 0:
		return nil, &NotFoundError{repairinvoice.Label}
	default:
		return nil, &NotSingularError{repairinvoice.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (riq *RepairInvoiceQuery) OnlyX(ctx context.Context) *RepairInvoice {
	ri, err := riq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return ri
}

// OnlyID returns the only RepairInvoice id in the query, returns an error if not exactly one id was returned.
func (riq *RepairInvoiceQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = riq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{repairinvoice.Label}
	default:
		err = &NotSingularError{repairinvoice.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (riq *RepairInvoiceQuery) OnlyIDX(ctx context.Context) int {
	id, err := riq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of RepairInvoices.
func (riq *RepairInvoiceQuery) All(ctx context.Context) ([]*RepairInvoice, error) {
	if err := riq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return riq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (riq *RepairInvoiceQuery) AllX(ctx context.Context) []*RepairInvoice {
	ris, err := riq.All(ctx)
	if err != nil {
		panic(err)
	}
	return ris
}

// IDs executes the query and returns a list of RepairInvoice ids.
func (riq *RepairInvoiceQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := riq.Select(repairinvoice.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (riq *RepairInvoiceQuery) IDsX(ctx context.Context) []int {
	ids, err := riq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (riq *RepairInvoiceQuery) Count(ctx context.Context) (int, error) {
	if err := riq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return riq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (riq *RepairInvoiceQuery) CountX(ctx context.Context) int {
	count, err := riq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (riq *RepairInvoiceQuery) Exist(ctx context.Context) (bool, error) {
	if err := riq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return riq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (riq *RepairInvoiceQuery) ExistX(ctx context.Context) bool {
	exist, err := riq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (riq *RepairInvoiceQuery) Clone() *RepairInvoiceQuery {
	return &RepairInvoiceQuery{
		config:     riq.config,
		limit:      riq.limit,
		offset:     riq.offset,
		order:      append([]OrderFunc{}, riq.order...),
		unique:     append([]string{}, riq.unique...),
		predicates: append([]predicate.RepairInvoice{}, riq.predicates...),
		// clone intermediate query.
		sql:  riq.sql.Clone(),
		path: riq.path,
	}
}

//  WithDevice tells the query-builder to eager-loads the nodes that are connected to
// the "device" edge. The optional arguments used to configure the query builder of the edge.
func (riq *RepairInvoiceQuery) WithDevice(opts ...func(*DeviceQuery)) *RepairInvoiceQuery {
	query := &DeviceQuery{config: riq.config}
	for _, opt := range opts {
		opt(query)
	}
	riq.withDevice = query
	return riq
}

//  WithStatus tells the query-builder to eager-loads the nodes that are connected to
// the "status" edge. The optional arguments used to configure the query builder of the edge.
func (riq *RepairInvoiceQuery) WithStatus(opts ...func(*StatusRQuery)) *RepairInvoiceQuery {
	query := &StatusRQuery{config: riq.config}
	for _, opt := range opts {
		opt(query)
	}
	riq.withStatus = query
	return riq
}

//  WithSymptom tells the query-builder to eager-loads the nodes that are connected to
// the "symptom" edge. The optional arguments used to configure the query builder of the edge.
func (riq *RepairInvoiceQuery) WithSymptom(opts ...func(*SymptomQuery)) *RepairInvoiceQuery {
	query := &SymptomQuery{config: riq.config}
	for _, opt := range opts {
		opt(query)
	}
	riq.withSymptom = query
	return riq
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Rename string `json:"Rename,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.RepairInvoice.Query().
//		GroupBy(repairinvoice.FieldRename).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (riq *RepairInvoiceQuery) GroupBy(field string, fields ...string) *RepairInvoiceGroupBy {
	group := &RepairInvoiceGroupBy{config: riq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := riq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return riq.sqlQuery(), nil
	}
	return group
}

// Select one or more fields from the given query.
//
// Example:
//
//	var v []struct {
//		Rename string `json:"Rename,omitempty"`
//	}
//
//	client.RepairInvoice.Query().
//		Select(repairinvoice.FieldRename).
//		Scan(ctx, &v)
//
func (riq *RepairInvoiceQuery) Select(field string, fields ...string) *RepairInvoiceSelect {
	selector := &RepairInvoiceSelect{config: riq.config}
	selector.fields = append([]string{field}, fields...)
	selector.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := riq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return riq.sqlQuery(), nil
	}
	return selector
}

func (riq *RepairInvoiceQuery) prepareQuery(ctx context.Context) error {
	if riq.path != nil {
		prev, err := riq.path(ctx)
		if err != nil {
			return err
		}
		riq.sql = prev
	}
	return nil
}

func (riq *RepairInvoiceQuery) sqlAll(ctx context.Context) ([]*RepairInvoice, error) {
	var (
		nodes       = []*RepairInvoice{}
		withFKs     = riq.withFKs
		_spec       = riq.querySpec()
		loadedTypes = [3]bool{
			riq.withDevice != nil,
			riq.withStatus != nil,
			riq.withSymptom != nil,
		}
	)
	if riq.withDevice != nil || riq.withStatus != nil || riq.withSymptom != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, repairinvoice.ForeignKeys...)
	}
	_spec.ScanValues = func() []interface{} {
		node := &RepairInvoice{config: riq.config}
		nodes = append(nodes, node)
		values := node.scanValues()
		if withFKs {
			values = append(values, node.fkValues()...)
		}
		return values
	}
	_spec.Assign = func(values ...interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(values...)
	}
	if err := sqlgraph.QueryNodes(ctx, riq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := riq.withDevice; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*RepairInvoice)
		for i := range nodes {
			if fk := nodes[i].device_id; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(device.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "device_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Device = n
			}
		}
	}

	if query := riq.withStatus; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*RepairInvoice)
		for i := range nodes {
			if fk := nodes[i].statusr_id; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(statusr.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "statusr_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Status = n
			}
		}
	}

	if query := riq.withSymptom; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*RepairInvoice)
		for i := range nodes {
			if fk := nodes[i].symptom_id; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(symptom.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "symptom_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Symptom = n
			}
		}
	}

	return nodes, nil
}

func (riq *RepairInvoiceQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := riq.querySpec()
	return sqlgraph.CountNodes(ctx, riq.driver, _spec)
}

func (riq *RepairInvoiceQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := riq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (riq *RepairInvoiceQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   repairinvoice.Table,
			Columns: repairinvoice.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: repairinvoice.FieldID,
			},
		},
		From:   riq.sql,
		Unique: true,
	}
	if ps := riq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := riq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := riq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := riq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (riq *RepairInvoiceQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(riq.driver.Dialect())
	t1 := builder.Table(repairinvoice.Table)
	selector := builder.Select(t1.Columns(repairinvoice.Columns...)...).From(t1)
	if riq.sql != nil {
		selector = riq.sql
		selector.Select(selector.Columns(repairinvoice.Columns...)...)
	}
	for _, p := range riq.predicates {
		p(selector)
	}
	for _, p := range riq.order {
		p(selector)
	}
	if offset := riq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := riq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// RepairInvoiceGroupBy is the builder for group-by RepairInvoice entities.
type RepairInvoiceGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (rigb *RepairInvoiceGroupBy) Aggregate(fns ...AggregateFunc) *RepairInvoiceGroupBy {
	rigb.fns = append(rigb.fns, fns...)
	return rigb
}

// Scan applies the group-by query and scan the result into the given value.
func (rigb *RepairInvoiceGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := rigb.path(ctx)
	if err != nil {
		return err
	}
	rigb.sql = query
	return rigb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (rigb *RepairInvoiceGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := rigb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (rigb *RepairInvoiceGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(rigb.fields) > 1 {
		return nil, errors.New("ent: RepairInvoiceGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := rigb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (rigb *RepairInvoiceGroupBy) StringsX(ctx context.Context) []string {
	v, err := rigb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from group-by. It is only allowed when querying group-by with one field.
func (rigb *RepairInvoiceGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = rigb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{repairinvoice.Label}
	default:
		err = fmt.Errorf("ent: RepairInvoiceGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (rigb *RepairInvoiceGroupBy) StringX(ctx context.Context) string {
	v, err := rigb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (rigb *RepairInvoiceGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(rigb.fields) > 1 {
		return nil, errors.New("ent: RepairInvoiceGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := rigb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (rigb *RepairInvoiceGroupBy) IntsX(ctx context.Context) []int {
	v, err := rigb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from group-by. It is only allowed when querying group-by with one field.
func (rigb *RepairInvoiceGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = rigb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{repairinvoice.Label}
	default:
		err = fmt.Errorf("ent: RepairInvoiceGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (rigb *RepairInvoiceGroupBy) IntX(ctx context.Context) int {
	v, err := rigb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (rigb *RepairInvoiceGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(rigb.fields) > 1 {
		return nil, errors.New("ent: RepairInvoiceGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := rigb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (rigb *RepairInvoiceGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := rigb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from group-by. It is only allowed when querying group-by with one field.
func (rigb *RepairInvoiceGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = rigb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{repairinvoice.Label}
	default:
		err = fmt.Errorf("ent: RepairInvoiceGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (rigb *RepairInvoiceGroupBy) Float64X(ctx context.Context) float64 {
	v, err := rigb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (rigb *RepairInvoiceGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(rigb.fields) > 1 {
		return nil, errors.New("ent: RepairInvoiceGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := rigb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (rigb *RepairInvoiceGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := rigb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from group-by. It is only allowed when querying group-by with one field.
func (rigb *RepairInvoiceGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = rigb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{repairinvoice.Label}
	default:
		err = fmt.Errorf("ent: RepairInvoiceGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (rigb *RepairInvoiceGroupBy) BoolX(ctx context.Context) bool {
	v, err := rigb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (rigb *RepairInvoiceGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := rigb.sqlQuery().Query()
	if err := rigb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (rigb *RepairInvoiceGroupBy) sqlQuery() *sql.Selector {
	selector := rigb.sql
	columns := make([]string, 0, len(rigb.fields)+len(rigb.fns))
	columns = append(columns, rigb.fields...)
	for _, fn := range rigb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(rigb.fields...)
}

// RepairInvoiceSelect is the builder for select fields of RepairInvoice entities.
type RepairInvoiceSelect struct {
	config
	fields []string
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Scan applies the selector query and scan the result into the given value.
func (ris *RepairInvoiceSelect) Scan(ctx context.Context, v interface{}) error {
	query, err := ris.path(ctx)
	if err != nil {
		return err
	}
	ris.sql = query
	return ris.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ris *RepairInvoiceSelect) ScanX(ctx context.Context, v interface{}) {
	if err := ris.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (ris *RepairInvoiceSelect) Strings(ctx context.Context) ([]string, error) {
	if len(ris.fields) > 1 {
		return nil, errors.New("ent: RepairInvoiceSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := ris.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ris *RepairInvoiceSelect) StringsX(ctx context.Context) []string {
	v, err := ris.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from selector. It is only allowed when selecting one field.
func (ris *RepairInvoiceSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ris.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{repairinvoice.Label}
	default:
		err = fmt.Errorf("ent: RepairInvoiceSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ris *RepairInvoiceSelect) StringX(ctx context.Context) string {
	v, err := ris.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (ris *RepairInvoiceSelect) Ints(ctx context.Context) ([]int, error) {
	if len(ris.fields) > 1 {
		return nil, errors.New("ent: RepairInvoiceSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := ris.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ris *RepairInvoiceSelect) IntsX(ctx context.Context) []int {
	v, err := ris.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from selector. It is only allowed when selecting one field.
func (ris *RepairInvoiceSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ris.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{repairinvoice.Label}
	default:
		err = fmt.Errorf("ent: RepairInvoiceSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ris *RepairInvoiceSelect) IntX(ctx context.Context) int {
	v, err := ris.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (ris *RepairInvoiceSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(ris.fields) > 1 {
		return nil, errors.New("ent: RepairInvoiceSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := ris.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ris *RepairInvoiceSelect) Float64sX(ctx context.Context) []float64 {
	v, err := ris.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from selector. It is only allowed when selecting one field.
func (ris *RepairInvoiceSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ris.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{repairinvoice.Label}
	default:
		err = fmt.Errorf("ent: RepairInvoiceSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ris *RepairInvoiceSelect) Float64X(ctx context.Context) float64 {
	v, err := ris.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (ris *RepairInvoiceSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(ris.fields) > 1 {
		return nil, errors.New("ent: RepairInvoiceSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := ris.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ris *RepairInvoiceSelect) BoolsX(ctx context.Context) []bool {
	v, err := ris.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from selector. It is only allowed when selecting one field.
func (ris *RepairInvoiceSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ris.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{repairinvoice.Label}
	default:
		err = fmt.Errorf("ent: RepairInvoiceSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ris *RepairInvoiceSelect) BoolX(ctx context.Context) bool {
	v, err := ris.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ris *RepairInvoiceSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ris.sqlQuery().Query()
	if err := ris.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ris *RepairInvoiceSelect) sqlQuery() sql.Querier {
	selector := ris.sql
	selector.Select(selector.Columns(ris.fields...)...)
	return selector
}
