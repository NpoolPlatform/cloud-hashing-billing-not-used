// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/db/ent/userwithdrawitem"
	"github.com/google/uuid"
)

// UserWithdrawItemCreate is the builder for creating a UserWithdrawItem entity.
type UserWithdrawItemCreate struct {
	config
	mutation *UserWithdrawItemMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetAppID sets the "app_id" field.
func (uwic *UserWithdrawItemCreate) SetAppID(u uuid.UUID) *UserWithdrawItemCreate {
	uwic.mutation.SetAppID(u)
	return uwic
}

// SetUserID sets the "user_id" field.
func (uwic *UserWithdrawItemCreate) SetUserID(u uuid.UUID) *UserWithdrawItemCreate {
	uwic.mutation.SetUserID(u)
	return uwic
}

// SetCoinTypeID sets the "coin_type_id" field.
func (uwic *UserWithdrawItemCreate) SetCoinTypeID(u uuid.UUID) *UserWithdrawItemCreate {
	uwic.mutation.SetCoinTypeID(u)
	return uwic
}

// SetWithdrawToAccountID sets the "withdraw_to_account_id" field.
func (uwic *UserWithdrawItemCreate) SetWithdrawToAccountID(u uuid.UUID) *UserWithdrawItemCreate {
	uwic.mutation.SetWithdrawToAccountID(u)
	return uwic
}

// SetAmount sets the "amount" field.
func (uwic *UserWithdrawItemCreate) SetAmount(u uint64) *UserWithdrawItemCreate {
	uwic.mutation.SetAmount(u)
	return uwic
}

// SetPlatformTransactionID sets the "platform_transaction_id" field.
func (uwic *UserWithdrawItemCreate) SetPlatformTransactionID(u uuid.UUID) *UserWithdrawItemCreate {
	uwic.mutation.SetPlatformTransactionID(u)
	return uwic
}

// SetCreateAt sets the "create_at" field.
func (uwic *UserWithdrawItemCreate) SetCreateAt(u uint32) *UserWithdrawItemCreate {
	uwic.mutation.SetCreateAt(u)
	return uwic
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (uwic *UserWithdrawItemCreate) SetNillableCreateAt(u *uint32) *UserWithdrawItemCreate {
	if u != nil {
		uwic.SetCreateAt(*u)
	}
	return uwic
}

// SetUpdateAt sets the "update_at" field.
func (uwic *UserWithdrawItemCreate) SetUpdateAt(u uint32) *UserWithdrawItemCreate {
	uwic.mutation.SetUpdateAt(u)
	return uwic
}

// SetNillableUpdateAt sets the "update_at" field if the given value is not nil.
func (uwic *UserWithdrawItemCreate) SetNillableUpdateAt(u *uint32) *UserWithdrawItemCreate {
	if u != nil {
		uwic.SetUpdateAt(*u)
	}
	return uwic
}

// SetDeleteAt sets the "delete_at" field.
func (uwic *UserWithdrawItemCreate) SetDeleteAt(u uint32) *UserWithdrawItemCreate {
	uwic.mutation.SetDeleteAt(u)
	return uwic
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (uwic *UserWithdrawItemCreate) SetNillableDeleteAt(u *uint32) *UserWithdrawItemCreate {
	if u != nil {
		uwic.SetDeleteAt(*u)
	}
	return uwic
}

// SetID sets the "id" field.
func (uwic *UserWithdrawItemCreate) SetID(u uuid.UUID) *UserWithdrawItemCreate {
	uwic.mutation.SetID(u)
	return uwic
}

// SetNillableID sets the "id" field if the given value is not nil.
func (uwic *UserWithdrawItemCreate) SetNillableID(u *uuid.UUID) *UserWithdrawItemCreate {
	if u != nil {
		uwic.SetID(*u)
	}
	return uwic
}

// Mutation returns the UserWithdrawItemMutation object of the builder.
func (uwic *UserWithdrawItemCreate) Mutation() *UserWithdrawItemMutation {
	return uwic.mutation
}

// Save creates the UserWithdrawItem in the database.
func (uwic *UserWithdrawItemCreate) Save(ctx context.Context) (*UserWithdrawItem, error) {
	var (
		err  error
		node *UserWithdrawItem
	)
	uwic.defaults()
	if len(uwic.hooks) == 0 {
		if err = uwic.check(); err != nil {
			return nil, err
		}
		node, err = uwic.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserWithdrawItemMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = uwic.check(); err != nil {
				return nil, err
			}
			uwic.mutation = mutation
			if node, err = uwic.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(uwic.hooks) - 1; i >= 0; i-- {
			if uwic.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = uwic.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uwic.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (uwic *UserWithdrawItemCreate) SaveX(ctx context.Context) *UserWithdrawItem {
	v, err := uwic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uwic *UserWithdrawItemCreate) Exec(ctx context.Context) error {
	_, err := uwic.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uwic *UserWithdrawItemCreate) ExecX(ctx context.Context) {
	if err := uwic.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uwic *UserWithdrawItemCreate) defaults() {
	if _, ok := uwic.mutation.CreateAt(); !ok {
		v := userwithdrawitem.DefaultCreateAt()
		uwic.mutation.SetCreateAt(v)
	}
	if _, ok := uwic.mutation.UpdateAt(); !ok {
		v := userwithdrawitem.DefaultUpdateAt()
		uwic.mutation.SetUpdateAt(v)
	}
	if _, ok := uwic.mutation.DeleteAt(); !ok {
		v := userwithdrawitem.DefaultDeleteAt()
		uwic.mutation.SetDeleteAt(v)
	}
	if _, ok := uwic.mutation.ID(); !ok {
		v := userwithdrawitem.DefaultID()
		uwic.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uwic *UserWithdrawItemCreate) check() error {
	if _, ok := uwic.mutation.AppID(); !ok {
		return &ValidationError{Name: "app_id", err: errors.New(`ent: missing required field "UserWithdrawItem.app_id"`)}
	}
	if _, ok := uwic.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "UserWithdrawItem.user_id"`)}
	}
	if _, ok := uwic.mutation.CoinTypeID(); !ok {
		return &ValidationError{Name: "coin_type_id", err: errors.New(`ent: missing required field "UserWithdrawItem.coin_type_id"`)}
	}
	if _, ok := uwic.mutation.WithdrawToAccountID(); !ok {
		return &ValidationError{Name: "withdraw_to_account_id", err: errors.New(`ent: missing required field "UserWithdrawItem.withdraw_to_account_id"`)}
	}
	if _, ok := uwic.mutation.Amount(); !ok {
		return &ValidationError{Name: "amount", err: errors.New(`ent: missing required field "UserWithdrawItem.amount"`)}
	}
	if _, ok := uwic.mutation.PlatformTransactionID(); !ok {
		return &ValidationError{Name: "platform_transaction_id", err: errors.New(`ent: missing required field "UserWithdrawItem.platform_transaction_id"`)}
	}
	if _, ok := uwic.mutation.CreateAt(); !ok {
		return &ValidationError{Name: "create_at", err: errors.New(`ent: missing required field "UserWithdrawItem.create_at"`)}
	}
	if _, ok := uwic.mutation.UpdateAt(); !ok {
		return &ValidationError{Name: "update_at", err: errors.New(`ent: missing required field "UserWithdrawItem.update_at"`)}
	}
	if _, ok := uwic.mutation.DeleteAt(); !ok {
		return &ValidationError{Name: "delete_at", err: errors.New(`ent: missing required field "UserWithdrawItem.delete_at"`)}
	}
	return nil
}

func (uwic *UserWithdrawItemCreate) sqlSave(ctx context.Context) (*UserWithdrawItem, error) {
	_node, _spec := uwic.createSpec()
	if err := sqlgraph.CreateNode(ctx, uwic.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
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

func (uwic *UserWithdrawItemCreate) createSpec() (*UserWithdrawItem, *sqlgraph.CreateSpec) {
	var (
		_node = &UserWithdrawItem{config: uwic.config}
		_spec = &sqlgraph.CreateSpec{
			Table: userwithdrawitem.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: userwithdrawitem.FieldID,
			},
		}
	)
	_spec.OnConflict = uwic.conflict
	if id, ok := uwic.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := uwic.mutation.AppID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: userwithdrawitem.FieldAppID,
		})
		_node.AppID = value
	}
	if value, ok := uwic.mutation.UserID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: userwithdrawitem.FieldUserID,
		})
		_node.UserID = value
	}
	if value, ok := uwic.mutation.CoinTypeID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: userwithdrawitem.FieldCoinTypeID,
		})
		_node.CoinTypeID = value
	}
	if value, ok := uwic.mutation.WithdrawToAccountID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: userwithdrawitem.FieldWithdrawToAccountID,
		})
		_node.WithdrawToAccountID = value
	}
	if value, ok := uwic.mutation.Amount(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: userwithdrawitem.FieldAmount,
		})
		_node.Amount = value
	}
	if value, ok := uwic.mutation.PlatformTransactionID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: userwithdrawitem.FieldPlatformTransactionID,
		})
		_node.PlatformTransactionID = value
	}
	if value, ok := uwic.mutation.CreateAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: userwithdrawitem.FieldCreateAt,
		})
		_node.CreateAt = value
	}
	if value, ok := uwic.mutation.UpdateAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: userwithdrawitem.FieldUpdateAt,
		})
		_node.UpdateAt = value
	}
	if value, ok := uwic.mutation.DeleteAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: userwithdrawitem.FieldDeleteAt,
		})
		_node.DeleteAt = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.UserWithdrawItem.Create().
//		SetAppID(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.UserWithdrawItemUpsert) {
//			SetAppID(v+v).
//		}).
//		Exec(ctx)
//
func (uwic *UserWithdrawItemCreate) OnConflict(opts ...sql.ConflictOption) *UserWithdrawItemUpsertOne {
	uwic.conflict = opts
	return &UserWithdrawItemUpsertOne{
		create: uwic,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.UserWithdrawItem.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (uwic *UserWithdrawItemCreate) OnConflictColumns(columns ...string) *UserWithdrawItemUpsertOne {
	uwic.conflict = append(uwic.conflict, sql.ConflictColumns(columns...))
	return &UserWithdrawItemUpsertOne{
		create: uwic,
	}
}

type (
	// UserWithdrawItemUpsertOne is the builder for "upsert"-ing
	//  one UserWithdrawItem node.
	UserWithdrawItemUpsertOne struct {
		create *UserWithdrawItemCreate
	}

	// UserWithdrawItemUpsert is the "OnConflict" setter.
	UserWithdrawItemUpsert struct {
		*sql.UpdateSet
	}
)

// SetAppID sets the "app_id" field.
func (u *UserWithdrawItemUpsert) SetAppID(v uuid.UUID) *UserWithdrawItemUpsert {
	u.Set(userwithdrawitem.FieldAppID, v)
	return u
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *UserWithdrawItemUpsert) UpdateAppID() *UserWithdrawItemUpsert {
	u.SetExcluded(userwithdrawitem.FieldAppID)
	return u
}

// SetUserID sets the "user_id" field.
func (u *UserWithdrawItemUpsert) SetUserID(v uuid.UUID) *UserWithdrawItemUpsert {
	u.Set(userwithdrawitem.FieldUserID, v)
	return u
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *UserWithdrawItemUpsert) UpdateUserID() *UserWithdrawItemUpsert {
	u.SetExcluded(userwithdrawitem.FieldUserID)
	return u
}

// SetCoinTypeID sets the "coin_type_id" field.
func (u *UserWithdrawItemUpsert) SetCoinTypeID(v uuid.UUID) *UserWithdrawItemUpsert {
	u.Set(userwithdrawitem.FieldCoinTypeID, v)
	return u
}

// UpdateCoinTypeID sets the "coin_type_id" field to the value that was provided on create.
func (u *UserWithdrawItemUpsert) UpdateCoinTypeID() *UserWithdrawItemUpsert {
	u.SetExcluded(userwithdrawitem.FieldCoinTypeID)
	return u
}

// SetWithdrawToAccountID sets the "withdraw_to_account_id" field.
func (u *UserWithdrawItemUpsert) SetWithdrawToAccountID(v uuid.UUID) *UserWithdrawItemUpsert {
	u.Set(userwithdrawitem.FieldWithdrawToAccountID, v)
	return u
}

// UpdateWithdrawToAccountID sets the "withdraw_to_account_id" field to the value that was provided on create.
func (u *UserWithdrawItemUpsert) UpdateWithdrawToAccountID() *UserWithdrawItemUpsert {
	u.SetExcluded(userwithdrawitem.FieldWithdrawToAccountID)
	return u
}

// SetAmount sets the "amount" field.
func (u *UserWithdrawItemUpsert) SetAmount(v uint64) *UserWithdrawItemUpsert {
	u.Set(userwithdrawitem.FieldAmount, v)
	return u
}

// UpdateAmount sets the "amount" field to the value that was provided on create.
func (u *UserWithdrawItemUpsert) UpdateAmount() *UserWithdrawItemUpsert {
	u.SetExcluded(userwithdrawitem.FieldAmount)
	return u
}

// AddAmount adds v to the "amount" field.
func (u *UserWithdrawItemUpsert) AddAmount(v uint64) *UserWithdrawItemUpsert {
	u.Add(userwithdrawitem.FieldAmount, v)
	return u
}

// SetPlatformTransactionID sets the "platform_transaction_id" field.
func (u *UserWithdrawItemUpsert) SetPlatformTransactionID(v uuid.UUID) *UserWithdrawItemUpsert {
	u.Set(userwithdrawitem.FieldPlatformTransactionID, v)
	return u
}

// UpdatePlatformTransactionID sets the "platform_transaction_id" field to the value that was provided on create.
func (u *UserWithdrawItemUpsert) UpdatePlatformTransactionID() *UserWithdrawItemUpsert {
	u.SetExcluded(userwithdrawitem.FieldPlatformTransactionID)
	return u
}

// SetCreateAt sets the "create_at" field.
func (u *UserWithdrawItemUpsert) SetCreateAt(v uint32) *UserWithdrawItemUpsert {
	u.Set(userwithdrawitem.FieldCreateAt, v)
	return u
}

// UpdateCreateAt sets the "create_at" field to the value that was provided on create.
func (u *UserWithdrawItemUpsert) UpdateCreateAt() *UserWithdrawItemUpsert {
	u.SetExcluded(userwithdrawitem.FieldCreateAt)
	return u
}

// AddCreateAt adds v to the "create_at" field.
func (u *UserWithdrawItemUpsert) AddCreateAt(v uint32) *UserWithdrawItemUpsert {
	u.Add(userwithdrawitem.FieldCreateAt, v)
	return u
}

// SetUpdateAt sets the "update_at" field.
func (u *UserWithdrawItemUpsert) SetUpdateAt(v uint32) *UserWithdrawItemUpsert {
	u.Set(userwithdrawitem.FieldUpdateAt, v)
	return u
}

// UpdateUpdateAt sets the "update_at" field to the value that was provided on create.
func (u *UserWithdrawItemUpsert) UpdateUpdateAt() *UserWithdrawItemUpsert {
	u.SetExcluded(userwithdrawitem.FieldUpdateAt)
	return u
}

// AddUpdateAt adds v to the "update_at" field.
func (u *UserWithdrawItemUpsert) AddUpdateAt(v uint32) *UserWithdrawItemUpsert {
	u.Add(userwithdrawitem.FieldUpdateAt, v)
	return u
}

// SetDeleteAt sets the "delete_at" field.
func (u *UserWithdrawItemUpsert) SetDeleteAt(v uint32) *UserWithdrawItemUpsert {
	u.Set(userwithdrawitem.FieldDeleteAt, v)
	return u
}

// UpdateDeleteAt sets the "delete_at" field to the value that was provided on create.
func (u *UserWithdrawItemUpsert) UpdateDeleteAt() *UserWithdrawItemUpsert {
	u.SetExcluded(userwithdrawitem.FieldDeleteAt)
	return u
}

// AddDeleteAt adds v to the "delete_at" field.
func (u *UserWithdrawItemUpsert) AddDeleteAt(v uint32) *UserWithdrawItemUpsert {
	u.Add(userwithdrawitem.FieldDeleteAt, v)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.UserWithdrawItem.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(userwithdrawitem.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *UserWithdrawItemUpsertOne) UpdateNewValues() *UserWithdrawItemUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(userwithdrawitem.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.UserWithdrawItem.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *UserWithdrawItemUpsertOne) Ignore() *UserWithdrawItemUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *UserWithdrawItemUpsertOne) DoNothing() *UserWithdrawItemUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the UserWithdrawItemCreate.OnConflict
// documentation for more info.
func (u *UserWithdrawItemUpsertOne) Update(set func(*UserWithdrawItemUpsert)) *UserWithdrawItemUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&UserWithdrawItemUpsert{UpdateSet: update})
	}))
	return u
}

// SetAppID sets the "app_id" field.
func (u *UserWithdrawItemUpsertOne) SetAppID(v uuid.UUID) *UserWithdrawItemUpsertOne {
	return u.Update(func(s *UserWithdrawItemUpsert) {
		s.SetAppID(v)
	})
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *UserWithdrawItemUpsertOne) UpdateAppID() *UserWithdrawItemUpsertOne {
	return u.Update(func(s *UserWithdrawItemUpsert) {
		s.UpdateAppID()
	})
}

// SetUserID sets the "user_id" field.
func (u *UserWithdrawItemUpsertOne) SetUserID(v uuid.UUID) *UserWithdrawItemUpsertOne {
	return u.Update(func(s *UserWithdrawItemUpsert) {
		s.SetUserID(v)
	})
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *UserWithdrawItemUpsertOne) UpdateUserID() *UserWithdrawItemUpsertOne {
	return u.Update(func(s *UserWithdrawItemUpsert) {
		s.UpdateUserID()
	})
}

// SetCoinTypeID sets the "coin_type_id" field.
func (u *UserWithdrawItemUpsertOne) SetCoinTypeID(v uuid.UUID) *UserWithdrawItemUpsertOne {
	return u.Update(func(s *UserWithdrawItemUpsert) {
		s.SetCoinTypeID(v)
	})
}

// UpdateCoinTypeID sets the "coin_type_id" field to the value that was provided on create.
func (u *UserWithdrawItemUpsertOne) UpdateCoinTypeID() *UserWithdrawItemUpsertOne {
	return u.Update(func(s *UserWithdrawItemUpsert) {
		s.UpdateCoinTypeID()
	})
}

// SetWithdrawToAccountID sets the "withdraw_to_account_id" field.
func (u *UserWithdrawItemUpsertOne) SetWithdrawToAccountID(v uuid.UUID) *UserWithdrawItemUpsertOne {
	return u.Update(func(s *UserWithdrawItemUpsert) {
		s.SetWithdrawToAccountID(v)
	})
}

// UpdateWithdrawToAccountID sets the "withdraw_to_account_id" field to the value that was provided on create.
func (u *UserWithdrawItemUpsertOne) UpdateWithdrawToAccountID() *UserWithdrawItemUpsertOne {
	return u.Update(func(s *UserWithdrawItemUpsert) {
		s.UpdateWithdrawToAccountID()
	})
}

// SetAmount sets the "amount" field.
func (u *UserWithdrawItemUpsertOne) SetAmount(v uint64) *UserWithdrawItemUpsertOne {
	return u.Update(func(s *UserWithdrawItemUpsert) {
		s.SetAmount(v)
	})
}

// AddAmount adds v to the "amount" field.
func (u *UserWithdrawItemUpsertOne) AddAmount(v uint64) *UserWithdrawItemUpsertOne {
	return u.Update(func(s *UserWithdrawItemUpsert) {
		s.AddAmount(v)
	})
}

// UpdateAmount sets the "amount" field to the value that was provided on create.
func (u *UserWithdrawItemUpsertOne) UpdateAmount() *UserWithdrawItemUpsertOne {
	return u.Update(func(s *UserWithdrawItemUpsert) {
		s.UpdateAmount()
	})
}

// SetPlatformTransactionID sets the "platform_transaction_id" field.
func (u *UserWithdrawItemUpsertOne) SetPlatformTransactionID(v uuid.UUID) *UserWithdrawItemUpsertOne {
	return u.Update(func(s *UserWithdrawItemUpsert) {
		s.SetPlatformTransactionID(v)
	})
}

// UpdatePlatformTransactionID sets the "platform_transaction_id" field to the value that was provided on create.
func (u *UserWithdrawItemUpsertOne) UpdatePlatformTransactionID() *UserWithdrawItemUpsertOne {
	return u.Update(func(s *UserWithdrawItemUpsert) {
		s.UpdatePlatformTransactionID()
	})
}

// SetCreateAt sets the "create_at" field.
func (u *UserWithdrawItemUpsertOne) SetCreateAt(v uint32) *UserWithdrawItemUpsertOne {
	return u.Update(func(s *UserWithdrawItemUpsert) {
		s.SetCreateAt(v)
	})
}

// AddCreateAt adds v to the "create_at" field.
func (u *UserWithdrawItemUpsertOne) AddCreateAt(v uint32) *UserWithdrawItemUpsertOne {
	return u.Update(func(s *UserWithdrawItemUpsert) {
		s.AddCreateAt(v)
	})
}

// UpdateCreateAt sets the "create_at" field to the value that was provided on create.
func (u *UserWithdrawItemUpsertOne) UpdateCreateAt() *UserWithdrawItemUpsertOne {
	return u.Update(func(s *UserWithdrawItemUpsert) {
		s.UpdateCreateAt()
	})
}

// SetUpdateAt sets the "update_at" field.
func (u *UserWithdrawItemUpsertOne) SetUpdateAt(v uint32) *UserWithdrawItemUpsertOne {
	return u.Update(func(s *UserWithdrawItemUpsert) {
		s.SetUpdateAt(v)
	})
}

// AddUpdateAt adds v to the "update_at" field.
func (u *UserWithdrawItemUpsertOne) AddUpdateAt(v uint32) *UserWithdrawItemUpsertOne {
	return u.Update(func(s *UserWithdrawItemUpsert) {
		s.AddUpdateAt(v)
	})
}

// UpdateUpdateAt sets the "update_at" field to the value that was provided on create.
func (u *UserWithdrawItemUpsertOne) UpdateUpdateAt() *UserWithdrawItemUpsertOne {
	return u.Update(func(s *UserWithdrawItemUpsert) {
		s.UpdateUpdateAt()
	})
}

// SetDeleteAt sets the "delete_at" field.
func (u *UserWithdrawItemUpsertOne) SetDeleteAt(v uint32) *UserWithdrawItemUpsertOne {
	return u.Update(func(s *UserWithdrawItemUpsert) {
		s.SetDeleteAt(v)
	})
}

// AddDeleteAt adds v to the "delete_at" field.
func (u *UserWithdrawItemUpsertOne) AddDeleteAt(v uint32) *UserWithdrawItemUpsertOne {
	return u.Update(func(s *UserWithdrawItemUpsert) {
		s.AddDeleteAt(v)
	})
}

// UpdateDeleteAt sets the "delete_at" field to the value that was provided on create.
func (u *UserWithdrawItemUpsertOne) UpdateDeleteAt() *UserWithdrawItemUpsertOne {
	return u.Update(func(s *UserWithdrawItemUpsert) {
		s.UpdateDeleteAt()
	})
}

// Exec executes the query.
func (u *UserWithdrawItemUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for UserWithdrawItemCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *UserWithdrawItemUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *UserWithdrawItemUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: UserWithdrawItemUpsertOne.ID is not supported by MySQL driver. Use UserWithdrawItemUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *UserWithdrawItemUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// UserWithdrawItemCreateBulk is the builder for creating many UserWithdrawItem entities in bulk.
type UserWithdrawItemCreateBulk struct {
	config
	builders []*UserWithdrawItemCreate
	conflict []sql.ConflictOption
}

// Save creates the UserWithdrawItem entities in the database.
func (uwicb *UserWithdrawItemCreateBulk) Save(ctx context.Context) ([]*UserWithdrawItem, error) {
	specs := make([]*sqlgraph.CreateSpec, len(uwicb.builders))
	nodes := make([]*UserWithdrawItem, len(uwicb.builders))
	mutators := make([]Mutator, len(uwicb.builders))
	for i := range uwicb.builders {
		func(i int, root context.Context) {
			builder := uwicb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserWithdrawItemMutation)
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
					_, err = mutators[i+1].Mutate(root, uwicb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = uwicb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, uwicb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
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
		if _, err := mutators[0].Mutate(ctx, uwicb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (uwicb *UserWithdrawItemCreateBulk) SaveX(ctx context.Context) []*UserWithdrawItem {
	v, err := uwicb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uwicb *UserWithdrawItemCreateBulk) Exec(ctx context.Context) error {
	_, err := uwicb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uwicb *UserWithdrawItemCreateBulk) ExecX(ctx context.Context) {
	if err := uwicb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.UserWithdrawItem.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.UserWithdrawItemUpsert) {
//			SetAppID(v+v).
//		}).
//		Exec(ctx)
//
func (uwicb *UserWithdrawItemCreateBulk) OnConflict(opts ...sql.ConflictOption) *UserWithdrawItemUpsertBulk {
	uwicb.conflict = opts
	return &UserWithdrawItemUpsertBulk{
		create: uwicb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.UserWithdrawItem.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (uwicb *UserWithdrawItemCreateBulk) OnConflictColumns(columns ...string) *UserWithdrawItemUpsertBulk {
	uwicb.conflict = append(uwicb.conflict, sql.ConflictColumns(columns...))
	return &UserWithdrawItemUpsertBulk{
		create: uwicb,
	}
}

// UserWithdrawItemUpsertBulk is the builder for "upsert"-ing
// a bulk of UserWithdrawItem nodes.
type UserWithdrawItemUpsertBulk struct {
	create *UserWithdrawItemCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.UserWithdrawItem.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(userwithdrawitem.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *UserWithdrawItemUpsertBulk) UpdateNewValues() *UserWithdrawItemUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(userwithdrawitem.FieldID)
				return
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.UserWithdrawItem.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *UserWithdrawItemUpsertBulk) Ignore() *UserWithdrawItemUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *UserWithdrawItemUpsertBulk) DoNothing() *UserWithdrawItemUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the UserWithdrawItemCreateBulk.OnConflict
// documentation for more info.
func (u *UserWithdrawItemUpsertBulk) Update(set func(*UserWithdrawItemUpsert)) *UserWithdrawItemUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&UserWithdrawItemUpsert{UpdateSet: update})
	}))
	return u
}

// SetAppID sets the "app_id" field.
func (u *UserWithdrawItemUpsertBulk) SetAppID(v uuid.UUID) *UserWithdrawItemUpsertBulk {
	return u.Update(func(s *UserWithdrawItemUpsert) {
		s.SetAppID(v)
	})
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *UserWithdrawItemUpsertBulk) UpdateAppID() *UserWithdrawItemUpsertBulk {
	return u.Update(func(s *UserWithdrawItemUpsert) {
		s.UpdateAppID()
	})
}

// SetUserID sets the "user_id" field.
func (u *UserWithdrawItemUpsertBulk) SetUserID(v uuid.UUID) *UserWithdrawItemUpsertBulk {
	return u.Update(func(s *UserWithdrawItemUpsert) {
		s.SetUserID(v)
	})
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *UserWithdrawItemUpsertBulk) UpdateUserID() *UserWithdrawItemUpsertBulk {
	return u.Update(func(s *UserWithdrawItemUpsert) {
		s.UpdateUserID()
	})
}

// SetCoinTypeID sets the "coin_type_id" field.
func (u *UserWithdrawItemUpsertBulk) SetCoinTypeID(v uuid.UUID) *UserWithdrawItemUpsertBulk {
	return u.Update(func(s *UserWithdrawItemUpsert) {
		s.SetCoinTypeID(v)
	})
}

// UpdateCoinTypeID sets the "coin_type_id" field to the value that was provided on create.
func (u *UserWithdrawItemUpsertBulk) UpdateCoinTypeID() *UserWithdrawItemUpsertBulk {
	return u.Update(func(s *UserWithdrawItemUpsert) {
		s.UpdateCoinTypeID()
	})
}

// SetWithdrawToAccountID sets the "withdraw_to_account_id" field.
func (u *UserWithdrawItemUpsertBulk) SetWithdrawToAccountID(v uuid.UUID) *UserWithdrawItemUpsertBulk {
	return u.Update(func(s *UserWithdrawItemUpsert) {
		s.SetWithdrawToAccountID(v)
	})
}

// UpdateWithdrawToAccountID sets the "withdraw_to_account_id" field to the value that was provided on create.
func (u *UserWithdrawItemUpsertBulk) UpdateWithdrawToAccountID() *UserWithdrawItemUpsertBulk {
	return u.Update(func(s *UserWithdrawItemUpsert) {
		s.UpdateWithdrawToAccountID()
	})
}

// SetAmount sets the "amount" field.
func (u *UserWithdrawItemUpsertBulk) SetAmount(v uint64) *UserWithdrawItemUpsertBulk {
	return u.Update(func(s *UserWithdrawItemUpsert) {
		s.SetAmount(v)
	})
}

// AddAmount adds v to the "amount" field.
func (u *UserWithdrawItemUpsertBulk) AddAmount(v uint64) *UserWithdrawItemUpsertBulk {
	return u.Update(func(s *UserWithdrawItemUpsert) {
		s.AddAmount(v)
	})
}

// UpdateAmount sets the "amount" field to the value that was provided on create.
func (u *UserWithdrawItemUpsertBulk) UpdateAmount() *UserWithdrawItemUpsertBulk {
	return u.Update(func(s *UserWithdrawItemUpsert) {
		s.UpdateAmount()
	})
}

// SetPlatformTransactionID sets the "platform_transaction_id" field.
func (u *UserWithdrawItemUpsertBulk) SetPlatformTransactionID(v uuid.UUID) *UserWithdrawItemUpsertBulk {
	return u.Update(func(s *UserWithdrawItemUpsert) {
		s.SetPlatformTransactionID(v)
	})
}

// UpdatePlatformTransactionID sets the "platform_transaction_id" field to the value that was provided on create.
func (u *UserWithdrawItemUpsertBulk) UpdatePlatformTransactionID() *UserWithdrawItemUpsertBulk {
	return u.Update(func(s *UserWithdrawItemUpsert) {
		s.UpdatePlatformTransactionID()
	})
}

// SetCreateAt sets the "create_at" field.
func (u *UserWithdrawItemUpsertBulk) SetCreateAt(v uint32) *UserWithdrawItemUpsertBulk {
	return u.Update(func(s *UserWithdrawItemUpsert) {
		s.SetCreateAt(v)
	})
}

// AddCreateAt adds v to the "create_at" field.
func (u *UserWithdrawItemUpsertBulk) AddCreateAt(v uint32) *UserWithdrawItemUpsertBulk {
	return u.Update(func(s *UserWithdrawItemUpsert) {
		s.AddCreateAt(v)
	})
}

// UpdateCreateAt sets the "create_at" field to the value that was provided on create.
func (u *UserWithdrawItemUpsertBulk) UpdateCreateAt() *UserWithdrawItemUpsertBulk {
	return u.Update(func(s *UserWithdrawItemUpsert) {
		s.UpdateCreateAt()
	})
}

// SetUpdateAt sets the "update_at" field.
func (u *UserWithdrawItemUpsertBulk) SetUpdateAt(v uint32) *UserWithdrawItemUpsertBulk {
	return u.Update(func(s *UserWithdrawItemUpsert) {
		s.SetUpdateAt(v)
	})
}

// AddUpdateAt adds v to the "update_at" field.
func (u *UserWithdrawItemUpsertBulk) AddUpdateAt(v uint32) *UserWithdrawItemUpsertBulk {
	return u.Update(func(s *UserWithdrawItemUpsert) {
		s.AddUpdateAt(v)
	})
}

// UpdateUpdateAt sets the "update_at" field to the value that was provided on create.
func (u *UserWithdrawItemUpsertBulk) UpdateUpdateAt() *UserWithdrawItemUpsertBulk {
	return u.Update(func(s *UserWithdrawItemUpsert) {
		s.UpdateUpdateAt()
	})
}

// SetDeleteAt sets the "delete_at" field.
func (u *UserWithdrawItemUpsertBulk) SetDeleteAt(v uint32) *UserWithdrawItemUpsertBulk {
	return u.Update(func(s *UserWithdrawItemUpsert) {
		s.SetDeleteAt(v)
	})
}

// AddDeleteAt adds v to the "delete_at" field.
func (u *UserWithdrawItemUpsertBulk) AddDeleteAt(v uint32) *UserWithdrawItemUpsertBulk {
	return u.Update(func(s *UserWithdrawItemUpsert) {
		s.AddDeleteAt(v)
	})
}

// UpdateDeleteAt sets the "delete_at" field to the value that was provided on create.
func (u *UserWithdrawItemUpsertBulk) UpdateDeleteAt() *UserWithdrawItemUpsertBulk {
	return u.Update(func(s *UserWithdrawItemUpsert) {
		s.UpdateDeleteAt()
	})
}

// Exec executes the query.
func (u *UserWithdrawItemUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the UserWithdrawItemCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for UserWithdrawItemCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *UserWithdrawItemUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}