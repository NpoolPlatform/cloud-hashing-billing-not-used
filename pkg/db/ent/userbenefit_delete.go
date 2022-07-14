// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/db/ent/predicate"
	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/db/ent/userbenefit"
)

// UserBenefitDelete is the builder for deleting a UserBenefit entity.
type UserBenefitDelete struct {
	config
	hooks    []Hook
	mutation *UserBenefitMutation
}

// Where appends a list predicates to the UserBenefitDelete builder.
func (ubd *UserBenefitDelete) Where(ps ...predicate.UserBenefit) *UserBenefitDelete {
	ubd.mutation.Where(ps...)
	return ubd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ubd *UserBenefitDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ubd.hooks) == 0 {
		affected, err = ubd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserBenefitMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ubd.mutation = mutation
			affected, err = ubd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ubd.hooks) - 1; i >= 0; i-- {
			if ubd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ubd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ubd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (ubd *UserBenefitDelete) ExecX(ctx context.Context) int {
	n, err := ubd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ubd *UserBenefitDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: userbenefit.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: userbenefit.FieldID,
			},
		},
	}
	if ps := ubd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, ubd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// UserBenefitDeleteOne is the builder for deleting a single UserBenefit entity.
type UserBenefitDeleteOne struct {
	ubd *UserBenefitDelete
}

// Exec executes the deletion query.
func (ubdo *UserBenefitDeleteOne) Exec(ctx context.Context) error {
	n, err := ubdo.ubd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{userbenefit.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ubdo *UserBenefitDeleteOne) ExecX(ctx context.Context) {
	ubdo.ubd.ExecX(ctx)
}
