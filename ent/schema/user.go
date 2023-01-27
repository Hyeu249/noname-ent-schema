package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/google/uuid"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Learn mixin: https://entgo.io/docs/schema-mixin
func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		SoftDeleteMixin{},
	}
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).Comment("UUID of invoice (UUID nguoi dung)"),

		field.String("username").
			NotEmpty().
			Unique().
			Comment("Username"),
		field.Bool("is_active").
			Default(true).
			Comment("Is user active?"),
		field.String("hashed_pwd").
			NotEmpty().
			Comment("Hashed password of user"),
		field.Bool("is_superuser").
			Default(false).
			Comment("Is user a super user?"),
		field.String("name").
			NotEmpty().
			Comment("Name of user (Ten nguoi dung)"),
		field.Text("email").
			Unique().
			Comment("Email of user (sdt)"),
		field.Text("phone").
			Optional().
			Comment("Phone no. of user (sdt)"),
		field.Text("address").
			Optional().
			Comment("Address of user (dia chi)"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("image_of_user", ImageOfUser.Type).
			Ref("user"),
		edge.From("product_quote", ProductQuote.Type).
			Ref("user"),
		edge.From("item_has_owner", ItemHasOwner.Type).
			Ref("user"),
	}
}
