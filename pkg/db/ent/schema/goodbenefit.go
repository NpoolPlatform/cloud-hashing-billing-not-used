package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"

	"github.com/google/uuid"
)

// GoodBenefit holds the schema definition for the GoodBenefit entity.
type GoodBenefit struct {
	ent.Schema
}

// Fields of the GoodBenefit.
func (GoodBenefit) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Unique(),
		field.UUID("good_id", uuid.UUID{}).Unique(),
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

// Edges of the GoodBenefit.
func (GoodBenefit) Edges() []ent.Edge {
	return nil
}
