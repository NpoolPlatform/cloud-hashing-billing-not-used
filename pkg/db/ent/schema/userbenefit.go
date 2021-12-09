package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

	"github.com/google/uuid"
)

// UserBenefit holds the schema definition for the UserBenefit entity.
type UserBenefit struct {
	ent.Schema
}

// Fields of the UserBenefit.
func (UserBenefit) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.UUID("good_id", uuid.UUID{}),
		field.UUID("app_id", uuid.UUID{}),
		field.UUID("user_id", uuid.UUID{}),
		field.UUID("order_id", uuid.UUID{}),
		field.Uint64("amount"),
		field.Uint32("last_benefit_timestamp"),
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

// Edges of the UserBenefit.
func (UserBenefit) Edges() []ent.Edge {
	return nil
}

// Indexes of the PlatformBenefit
func (UserBenefit) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("good_id", "last_benefit_timestamp", "app_id", "user_id", "order_id").
			Unique(),
	}
}
