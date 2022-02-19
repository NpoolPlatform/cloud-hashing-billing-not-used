// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/db/ent/userbenefit"
	"github.com/google/uuid"
)

// UserBenefit is the model entity for the UserBenefit schema.
type UserBenefit struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// AppID holds the value of the "app_id" field.
	AppID uuid.UUID `json:"app_id,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID uuid.UUID `json:"user_id,omitempty"`
	// GoodID holds the value of the "good_id" field.
	GoodID uuid.UUID `json:"good_id,omitempty"`
	// OrderID holds the value of the "order_id" field.
	OrderID uuid.UUID `json:"order_id,omitempty"`
	// CoinTypeID holds the value of the "coin_type_id" field.
	CoinTypeID uuid.UUID `json:"coin_type_id,omitempty"`
	// Amount holds the value of the "amount" field.
	Amount uint64 `json:"amount,omitempty"`
	// LastBenefitTimestamp holds the value of the "last_benefit_timestamp" field.
	LastBenefitTimestamp uint32 `json:"last_benefit_timestamp,omitempty"`
	// CreateAt holds the value of the "create_at" field.
	CreateAt uint32 `json:"create_at,omitempty"`
	// UpdateAt holds the value of the "update_at" field.
	UpdateAt uint32 `json:"update_at,omitempty"`
	// DeleteAt holds the value of the "delete_at" field.
	DeleteAt uint32 `json:"delete_at,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*UserBenefit) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case userbenefit.FieldAmount, userbenefit.FieldLastBenefitTimestamp, userbenefit.FieldCreateAt, userbenefit.FieldUpdateAt, userbenefit.FieldDeleteAt:
			values[i] = new(sql.NullInt64)
		case userbenefit.FieldID, userbenefit.FieldAppID, userbenefit.FieldUserID, userbenefit.FieldGoodID, userbenefit.FieldOrderID, userbenefit.FieldCoinTypeID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type UserBenefit", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the UserBenefit fields.
func (ub *UserBenefit) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case userbenefit.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				ub.ID = *value
			}
		case userbenefit.FieldAppID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field app_id", values[i])
			} else if value != nil {
				ub.AppID = *value
			}
		case userbenefit.FieldUserID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value != nil {
				ub.UserID = *value
			}
		case userbenefit.FieldGoodID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field good_id", values[i])
			} else if value != nil {
				ub.GoodID = *value
			}
		case userbenefit.FieldOrderID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field order_id", values[i])
			} else if value != nil {
				ub.OrderID = *value
			}
		case userbenefit.FieldCoinTypeID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field coin_type_id", values[i])
			} else if value != nil {
				ub.CoinTypeID = *value
			}
		case userbenefit.FieldAmount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field amount", values[i])
			} else if value.Valid {
				ub.Amount = uint64(value.Int64)
			}
		case userbenefit.FieldLastBenefitTimestamp:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field last_benefit_timestamp", values[i])
			} else if value.Valid {
				ub.LastBenefitTimestamp = uint32(value.Int64)
			}
		case userbenefit.FieldCreateAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field create_at", values[i])
			} else if value.Valid {
				ub.CreateAt = uint32(value.Int64)
			}
		case userbenefit.FieldUpdateAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field update_at", values[i])
			} else if value.Valid {
				ub.UpdateAt = uint32(value.Int64)
			}
		case userbenefit.FieldDeleteAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field delete_at", values[i])
			} else if value.Valid {
				ub.DeleteAt = uint32(value.Int64)
			}
		}
	}
	return nil
}

// Update returns a builder for updating this UserBenefit.
// Note that you need to call UserBenefit.Unwrap() before calling this method if this UserBenefit
// was returned from a transaction, and the transaction was committed or rolled back.
func (ub *UserBenefit) Update() *UserBenefitUpdateOne {
	return (&UserBenefitClient{config: ub.config}).UpdateOne(ub)
}

// Unwrap unwraps the UserBenefit entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ub *UserBenefit) Unwrap() *UserBenefit {
	tx, ok := ub.config.driver.(*txDriver)
	if !ok {
		panic("ent: UserBenefit is not a transactional entity")
	}
	ub.config.driver = tx.drv
	return ub
}

// String implements the fmt.Stringer.
func (ub *UserBenefit) String() string {
	var builder strings.Builder
	builder.WriteString("UserBenefit(")
	builder.WriteString(fmt.Sprintf("id=%v", ub.ID))
	builder.WriteString(", app_id=")
	builder.WriteString(fmt.Sprintf("%v", ub.AppID))
	builder.WriteString(", user_id=")
	builder.WriteString(fmt.Sprintf("%v", ub.UserID))
	builder.WriteString(", good_id=")
	builder.WriteString(fmt.Sprintf("%v", ub.GoodID))
	builder.WriteString(", order_id=")
	builder.WriteString(fmt.Sprintf("%v", ub.OrderID))
	builder.WriteString(", coin_type_id=")
	builder.WriteString(fmt.Sprintf("%v", ub.CoinTypeID))
	builder.WriteString(", amount=")
	builder.WriteString(fmt.Sprintf("%v", ub.Amount))
	builder.WriteString(", last_benefit_timestamp=")
	builder.WriteString(fmt.Sprintf("%v", ub.LastBenefitTimestamp))
	builder.WriteString(", create_at=")
	builder.WriteString(fmt.Sprintf("%v", ub.CreateAt))
	builder.WriteString(", update_at=")
	builder.WriteString(fmt.Sprintf("%v", ub.UpdateAt))
	builder.WriteString(", delete_at=")
	builder.WriteString(fmt.Sprintf("%v", ub.DeleteAt))
	builder.WriteByte(')')
	return builder.String()
}

// UserBenefits is a parsable slice of UserBenefit.
type UserBenefits []*UserBenefit

func (ub UserBenefits) config(cfg config) {
	for _i := range ub {
		ub[_i].config = cfg
	}
}
