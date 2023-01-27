package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/google/uuid"
)

// Goods holds the schema definition for the Goods entity.
type Goods struct {
	ent.Schema
}

// Learn mixin: https://entgo.io/docs/schema-mixin
func (Goods) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		SoftDeleteMixin{},
	}
}

// Fields of the Goods.
func (Goods) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Comment("UUID of Goods (UUID Mặt hàng)"),

		field.String("description").
			Comment("Description of Goods (Mô tả mặt hàng)"),
	}
}

// Edges of the Goods.
func (Goods) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("type_of_goods", TypeOfGoods.Type).
			Unique().
			Required(),
		edge.To("image_of_user", ImageOfUser.Type).
			Unique().
			Required(),

		edge.From("product_quote", ProductQuote.Type).
			Ref("goods"),
	}
}
