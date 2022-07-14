// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/db/ent/goodpayment"
	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/db/ent/predicate"
)

// GoodPaymentDelete is the builder for deleting a GoodPayment entity.
type GoodPaymentDelete struct {
	config
	hooks    []Hook
	mutation *GoodPaymentMutation
}

// Where appends a list predicates to the GoodPaymentDelete builder.
func (gpd *GoodPaymentDelete) Where(ps ...predicate.GoodPayment) *GoodPaymentDelete {
	gpd.mutation.Where(ps...)
	return gpd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (gpd *GoodPaymentDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(gpd.hooks) == 0 {
		affected, err = gpd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GoodPaymentMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			gpd.mutation = mutation
			affected, err = gpd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(gpd.hooks) - 1; i >= 0; i-- {
			if gpd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = gpd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, gpd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (gpd *GoodPaymentDelete) ExecX(ctx context.Context) int {
	n, err := gpd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (gpd *GoodPaymentDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: goodpayment.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: goodpayment.FieldID,
			},
		},
	}
	if ps := gpd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, gpd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// GoodPaymentDeleteOne is the builder for deleting a single GoodPayment entity.
type GoodPaymentDeleteOne struct {
	gpd *GoodPaymentDelete
}

// Exec executes the deletion query.
func (gpdo *GoodPaymentDeleteOne) Exec(ctx context.Context) error {
	n, err := gpdo.gpd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{goodpayment.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (gpdo *GoodPaymentDeleteOne) ExecX(ctx context.Context) {
	gpdo.gpd.ExecX(ctx)
}
