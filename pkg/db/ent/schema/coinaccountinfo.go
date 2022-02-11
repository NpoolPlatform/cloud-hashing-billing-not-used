package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

	"github.com/google/uuid"
)

// CoinAccountInfo holds the schema definition for the CoinAccountInfo entity.
type CoinAccountInfo struct {
	ent.Schema
}

// Fields of the CoinAccountInfo.
func (CoinAccountInfo) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.UUID("coin_type_id", uuid.UUID{}),
		field.String("address"),
		field.Bool("platform_hold_private_key"),
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

// Edges of the CoinAccountInfo.
func (CoinAccountInfo) Edges() []ent.Edge {
	return nil
}

// Indexes of the CoinAccountInfo.
func (CoinAccountInfo) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("coin_type_id", "address").
			Unique(),
	}
}
