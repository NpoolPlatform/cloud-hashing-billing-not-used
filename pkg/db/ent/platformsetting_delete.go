// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/db/ent/platformsetting"
	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/db/ent/predicate"
)

// PlatformSettingDelete is the builder for deleting a PlatformSetting entity.
type PlatformSettingDelete struct {
	config
	hooks    []Hook
	mutation *PlatformSettingMutation
}

// Where appends a list predicates to the PlatformSettingDelete builder.
func (psd *PlatformSettingDelete) Where(ps ...predicate.PlatformSetting) *PlatformSettingDelete {
	psd.mutation.Where(ps...)
	return psd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (psd *PlatformSettingDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(psd.hooks) == 0 {
		affected, err = psd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PlatformSettingMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			psd.mutation = mutation
			affected, err = psd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(psd.hooks) - 1; i >= 0; i-- {
			if psd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = psd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, psd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (psd *PlatformSettingDelete) ExecX(ctx context.Context) int {
	n, err := psd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (psd *PlatformSettingDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: platformsetting.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: platformsetting.FieldID,
			},
		},
	}
	if ps := psd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, psd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// PlatformSettingDeleteOne is the builder for deleting a single PlatformSetting entity.
type PlatformSettingDeleteOne struct {
	psd *PlatformSettingDelete
}

// Exec executes the deletion query.
func (psdo *PlatformSettingDeleteOne) Exec(ctx context.Context) error {
	n, err := psdo.psd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{platformsetting.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (psdo *PlatformSettingDeleteOne) ExecX(ctx context.Context) {
	psdo.psd.ExecX(ctx)
}
