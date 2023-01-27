package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/google/uuid"
)

// TypeOfGoods holds the schema definition for the TypeOfGoods entity.
type TypeOfGoods struct {
	ent.Schema
}

// Learn mixin: https://entgo.io/docs/schema-mixin
func (TypeOfGoods) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		SoftDeleteMixin{},
	}
}

// Fields of the TypeOfGoods.
func (TypeOfGoods) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Comment("UUID of Type Of Goods (UUID Loại mặt hàng)"),

		field.String("name").
			NotEmpty().
			Comment("Name of Goods (Tên của mặt hàng)"),
		field.String("type").
			NotEmpty().
			Comment("Type of Goods (Loại mặt hàng)"),
		field.String("description").
			Comment("Description of Type of Goods (Mô tả loại mặt hàng)"),
	}
}

// Edges of the TypeOfGoods.
func (TypeOfGoods) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("goods", Goods.Type).
			Ref("type_of_goods").
			Unique(),
	}
}
