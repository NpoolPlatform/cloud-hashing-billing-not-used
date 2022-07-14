package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"

	"github.com/google/uuid"
)

// UserWithdrawItem holds the schema definition for the UserWithdrawItem entity.
type UserWithdrawItem struct {
	ent.Schema
}

// Fields of the UserWithdrawItem.
func (UserWithdrawItem) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.UUID("app_id", uuid.UUID{}),
		field.UUID("user_id", uuid.UUID{}),
		field.UUID("coin_type_id", uuid.UUID{}),
		field.UUID("withdraw_to_account_id", uuid.UUID{}),
		field.Uint64("amount"),
		field.UUID("platform_transaction_id", uuid.UUID{}),
		field.String("withdraw_type"),
		field.Bool("exempt_fee").Default(false),
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

// Edges of the UserWithdrawItem.
func (UserWithdrawItem) Edges() []ent.Edge {
	return nil
}
