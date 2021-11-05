package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"

	"github.com/google/uuid"
)

// CoinAccountTransaction holds the schema definition for the CoinAccountTransaction entity.
type CoinAccountTransaction struct {
	ent.Schema
}

// Fields of the CoinAccountTransaction.
func (CoinAccountTransaction) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
		field.UUID("user_id", uuid.UUID{}),
		field.UUID("app_id", uuid.UUID{}),
		field.UUID("from_address_id", uuid.UUID{}),
		field.UUID("to_address_id", uuid.UUID{}),
		field.UUID("coin_type_id", uuid.UUID{}),
		field.Int64("amount"),
		field.String("message"),
		field.Enum("state").
			Values("wait", "paying", "successful", "fail"),
		field.String("chain_transaction_id"),
		field.UUID("platform_transaction_id", uuid.UUID{}),
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

// Edges of the CoinAccountTransaction.
func (CoinAccountTransaction) Edges() []ent.Edge {
	return nil
}
