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
			Default(uuid.New),
		field.UUID("coin_type_id", uuid.UUID{}),
		field.String("address"),
		field.Enum("generated_by").
			Values("platform", "user"),
		field.Enum("used_for").
			Values("benefit", "offline", "user", "paying"),
		field.Bool("idle"),
		field.UUID("app_id", uuid.UUID{}),
		field.UUID("user_id", uuid.UUID{}),
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

// Indexs of the CoinAccountInfo.
func (CoinAccountInfo) Indexs() []ent.Index {
	return []ent.Index{
		index.Fields("coin_type_id", "address").
			Unique(),
	}
}