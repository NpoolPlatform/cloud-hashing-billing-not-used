package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"

	"github.com/google/uuid"
)

// PlatformSetting holds the schema definition for the PlatformSetting entity.
type PlatformSetting struct {
	ent.Schema
}

// Fields of the PlatformSetting.
func (PlatformSetting) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
		field.UUID("good_id", uuid.UUID{}).
			Unique(),
		field.UUID("benefit_account_id", uuid.UUID{}),
		field.UUID("platform_offline_account_id", uuid.UUID{}),
		field.UUID("user_online_account_id", uuid.UUID{}),
		field.UUID("user_offline_account_id", uuid.UUID{}),
		field.Int32("benefit_interval_hours"),
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

// Edges of the PlatformSetting.
func (PlatformSetting) Edges() []ent.Edge {
	return nil
}
