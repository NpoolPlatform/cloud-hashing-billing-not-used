package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

	constant "github.com/NpoolPlatform/cloud-hashing-billing/pkg/const"

	"github.com/google/uuid"
)

// GoodPayment holds the schema definition for the GoodPayment entity.
type GoodPayment struct {
	ent.Schema
}

// Fields of the GoodPayment.
func (GoodPayment) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Unique(),
		field.UUID("good_id", uuid.UUID{}),
		field.UUID("payment_coin_type_id", uuid.UUID{}),
		field.UUID("account_id", uuid.UUID{}).Unique(),
		field.Bool("idle"),
		field.String("occupied_by"),
		field.Uint32("available_at").Default(uint32(time.Now().Unix())),
		field.UUID("collecting_tid", uuid.UUID{}).Optional().Default(func() uuid.UUID {
			return uuid.UUID{}
		}),
		field.String("used_for").Optional().Default(constant.TransactionForNotUsed),
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

// Edges of the GoodPayment.
func (GoodPayment) Edges() []ent.Edge {
	return nil
}

// Indexes of the GoodPayment.
func (GoodPayment) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("good_id", "payment_coin_type_id", "account_id").
			Unique(),
	}
}
