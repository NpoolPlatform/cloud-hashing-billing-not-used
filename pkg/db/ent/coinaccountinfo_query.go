// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/db/ent/coinaccountinfo"
	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// CoinAccountInfoQuery is the builder for querying CoinAccountInfo entities.
type CoinAccountInfoQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.CoinAccountInfo
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the CoinAccountInfoQuery builder.
func (caiq *CoinAccountInfoQuery) Where(ps ...predicate.CoinAccountInfo) *CoinAccountInfoQuery {
	caiq.predicates = append(caiq.predicates, ps...)
	return caiq
}

// Limit adds a limit step to the query.
func (caiq *CoinAccountInfoQuery) Limit(limit int) *CoinAccountInfoQuery {
	caiq.limit = &limit
	return caiq
}

// Offset adds an offset step to the query.
func (caiq *CoinAccountInfoQuery) Offset(offset int) *CoinAccountInfoQuery {
	caiq.offset = &offset
	return caiq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (caiq *CoinAccountInfoQuery) Unique(unique bool) *CoinAccountInfoQuery {
	caiq.unique = &unique
	return caiq
}

// Order adds an order step to the query.
func (caiq *CoinAccountInfoQuery) Order(o ...OrderFunc) *CoinAccountInfoQuery {
	caiq.order = append(caiq.order, o...)
	return caiq
}

// First returns the first CoinAccountInfo entity from the query.
// Returns a *NotFoundError when no CoinAccountInfo was found.
func (caiq *CoinAccountInfoQuery) First(ctx context.Context) (*CoinAccountInfo, error) {
	nodes, err := caiq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{coinaccountinfo.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (caiq *CoinAccountInfoQuery) FirstX(ctx context.Context) *CoinAccountInfo {
	node, err := caiq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first CoinAccountInfo ID from the query.
// Returns a *NotFoundError when no CoinAccountInfo ID was found.
func (caiq *CoinAccountInfoQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = caiq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{coinaccountinfo.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (caiq *CoinAccountInfoQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := caiq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single CoinAccountInfo entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one CoinAccountInfo entity is not found.
// Returns a *NotFoundError when no CoinAccountInfo entities are found.
func (caiq *CoinAccountInfoQuery) Only(ctx context.Context) (*CoinAccountInfo, error) {
	nodes, err := caiq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{coinaccountinfo.Label}
	default:
		return nil, &NotSingularError{coinaccountinfo.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (caiq *CoinAccountInfoQuery) OnlyX(ctx context.Context) *CoinAccountInfo {
	node, err := caiq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only CoinAccountInfo ID in the query.
// Returns a *NotSingularError when exactly one CoinAccountInfo ID is not found.
// Returns a *NotFoundError when no entities are found.
func (caiq *CoinAccountInfoQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = caiq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{coinaccountinfo.Label}
	default:
		err = &NotSingularError{coinaccountinfo.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (caiq *CoinAccountInfoQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := caiq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of CoinAccountInfos.
func (caiq *CoinAccountInfoQuery) All(ctx context.Context) ([]*CoinAccountInfo, error) {
	if err := caiq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return caiq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (caiq *CoinAccountInfoQuery) AllX(ctx context.Context) []*CoinAccountInfo {
	nodes, err := caiq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of CoinAccountInfo IDs.
func (caiq *CoinAccountInfoQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := caiq.Select(coinaccountinfo.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (caiq *CoinAccountInfoQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := caiq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (caiq *CoinAccountInfoQuery) Count(ctx context.Context) (int, error) {
	if err := caiq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return caiq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (caiq *CoinAccountInfoQuery) CountX(ctx context.Context) int {
	count, err := caiq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (caiq *CoinAccountInfoQuery) Exist(ctx context.Context) (bool, error) {
	if err := caiq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return caiq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (caiq *CoinAccountInfoQuery) ExistX(ctx context.Context) bool {
	exist, err := caiq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the CoinAccountInfoQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (caiq *CoinAccountInfoQuery) Clone() *CoinAccountInfoQuery {
	if caiq == nil {
		return nil
	}
	return &CoinAccountInfoQuery{
		config:     caiq.config,
		limit:      caiq.limit,
		offset:     caiq.offset,
		order:      append([]OrderFunc{}, caiq.order...),
		predicates: append([]predicate.CoinAccountInfo{}, caiq.predicates...),
		// clone intermediate query.
		sql:  caiq.sql.Clone(),
		path: caiq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CoinTypeID uuid.UUID `json:"coin_type_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.CoinAccountInfo.Query().
//		GroupBy(coinaccountinfo.FieldCoinTypeID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (caiq *CoinAccountInfoQuery) GroupBy(field string, fields ...string) *CoinAccountInfoGroupBy {
	group := &CoinAccountInfoGroupBy{config: caiq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := caiq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return caiq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CoinTypeID uuid.UUID `json:"coin_type_id,omitempty"`
//	}
//
//	client.CoinAccountInfo.Query().
//		Select(coinaccountinfo.FieldCoinTypeID).
//		Scan(ctx, &v)
//
func (caiq *CoinAccountInfoQuery) Select(fields ...string) *CoinAccountInfoSelect {
	caiq.fields = append(caiq.fields, fields...)
	return &CoinAccountInfoSelect{CoinAccountInfoQuery: caiq}
}

func (caiq *CoinAccountInfoQuery) prepareQuery(ctx context.Context) error {
	for _, f := range caiq.fields {
		if !coinaccountinfo.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if caiq.path != nil {
		prev, err := caiq.path(ctx)
		if err != nil {
			return err
		}
		caiq.sql = prev
	}
	return nil
}

func (caiq *CoinAccountInfoQuery) sqlAll(ctx context.Context) ([]*CoinAccountInfo, error) {
	var (
		nodes = []*CoinAccountInfo{}
		_spec = caiq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &CoinAccountInfo{config: caiq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, caiq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (caiq *CoinAccountInfoQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := caiq.querySpec()
	return sqlgraph.CountNodes(ctx, caiq.driver, _spec)
}

func (caiq *CoinAccountInfoQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := caiq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (caiq *CoinAccountInfoQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   coinaccountinfo.Table,
			Columns: coinaccountinfo.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: coinaccountinfo.FieldID,
			},
		},
		From:   caiq.sql,
		Unique: true,
	}
	if unique := caiq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := caiq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, coinaccountinfo.FieldID)
		for i := range fields {
			if fields[i] != coinaccountinfo.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := caiq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := caiq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := caiq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := caiq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (caiq *CoinAccountInfoQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(caiq.driver.Dialect())
	t1 := builder.Table(coinaccountinfo.Table)
	columns := caiq.fields
	if len(columns) == 0 {
		columns = coinaccountinfo.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if caiq.sql != nil {
		selector = caiq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	for _, p := range caiq.predicates {
		p(selector)
	}
	for _, p := range caiq.order {
		p(selector)
	}
	if offset := caiq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := caiq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// CoinAccountInfoGroupBy is the group-by builder for CoinAccountInfo entities.
type CoinAccountInfoGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (caigb *CoinAccountInfoGroupBy) Aggregate(fns ...AggregateFunc) *CoinAccountInfoGroupBy {
	caigb.fns = append(caigb.fns, fns...)
	return caigb
}

// Scan applies the group-by query and scans the result into the given value.
func (caigb *CoinAccountInfoGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := caigb.path(ctx)
	if err != nil {
		return err
	}
	caigb.sql = query
	return caigb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (caigb *CoinAccountInfoGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := caigb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (caigb *CoinAccountInfoGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(caigb.fields) > 1 {
		return nil, errors.New("ent: CoinAccountInfoGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := caigb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (caigb *CoinAccountInfoGroupBy) StringsX(ctx context.Context) []string {
	v, err := caigb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (caigb *CoinAccountInfoGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = caigb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{coinaccountinfo.Label}
	default:
		err = fmt.Errorf("ent: CoinAccountInfoGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (caigb *CoinAccountInfoGroupBy) StringX(ctx context.Context) string {
	v, err := caigb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (caigb *CoinAccountInfoGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(caigb.fields) > 1 {
		return nil, errors.New("ent: CoinAccountInfoGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := caigb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (caigb *CoinAccountInfoGroupBy) IntsX(ctx context.Context) []int {
	v, err := caigb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (caigb *CoinAccountInfoGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = caigb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{coinaccountinfo.Label}
	default:
		err = fmt.Errorf("ent: CoinAccountInfoGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (caigb *CoinAccountInfoGroupBy) IntX(ctx context.Context) int {
	v, err := caigb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (caigb *CoinAccountInfoGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(caigb.fields) > 1 {
		return nil, errors.New("ent: CoinAccountInfoGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := caigb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (caigb *CoinAccountInfoGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := caigb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (caigb *CoinAccountInfoGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = caigb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{coinaccountinfo.Label}
	default:
		err = fmt.Errorf("ent: CoinAccountInfoGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (caigb *CoinAccountInfoGroupBy) Float64X(ctx context.Context) float64 {
	v, err := caigb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (caigb *CoinAccountInfoGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(caigb.fields) > 1 {
		return nil, errors.New("ent: CoinAccountInfoGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := caigb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (caigb *CoinAccountInfoGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := caigb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (caigb *CoinAccountInfoGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = caigb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{coinaccountinfo.Label}
	default:
		err = fmt.Errorf("ent: CoinAccountInfoGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (caigb *CoinAccountInfoGroupBy) BoolX(ctx context.Context) bool {
	v, err := caigb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (caigb *CoinAccountInfoGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range caigb.fields {
		if !coinaccountinfo.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := caigb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := caigb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (caigb *CoinAccountInfoGroupBy) sqlQuery() *sql.Selector {
	selector := caigb.sql.Select()
	aggregation := make([]string, 0, len(caigb.fns))
	for _, fn := range caigb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(caigb.fields)+len(caigb.fns))
		for _, f := range caigb.fields {
			columns = append(columns, selector.C(f))
		}
		for _, c := range aggregation {
			columns = append(columns, c)
		}
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(caigb.fields...)...)
}

// CoinAccountInfoSelect is the builder for selecting fields of CoinAccountInfo entities.
type CoinAccountInfoSelect struct {
	*CoinAccountInfoQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (cais *CoinAccountInfoSelect) Scan(ctx context.Context, v interface{}) error {
	if err := cais.prepareQuery(ctx); err != nil {
		return err
	}
	cais.sql = cais.CoinAccountInfoQuery.sqlQuery(ctx)
	return cais.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (cais *CoinAccountInfoSelect) ScanX(ctx context.Context, v interface{}) {
	if err := cais.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (cais *CoinAccountInfoSelect) Strings(ctx context.Context) ([]string, error) {
	if len(cais.fields) > 1 {
		return nil, errors.New("ent: CoinAccountInfoSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := cais.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (cais *CoinAccountInfoSelect) StringsX(ctx context.Context) []string {
	v, err := cais.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (cais *CoinAccountInfoSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = cais.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{coinaccountinfo.Label}
	default:
		err = fmt.Errorf("ent: CoinAccountInfoSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (cais *CoinAccountInfoSelect) StringX(ctx context.Context) string {
	v, err := cais.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (cais *CoinAccountInfoSelect) Ints(ctx context.Context) ([]int, error) {
	if len(cais.fields) > 1 {
		return nil, errors.New("ent: CoinAccountInfoSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := cais.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (cais *CoinAccountInfoSelect) IntsX(ctx context.Context) []int {
	v, err := cais.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (cais *CoinAccountInfoSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = cais.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{coinaccountinfo.Label}
	default:
		err = fmt.Errorf("ent: CoinAccountInfoSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (cais *CoinAccountInfoSelect) IntX(ctx context.Context) int {
	v, err := cais.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (cais *CoinAccountInfoSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(cais.fields) > 1 {
		return nil, errors.New("ent: CoinAccountInfoSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := cais.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (cais *CoinAccountInfoSelect) Float64sX(ctx context.Context) []float64 {
	v, err := cais.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (cais *CoinAccountInfoSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = cais.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{coinaccountinfo.Label}
	default:
		err = fmt.Errorf("ent: CoinAccountInfoSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (cais *CoinAccountInfoSelect) Float64X(ctx context.Context) float64 {
	v, err := cais.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (cais *CoinAccountInfoSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(cais.fields) > 1 {
		return nil, errors.New("ent: CoinAccountInfoSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := cais.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (cais *CoinAccountInfoSelect) BoolsX(ctx context.Context) []bool {
	v, err := cais.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (cais *CoinAccountInfoSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = cais.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{coinaccountinfo.Label}
	default:
		err = fmt.Errorf("ent: CoinAccountInfoSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (cais *CoinAccountInfoSelect) BoolX(ctx context.Context) bool {
	v, err := cais.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (cais *CoinAccountInfoSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := cais.sql.Query()
	if err := cais.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
