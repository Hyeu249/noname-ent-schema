package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/google/uuid"
)

// Order holds the schema definition for the Order entity.
type Order struct {
	ent.Schema
}

// Learn mixin: https://entgo.io/docs/schema-mixin
func (Order) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		SoftDeleteMixin{},
	}
}

// Fields of the Order.
func (Order) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Comment("UUID of Order (UUID Đơn hàng)"),

		field.String("name").
			NotEmpty().
			Comment("Name of buyer (Tên của người mua)"),
		field.String("address").
			NotEmpty().
			Comment("Address of buyer (Địa chỉ của người mua)"),
		field.String("phone").
			NotEmpty().
			Comment("Phone of buyer (Số điện thoại của người mua)"),
		field.String("email").
			Optional().
			Comment("Email of buyer (Email của người mua)"),
	}
}

// Edges of the Order.
func (Order) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("item_has_owner", ItemHasOwner.Type).
			Unique().
			Required(),
	}
}
