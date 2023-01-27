package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/google/uuid"
)

// Image holds the schema definition for the Image entity.
type Image struct {
	ent.Schema
}

// Learn mixin: https://entgo.io/docs/schema-mixin
func (Image) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		SoftDeleteMixin{},
	}
}

// Fields of the Image.
func (Image) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Comment("UUID of image (UUID anh)"),

		field.String("name").
			Optional().
			Comment("Name of image (Ten anh)"),
		field.String("file_extension").
			NotEmpty().
			Comment("file extension of the upload file (incl. file extension)"),
		field.String("location").
			NotEmpty().
			Comment("Location of the image. It can be a path/to/file or an URL."),
		field.String("description").
			Optional().
			Comment("Description of image (Mo ta)"),
	}
}

// Edges of the Image.
func (Image) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("image_of_user", ImageOfUser.Type).
			Ref("image").
			Unique(),
	}
}
