// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/db/ent/predicate"
	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/db/ent/userwithdraw"
)

// UserWithdrawDelete is the builder for deleting a UserWithdraw entity.
type UserWithdrawDelete struct {
	config
	hooks    []Hook
	mutation *UserWithdrawMutation
}

// Where appends a list predicates to the UserWithdrawDelete builder.
func (uwd *UserWithdrawDelete) Where(ps ...predicate.UserWithdraw) *UserWithdrawDelete {
	uwd.mutation.Where(ps...)
	return uwd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (uwd *UserWithdrawDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(uwd.hooks) == 0 {
		affected, err = uwd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserWithdrawMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			uwd.mutation = mutation
			affected, err = uwd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(uwd.hooks) - 1; i >= 0; i-- {
			if uwd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = uwd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uwd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (uwd *UserWithdrawDelete) ExecX(ctx context.Context) int {
	n, err := uwd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (uwd *UserWithdrawDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: userwithdraw.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: userwithdraw.FieldID,
			},
		},
	}
	if ps := uwd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, uwd.driver, _spec)
}

// UserWithdrawDeleteOne is the builder for deleting a single UserWithdraw entity.
type UserWithdrawDeleteOne struct {
	uwd *UserWithdrawDelete
}

// Exec executes the deletion query.
func (uwdo *UserWithdrawDeleteOne) Exec(ctx context.Context) error {
	n, err := uwdo.uwd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{userwithdraw.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (uwdo *UserWithdrawDeleteOne) ExecX(ctx context.Context) {
	uwdo.uwd.ExecX(ctx)
}
