package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/google/uuid"
)

// ImageOfUser holds the schema definition for the ImageOfUser entity.
type ImageOfUser struct {
	ent.Schema
}

// Learn mixin: https://entgo.io/docs/schema-mixin
func (ImageOfUser) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		SoftDeleteMixin{},
	}
}

// Fields of the ImageOfUser.
func (ImageOfUser) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Comment("UUID of Image Of User (UUID Ảnh của người dùng)"),

		field.Float32("royalty").
			Optional().
			Comment("Royalty for owner (Tiền bản quyền cho chủ sở hữu)"),
		field.Bool("is_image_published").
			Default(false).
			Comment("Is image published (Liệu hình ảnh được cho phép người khác sử dụng không?)"),
	}
}

// Edges of the ImageOfUser.
func (ImageOfUser) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("image", Image.Type).
			Unique().
			Required(),
		edge.To("user", User.Type).
			Unique().
			Required(),

		edge.From("goods", Goods.Type).
			Ref("image_of_user").
			Unique(),
		edge.From("product_quote", ProductQuote.Type).
			Ref("image_of_user"),
	}
}
