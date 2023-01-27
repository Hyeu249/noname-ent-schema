package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/google/uuid"
)

// ProductQuote holds the schema definition for the ProductQuote entity.
type ProductQuote struct {
	ent.Schema
}

// Learn mixin: https://entgo.io/docs/schema-mixin
func (ProductQuote) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		SoftDeleteMixin{},
	}
}

// Fields of the ProductQuote.
func (ProductQuote) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Comment("UUID of Product Quote (UUID Báo giá mặt hàng)"),

		field.Int("price").
			Comment("Price of Goods (Giá của mặt hàng)"),
		field.Bool("is_active").
			Default(false).
			Comment("is product quote actived (Liệu báo giá mặt hàng được kích hoạt?)"),
	}
}

// Edges of the ProductQuote.
func (ProductQuote) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", User.Type).
			Unique().
			Required(),
		edge.To("goods", Goods.Type).
			Unique().
			Required(),
		edge.To("image_of_user", ImageOfUser.Type).
			Unique().
			Required(),

		edge.From("item_has_owner", ItemHasOwner.Type).
			Ref("product_quote"),
	}
}
