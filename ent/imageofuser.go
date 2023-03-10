// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/Hyeu249/noname-ent-schema/ent/goods"
	"github.com/Hyeu249/noname-ent-schema/ent/image"
	"github.com/Hyeu249/noname-ent-schema/ent/imageofuser"
	"github.com/Hyeu249/noname-ent-schema/ent/user"
	"github.com/google/uuid"
)

// ImageOfUser is the model entity for the ImageOfUser schema.
type ImageOfUser struct {
	config `json:"-"`
	// ID of the ent.
	// UUID of Image Of User (UUID Ảnh của người dùng)
	ID uuid.UUID `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	// Royalty for owner (Tiền bản quyền cho chủ sở hữu)
	Royalty float32 `json:"royalty,omitempty"`
	// Is image published (Liệu hình ảnh được cho phép người khác sử dụng không?)
	IsImagePublished bool `json:"is_image_published,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ImageOfUserQuery when eager-loading is set.
	Edges               ImageOfUserEdges `json:"edges"`
	goods_image_of_user *uuid.UUID
	image_of_user_user  *uuid.UUID
}

// ImageOfUserEdges holds the relations/edges for other nodes in the graph.
type ImageOfUserEdges struct {
	// Image holds the value of the image edge.
	Image *Image `json:"image,omitempty"`
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// Goods holds the value of the goods edge.
	Goods *Goods `json:"goods,omitempty"`
	// ProductQuote holds the value of the product_quote edge.
	ProductQuote []*ProductQuote `json:"product_quote,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// ImageOrErr returns the Image value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ImageOfUserEdges) ImageOrErr() (*Image, error) {
	if e.loadedTypes[0] {
		if e.Image == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: image.Label}
		}
		return e.Image, nil
	}
	return nil, &NotLoadedError{edge: "image"}
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ImageOfUserEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[1] {
		if e.User == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// GoodsOrErr returns the Goods value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ImageOfUserEdges) GoodsOrErr() (*Goods, error) {
	if e.loadedTypes[2] {
		if e.Goods == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: goods.Label}
		}
		return e.Goods, nil
	}
	return nil, &NotLoadedError{edge: "goods"}
}

// ProductQuoteOrErr returns the ProductQuote value or an error if the edge
// was not loaded in eager-loading.
func (e ImageOfUserEdges) ProductQuoteOrErr() ([]*ProductQuote, error) {
	if e.loadedTypes[3] {
		return e.ProductQuote, nil
	}
	return nil, &NotLoadedError{edge: "product_quote"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ImageOfUser) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case imageofuser.FieldIsImagePublished:
			values[i] = new(sql.NullBool)
		case imageofuser.FieldRoyalty:
			values[i] = new(sql.NullFloat64)
		case imageofuser.FieldCreateTime, imageofuser.FieldUpdateTime, imageofuser.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		case imageofuser.FieldID:
			values[i] = new(uuid.UUID)
		case imageofuser.ForeignKeys[0]: // goods_image_of_user
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		case imageofuser.ForeignKeys[1]: // image_of_user_user
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			return nil, fmt.Errorf("unexpected column %q for type ImageOfUser", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ImageOfUser fields.
func (iou *ImageOfUser) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case imageofuser.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				iou.ID = *value
			}
		case imageofuser.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				iou.CreateTime = value.Time
			}
		case imageofuser.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				iou.UpdateTime = value.Time
			}
		case imageofuser.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				iou.DeletedAt = new(time.Time)
				*iou.DeletedAt = value.Time
			}
		case imageofuser.FieldRoyalty:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field royalty", values[i])
			} else if value.Valid {
				iou.Royalty = float32(value.Float64)
			}
		case imageofuser.FieldIsImagePublished:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_image_published", values[i])
			} else if value.Valid {
				iou.IsImagePublished = value.Bool
			}
		case imageofuser.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field goods_image_of_user", values[i])
			} else if value.Valid {
				iou.goods_image_of_user = new(uuid.UUID)
				*iou.goods_image_of_user = *value.S.(*uuid.UUID)
			}
		case imageofuser.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field image_of_user_user", values[i])
			} else if value.Valid {
				iou.image_of_user_user = new(uuid.UUID)
				*iou.image_of_user_user = *value.S.(*uuid.UUID)
			}
		}
	}
	return nil
}

// QueryImage queries the "image" edge of the ImageOfUser entity.
func (iou *ImageOfUser) QueryImage() *ImageQuery {
	return NewImageOfUserClient(iou.config).QueryImage(iou)
}

// QueryUser queries the "user" edge of the ImageOfUser entity.
func (iou *ImageOfUser) QueryUser() *UserQuery {
	return NewImageOfUserClient(iou.config).QueryUser(iou)
}

// QueryGoods queries the "goods" edge of the ImageOfUser entity.
func (iou *ImageOfUser) QueryGoods() *GoodsQuery {
	return NewImageOfUserClient(iou.config).QueryGoods(iou)
}

// QueryProductQuote queries the "product_quote" edge of the ImageOfUser entity.
func (iou *ImageOfUser) QueryProductQuote() *ProductQuoteQuery {
	return NewImageOfUserClient(iou.config).QueryProductQuote(iou)
}

// Update returns a builder for updating this ImageOfUser.
// Note that you need to call ImageOfUser.Unwrap() before calling this method if this ImageOfUser
// was returned from a transaction, and the transaction was committed or rolled back.
func (iou *ImageOfUser) Update() *ImageOfUserUpdateOne {
	return NewImageOfUserClient(iou.config).UpdateOne(iou)
}

// Unwrap unwraps the ImageOfUser entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (iou *ImageOfUser) Unwrap() *ImageOfUser {
	_tx, ok := iou.config.driver.(*txDriver)
	if !ok {
		panic("ent: ImageOfUser is not a transactional entity")
	}
	iou.config.driver = _tx.drv
	return iou
}

// String implements the fmt.Stringer.
func (iou *ImageOfUser) String() string {
	var builder strings.Builder
	builder.WriteString("ImageOfUser(")
	builder.WriteString(fmt.Sprintf("id=%v, ", iou.ID))
	builder.WriteString("create_time=")
	builder.WriteString(iou.CreateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("update_time=")
	builder.WriteString(iou.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	if v := iou.DeletedAt; v != nil {
		builder.WriteString("deleted_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	builder.WriteString("royalty=")
	builder.WriteString(fmt.Sprintf("%v", iou.Royalty))
	builder.WriteString(", ")
	builder.WriteString("is_image_published=")
	builder.WriteString(fmt.Sprintf("%v", iou.IsImagePublished))
	builder.WriteByte(')')
	return builder.String()
}

// ImageOfUsers is a parsable slice of ImageOfUser.
type ImageOfUsers []*ImageOfUser

func (iou ImageOfUsers) config(cfg config) {
	for _i := range iou {
		iou[_i].config = cfg
	}
}
