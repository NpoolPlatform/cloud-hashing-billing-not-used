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
	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/db/ent/userdirectbenefit"
	"github.com/google/uuid"
)

// UserDirectBenefitCreate is the builder for creating a UserDirectBenefit entity.
type UserDirectBenefitCreate struct {
	config
	mutation *UserDirectBenefitMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetAppID sets the "app_id" field.
func (udbc *UserDirectBenefitCreate) SetAppID(u uuid.UUID) *UserDirectBenefitCreate {
	udbc.mutation.SetAppID(u)
	return udbc
}

// SetUserID sets the "user_id" field.
func (udbc *UserDirectBenefitCreate) SetUserID(u uuid.UUID) *UserDirectBenefitCreate {
	udbc.mutation.SetUserID(u)
	return udbc
}

// SetCoinTypeID sets the "coin_type_id" field.
func (udbc *UserDirectBenefitCreate) SetCoinTypeID(u uuid.UUID) *UserDirectBenefitCreate {
	udbc.mutation.SetCoinTypeID(u)
	return udbc
}

// SetAccountID sets the "account_id" field.
func (udbc *UserDirectBenefitCreate) SetAccountID(u uuid.UUID) *UserDirectBenefitCreate {
	udbc.mutation.SetAccountID(u)
	return udbc
}

// SetCreateAt sets the "create_at" field.
func (udbc *UserDirectBenefitCreate) SetCreateAt(u uint32) *UserDirectBenefitCreate {
	udbc.mutation.SetCreateAt(u)
	return udbc
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (udbc *UserDirectBenefitCreate) SetNillableCreateAt(u *uint32) *UserDirectBenefitCreate {
	if u != nil {
		udbc.SetCreateAt(*u)
	}
	return udbc
}

// SetUpdateAt sets the "update_at" field.
func (udbc *UserDirectBenefitCreate) SetUpdateAt(u uint32) *UserDirectBenefitCreate {
	udbc.mutation.SetUpdateAt(u)
	return udbc
}

// SetNillableUpdateAt sets the "update_at" field if the given value is not nil.
func (udbc *UserDirectBenefitCreate) SetNillableUpdateAt(u *uint32) *UserDirectBenefitCreate {
	if u != nil {
		udbc.SetUpdateAt(*u)
	}
	return udbc
}

// SetDeleteAt sets the "delete_at" field.
func (udbc *UserDirectBenefitCreate) SetDeleteAt(u uint32) *UserDirectBenefitCreate {
	udbc.mutation.SetDeleteAt(u)
	return udbc
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (udbc *UserDirectBenefitCreate) SetNillableDeleteAt(u *uint32) *UserDirectBenefitCreate {
	if u != nil {
		udbc.SetDeleteAt(*u)
	}
	return udbc
}

// SetID sets the "id" field.
func (udbc *UserDirectBenefitCreate) SetID(u uuid.UUID) *UserDirectBenefitCreate {
	udbc.mutation.SetID(u)
	return udbc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (udbc *UserDirectBenefitCreate) SetNillableID(u *uuid.UUID) *UserDirectBenefitCreate {
	if u != nil {
		udbc.SetID(*u)
	}
	return udbc
}

// Mutation returns the UserDirectBenefitMutation object of the builder.
func (udbc *UserDirectBenefitCreate) Mutation() *UserDirectBenefitMutation {
	return udbc.mutation
}

// Save creates the UserDirectBenefit in the database.
func (udbc *UserDirectBenefitCreate) Save(ctx context.Context) (*UserDirectBenefit, error) {
	var (
		err  error
		node *UserDirectBenefit
	)
	udbc.defaults()
	if len(udbc.hooks) == 0 {
		if err = udbc.check(); err != nil {
			return nil, err
		}
		node, err = udbc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserDirectBenefitMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = udbc.check(); err != nil {
				return nil, err
			}
			udbc.mutation = mutation
			if node, err = udbc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(udbc.hooks) - 1; i >= 0; i-- {
			if udbc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = udbc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, udbc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (udbc *UserDirectBenefitCreate) SaveX(ctx context.Context) *UserDirectBenefit {
	v, err := udbc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (udbc *UserDirectBenefitCreate) Exec(ctx context.Context) error {
	_, err := udbc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (udbc *UserDirectBenefitCreate) ExecX(ctx context.Context) {
	if err := udbc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (udbc *UserDirectBenefitCreate) defaults() {
	if _, ok := udbc.mutation.CreateAt(); !ok {
		v := userdirectbenefit.DefaultCreateAt()
		udbc.mutation.SetCreateAt(v)
	}
	if _, ok := udbc.mutation.UpdateAt(); !ok {
		v := userdirectbenefit.DefaultUpdateAt()
		udbc.mutation.SetUpdateAt(v)
	}
	if _, ok := udbc.mutation.DeleteAt(); !ok {
		v := userdirectbenefit.DefaultDeleteAt()
		udbc.mutation.SetDeleteAt(v)
	}
	if _, ok := udbc.mutation.ID(); !ok {
		v := userdirectbenefit.DefaultID()
		udbc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (udbc *UserDirectBenefitCreate) check() error {
	if _, ok := udbc.mutation.AppID(); !ok {
		return &ValidationError{Name: "app_id", err: errors.New(`ent: missing required field "UserDirectBenefit.app_id"`)}
	}
	if _, ok := udbc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "UserDirectBenefit.user_id"`)}
	}
	if _, ok := udbc.mutation.CoinTypeID(); !ok {
		return &ValidationError{Name: "coin_type_id", err: errors.New(`ent: missing required field "UserDirectBenefit.coin_type_id"`)}
	}
	if _, ok := udbc.mutation.AccountID(); !ok {
		return &ValidationError{Name: "account_id", err: errors.New(`ent: missing required field "UserDirectBenefit.account_id"`)}
	}
	if _, ok := udbc.mutation.CreateAt(); !ok {
		return &ValidationError{Name: "create_at", err: errors.New(`ent: missing required field "UserDirectBenefit.create_at"`)}
	}
	if _, ok := udbc.mutation.UpdateAt(); !ok {
		return &ValidationError{Name: "update_at", err: errors.New(`ent: missing required field "UserDirectBenefit.update_at"`)}
	}
	if _, ok := udbc.mutation.DeleteAt(); !ok {
		return &ValidationError{Name: "delete_at", err: errors.New(`ent: missing required field "UserDirectBenefit.delete_at"`)}
	}
	return nil
}

func (udbc *UserDirectBenefitCreate) sqlSave(ctx context.Context) (*UserDirectBenefit, error) {
	_node, _spec := udbc.createSpec()
	if err := sqlgraph.CreateNode(ctx, udbc.driver, _spec); err != nil {
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

func (udbc *UserDirectBenefitCreate) createSpec() (*UserDirectBenefit, *sqlgraph.CreateSpec) {
	var (
		_node = &UserDirectBenefit{config: udbc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: userdirectbenefit.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: userdirectbenefit.FieldID,
			},
		}
	)
	_spec.OnConflict = udbc.conflict
	if id, ok := udbc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := udbc.mutation.AppID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: userdirectbenefit.FieldAppID,
		})
		_node.AppID = value
	}
	if value, ok := udbc.mutation.UserID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: userdirectbenefit.FieldUserID,
		})
		_node.UserID = value
	}
	if value, ok := udbc.mutation.CoinTypeID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: userdirectbenefit.FieldCoinTypeID,
		})
		_node.CoinTypeID = value
	}
	if value, ok := udbc.mutation.AccountID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: userdirectbenefit.FieldAccountID,
		})
		_node.AccountID = value
	}
	if value, ok := udbc.mutation.CreateAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: userdirectbenefit.FieldCreateAt,
		})
		_node.CreateAt = value
	}
	if value, ok := udbc.mutation.UpdateAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: userdirectbenefit.FieldUpdateAt,
		})
		_node.UpdateAt = value
	}
	if value, ok := udbc.mutation.DeleteAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: userdirectbenefit.FieldDeleteAt,
		})
		_node.DeleteAt = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.UserDirectBenefit.Create().
//		SetAppID(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.UserDirectBenefitUpsert) {
//			SetAppID(v+v).
//		}).
//		Exec(ctx)
//
func (udbc *UserDirectBenefitCreate) OnConflict(opts ...sql.ConflictOption) *UserDirectBenefitUpsertOne {
	udbc.conflict = opts
	return &UserDirectBenefitUpsertOne{
		create: udbc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.UserDirectBenefit.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (udbc *UserDirectBenefitCreate) OnConflictColumns(columns ...string) *UserDirectBenefitUpsertOne {
	udbc.conflict = append(udbc.conflict, sql.ConflictColumns(columns...))
	return &UserDirectBenefitUpsertOne{
		create: udbc,
	}
}

type (
	// UserDirectBenefitUpsertOne is the builder for "upsert"-ing
	//  one UserDirectBenefit node.
	UserDirectBenefitUpsertOne struct {
		create *UserDirectBenefitCreate
	}

	// UserDirectBenefitUpsert is the "OnConflict" setter.
	UserDirectBenefitUpsert struct {
		*sql.UpdateSet
	}
)

// SetAppID sets the "app_id" field.
func (u *UserDirectBenefitUpsert) SetAppID(v uuid.UUID) *UserDirectBenefitUpsert {
	u.Set(userdirectbenefit.FieldAppID, v)
	return u
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *UserDirectBenefitUpsert) UpdateAppID() *UserDirectBenefitUpsert {
	u.SetExcluded(userdirectbenefit.FieldAppID)
	return u
}

// SetUserID sets the "user_id" field.
func (u *UserDirectBenefitUpsert) SetUserID(v uuid.UUID) *UserDirectBenefitUpsert {
	u.Set(userdirectbenefit.FieldUserID, v)
	return u
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *UserDirectBenefitUpsert) UpdateUserID() *UserDirectBenefitUpsert {
	u.SetExcluded(userdirectbenefit.FieldUserID)
	return u
}

// SetCoinTypeID sets the "coin_type_id" field.
func (u *UserDirectBenefitUpsert) SetCoinTypeID(v uuid.UUID) *UserDirectBenefitUpsert {
	u.Set(userdirectbenefit.FieldCoinTypeID, v)
	return u
}

// UpdateCoinTypeID sets the "coin_type_id" field to the value that was provided on create.
func (u *UserDirectBenefitUpsert) UpdateCoinTypeID() *UserDirectBenefitUpsert {
	u.SetExcluded(userdirectbenefit.FieldCoinTypeID)
	return u
}

// SetAccountID sets the "account_id" field.
func (u *UserDirectBenefitUpsert) SetAccountID(v uuid.UUID) *UserDirectBenefitUpsert {
	u.Set(userdirectbenefit.FieldAccountID, v)
	return u
}

// UpdateAccountID sets the "account_id" field to the value that was provided on create.
func (u *UserDirectBenefitUpsert) UpdateAccountID() *UserDirectBenefitUpsert {
	u.SetExcluded(userdirectbenefit.FieldAccountID)
	return u
}

// SetCreateAt sets the "create_at" field.
func (u *UserDirectBenefitUpsert) SetCreateAt(v uint32) *UserDirectBenefitUpsert {
	u.Set(userdirectbenefit.FieldCreateAt, v)
	return u
}

// UpdateCreateAt sets the "create_at" field to the value that was provided on create.
func (u *UserDirectBenefitUpsert) UpdateCreateAt() *UserDirectBenefitUpsert {
	u.SetExcluded(userdirectbenefit.FieldCreateAt)
	return u
}

// AddCreateAt adds v to the "create_at" field.
func (u *UserDirectBenefitUpsert) AddCreateAt(v uint32) *UserDirectBenefitUpsert {
	u.Add(userdirectbenefit.FieldCreateAt, v)
	return u
}

// SetUpdateAt sets the "update_at" field.
func (u *UserDirectBenefitUpsert) SetUpdateAt(v uint32) *UserDirectBenefitUpsert {
	u.Set(userdirectbenefit.FieldUpdateAt, v)
	return u
}

// UpdateUpdateAt sets the "update_at" field to the value that was provided on create.
func (u *UserDirectBenefitUpsert) UpdateUpdateAt() *UserDirectBenefitUpsert {
	u.SetExcluded(userdirectbenefit.FieldUpdateAt)
	return u
}

// AddUpdateAt adds v to the "update_at" field.
func (u *UserDirectBenefitUpsert) AddUpdateAt(v uint32) *UserDirectBenefitUpsert {
	u.Add(userdirectbenefit.FieldUpdateAt, v)
	return u
}

// SetDeleteAt sets the "delete_at" field.
func (u *UserDirectBenefitUpsert) SetDeleteAt(v uint32) *UserDirectBenefitUpsert {
	u.Set(userdirectbenefit.FieldDeleteAt, v)
	return u
}

// UpdateDeleteAt sets the "delete_at" field to the value that was provided on create.
func (u *UserDirectBenefitUpsert) UpdateDeleteAt() *UserDirectBenefitUpsert {
	u.SetExcluded(userdirectbenefit.FieldDeleteAt)
	return u
}

// AddDeleteAt adds v to the "delete_at" field.
func (u *UserDirectBenefitUpsert) AddDeleteAt(v uint32) *UserDirectBenefitUpsert {
	u.Add(userdirectbenefit.FieldDeleteAt, v)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.UserDirectBenefit.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(userdirectbenefit.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *UserDirectBenefitUpsertOne) UpdateNewValues() *UserDirectBenefitUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(userdirectbenefit.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.UserDirectBenefit.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *UserDirectBenefitUpsertOne) Ignore() *UserDirectBenefitUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *UserDirectBenefitUpsertOne) DoNothing() *UserDirectBenefitUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the UserDirectBenefitCreate.OnConflict
// documentation for more info.
func (u *UserDirectBenefitUpsertOne) Update(set func(*UserDirectBenefitUpsert)) *UserDirectBenefitUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&UserDirectBenefitUpsert{UpdateSet: update})
	}))
	return u
}

// SetAppID sets the "app_id" field.
func (u *UserDirectBenefitUpsertOne) SetAppID(v uuid.UUID) *UserDirectBenefitUpsertOne {
	return u.Update(func(s *UserDirectBenefitUpsert) {
		s.SetAppID(v)
	})
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *UserDirectBenefitUpsertOne) UpdateAppID() *UserDirectBenefitUpsertOne {
	return u.Update(func(s *UserDirectBenefitUpsert) {
		s.UpdateAppID()
	})
}

// SetUserID sets the "user_id" field.
func (u *UserDirectBenefitUpsertOne) SetUserID(v uuid.UUID) *UserDirectBenefitUpsertOne {
	return u.Update(func(s *UserDirectBenefitUpsert) {
		s.SetUserID(v)
	})
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *UserDirectBenefitUpsertOne) UpdateUserID() *UserDirectBenefitUpsertOne {
	return u.Update(func(s *UserDirectBenefitUpsert) {
		s.UpdateUserID()
	})
}

// SetCoinTypeID sets the "coin_type_id" field.
func (u *UserDirectBenefitUpsertOne) SetCoinTypeID(v uuid.UUID) *UserDirectBenefitUpsertOne {
	return u.Update(func(s *UserDirectBenefitUpsert) {
		s.SetCoinTypeID(v)
	})
}

// UpdateCoinTypeID sets the "coin_type_id" field to the value that was provided on create.
func (u *UserDirectBenefitUpsertOne) UpdateCoinTypeID() *UserDirectBenefitUpsertOne {
	return u.Update(func(s *UserDirectBenefitUpsert) {
		s.UpdateCoinTypeID()
	})
}

// SetAccountID sets the "account_id" field.
func (u *UserDirectBenefitUpsertOne) SetAccountID(v uuid.UUID) *UserDirectBenefitUpsertOne {
	return u.Update(func(s *UserDirectBenefitUpsert) {
		s.SetAccountID(v)
	})
}

// UpdateAccountID sets the "account_id" field to the value that was provided on create.
func (u *UserDirectBenefitUpsertOne) UpdateAccountID() *UserDirectBenefitUpsertOne {
	return u.Update(func(s *UserDirectBenefitUpsert) {
		s.UpdateAccountID()
	})
}

// SetCreateAt sets the "create_at" field.
func (u *UserDirectBenefitUpsertOne) SetCreateAt(v uint32) *UserDirectBenefitUpsertOne {
	return u.Update(func(s *UserDirectBenefitUpsert) {
		s.SetCreateAt(v)
	})
}

// AddCreateAt adds v to the "create_at" field.
func (u *UserDirectBenefitUpsertOne) AddCreateAt(v uint32) *UserDirectBenefitUpsertOne {
	return u.Update(func(s *UserDirectBenefitUpsert) {
		s.AddCreateAt(v)
	})
}

// UpdateCreateAt sets the "create_at" field to the value that was provided on create.
func (u *UserDirectBenefitUpsertOne) UpdateCreateAt() *UserDirectBenefitUpsertOne {
	return u.Update(func(s *UserDirectBenefitUpsert) {
		s.UpdateCreateAt()
	})
}

// SetUpdateAt sets the "update_at" field.
func (u *UserDirectBenefitUpsertOne) SetUpdateAt(v uint32) *UserDirectBenefitUpsertOne {
	return u.Update(func(s *UserDirectBenefitUpsert) {
		s.SetUpdateAt(v)
	})
}

// AddUpdateAt adds v to the "update_at" field.
func (u *UserDirectBenefitUpsertOne) AddUpdateAt(v uint32) *UserDirectBenefitUpsertOne {
	return u.Update(func(s *UserDirectBenefitUpsert) {
		s.AddUpdateAt(v)
	})
}

// UpdateUpdateAt sets the "update_at" field to the value that was provided on create.
func (u *UserDirectBenefitUpsertOne) UpdateUpdateAt() *UserDirectBenefitUpsertOne {
	return u.Update(func(s *UserDirectBenefitUpsert) {
		s.UpdateUpdateAt()
	})
}

// SetDeleteAt sets the "delete_at" field.
func (u *UserDirectBenefitUpsertOne) SetDeleteAt(v uint32) *UserDirectBenefitUpsertOne {
	return u.Update(func(s *UserDirectBenefitUpsert) {
		s.SetDeleteAt(v)
	})
}

// AddDeleteAt adds v to the "delete_at" field.
func (u *UserDirectBenefitUpsertOne) AddDeleteAt(v uint32) *UserDirectBenefitUpsertOne {
	return u.Update(func(s *UserDirectBenefitUpsert) {
		s.AddDeleteAt(v)
	})
}

// UpdateDeleteAt sets the "delete_at" field to the value that was provided on create.
func (u *UserDirectBenefitUpsertOne) UpdateDeleteAt() *UserDirectBenefitUpsertOne {
	return u.Update(func(s *UserDirectBenefitUpsert) {
		s.UpdateDeleteAt()
	})
}

// Exec executes the query.
func (u *UserDirectBenefitUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for UserDirectBenefitCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *UserDirectBenefitUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *UserDirectBenefitUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: UserDirectBenefitUpsertOne.ID is not supported by MySQL driver. Use UserDirectBenefitUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *UserDirectBenefitUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// UserDirectBenefitCreateBulk is the builder for creating many UserDirectBenefit entities in bulk.
type UserDirectBenefitCreateBulk struct {
	config
	builders []*UserDirectBenefitCreate
	conflict []sql.ConflictOption
}

// Save creates the UserDirectBenefit entities in the database.
func (udbcb *UserDirectBenefitCreateBulk) Save(ctx context.Context) ([]*UserDirectBenefit, error) {
	specs := make([]*sqlgraph.CreateSpec, len(udbcb.builders))
	nodes := make([]*UserDirectBenefit, len(udbcb.builders))
	mutators := make([]Mutator, len(udbcb.builders))
	for i := range udbcb.builders {
		func(i int, root context.Context) {
			builder := udbcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserDirectBenefitMutation)
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
					_, err = mutators[i+1].Mutate(root, udbcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = udbcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, udbcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, udbcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (udbcb *UserDirectBenefitCreateBulk) SaveX(ctx context.Context) []*UserDirectBenefit {
	v, err := udbcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (udbcb *UserDirectBenefitCreateBulk) Exec(ctx context.Context) error {
	_, err := udbcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (udbcb *UserDirectBenefitCreateBulk) ExecX(ctx context.Context) {
	if err := udbcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.UserDirectBenefit.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.UserDirectBenefitUpsert) {
//			SetAppID(v+v).
//		}).
//		Exec(ctx)
//
func (udbcb *UserDirectBenefitCreateBulk) OnConflict(opts ...sql.ConflictOption) *UserDirectBenefitUpsertBulk {
	udbcb.conflict = opts
	return &UserDirectBenefitUpsertBulk{
		create: udbcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.UserDirectBenefit.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (udbcb *UserDirectBenefitCreateBulk) OnConflictColumns(columns ...string) *UserDirectBenefitUpsertBulk {
	udbcb.conflict = append(udbcb.conflict, sql.ConflictColumns(columns...))
	return &UserDirectBenefitUpsertBulk{
		create: udbcb,
	}
}

// UserDirectBenefitUpsertBulk is the builder for "upsert"-ing
// a bulk of UserDirectBenefit nodes.
type UserDirectBenefitUpsertBulk struct {
	create *UserDirectBenefitCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.UserDirectBenefit.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(userdirectbenefit.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *UserDirectBenefitUpsertBulk) UpdateNewValues() *UserDirectBenefitUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(userdirectbenefit.FieldID)
				return
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.UserDirectBenefit.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *UserDirectBenefitUpsertBulk) Ignore() *UserDirectBenefitUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *UserDirectBenefitUpsertBulk) DoNothing() *UserDirectBenefitUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the UserDirectBenefitCreateBulk.OnConflict
// documentation for more info.
func (u *UserDirectBenefitUpsertBulk) Update(set func(*UserDirectBenefitUpsert)) *UserDirectBenefitUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&UserDirectBenefitUpsert{UpdateSet: update})
	}))
	return u
}

// SetAppID sets the "app_id" field.
func (u *UserDirectBenefitUpsertBulk) SetAppID(v uuid.UUID) *UserDirectBenefitUpsertBulk {
	return u.Update(func(s *UserDirectBenefitUpsert) {
		s.SetAppID(v)
	})
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *UserDirectBenefitUpsertBulk) UpdateAppID() *UserDirectBenefitUpsertBulk {
	return u.Update(func(s *UserDirectBenefitUpsert) {
		s.UpdateAppID()
	})
}

// SetUserID sets the "user_id" field.
func (u *UserDirectBenefitUpsertBulk) SetUserID(v uuid.UUID) *UserDirectBenefitUpsertBulk {
	return u.Update(func(s *UserDirectBenefitUpsert) {
		s.SetUserID(v)
	})
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *UserDirectBenefitUpsertBulk) UpdateUserID() *UserDirectBenefitUpsertBulk {
	return u.Update(func(s *UserDirectBenefitUpsert) {
		s.UpdateUserID()
	})
}

// SetCoinTypeID sets the "coin_type_id" field.
func (u *UserDirectBenefitUpsertBulk) SetCoinTypeID(v uuid.UUID) *UserDirectBenefitUpsertBulk {
	return u.Update(func(s *UserDirectBenefitUpsert) {
		s.SetCoinTypeID(v)
	})
}

// UpdateCoinTypeID sets the "coin_type_id" field to the value that was provided on create.
func (u *UserDirectBenefitUpsertBulk) UpdateCoinTypeID() *UserDirectBenefitUpsertBulk {
	return u.Update(func(s *UserDirectBenefitUpsert) {
		s.UpdateCoinTypeID()
	})
}

// SetAccountID sets the "account_id" field.
func (u *UserDirectBenefitUpsertBulk) SetAccountID(v uuid.UUID) *UserDirectBenefitUpsertBulk {
	return u.Update(func(s *UserDirectBenefitUpsert) {
		s.SetAccountID(v)
	})
}

// UpdateAccountID sets the "account_id" field to the value that was provided on create.
func (u *UserDirectBenefitUpsertBulk) UpdateAccountID() *UserDirectBenefitUpsertBulk {
	return u.Update(func(s *UserDirectBenefitUpsert) {
		s.UpdateAccountID()
	})
}

// SetCreateAt sets the "create_at" field.
func (u *UserDirectBenefitUpsertBulk) SetCreateAt(v uint32) *UserDirectBenefitUpsertBulk {
	return u.Update(func(s *UserDirectBenefitUpsert) {
		s.SetCreateAt(v)
	})
}

// AddCreateAt adds v to the "create_at" field.
func (u *UserDirectBenefitUpsertBulk) AddCreateAt(v uint32) *UserDirectBenefitUpsertBulk {
	return u.Update(func(s *UserDirectBenefitUpsert) {
		s.AddCreateAt(v)
	})
}

// UpdateCreateAt sets the "create_at" field to the value that was provided on create.
func (u *UserDirectBenefitUpsertBulk) UpdateCreateAt() *UserDirectBenefitUpsertBulk {
	return u.Update(func(s *UserDirectBenefitUpsert) {
		s.UpdateCreateAt()
	})
}

// SetUpdateAt sets the "update_at" field.
func (u *UserDirectBenefitUpsertBulk) SetUpdateAt(v uint32) *UserDirectBenefitUpsertBulk {
	return u.Update(func(s *UserDirectBenefitUpsert) {
		s.SetUpdateAt(v)
	})
}

// AddUpdateAt adds v to the "update_at" field.
func (u *UserDirectBenefitUpsertBulk) AddUpdateAt(v uint32) *UserDirectBenefitUpsertBulk {
	return u.Update(func(s *UserDirectBenefitUpsert) {
		s.AddUpdateAt(v)
	})
}

// UpdateUpdateAt sets the "update_at" field to the value that was provided on create.
func (u *UserDirectBenefitUpsertBulk) UpdateUpdateAt() *UserDirectBenefitUpsertBulk {
	return u.Update(func(s *UserDirectBenefitUpsert) {
		s.UpdateUpdateAt()
	})
}

// SetDeleteAt sets the "delete_at" field.
func (u *UserDirectBenefitUpsertBulk) SetDeleteAt(v uint32) *UserDirectBenefitUpsertBulk {
	return u.Update(func(s *UserDirectBenefitUpsert) {
		s.SetDeleteAt(v)
	})
}

// AddDeleteAt adds v to the "delete_at" field.
func (u *UserDirectBenefitUpsertBulk) AddDeleteAt(v uint32) *UserDirectBenefitUpsertBulk {
	return u.Update(func(s *UserDirectBenefitUpsert) {
		s.AddDeleteAt(v)
	})
}

// UpdateDeleteAt sets the "delete_at" field to the value that was provided on create.
func (u *UserDirectBenefitUpsertBulk) UpdateDeleteAt() *UserDirectBenefitUpsertBulk {
	return u.Update(func(s *UserDirectBenefitUpsert) {
		s.UpdateDeleteAt()
	})
}

// Exec executes the query.
func (u *UserDirectBenefitUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the UserDirectBenefitCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for UserDirectBenefitCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *UserDirectBenefitUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}