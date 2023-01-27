// Code generated by ent, DO NOT EDIT.

package typeofgoods

import (
	"time"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the typeofgoods type in the database.
	Label = "type_of_goods"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// EdgeGoods holds the string denoting the goods edge name in mutations.
	EdgeGoods = "goods"
	// Table holds the table name of the typeofgoods in the database.
	Table = "type_of_goods"
	// GoodsTable is the table that holds the goods relation/edge.
	GoodsTable = "type_of_goods"
	// GoodsInverseTable is the table name for the Goods entity.
	// It exists in this package in order to avoid circular dependency with the "goods" package.
	GoodsInverseTable = "goods"
	// GoodsColumn is the table column denoting the goods relation/edge.
	GoodsColumn = "goods_type_of_goods"
)

// Columns holds all SQL columns for typeofgoods fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldDeletedAt,
	FieldName,
	FieldType,
	FieldDescription,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "type_of_goods"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"goods_type_of_goods",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreateTime holds the default value on creation for the "create_time" field.
	DefaultCreateTime func() time.Time
	// DefaultUpdateTime holds the default value on creation for the "update_time" field.
	DefaultUpdateTime func() time.Time
	// UpdateDefaultUpdateTime holds the default value on update for the "update_time" field.
	UpdateDefaultUpdateTime func() time.Time
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// TypeValidator is a validator for the "type" field. It is called by the builders before save.
	TypeValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)