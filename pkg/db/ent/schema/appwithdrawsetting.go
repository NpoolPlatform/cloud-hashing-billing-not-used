package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

	"github.com/google/uuid"
)

// AppWithdrawSetting holds the schema definition for the AppWithdrawSetting entity.
type AppWithdrawSetting struct {
	ent.Schema
}

// Fields of the AppWithdrawSetting.
func (AppWithdrawSetting) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Unique(),
		field.UUID("app_id", uuid.UUID{}).Unique(),
		field.UUID("coin_type_id", uuid.UUID{}).Unique(),
		field.Uint64("withdraw_auto_review_coin_amount"),
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

// Edges of the AppWithdrawSetting.
func (AppWithdrawSetting) Edges() []ent.Edge {
	return nil
}

// Indexes of the AppWithdrawSetting.
func (AppWithdrawSetting) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("app_id", "coin_type_id").
			Unique(),
	}
}
