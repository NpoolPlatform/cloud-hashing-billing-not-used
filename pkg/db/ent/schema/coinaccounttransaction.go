package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"

	"github.com/google/uuid"

	constant "github.com/NpoolPlatform/cloud-hashing-billing/pkg/const"
)

// CoinAccountTransaction holds the schema definition for the CoinAccountTransaction entity.
type CoinAccountTransaction struct {
	ent.Schema
}

// Fields of the CoinAccountTransaction.
func (CoinAccountTransaction) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.UUID("app_id", uuid.UUID{}),
		field.UUID("user_id", uuid.UUID{}),
		field.UUID("good_id", uuid.UUID{}),
		field.UUID("from_address_id", uuid.UUID{}),
		field.UUID("to_address_id", uuid.UUID{}),
		field.UUID("coin_type_id", uuid.UUID{}),
		field.Uint64("amount"),
		field.String("message"),
		field.Enum("state").
			Values(constant.CoinTransactionStateCreated,
				constant.CoinTransactionStateWait,
				constant.CoinTransactionStatePaying,
				constant.CoinTransactionStateSuccessful,
				constant.CoinTransactionStateRejected,
				constant.CoinTransactionStateFail),
		field.String("chain_transaction_id"),
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
