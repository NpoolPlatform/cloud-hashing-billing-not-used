// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/db/ent/platformbenefit"
	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// PlatformBenefitQuery is the builder for querying PlatformBenefit entities.
type PlatformBenefitQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.PlatformBenefit
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the PlatformBenefitQuery builder.
func (pbq *PlatformBenefitQuery) Where(ps ...predicate.PlatformBenefit) *PlatformBenefitQuery {
	pbq.predicates = append(pbq.predicates, ps...)
	return pbq
}

// Limit adds a limit step to the query.
func (pbq *PlatformBenefitQuery) Limit(limit int) *PlatformBenefitQuery {
	pbq.limit = &limit
	return pbq
}

// Offset adds an offset step to the query.
func (pbq *PlatformBenefitQuery) Offset(offset int) *PlatformBenefitQuery {
	pbq.offset = &offset
	return pbq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (pbq *PlatformBenefitQuery) Unique(unique bool) *PlatformBenefitQuery {
	pbq.unique = &unique
	return pbq
}

// Order adds an order step to the query.
func (pbq *PlatformBenefitQuery) Order(o ...OrderFunc) *PlatformBenefitQuery {
	pbq.order = append(pbq.order, o...)
	return pbq
}

// First returns the first PlatformBenefit entity from the query.
// Returns a *NotFoundError when no PlatformBenefit was found.
func (pbq *PlatformBenefitQuery) First(ctx context.Context) (*PlatformBenefit, error) {
	nodes, err := pbq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{platformbenefit.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (pbq *PlatformBenefitQuery) FirstX(ctx context.Context) *PlatformBenefit {
	node, err := pbq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first PlatformBenefit ID from the query.
// Returns a *NotFoundError when no PlatformBenefit ID was found.
func (pbq *PlatformBenefitQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = pbq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{platformbenefit.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (pbq *PlatformBenefitQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := pbq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single PlatformBenefit entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one PlatformBenefit entity is found.
// Returns a *NotFoundError when no PlatformBenefit entities are found.
func (pbq *PlatformBenefitQuery) Only(ctx context.Context) (*PlatformBenefit, error) {
	nodes, err := pbq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{platformbenefit.Label}
	default:
		return nil, &NotSingularError{platformbenefit.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (pbq *PlatformBenefitQuery) OnlyX(ctx context.Context) *PlatformBenefit {
	node, err := pbq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only PlatformBenefit ID in the query.
// Returns a *NotSingularError when more than one PlatformBenefit ID is found.
// Returns a *NotFoundError when no entities are found.
func (pbq *PlatformBenefitQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = pbq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{platformbenefit.Label}
	default:
		err = &NotSingularError{platformbenefit.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (pbq *PlatformBenefitQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := pbq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of PlatformBenefits.
func (pbq *PlatformBenefitQuery) All(ctx context.Context) ([]*PlatformBenefit, error) {
	if err := pbq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return pbq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (pbq *PlatformBenefitQuery) AllX(ctx context.Context) []*PlatformBenefit {
	nodes, err := pbq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of PlatformBenefit IDs.
func (pbq *PlatformBenefitQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := pbq.Select(platformbenefit.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (pbq *PlatformBenefitQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := pbq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (pbq *PlatformBenefitQuery) Count(ctx context.Context) (int, error) {
	if err := pbq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return pbq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (pbq *PlatformBenefitQuery) CountX(ctx context.Context) int {
	count, err := pbq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (pbq *PlatformBenefitQuery) Exist(ctx context.Context) (bool, error) {
	if err := pbq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return pbq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (pbq *PlatformBenefitQuery) ExistX(ctx context.Context) bool {
	exist, err := pbq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the PlatformBenefitQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (pbq *PlatformBenefitQuery) Clone() *PlatformBenefitQuery {
	if pbq == nil {
		return nil
	}
	return &PlatformBenefitQuery{
		config:     pbq.config,
		limit:      pbq.limit,
		offset:     pbq.offset,
		order:      append([]OrderFunc{}, pbq.order...),
		predicates: append([]predicate.PlatformBenefit{}, pbq.predicates...),
		// clone intermediate query.
		sql:    pbq.sql.Clone(),
		path:   pbq.path,
		unique: pbq.unique,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		GoodID uuid.UUID `json:"good_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.PlatformBenefit.Query().
//		GroupBy(platformbenefit.FieldGoodID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (pbq *PlatformBenefitQuery) GroupBy(field string, fields ...string) *PlatformBenefitGroupBy {
	grbuild := &PlatformBenefitGroupBy{config: pbq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := pbq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return pbq.sqlQuery(ctx), nil
	}
	grbuild.label = platformbenefit.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		GoodID uuid.UUID `json:"good_id,omitempty"`
//	}
//
//	client.PlatformBenefit.Query().
//		Select(platformbenefit.FieldGoodID).
//		Scan(ctx, &v)
//
func (pbq *PlatformBenefitQuery) Select(fields ...string) *PlatformBenefitSelect {
	pbq.fields = append(pbq.fields, fields...)
	selbuild := &PlatformBenefitSelect{PlatformBenefitQuery: pbq}
	selbuild.label = platformbenefit.Label
	selbuild.flds, selbuild.scan = &pbq.fields, selbuild.Scan
	return selbuild
}

func (pbq *PlatformBenefitQuery) prepareQuery(ctx context.Context) error {
	for _, f := range pbq.fields {
		if !platformbenefit.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if pbq.path != nil {
		prev, err := pbq.path(ctx)
		if err != nil {
			return err
		}
		pbq.sql = prev
	}
	return nil
}

func (pbq *PlatformBenefitQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*PlatformBenefit, error) {
	var (
		nodes = []*PlatformBenefit{}
		_spec = pbq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*PlatformBenefit).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &PlatformBenefit{config: pbq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, pbq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (pbq *PlatformBenefitQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := pbq.querySpec()
	_spec.Node.Columns = pbq.fields
	if len(pbq.fields) > 0 {
		_spec.Unique = pbq.unique != nil && *pbq.unique
	}
	return sqlgraph.CountNodes(ctx, pbq.driver, _spec)
}

func (pbq *PlatformBenefitQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := pbq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (pbq *PlatformBenefitQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   platformbenefit.Table,
			Columns: platformbenefit.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: platformbenefit.FieldID,
			},
		},
		From:   pbq.sql,
		Unique: true,
	}
	if unique := pbq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := pbq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, platformbenefit.FieldID)
		for i := range fields {
			if fields[i] != platformbenefit.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := pbq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := pbq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := pbq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := pbq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (pbq *PlatformBenefitQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(pbq.driver.Dialect())
	t1 := builder.Table(platformbenefit.Table)
	columns := pbq.fields
	if len(columns) == 0 {
		columns = platformbenefit.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if pbq.sql != nil {
		selector = pbq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if pbq.unique != nil && *pbq.unique {
		selector.Distinct()
	}
	for _, p := range pbq.predicates {
		p(selector)
	}
	for _, p := range pbq.order {
		p(selector)
	}
	if offset := pbq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := pbq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// PlatformBenefitGroupBy is the group-by builder for PlatformBenefit entities.
type PlatformBenefitGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (pbgb *PlatformBenefitGroupBy) Aggregate(fns ...AggregateFunc) *PlatformBenefitGroupBy {
	pbgb.fns = append(pbgb.fns, fns...)
	return pbgb
}

// Scan applies the group-by query and scans the result into the given value.
func (pbgb *PlatformBenefitGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := pbgb.path(ctx)
	if err != nil {
		return err
	}
	pbgb.sql = query
	return pbgb.sqlScan(ctx, v)
}

func (pbgb *PlatformBenefitGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range pbgb.fields {
		if !platformbenefit.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := pbgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pbgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (pbgb *PlatformBenefitGroupBy) sqlQuery() *sql.Selector {
	selector := pbgb.sql.Select()
	aggregation := make([]string, 0, len(pbgb.fns))
	for _, fn := range pbgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(pbgb.fields)+len(pbgb.fns))
		for _, f := range pbgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(pbgb.fields...)...)
}

// PlatformBenefitSelect is the builder for selecting fields of PlatformBenefit entities.
type PlatformBenefitSelect struct {
	*PlatformBenefitQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (pbs *PlatformBenefitSelect) Scan(ctx context.Context, v interface{}) error {
	if err := pbs.prepareQuery(ctx); err != nil {
		return err
	}
	pbs.sql = pbs.PlatformBenefitQuery.sqlQuery(ctx)
	return pbs.sqlScan(ctx, v)
}

func (pbs *PlatformBenefitSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := pbs.sql.Query()
	if err := pbs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
