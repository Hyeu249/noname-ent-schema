// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// GoodsColumns holds the columns for the "goods" table.
	GoodsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "description", Type: field.TypeString},
	}
	// GoodsTable holds the schema information for the "goods" table.
	GoodsTable = &schema.Table{
		Name:       "goods",
		Columns:    GoodsColumns,
		PrimaryKey: []*schema.Column{GoodsColumns[0]},
	}
	// ImagesColumns holds the columns for the "images" table.
	ImagesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "name", Type: field.TypeString, Nullable: true},
		{Name: "file_extension", Type: field.TypeString},
		{Name: "location", Type: field.TypeString},
		{Name: "description", Type: field.TypeString, Nullable: true},
		{Name: "image_of_user_image", Type: field.TypeUUID, Unique: true, Nullable: true},
	}
	// ImagesTable holds the schema information for the "images" table.
	ImagesTable = &schema.Table{
		Name:       "images",
		Columns:    ImagesColumns,
		PrimaryKey: []*schema.Column{ImagesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "images_image_of_users_image",
				Columns:    []*schema.Column{ImagesColumns[8]},
				RefColumns: []*schema.Column{ImageOfUsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// ImageOfUsersColumns holds the columns for the "image_of_users" table.
	ImageOfUsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "royalty", Type: field.TypeFloat32, Nullable: true},
		{Name: "is_image_published", Type: field.TypeBool, Default: false},
		{Name: "goods_image_of_user", Type: field.TypeUUID, Unique: true, Nullable: true},
		{Name: "image_of_user_user", Type: field.TypeUUID},
	}
	// ImageOfUsersTable holds the schema information for the "image_of_users" table.
	ImageOfUsersTable = &schema.Table{
		Name:       "image_of_users",
		Columns:    ImageOfUsersColumns,
		PrimaryKey: []*schema.Column{ImageOfUsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "image_of_users_goods_image_of_user",
				Columns:    []*schema.Column{ImageOfUsersColumns[6]},
				RefColumns: []*schema.Column{GoodsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "image_of_users_users_user",
				Columns:    []*schema.Column{ImageOfUsersColumns[7]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// ItemHasOwnersColumns holds the columns for the "item_has_owners" table.
	ItemHasOwnersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "item_has_owner_product_quote", Type: field.TypeUUID},
		{Name: "item_has_owner_user", Type: field.TypeUUID},
	}
	// ItemHasOwnersTable holds the schema information for the "item_has_owners" table.
	ItemHasOwnersTable = &schema.Table{
		Name:       "item_has_owners",
		Columns:    ItemHasOwnersColumns,
		PrimaryKey: []*schema.Column{ItemHasOwnersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "item_has_owners_product_quotes_product_quote",
				Columns:    []*schema.Column{ItemHasOwnersColumns[4]},
				RefColumns: []*schema.Column{ProductQuotesColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "item_has_owners_users_user",
				Columns:    []*schema.Column{ItemHasOwnersColumns[5]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// OrdersColumns holds the columns for the "orders" table.
	OrdersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "name", Type: field.TypeString},
		{Name: "address", Type: field.TypeString},
		{Name: "phone", Type: field.TypeString},
		{Name: "email", Type: field.TypeString, Nullable: true},
		{Name: "order_item_has_owner", Type: field.TypeUUID},
	}
	// OrdersTable holds the schema information for the "orders" table.
	OrdersTable = &schema.Table{
		Name:       "orders",
		Columns:    OrdersColumns,
		PrimaryKey: []*schema.Column{OrdersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "orders_item_has_owners_item_has_owner",
				Columns:    []*schema.Column{OrdersColumns[8]},
				RefColumns: []*schema.Column{ItemHasOwnersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// ProductQuotesColumns holds the columns for the "product_quotes" table.
	ProductQuotesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "price", Type: field.TypeInt},
		{Name: "is_active", Type: field.TypeBool, Default: false},
		{Name: "product_quote_user", Type: field.TypeUUID},
		{Name: "product_quote_goods", Type: field.TypeUUID},
		{Name: "product_quote_image_of_user", Type: field.TypeUUID},
	}
	// ProductQuotesTable holds the schema information for the "product_quotes" table.
	ProductQuotesTable = &schema.Table{
		Name:       "product_quotes",
		Columns:    ProductQuotesColumns,
		PrimaryKey: []*schema.Column{ProductQuotesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "product_quotes_users_user",
				Columns:    []*schema.Column{ProductQuotesColumns[6]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "product_quotes_goods_goods",
				Columns:    []*schema.Column{ProductQuotesColumns[7]},
				RefColumns: []*schema.Column{GoodsColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "product_quotes_image_of_users_image_of_user",
				Columns:    []*schema.Column{ProductQuotesColumns[8]},
				RefColumns: []*schema.Column{ImageOfUsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// TypeOfGoodsColumns holds the columns for the "type_of_goods" table.
	TypeOfGoodsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "name", Type: field.TypeString},
		{Name: "type", Type: field.TypeString},
		{Name: "description", Type: field.TypeString},
		{Name: "goods_type_of_goods", Type: field.TypeUUID, Unique: true, Nullable: true},
	}
	// TypeOfGoodsTable holds the schema information for the "type_of_goods" table.
	TypeOfGoodsTable = &schema.Table{
		Name:       "type_of_goods",
		Columns:    TypeOfGoodsColumns,
		PrimaryKey: []*schema.Column{TypeOfGoodsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "type_of_goods_goods_type_of_goods",
				Columns:    []*schema.Column{TypeOfGoodsColumns[7]},
				RefColumns: []*schema.Column{GoodsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "username", Type: field.TypeString, Unique: true},
		{Name: "is_active", Type: field.TypeBool, Default: true},
		{Name: "hashed_pwd", Type: field.TypeString},
		{Name: "is_superuser", Type: field.TypeBool, Default: false},
		{Name: "name", Type: field.TypeString},
		{Name: "email", Type: field.TypeString, Unique: true, Size: 2147483647},
		{Name: "phone", Type: field.TypeString, Nullable: true, Size: 2147483647},
		{Name: "address", Type: field.TypeString, Nullable: true, Size: 2147483647},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		GoodsTable,
		ImagesTable,
		ImageOfUsersTable,
		ItemHasOwnersTable,
		OrdersTable,
		ProductQuotesTable,
		TypeOfGoodsTable,
		UsersTable,
	}
)

func init() {
	ImagesTable.ForeignKeys[0].RefTable = ImageOfUsersTable
	ImageOfUsersTable.ForeignKeys[0].RefTable = GoodsTable
	ImageOfUsersTable.ForeignKeys[1].RefTable = UsersTable
	ItemHasOwnersTable.ForeignKeys[0].RefTable = ProductQuotesTable
	ItemHasOwnersTable.ForeignKeys[1].RefTable = UsersTable
	OrdersTable.ForeignKeys[0].RefTable = ItemHasOwnersTable
	ProductQuotesTable.ForeignKeys[0].RefTable = UsersTable
	ProductQuotesTable.ForeignKeys[1].RefTable = GoodsTable
	ProductQuotesTable.ForeignKeys[2].RefTable = ImageOfUsersTable
	TypeOfGoodsTable.ForeignKeys[0].RefTable = GoodsTable
}