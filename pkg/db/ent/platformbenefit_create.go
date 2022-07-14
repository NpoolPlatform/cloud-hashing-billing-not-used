// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/db/ent/platformbenefit"
	"github.com/google/uuid"
)

// PlatformBenefitCreate is the builder for creating a PlatformBenefit entity.
type PlatformBenefitCreate struct {
	config
	mutation *PlatformBenefitMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetGoodID sets the "good_id" field.
func (pbc *PlatformBenefitCreate) SetGoodID(u uuid.UUID) *PlatformBenefitCreate {
	pbc.mutation.SetGoodID(u)
	return pbc
}

// SetBenefitAccountID sets the "benefit_account_id" field.
func (pbc *PlatformBenefitCreate) SetBenefitAccountID(u uuid.UUID) *PlatformBenefitCreate {
	pbc.mutation.SetBenefitAccountID(u)
	return pbc
}

// SetAmount sets the "amount" field.
func (pbc *PlatformBenefitCreate) SetAmount(u uint64) *PlatformBenefitCreate {
	pbc.mutation.SetAmount(u)
	return pbc
}

// SetLastBenefitTimestamp sets the "last_benefit_timestamp" field.
func (pbc *PlatformBenefitCreate) SetLastBenefitTimestamp(u uint32) *PlatformBenefitCreate {
	pbc.mutation.SetLastBenefitTimestamp(u)
	return pbc
}

// SetChainTransactionID sets the "chain_transaction_id" field.
func (pbc *PlatformBenefitCreate) SetChainTransactionID(s string) *PlatformBenefitCreate {
	pbc.mutation.SetChainTransactionID(s)
	return pbc
}

// SetCreateAt sets the "create_at" field.
func (pbc *PlatformBenefitCreate) SetCreateAt(u uint32) *PlatformBenefitCreate {
	pbc.mutation.SetCreateAt(u)
	return pbc
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (pbc *PlatformBenefitCreate) SetNillableCreateAt(u *uint32) *PlatformBenefitCreate {
	if u != nil {
		pbc.SetCreateAt(*u)
	}
	return pbc
}

// SetUpdateAt sets the "update_at" field.
func (pbc *PlatformBenefitCreate) SetUpdateAt(u uint32) *PlatformBenefitCreate {
	pbc.mutation.SetUpdateAt(u)
	return pbc
}

// SetNillableUpdateAt sets the "update_at" field if the given value is not nil.
func (pbc *PlatformBenefitCreate) SetNillableUpdateAt(u *uint32) *PlatformBenefitCreate {
	if u != nil {
		pbc.SetUpdateAt(*u)
	}
	return pbc
}

// SetDeleteAt sets the "delete_at" field.
func (pbc *PlatformBenefitCreate) SetDeleteAt(u uint32) *PlatformBenefitCreate {
	pbc.mutation.SetDeleteAt(u)
	return pbc
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (pbc *PlatformBenefitCreate) SetNillableDeleteAt(u *uint32) *PlatformBenefitCreate {
	if u != nil {
		pbc.SetDeleteAt(*u)
	}
	return pbc
}

// SetID sets the "id" field.
func (pbc *PlatformBenefitCreate) SetID(u uuid.UUID) *PlatformBenefitCreate {
	pbc.mutation.SetID(u)
	return pbc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (pbc *PlatformBenefitCreate) SetNillableID(u *uuid.UUID) *PlatformBenefitCreate {
	if u != nil {
		pbc.SetID(*u)
	}
	return pbc
}

// Mutation returns the PlatformBenefitMutation object of the builder.
func (pbc *PlatformBenefitCreate) Mutation() *PlatformBenefitMutation {
	return pbc.mutation
}

// Save creates the PlatformBenefit in the database.
func (pbc *PlatformBenefitCreate) Save(ctx context.Context) (*PlatformBenefit, error) {
	var (
		err  error
		node *PlatformBenefit
	)
	pbc.defaults()
	if len(pbc.hooks) == 0 {
		if err = pbc.check(); err != nil {
			return nil, err
		}
		node, err = pbc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PlatformBenefitMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = pbc.check(); err != nil {
				return nil, err
			}
			pbc.mutation = mutation
			if node, err = pbc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(pbc.hooks) - 1; i >= 0; i-- {
			if pbc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = pbc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, pbc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*PlatformBenefit)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from PlatformBenefitMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (pbc *PlatformBenefitCreate) SaveX(ctx context.Context) *PlatformBenefit {
	v, err := pbc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pbc *PlatformBenefitCreate) Exec(ctx context.Context) error {
	_, err := pbc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pbc *PlatformBenefitCreate) ExecX(ctx context.Context) {
	if err := pbc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pbc *PlatformBenefitCreate) defaults() {
	if _, ok := pbc.mutation.CreateAt(); !ok {
		v := platformbenefit.DefaultCreateAt()
		pbc.mutation.SetCreateAt(v)
	}
	if _, ok := pbc.mutation.UpdateAt(); !ok {
		v := platformbenefit.DefaultUpdateAt()
		pbc.mutation.SetUpdateAt(v)
	}
	if _, ok := pbc.mutation.DeleteAt(); !ok {
		v := platformbenefit.DefaultDeleteAt()
		pbc.mutation.SetDeleteAt(v)
	}
	if _, ok := pbc.mutation.ID(); !ok {
		v := platformbenefit.DefaultID()
		pbc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pbc *PlatformBenefitCreate) check() error {
	if _, ok := pbc.mutation.GoodID(); !ok {
		return &ValidationError{Name: "good_id", err: errors.New(`ent: missing required field "PlatformBenefit.good_id"`)}
	}
	if _, ok := pbc.mutation.BenefitAccountID(); !ok {
		return &ValidationError{Name: "benefit_account_id", err: errors.New(`ent: missing required field "PlatformBenefit.benefit_account_id"`)}
	}
	if _, ok := pbc.mutation.Amount(); !ok {
		return &ValidationError{Name: "amount", err: errors.New(`ent: missing required field "PlatformBenefit.amount"`)}
	}
	if _, ok := pbc.mutation.LastBenefitTimestamp(); !ok {
		return &ValidationError{Name: "last_benefit_timestamp", err: errors.New(`ent: missing required field "PlatformBenefit.last_benefit_timestamp"`)}
	}
	if _, ok := pbc.mutation.ChainTransactionID(); !ok {
		return &ValidationError{Name: "chain_transaction_id", err: errors.New(`ent: missing required field "PlatformBenefit.chain_transaction_id"`)}
	}
	if _, ok := pbc.mutation.CreateAt(); !ok {
		return &ValidationError{Name: "create_at", err: errors.New(`ent: missing required field "PlatformBenefit.create_at"`)}
	}
	if _, ok := pbc.mutation.UpdateAt(); !ok {
		return &ValidationError{Name: "update_at", err: errors.New(`ent: missing required field "PlatformBenefit.update_at"`)}
	}
	if _, ok := pbc.mutation.DeleteAt(); !ok {
		return &ValidationError{Name: "delete_at", err: errors.New(`ent: missing required field "PlatformBenefit.delete_at"`)}
	}
	return nil
}

func (pbc *PlatformBenefitCreate) sqlSave(ctx context.Context) (*PlatformBenefit, error) {
	_node, _spec := pbc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pbc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	return _node, nil
}

func (pbc *PlatformBenefitCreate) createSpec() (*PlatformBenefit, *sqlgraph.CreateSpec) {
	var (
		_node = &PlatformBenefit{config: pbc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: platformbenefit.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: platformbenefit.FieldID,
			},
		}
	)
	_spec.OnConflict = pbc.conflict
	if id, ok := pbc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := pbc.mutation.GoodID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: platformbenefit.FieldGoodID,
		})
		_node.GoodID = value
	}
	if value, ok := pbc.mutation.BenefitAccountID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: platformbenefit.FieldBenefitAccountID,
		})
		_node.BenefitAccountID = value
	}
	if value, ok := pbc.mutation.Amount(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: platformbenefit.FieldAmount,
		})
		_node.Amount = value
	}
	if value, ok := pbc.mutation.LastBenefitTimestamp(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: platformbenefit.FieldLastBenefitTimestamp,
		})
		_node.LastBenefitTimestamp = value
	}
	if value, ok := pbc.mutation.ChainTransactionID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: platformbenefit.FieldChainTransactionID,
		})
		_node.ChainTransactionID = value
	}
	if value, ok := pbc.mutation.CreateAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: platformbenefit.FieldCreateAt,
		})
		_node.CreateAt = value
	}
	if value, ok := pbc.mutation.UpdateAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: platformbenefit.FieldUpdateAt,
		})
		_node.UpdateAt = value
	}
	if value, ok := pbc.mutation.DeleteAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: platformbenefit.FieldDeleteAt,
		})
		_node.DeleteAt = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.PlatformBenefit.Create().
//		SetGoodID(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.PlatformBenefitUpsert) {
//			SetGoodID(v+v).
//		}).
//		Exec(ctx)
//
func (pbc *PlatformBenefitCreate) OnConflict(opts ...sql.ConflictOption) *PlatformBenefitUpsertOne {
	pbc.conflict = opts
	return &PlatformBenefitUpsertOne{
		create: pbc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.PlatformBenefit.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (pbc *PlatformBenefitCreate) OnConflictColumns(columns ...string) *PlatformBenefitUpsertOne {
	pbc.conflict = append(pbc.conflict, sql.ConflictColumns(columns...))
	return &PlatformBenefitUpsertOne{
		create: pbc,
	}
}

type (
	// PlatformBenefitUpsertOne is the builder for "upsert"-ing
	//  one PlatformBenefit node.
	PlatformBenefitUpsertOne struct {
		create *PlatformBenefitCreate
	}

	// PlatformBenefitUpsert is the "OnConflict" setter.
	PlatformBenefitUpsert struct {
		*sql.UpdateSet
	}
)

// SetGoodID sets the "good_id" field.
func (u *PlatformBenefitUpsert) SetGoodID(v uuid.UUID) *PlatformBenefitUpsert {
	u.Set(platformbenefit.FieldGoodID, v)
	return u
}

// UpdateGoodID sets the "good_id" field to the value that was provided on create.
func (u *PlatformBenefitUpsert) UpdateGoodID() *PlatformBenefitUpsert {
	u.SetExcluded(platformbenefit.FieldGoodID)
	return u
}

// SetBenefitAccountID sets the "benefit_account_id" field.
func (u *PlatformBenefitUpsert) SetBenefitAccountID(v uuid.UUID) *PlatformBenefitUpsert {
	u.Set(platformbenefit.FieldBenefitAccountID, v)
	return u
}

// UpdateBenefitAccountID sets the "benefit_account_id" field to the value that was provided on create.
func (u *PlatformBenefitUpsert) UpdateBenefitAccountID() *PlatformBenefitUpsert {
	u.SetExcluded(platformbenefit.FieldBenefitAccountID)
	return u
}

// SetAmount sets the "amount" field.
func (u *PlatformBenefitUpsert) SetAmount(v uint64) *PlatformBenefitUpsert {
	u.Set(platformbenefit.FieldAmount, v)
	return u
}

// UpdateAmount sets the "amount" field to the value that was provided on create.
func (u *PlatformBenefitUpsert) UpdateAmount() *PlatformBenefitUpsert {
	u.SetExcluded(platformbenefit.FieldAmount)
	return u
}

// AddAmount adds v to the "amount" field.
func (u *PlatformBenefitUpsert) AddAmount(v uint64) *PlatformBenefitUpsert {
	u.Add(platformbenefit.FieldAmount, v)
	return u
}

// SetLastBenefitTimestamp sets the "last_benefit_timestamp" field.
func (u *PlatformBenefitUpsert) SetLastBenefitTimestamp(v uint32) *PlatformBenefitUpsert {
	u.Set(platformbenefit.FieldLastBenefitTimestamp, v)
	return u
}

// UpdateLastBenefitTimestamp sets the "last_benefit_timestamp" field to the value that was provided on create.
func (u *PlatformBenefitUpsert) UpdateLastBenefitTimestamp() *PlatformBenefitUpsert {
	u.SetExcluded(platformbenefit.FieldLastBenefitTimestamp)
	return u
}

// AddLastBenefitTimestamp adds v to the "last_benefit_timestamp" field.
func (u *PlatformBenefitUpsert) AddLastBenefitTimestamp(v uint32) *PlatformBenefitUpsert {
	u.Add(platformbenefit.FieldLastBenefitTimestamp, v)
	return u
}

// SetChainTransactionID sets the "chain_transaction_id" field.
func (u *PlatformBenefitUpsert) SetChainTransactionID(v string) *PlatformBenefitUpsert {
	u.Set(platformbenefit.FieldChainTransactionID, v)
	return u
}

// UpdateChainTransactionID sets the "chain_transaction_id" field to the value that was provided on create.
func (u *PlatformBenefitUpsert) UpdateChainTransactionID() *PlatformBenefitUpsert {
	u.SetExcluded(platformbenefit.FieldChainTransactionID)
	return u
}

// SetCreateAt sets the "create_at" field.
func (u *PlatformBenefitUpsert) SetCreateAt(v uint32) *PlatformBenefitUpsert {
	u.Set(platformbenefit.FieldCreateAt, v)
	return u
}

// UpdateCreateAt sets the "create_at" field to the value that was provided on create.
func (u *PlatformBenefitUpsert) UpdateCreateAt() *PlatformBenefitUpsert {
	u.SetExcluded(platformbenefit.FieldCreateAt)
	return u
}

// AddCreateAt adds v to the "create_at" field.
func (u *PlatformBenefitUpsert) AddCreateAt(v uint32) *PlatformBenefitUpsert {
	u.Add(platformbenefit.FieldCreateAt, v)
	return u
}

// SetUpdateAt sets the "update_at" field.
func (u *PlatformBenefitUpsert) SetUpdateAt(v uint32) *PlatformBenefitUpsert {
	u.Set(platformbenefit.FieldUpdateAt, v)
	return u
}

// UpdateUpdateAt sets the "update_at" field to the value that was provided on create.
func (u *PlatformBenefitUpsert) UpdateUpdateAt() *PlatformBenefitUpsert {
	u.SetExcluded(platformbenefit.FieldUpdateAt)
	return u
}

// AddUpdateAt adds v to the "update_at" field.
func (u *PlatformBenefitUpsert) AddUpdateAt(v uint32) *PlatformBenefitUpsert {
	u.Add(platformbenefit.FieldUpdateAt, v)
	return u
}

// SetDeleteAt sets the "delete_at" field.
func (u *PlatformBenefitUpsert) SetDeleteAt(v uint32) *PlatformBenefitUpsert {
	u.Set(platformbenefit.FieldDeleteAt, v)
	return u
}

// UpdateDeleteAt sets the "delete_at" field to the value that was provided on create.
func (u *PlatformBenefitUpsert) UpdateDeleteAt() *PlatformBenefitUpsert {
	u.SetExcluded(platformbenefit.FieldDeleteAt)
	return u
}

// AddDeleteAt adds v to the "delete_at" field.
func (u *PlatformBenefitUpsert) AddDeleteAt(v uint32) *PlatformBenefitUpsert {
	u.Add(platformbenefit.FieldDeleteAt, v)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.PlatformBenefit.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(platformbenefit.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *PlatformBenefitUpsertOne) UpdateNewValues() *PlatformBenefitUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(platformbenefit.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.PlatformBenefit.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *PlatformBenefitUpsertOne) Ignore() *PlatformBenefitUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *PlatformBenefitUpsertOne) DoNothing() *PlatformBenefitUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the PlatformBenefitCreate.OnConflict
// documentation for more info.
func (u *PlatformBenefitUpsertOne) Update(set func(*PlatformBenefitUpsert)) *PlatformBenefitUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&PlatformBenefitUpsert{UpdateSet: update})
	}))
	return u
}

// SetGoodID sets the "good_id" field.
func (u *PlatformBenefitUpsertOne) SetGoodID(v uuid.UUID) *PlatformBenefitUpsertOne {
	return u.Update(func(s *PlatformBenefitUpsert) {
		s.SetGoodID(v)
	})
}

// UpdateGoodID sets the "good_id" field to the value that was provided on create.
func (u *PlatformBenefitUpsertOne) UpdateGoodID() *PlatformBenefitUpsertOne {
	return u.Update(func(s *PlatformBenefitUpsert) {
		s.UpdateGoodID()
	})
}

// SetBenefitAccountID sets the "benefit_account_id" field.
func (u *PlatformBenefitUpsertOne) SetBenefitAccountID(v uuid.UUID) *PlatformBenefitUpsertOne {
	return u.Update(func(s *PlatformBenefitUpsert) {
		s.SetBenefitAccountID(v)
	})
}

// UpdateBenefitAccountID sets the "benefit_account_id" field to the value that was provided on create.
func (u *PlatformBenefitUpsertOne) UpdateBenefitAccountID() *PlatformBenefitUpsertOne {
	return u.Update(func(s *PlatformBenefitUpsert) {
		s.UpdateBenefitAccountID()
	})
}

// SetAmount sets the "amount" field.
func (u *PlatformBenefitUpsertOne) SetAmount(v uint64) *PlatformBenefitUpsertOne {
	return u.Update(func(s *PlatformBenefitUpsert) {
		s.SetAmount(v)
	})
}

// AddAmount adds v to the "amount" field.
func (u *PlatformBenefitUpsertOne) AddAmount(v uint64) *PlatformBenefitUpsertOne {
	return u.Update(func(s *PlatformBenefitUpsert) {
		s.AddAmount(v)
	})
}

// UpdateAmount sets the "amount" field to the value that was provided on create.
func (u *PlatformBenefitUpsertOne) UpdateAmount() *PlatformBenefitUpsertOne {
	return u.Update(func(s *PlatformBenefitUpsert) {
		s.UpdateAmount()
	})
}

// SetLastBenefitTimestamp sets the "last_benefit_timestamp" field.
func (u *PlatformBenefitUpsertOne) SetLastBenefitTimestamp(v uint32) *PlatformBenefitUpsertOne {
	return u.Update(func(s *PlatformBenefitUpsert) {
		s.SetLastBenefitTimestamp(v)
	})
}

// AddLastBenefitTimestamp adds v to the "last_benefit_timestamp" field.
func (u *PlatformBenefitUpsertOne) AddLastBenefitTimestamp(v uint32) *PlatformBenefitUpsertOne {
	return u.Update(func(s *PlatformBenefitUpsert) {
		s.AddLastBenefitTimestamp(v)
	})
}

// UpdateLastBenefitTimestamp sets the "last_benefit_timestamp" field to the value that was provided on create.
func (u *PlatformBenefitUpsertOne) UpdateLastBenefitTimestamp() *PlatformBenefitUpsertOne {
	return u.Update(func(s *PlatformBenefitUpsert) {
		s.UpdateLastBenefitTimestamp()
	})
}

// SetChainTransactionID sets the "chain_transaction_id" field.
func (u *PlatformBenefitUpsertOne) SetChainTransactionID(v string) *PlatformBenefitUpsertOne {
	return u.Update(func(s *PlatformBenefitUpsert) {
		s.SetChainTransactionID(v)
	})
}

// UpdateChainTransactionID sets the "chain_transaction_id" field to the value that was provided on create.
func (u *PlatformBenefitUpsertOne) UpdateChainTransactionID() *PlatformBenefitUpsertOne {
	return u.Update(func(s *PlatformBenefitUpsert) {
		s.UpdateChainTransactionID()
	})
}

// SetCreateAt sets the "create_at" field.
func (u *PlatformBenefitUpsertOne) SetCreateAt(v uint32) *PlatformBenefitUpsertOne {
	return u.Update(func(s *PlatformBenefitUpsert) {
		s.SetCreateAt(v)
	})
}

// AddCreateAt adds v to the "create_at" field.
func (u *PlatformBenefitUpsertOne) AddCreateAt(v uint32) *PlatformBenefitUpsertOne {
	return u.Update(func(s *PlatformBenefitUpsert) {
		s.AddCreateAt(v)
	})
}

// UpdateCreateAt sets the "create_at" field to the value that was provided on create.
func (u *PlatformBenefitUpsertOne) UpdateCreateAt() *PlatformBenefitUpsertOne {
	return u.Update(func(s *PlatformBenefitUpsert) {
		s.UpdateCreateAt()
	})
}

// SetUpdateAt sets the "update_at" field.
func (u *PlatformBenefitUpsertOne) SetUpdateAt(v uint32) *PlatformBenefitUpsertOne {
	return u.Update(func(s *PlatformBenefitUpsert) {
		s.SetUpdateAt(v)
	})
}

// AddUpdateAt adds v to the "update_at" field.
func (u *PlatformBenefitUpsertOne) AddUpdateAt(v uint32) *PlatformBenefitUpsertOne {
	return u.Update(func(s *PlatformBenefitUpsert) {
		s.AddUpdateAt(v)
	})
}

// UpdateUpdateAt sets the "update_at" field to the value that was provided on create.
func (u *PlatformBenefitUpsertOne) UpdateUpdateAt() *PlatformBenefitUpsertOne {
	return u.Update(func(s *PlatformBenefitUpsert) {
		s.UpdateUpdateAt()
	})
}

// SetDeleteAt sets the "delete_at" field.
func (u *PlatformBenefitUpsertOne) SetDeleteAt(v uint32) *PlatformBenefitUpsertOne {
	return u.Update(func(s *PlatformBenefitUpsert) {
		s.SetDeleteAt(v)
	})
}

// AddDeleteAt adds v to the "delete_at" field.
func (u *PlatformBenefitUpsertOne) AddDeleteAt(v uint32) *PlatformBenefitUpsertOne {
	return u.Update(func(s *PlatformBenefitUpsert) {
		s.AddDeleteAt(v)
	})
}

// UpdateDeleteAt sets the "delete_at" field to the value that was provided on create.
func (u *PlatformBenefitUpsertOne) UpdateDeleteAt() *PlatformBenefitUpsertOne {
	return u.Update(func(s *PlatformBenefitUpsert) {
		s.UpdateDeleteAt()
	})
}

// Exec executes the query.
func (u *PlatformBenefitUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for PlatformBenefitCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *PlatformBenefitUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *PlatformBenefitUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: PlatformBenefitUpsertOne.ID is not supported by MySQL driver. Use PlatformBenefitUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *PlatformBenefitUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// PlatformBenefitCreateBulk is the builder for creating many PlatformBenefit entities in bulk.
type PlatformBenefitCreateBulk struct {
	config
	builders []*PlatformBenefitCreate
	conflict []sql.ConflictOption
}

// Save creates the PlatformBenefit entities in the database.
func (pbcb *PlatformBenefitCreateBulk) Save(ctx context.Context) ([]*PlatformBenefit, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pbcb.builders))
	nodes := make([]*PlatformBenefit, len(pbcb.builders))
	mutators := make([]Mutator, len(pbcb.builders))
	for i := range pbcb.builders {
		func(i int, root context.Context) {
			builder := pbcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PlatformBenefitMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, pbcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = pbcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pbcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, pbcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pbcb *PlatformBenefitCreateBulk) SaveX(ctx context.Context) []*PlatformBenefit {
	v, err := pbcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pbcb *PlatformBenefitCreateBulk) Exec(ctx context.Context) error {
	_, err := pbcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pbcb *PlatformBenefitCreateBulk) ExecX(ctx context.Context) {
	if err := pbcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.PlatformBenefit.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.PlatformBenefitUpsert) {
//			SetGoodID(v+v).
//		}).
//		Exec(ctx)
//
func (pbcb *PlatformBenefitCreateBulk) OnConflict(opts ...sql.ConflictOption) *PlatformBenefitUpsertBulk {
	pbcb.conflict = opts
	return &PlatformBenefitUpsertBulk{
		create: pbcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.PlatformBenefit.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (pbcb *PlatformBenefitCreateBulk) OnConflictColumns(columns ...string) *PlatformBenefitUpsertBulk {
	pbcb.conflict = append(pbcb.conflict, sql.ConflictColumns(columns...))
	return &PlatformBenefitUpsertBulk{
		create: pbcb,
	}
}

// PlatformBenefitUpsertBulk is the builder for "upsert"-ing
// a bulk of PlatformBenefit nodes.
type PlatformBenefitUpsertBulk struct {
	create *PlatformBenefitCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.PlatformBenefit.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(platformbenefit.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *PlatformBenefitUpsertBulk) UpdateNewValues() *PlatformBenefitUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(platformbenefit.FieldID)
				return
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.PlatformBenefit.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *PlatformBenefitUpsertBulk) Ignore() *PlatformBenefitUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *PlatformBenefitUpsertBulk) DoNothing() *PlatformBenefitUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the PlatformBenefitCreateBulk.OnConflict
// documentation for more info.
func (u *PlatformBenefitUpsertBulk) Update(set func(*PlatformBenefitUpsert)) *PlatformBenefitUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&PlatformBenefitUpsert{UpdateSet: update})
	}))
	return u
}

// SetGoodID sets the "good_id" field.
func (u *PlatformBenefitUpsertBulk) SetGoodID(v uuid.UUID) *PlatformBenefitUpsertBulk {
	return u.Update(func(s *PlatformBenefitUpsert) {
		s.SetGoodID(v)
	})
}

// UpdateGoodID sets the "good_id" field to the value that was provided on create.
func (u *PlatformBenefitUpsertBulk) UpdateGoodID() *PlatformBenefitUpsertBulk {
	return u.Update(func(s *PlatformBenefitUpsert) {
		s.UpdateGoodID()
	})
}

// SetBenefitAccountID sets the "benefit_account_id" field.
func (u *PlatformBenefitUpsertBulk) SetBenefitAccountID(v uuid.UUID) *PlatformBenefitUpsertBulk {
	return u.Update(func(s *PlatformBenefitUpsert) {
		s.SetBenefitAccountID(v)
	})
}

// UpdateBenefitAccountID sets the "benefit_account_id" field to the value that was provided on create.
func (u *PlatformBenefitUpsertBulk) UpdateBenefitAccountID() *PlatformBenefitUpsertBulk {
	return u.Update(func(s *PlatformBenefitUpsert) {
		s.UpdateBenefitAccountID()
	})
}

// SetAmount sets the "amount" field.
func (u *PlatformBenefitUpsertBulk) SetAmount(v uint64) *PlatformBenefitUpsertBulk {
	return u.Update(func(s *PlatformBenefitUpsert) {
		s.SetAmount(v)
	})
}

// AddAmount adds v to the "amount" field.
func (u *PlatformBenefitUpsertBulk) AddAmount(v uint64) *PlatformBenefitUpsertBulk {
	return u.Update(func(s *PlatformBenefitUpsert) {
		s.AddAmount(v)
	})
}

// UpdateAmount sets the "amount" field to the value that was provided on create.
func (u *PlatformBenefitUpsertBulk) UpdateAmount() *PlatformBenefitUpsertBulk {
	return u.Update(func(s *PlatformBenefitUpsert) {
		s.UpdateAmount()
	})
}

// SetLastBenefitTimestamp sets the "last_benefit_timestamp" field.
func (u *PlatformBenefitUpsertBulk) SetLastBenefitTimestamp(v uint32) *PlatformBenefitUpsertBulk {
	return u.Update(func(s *PlatformBenefitUpsert) {
		s.SetLastBenefitTimestamp(v)
	})
}

// AddLastBenefitTimestamp adds v to the "last_benefit_timestamp" field.
func (u *PlatformBenefitUpsertBulk) AddLastBenefitTimestamp(v uint32) *PlatformBenefitUpsertBulk {
	return u.Update(func(s *PlatformBenefitUpsert) {
		s.AddLastBenefitTimestamp(v)
	})
}

// UpdateLastBenefitTimestamp sets the "last_benefit_timestamp" field to the value that was provided on create.
func (u *PlatformBenefitUpsertBulk) UpdateLastBenefitTimestamp() *PlatformBenefitUpsertBulk {
	return u.Update(func(s *PlatformBenefitUpsert) {
		s.UpdateLastBenefitTimestamp()
	})
}

// SetChainTransactionID sets the "chain_transaction_id" field.
func (u *PlatformBenefitUpsertBulk) SetChainTransactionID(v string) *PlatformBenefitUpsertBulk {
	return u.Update(func(s *PlatformBenefitUpsert) {
		s.SetChainTransactionID(v)
	})
}

// UpdateChainTransactionID sets the "chain_transaction_id" field to the value that was provided on create.
func (u *PlatformBenefitUpsertBulk) UpdateChainTransactionID() *PlatformBenefitUpsertBulk {
	return u.Update(func(s *PlatformBenefitUpsert) {
		s.UpdateChainTransactionID()
	})
}

// SetCreateAt sets the "create_at" field.
func (u *PlatformBenefitUpsertBulk) SetCreateAt(v uint32) *PlatformBenefitUpsertBulk {
	return u.Update(func(s *PlatformBenefitUpsert) {
		s.SetCreateAt(v)
	})
}

// AddCreateAt adds v to the "create_at" field.
func (u *PlatformBenefitUpsertBulk) AddCreateAt(v uint32) *PlatformBenefitUpsertBulk {
	return u.Update(func(s *PlatformBenefitUpsert) {
		s.AddCreateAt(v)
	})
}

// UpdateCreateAt sets the "create_at" field to the value that was provided on create.
func (u *PlatformBenefitUpsertBulk) UpdateCreateAt() *PlatformBenefitUpsertBulk {
	return u.Update(func(s *PlatformBenefitUpsert) {
		s.UpdateCreateAt()
	})
}

// SetUpdateAt sets the "update_at" field.
func (u *PlatformBenefitUpsertBulk) SetUpdateAt(v uint32) *PlatformBenefitUpsertBulk {
	return u.Update(func(s *PlatformBenefitUpsert) {
		s.SetUpdateAt(v)
	})
}

// AddUpdateAt adds v to the "update_at" field.
func (u *PlatformBenefitUpsertBulk) AddUpdateAt(v uint32) *PlatformBenefitUpsertBulk {
	return u.Update(func(s *PlatformBenefitUpsert) {
		s.AddUpdateAt(v)
	})
}

// UpdateUpdateAt sets the "update_at" field to the value that was provided on create.
func (u *PlatformBenefitUpsertBulk) UpdateUpdateAt() *PlatformBenefitUpsertBulk {
	return u.Update(func(s *PlatformBenefitUpsert) {
		s.UpdateUpdateAt()
	})
}

// SetDeleteAt sets the "delete_at" field.
func (u *PlatformBenefitUpsertBulk) SetDeleteAt(v uint32) *PlatformBenefitUpsertBulk {
	return u.Update(func(s *PlatformBenefitUpsert) {
		s.SetDeleteAt(v)
	})
}

// AddDeleteAt adds v to the "delete_at" field.
func (u *PlatformBenefitUpsertBulk) AddDeleteAt(v uint32) *PlatformBenefitUpsertBulk {
	return u.Update(func(s *PlatformBenefitUpsert) {
		s.AddDeleteAt(v)
	})
}

// UpdateDeleteAt sets the "delete_at" field to the value that was provided on create.
func (u *PlatformBenefitUpsertBulk) UpdateDeleteAt() *PlatformBenefitUpsertBulk {
	return u.Update(func(s *PlatformBenefitUpsert) {
		s.UpdateDeleteAt()
	})
}

// Exec executes the query.
func (u *PlatformBenefitUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the PlatformBenefitCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for PlatformBenefitCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *PlatformBenefitUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
