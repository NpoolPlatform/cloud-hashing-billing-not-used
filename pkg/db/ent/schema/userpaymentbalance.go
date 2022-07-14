package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

	"github.com/google/uuid"
)

// UserPaymentBalance holds the schema definition for the UserPaymentBalance entity.
type UserPaymentBalance struct {
	ent.Schema
}

// Fields of the UserPaymentBalance.
func (UserPaymentBalance) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.UUID("app_id", uuid.UUID{}),
		field.UUID("user_id", uuid.UUID{}),
		field.UUID("payment_id", uuid.UUID{}),
		field.UUID("used_by_payment_id", uuid.UUID{}),
		field.Uint64("amount"),
		field.Uint64("coin_usd_currency"),
		field.UUID("coin_type_id", uuid.UUID{}).
			Default(func() uuid.UUID {
				return uuid.UUID{}
			}),
		field.Uint32("create_at").
			DefaultFunc(func() uint32 {
				return uint32(time.Now().Unix())
			}),
		field.Uint32("update_at").
			DefaultFunc(func() uint32 {
				return uint32(time.Now().Unix())
			}).
			UpdateDefault(func() uint32 {
				return uint32(time.Now().Unix())
			}),
		field.Uint32("delete_at").
			DefaultFunc(func() uint32 {
				return 0
			}),
	}
}

// Edges of the UserPaymentBalance.
func (UserPaymentBalance) Edges() []ent.Edge {
	return nil
}

// Indexes of the UserPaymentBalance
func (UserPaymentBalance) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("app_id", "user_id", "payment_id").
			Unique(),
	}
}
