// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/Hyeu249/noname-ent-schema/ent/itemhasowner"
	"github.com/Hyeu249/noname-ent-schema/ent/order"
	"github.com/google/uuid"
)

// Order is the model entity for the Order schema.
type Order struct {
	config `json:"-"`
	// ID of the ent.
	// UUID of Order (UUID Đơn hàng)
	ID uuid.UUID `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	// Name of buyer (Tên của người mua)
	Name string `json:"name,omitempty"`
	// Address of buyer (Địa chỉ của người mua)
	Address string `json:"address,omitempty"`
	// Phone of buyer (Số điện thoại của người mua)
	Phone string `json:"phone,omitempty"`
	// Email of buyer (Email của người mua)
	Email string `json:"email,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the OrderQuery when eager-loading is set.
	Edges                OrderEdges `json:"edges"`
	order_item_has_owner *uuid.UUID
}

// OrderEdges holds the relations/edges for other nodes in the graph.
type OrderEdges struct {
	// ItemHasOwner holds the value of the item_has_owner edge.
	ItemHasOwner *ItemHasOwner `json:"item_has_owner,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// ItemHasOwnerOrErr returns the ItemHasOwner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e OrderEdges) ItemHasOwnerOrErr() (*ItemHasOwner, error) {
	if e.loadedTypes[0] {
		if e.ItemHasOwner == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: itemhasowner.Label}
		}
		return e.ItemHasOwner, nil
	}
	return nil, &NotLoadedError{edge: "item_has_owner"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Order) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case order.FieldName, order.FieldAddress, order.FieldPhone, order.FieldEmail:
			values[i] = new(sql.NullString)
		case order.FieldCreateTime, order.FieldUpdateTime, order.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		case order.FieldID:
			values[i] = new(uuid.UUID)
		case order.ForeignKeys[0]: // order_item_has_owner
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			return nil, fmt.Errorf("unexpected column %q for type Order", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Order fields.
func (o *Order) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case order.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				o.ID = *value
			}
		case order.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				o.CreateTime = value.Time
			}
		case order.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				o.UpdateTime = value.Time
			}
		case order.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				o.DeletedAt = new(time.Time)
				*o.DeletedAt = value.Time
			}
		case order.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				o.Name = value.String
			}
		case order.FieldAddress:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field address", values[i])
			} else if value.Valid {
				o.Address = value.String
			}
		case order.FieldPhone:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field phone", values[i])
			} else if value.Valid {
				o.Phone = value.String
			}
		case order.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email", values[i])
			} else if value.Valid {
				o.Email = value.String
			}
		case order.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field order_item_has_owner", values[i])
			} else if value.Valid {
				o.order_item_has_owner = new(uuid.UUID)
				*o.order_item_has_owner = *value.S.(*uuid.UUID)
			}
		}
	}
	return nil
}

// QueryItemHasOwner queries the "item_has_owner" edge of the Order entity.
func (o *Order) QueryItemHasOwner() *ItemHasOwnerQuery {
	return NewOrderClient(o.config).QueryItemHasOwner(o)
}

// Update returns a builder for updating this Order.
// Note that you need to call Order.Unwrap() before calling this method if this Order
// was returned from a transaction, and the transaction was committed or rolled back.
func (o *Order) Update() *OrderUpdateOne {
	return NewOrderClient(o.config).UpdateOne(o)
}

// Unwrap unwraps the Order entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (o *Order) Unwrap() *Order {
	_tx, ok := o.config.driver.(*txDriver)
	if !ok {
		panic("ent: Order is not a transactional entity")
	}
	o.config.driver = _tx.drv
	return o
}

// String implements the fmt.Stringer.
func (o *Order) String() string {
	var builder strings.Builder
	builder.WriteString("Order(")
	builder.WriteString(fmt.Sprintf("id=%v, ", o.ID))
	builder.WriteString("create_time=")
	builder.WriteString(o.CreateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("update_time=")
	builder.WriteString(o.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	if v := o.DeletedAt; v != nil {
		builder.WriteString("deleted_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(o.Name)
	builder.WriteString(", ")
	builder.WriteString("address=")
	builder.WriteString(o.Address)
	builder.WriteString(", ")
	builder.WriteString("phone=")
	builder.WriteString(o.Phone)
	builder.WriteString(", ")
	builder.WriteString("email=")
	builder.WriteString(o.Email)
	builder.WriteByte(')')
	return builder.String()
}

// Orders is a parsable slice of Order.
type Orders []*Order

func (o Orders) config(cfg config) {
	for _i := range o {
		o[_i].config = cfg
	}
}
