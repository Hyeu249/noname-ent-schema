// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/Hyeu249/noname-ent-schema/ent/goods"
	"github.com/Hyeu249/noname-ent-schema/ent/imageofuser"
	"github.com/Hyeu249/noname-ent-schema/ent/typeofgoods"
	"github.com/google/uuid"
)

// Goods is the model entity for the Goods schema.
type Goods struct {
	config `json:"-"`
	// ID of the ent.
	// UUID of Goods (UUID Mặt hàng)
	ID uuid.UUID `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	// Description of Goods (Mô tả mặt hàng)
	Description string `json:"description,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the GoodsQuery when eager-loading is set.
	Edges GoodsEdges `json:"edges"`
}

// GoodsEdges holds the relations/edges for other nodes in the graph.
type GoodsEdges struct {
	// TypeOfGoods holds the value of the type_of_goods edge.
	TypeOfGoods *TypeOfGoods `json:"type_of_goods,omitempty"`
	// ImageOfUser holds the value of the image_of_user edge.
	ImageOfUser *ImageOfUser `json:"image_of_user,omitempty"`
	// ProductQuote holds the value of the product_quote edge.
	ProductQuote []*ProductQuote `json:"product_quote,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// TypeOfGoodsOrErr returns the TypeOfGoods value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e GoodsEdges) TypeOfGoodsOrErr() (*TypeOfGoods, error) {
	if e.loadedTypes[0] {
		if e.TypeOfGoods == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: typeofgoods.Label}
		}
		return e.TypeOfGoods, nil
	}
	return nil, &NotLoadedError{edge: "type_of_goods"}
}

// ImageOfUserOrErr returns the ImageOfUser value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e GoodsEdges) ImageOfUserOrErr() (*ImageOfUser, error) {
	if e.loadedTypes[1] {
		if e.ImageOfUser == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: imageofuser.Label}
		}
		return e.ImageOfUser, nil
	}
	return nil, &NotLoadedError{edge: "image_of_user"}
}

// ProductQuoteOrErr returns the ProductQuote value or an error if the edge
// was not loaded in eager-loading.
func (e GoodsEdges) ProductQuoteOrErr() ([]*ProductQuote, error) {
	if e.loadedTypes[2] {
		return e.ProductQuote, nil
	}
	return nil, &NotLoadedError{edge: "product_quote"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Goods) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case goods.FieldDescription:
			values[i] = new(sql.NullString)
		case goods.FieldCreateTime, goods.FieldUpdateTime, goods.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		case goods.FieldID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Goods", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Goods fields.
func (_go *Goods) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case goods.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				_go.ID = *value
			}
		case goods.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				_go.CreateTime = value.Time
			}
		case goods.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				_go.UpdateTime = value.Time
			}
		case goods.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				_go.DeletedAt = new(time.Time)
				*_go.DeletedAt = value.Time
			}
		case goods.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				_go.Description = value.String
			}
		}
	}
	return nil
}

// QueryTypeOfGoods queries the "type_of_goods" edge of the Goods entity.
func (_go *Goods) QueryTypeOfGoods() *TypeOfGoodsQuery {
	return NewGoodsClient(_go.config).QueryTypeOfGoods(_go)
}

// QueryImageOfUser queries the "image_of_user" edge of the Goods entity.
func (_go *Goods) QueryImageOfUser() *ImageOfUserQuery {
	return NewGoodsClient(_go.config).QueryImageOfUser(_go)
}

// QueryProductQuote queries the "product_quote" edge of the Goods entity.
func (_go *Goods) QueryProductQuote() *ProductQuoteQuery {
	return NewGoodsClient(_go.config).QueryProductQuote(_go)
}

// Update returns a builder for updating this Goods.
// Note that you need to call Goods.Unwrap() before calling this method if this Goods
// was returned from a transaction, and the transaction was committed or rolled back.
func (_go *Goods) Update() *GoodsUpdateOne {
	return NewGoodsClient(_go.config).UpdateOne(_go)
}

// Unwrap unwraps the Goods entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (_go *Goods) Unwrap() *Goods {
	_tx, ok := _go.config.driver.(*txDriver)
	if !ok {
		panic("ent: Goods is not a transactional entity")
	}
	_go.config.driver = _tx.drv
	return _go
}

// String implements the fmt.Stringer.
func (_go *Goods) String() string {
	var builder strings.Builder
	builder.WriteString("Goods(")
	builder.WriteString(fmt.Sprintf("id=%v, ", _go.ID))
	builder.WriteString("create_time=")
	builder.WriteString(_go.CreateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("update_time=")
	builder.WriteString(_go.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	if v := _go.DeletedAt; v != nil {
		builder.WriteString("deleted_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(_go.Description)
	builder.WriteByte(')')
	return builder.String()
}

// GoodsSlice is a parsable slice of Goods.
type GoodsSlice []*Goods

func (_go GoodsSlice) config(cfg config) {
	for _i := range _go {
		_go[_i].config = cfg
	}
}
