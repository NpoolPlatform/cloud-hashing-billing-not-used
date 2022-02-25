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
	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/db/ent/predicate"
	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/db/ent/userpaymentbalance"
	"github.com/google/uuid"
)

// UserPaymentBalanceQuery is the builder for querying UserPaymentBalance entities.
type UserPaymentBalanceQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.UserPaymentBalance
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the UserPaymentBalanceQuery builder.
func (upbq *UserPaymentBalanceQuery) Where(ps ...predicate.UserPaymentBalance) *UserPaymentBalanceQuery {
	upbq.predicates = append(upbq.predicates, ps...)
	return upbq
}

// Limit adds a limit step to the query.
func (upbq *UserPaymentBalanceQuery) Limit(limit int) *UserPaymentBalanceQuery {
	upbq.limit = &limit
	return upbq
}

// Offset adds an offset step to the query.
func (upbq *UserPaymentBalanceQuery) Offset(offset int) *UserPaymentBalanceQuery {
	upbq.offset = &offset
	return upbq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (upbq *UserPaymentBalanceQuery) Unique(unique bool) *UserPaymentBalanceQuery {
	upbq.unique = &unique
	return upbq
}

// Order adds an order step to the query.
func (upbq *UserPaymentBalanceQuery) Order(o ...OrderFunc) *UserPaymentBalanceQuery {
	upbq.order = append(upbq.order, o...)
	return upbq
}

// First returns the first UserPaymentBalance entity from the query.
// Returns a *NotFoundError when no UserPaymentBalance was found.
func (upbq *UserPaymentBalanceQuery) First(ctx context.Context) (*UserPaymentBalance, error) {
	nodes, err := upbq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{userpaymentbalance.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (upbq *UserPaymentBalanceQuery) FirstX(ctx context.Context) *UserPaymentBalance {
	node, err := upbq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first UserPaymentBalance ID from the query.
// Returns a *NotFoundError when no UserPaymentBalance ID was found.
func (upbq *UserPaymentBalanceQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = upbq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{userpaymentbalance.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (upbq *UserPaymentBalanceQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := upbq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single UserPaymentBalance entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one UserPaymentBalance entity is not found.
// Returns a *NotFoundError when no UserPaymentBalance entities are found.
func (upbq *UserPaymentBalanceQuery) Only(ctx context.Context) (*UserPaymentBalance, error) {
	nodes, err := upbq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{userpaymentbalance.Label}
	default:
		return nil, &NotSingularError{userpaymentbalance.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (upbq *UserPaymentBalanceQuery) OnlyX(ctx context.Context) *UserPaymentBalance {
	node, err := upbq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only UserPaymentBalance ID in the query.
// Returns a *NotSingularError when exactly one UserPaymentBalance ID is not found.
// Returns a *NotFoundError when no entities are found.
func (upbq *UserPaymentBalanceQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = upbq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{userpaymentbalance.Label}
	default:
		err = &NotSingularError{userpaymentbalance.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (upbq *UserPaymentBalanceQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := upbq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of UserPaymentBalances.
func (upbq *UserPaymentBalanceQuery) All(ctx context.Context) ([]*UserPaymentBalance, error) {
	if err := upbq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return upbq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (upbq *UserPaymentBalanceQuery) AllX(ctx context.Context) []*UserPaymentBalance {
	nodes, err := upbq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of UserPaymentBalance IDs.
func (upbq *UserPaymentBalanceQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := upbq.Select(userpaymentbalance.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (upbq *UserPaymentBalanceQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := upbq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (upbq *UserPaymentBalanceQuery) Count(ctx context.Context) (int, error) {
	if err := upbq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return upbq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (upbq *UserPaymentBalanceQuery) CountX(ctx context.Context) int {
	count, err := upbq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (upbq *UserPaymentBalanceQuery) Exist(ctx context.Context) (bool, error) {
	if err := upbq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return upbq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (upbq *UserPaymentBalanceQuery) ExistX(ctx context.Context) bool {
	exist, err := upbq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the UserPaymentBalanceQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (upbq *UserPaymentBalanceQuery) Clone() *UserPaymentBalanceQuery {
	if upbq == nil {
		return nil
	}
	return &UserPaymentBalanceQuery{
		config:     upbq.config,
		limit:      upbq.limit,
		offset:     upbq.offset,
		order:      append([]OrderFunc{}, upbq.order...),
		predicates: append([]predicate.UserPaymentBalance{}, upbq.predicates...),
		// clone intermediate query.
		sql:  upbq.sql.Clone(),
		path: upbq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		AppID uuid.UUID `json:"app_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.UserPaymentBalance.Query().
//		GroupBy(userpaymentbalance.FieldAppID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (upbq *UserPaymentBalanceQuery) GroupBy(field string, fields ...string) *UserPaymentBalanceGroupBy {
	group := &UserPaymentBalanceGroupBy{config: upbq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := upbq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return upbq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		AppID uuid.UUID `json:"app_id,omitempty"`
//	}
//
//	client.UserPaymentBalance.Query().
//		Select(userpaymentbalance.FieldAppID).
//		Scan(ctx, &v)
//
func (upbq *UserPaymentBalanceQuery) Select(fields ...string) *UserPaymentBalanceSelect {
	upbq.fields = append(upbq.fields, fields...)
	return &UserPaymentBalanceSelect{UserPaymentBalanceQuery: upbq}
}

func (upbq *UserPaymentBalanceQuery) prepareQuery(ctx context.Context) error {
	for _, f := range upbq.fields {
		if !userpaymentbalance.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if upbq.path != nil {
		prev, err := upbq.path(ctx)
		if err != nil {
			return err
		}
		upbq.sql = prev
	}
	return nil
}

func (upbq *UserPaymentBalanceQuery) sqlAll(ctx context.Context) ([]*UserPaymentBalance, error) {
	var (
		nodes = []*UserPaymentBalance{}
		_spec = upbq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &UserPaymentBalance{config: upbq.config}
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
	if err := sqlgraph.QueryNodes(ctx, upbq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (upbq *UserPaymentBalanceQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := upbq.querySpec()
	_spec.Node.Columns = upbq.fields
	if len(upbq.fields) > 0 {
		_spec.Unique = upbq.unique != nil && *upbq.unique
	}
	return sqlgraph.CountNodes(ctx, upbq.driver, _spec)
}

func (upbq *UserPaymentBalanceQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := upbq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (upbq *UserPaymentBalanceQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   userpaymentbalance.Table,
			Columns: userpaymentbalance.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: userpaymentbalance.FieldID,
			},
		},
		From:   upbq.sql,
		Unique: true,
	}
	if unique := upbq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := upbq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, userpaymentbalance.FieldID)
		for i := range fields {
			if fields[i] != userpaymentbalance.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := upbq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := upbq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := upbq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := upbq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (upbq *UserPaymentBalanceQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(upbq.driver.Dialect())
	t1 := builder.Table(userpaymentbalance.Table)
	columns := upbq.fields
	if len(columns) == 0 {
		columns = userpaymentbalance.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if upbq.sql != nil {
		selector = upbq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if upbq.unique != nil && *upbq.unique {
		selector.Distinct()
	}
	for _, p := range upbq.predicates {
		p(selector)
	}
	for _, p := range upbq.order {
		p(selector)
	}
	if offset := upbq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := upbq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// UserPaymentBalanceGroupBy is the group-by builder for UserPaymentBalance entities.
type UserPaymentBalanceGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (upbgb *UserPaymentBalanceGroupBy) Aggregate(fns ...AggregateFunc) *UserPaymentBalanceGroupBy {
	upbgb.fns = append(upbgb.fns, fns...)
	return upbgb
}

// Scan applies the group-by query and scans the result into the given value.
func (upbgb *UserPaymentBalanceGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := upbgb.path(ctx)
	if err != nil {
		return err
	}
	upbgb.sql = query
	return upbgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (upbgb *UserPaymentBalanceGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := upbgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (upbgb *UserPaymentBalanceGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(upbgb.fields) > 1 {
		return nil, errors.New("ent: UserPaymentBalanceGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := upbgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (upbgb *UserPaymentBalanceGroupBy) StringsX(ctx context.Context) []string {
	v, err := upbgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (upbgb *UserPaymentBalanceGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = upbgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{userpaymentbalance.Label}
	default:
		err = fmt.Errorf("ent: UserPaymentBalanceGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (upbgb *UserPaymentBalanceGroupBy) StringX(ctx context.Context) string {
	v, err := upbgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (upbgb *UserPaymentBalanceGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(upbgb.fields) > 1 {
		return nil, errors.New("ent: UserPaymentBalanceGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := upbgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (upbgb *UserPaymentBalanceGroupBy) IntsX(ctx context.Context) []int {
	v, err := upbgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (upbgb *UserPaymentBalanceGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = upbgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{userpaymentbalance.Label}
	default:
		err = fmt.Errorf("ent: UserPaymentBalanceGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (upbgb *UserPaymentBalanceGroupBy) IntX(ctx context.Context) int {
	v, err := upbgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (upbgb *UserPaymentBalanceGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(upbgb.fields) > 1 {
		return nil, errors.New("ent: UserPaymentBalanceGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := upbgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (upbgb *UserPaymentBalanceGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := upbgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (upbgb *UserPaymentBalanceGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = upbgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{userpaymentbalance.Label}
	default:
		err = fmt.Errorf("ent: UserPaymentBalanceGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (upbgb *UserPaymentBalanceGroupBy) Float64X(ctx context.Context) float64 {
	v, err := upbgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (upbgb *UserPaymentBalanceGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(upbgb.fields) > 1 {
		return nil, errors.New("ent: UserPaymentBalanceGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := upbgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (upbgb *UserPaymentBalanceGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := upbgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (upbgb *UserPaymentBalanceGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = upbgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{userpaymentbalance.Label}
	default:
		err = fmt.Errorf("ent: UserPaymentBalanceGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (upbgb *UserPaymentBalanceGroupBy) BoolX(ctx context.Context) bool {
	v, err := upbgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (upbgb *UserPaymentBalanceGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range upbgb.fields {
		if !userpaymentbalance.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := upbgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := upbgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (upbgb *UserPaymentBalanceGroupBy) sqlQuery() *sql.Selector {
	selector := upbgb.sql.Select()
	aggregation := make([]string, 0, len(upbgb.fns))
	for _, fn := range upbgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(upbgb.fields)+len(upbgb.fns))
		for _, f := range upbgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(upbgb.fields...)...)
}

// UserPaymentBalanceSelect is the builder for selecting fields of UserPaymentBalance entities.
type UserPaymentBalanceSelect struct {
	*UserPaymentBalanceQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (upbs *UserPaymentBalanceSelect) Scan(ctx context.Context, v interface{}) error {
	if err := upbs.prepareQuery(ctx); err != nil {
		return err
	}
	upbs.sql = upbs.UserPaymentBalanceQuery.sqlQuery(ctx)
	return upbs.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (upbs *UserPaymentBalanceSelect) ScanX(ctx context.Context, v interface{}) {
	if err := upbs.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (upbs *UserPaymentBalanceSelect) Strings(ctx context.Context) ([]string, error) {
	if len(upbs.fields) > 1 {
		return nil, errors.New("ent: UserPaymentBalanceSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := upbs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (upbs *UserPaymentBalanceSelect) StringsX(ctx context.Context) []string {
	v, err := upbs.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (upbs *UserPaymentBalanceSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = upbs.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{userpaymentbalance.Label}
	default:
		err = fmt.Errorf("ent: UserPaymentBalanceSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (upbs *UserPaymentBalanceSelect) StringX(ctx context.Context) string {
	v, err := upbs.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (upbs *UserPaymentBalanceSelect) Ints(ctx context.Context) ([]int, error) {
	if len(upbs.fields) > 1 {
		return nil, errors.New("ent: UserPaymentBalanceSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := upbs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (upbs *UserPaymentBalanceSelect) IntsX(ctx context.Context) []int {
	v, err := upbs.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (upbs *UserPaymentBalanceSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = upbs.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{userpaymentbalance.Label}
	default:
		err = fmt.Errorf("ent: UserPaymentBalanceSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (upbs *UserPaymentBalanceSelect) IntX(ctx context.Context) int {
	v, err := upbs.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (upbs *UserPaymentBalanceSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(upbs.fields) > 1 {
		return nil, errors.New("ent: UserPaymentBalanceSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := upbs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (upbs *UserPaymentBalanceSelect) Float64sX(ctx context.Context) []float64 {
	v, err := upbs.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (upbs *UserPaymentBalanceSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = upbs.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{userpaymentbalance.Label}
	default:
		err = fmt.Errorf("ent: UserPaymentBalanceSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (upbs *UserPaymentBalanceSelect) Float64X(ctx context.Context) float64 {
	v, err := upbs.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (upbs *UserPaymentBalanceSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(upbs.fields) > 1 {
		return nil, errors.New("ent: UserPaymentBalanceSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := upbs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (upbs *UserPaymentBalanceSelect) BoolsX(ctx context.Context) []bool {
	v, err := upbs.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (upbs *UserPaymentBalanceSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = upbs.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{userpaymentbalance.Label}
	default:
		err = fmt.Errorf("ent: UserPaymentBalanceSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (upbs *UserPaymentBalanceSelect) BoolX(ctx context.Context) bool {
	v, err := upbs.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (upbs *UserPaymentBalanceSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := upbs.sql.Query()
	if err := upbs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
