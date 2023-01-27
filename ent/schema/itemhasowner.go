package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/google/uuid"
)

// ItemHasOwner holds the schema definition for the ItemHasOwner entity.
type ItemHasOwner struct {
	ent.Schema
}

// Learn mixin: https://entgo.io/docs/schema-mixin
func (ItemHasOwner) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		SoftDeleteMixin{},
	}
}

// Fields of the ItemHasOwner.
func (ItemHasOwner) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Comment("UUID of Item Has Owner (UUID hàng đã có chủ)"),
	}
}

// Edges of the ItemHasOwner.
func (ItemHasOwner) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("product_quote", ProductQuote.Type).
			Unique().
			Required(),
		edge.To("user", User.Type).
			Unique().
			Required(),

		edge.From("order", Order.Type).
			Ref("item_has_owner"),
	}
}
