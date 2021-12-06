package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

	"github.com/google/uuid"
)

// PlatformBenefit holds the schema definition for the PlatformBenefit entity.
type PlatformBenefit struct {
	ent.Schema
}

// Fields of the PlatformBenefit.
func (PlatformBenefit) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.UUID("good_id", uuid.UUID{}),
		field.UUID("benefit_account_id", uuid.UUID{}),
		field.Uint64("amount"),
		field.Uint32("last_benefit_timestamp").
			Unique(),
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

// Edges of the PlatformBenefit.
func (PlatformBenefit) Edges() []ent.Edge {
	return nil
}

// Indexes of the PlatformBenefit
func (PlatformBenefit) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("good_id", "chain_transaction_id").
			Unique(),
	}
}
