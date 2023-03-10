// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/Hyeu249/noname-ent-schema/ent/goods"
	"github.com/Hyeu249/noname-ent-schema/ent/image"
	"github.com/Hyeu249/noname-ent-schema/ent/imageofuser"
	"github.com/Hyeu249/noname-ent-schema/ent/itemhasowner"
	"github.com/Hyeu249/noname-ent-schema/ent/order"
	"github.com/Hyeu249/noname-ent-schema/ent/productquote"
	"github.com/Hyeu249/noname-ent-schema/ent/schema"
	"github.com/Hyeu249/noname-ent-schema/ent/typeofgoods"
	"github.com/Hyeu249/noname-ent-schema/ent/user"
	"github.com/google/uuid"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	goodsMixin := schema.Goods{}.Mixin()
	goodsMixinFields0 := goodsMixin[0].Fields()
	_ = goodsMixinFields0
	goodsFields := schema.Goods{}.Fields()
	_ = goodsFields
	// goodsDescCreateTime is the schema descriptor for create_time field.
	goodsDescCreateTime := goodsMixinFields0[0].Descriptor()
	// goods.DefaultCreateTime holds the default value on creation for the create_time field.
	goods.DefaultCreateTime = goodsDescCreateTime.Default.(func() time.Time)
	// goodsDescUpdateTime is the schema descriptor for update_time field.
	goodsDescUpdateTime := goodsMixinFields0[1].Descriptor()
	// goods.DefaultUpdateTime holds the default value on creation for the update_time field.
	goods.DefaultUpdateTime = goodsDescUpdateTime.Default.(func() time.Time)
	// goods.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	goods.UpdateDefaultUpdateTime = goodsDescUpdateTime.UpdateDefault.(func() time.Time)
	// goodsDescID is the schema descriptor for id field.
	goodsDescID := goodsFields[0].Descriptor()
	// goods.DefaultID holds the default value on creation for the id field.
	goods.DefaultID = goodsDescID.Default.(func() uuid.UUID)
	imageMixin := schema.Image{}.Mixin()
	imageMixinFields0 := imageMixin[0].Fields()
	_ = imageMixinFields0
	imageFields := schema.Image{}.Fields()
	_ = imageFields
	// imageDescCreateTime is the schema descriptor for create_time field.
	imageDescCreateTime := imageMixinFields0[0].Descriptor()
	// image.DefaultCreateTime holds the default value on creation for the create_time field.
	image.DefaultCreateTime = imageDescCreateTime.Default.(func() time.Time)
	// imageDescUpdateTime is the schema descriptor for update_time field.
	imageDescUpdateTime := imageMixinFields0[1].Descriptor()
	// image.DefaultUpdateTime holds the default value on creation for the update_time field.
	image.DefaultUpdateTime = imageDescUpdateTime.Default.(func() time.Time)
	// image.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	image.UpdateDefaultUpdateTime = imageDescUpdateTime.UpdateDefault.(func() time.Time)
	// imageDescFileExtension is the schema descriptor for file_extension field.
	imageDescFileExtension := imageFields[2].Descriptor()
	// image.FileExtensionValidator is a validator for the "file_extension" field. It is called by the builders before save.
	image.FileExtensionValidator = imageDescFileExtension.Validators[0].(func(string) error)
	// imageDescLocation is the schema descriptor for location field.
	imageDescLocation := imageFields[3].Descriptor()
	// image.LocationValidator is a validator for the "location" field. It is called by the builders before save.
	image.LocationValidator = imageDescLocation.Validators[0].(func(string) error)
	// imageDescID is the schema descriptor for id field.
	imageDescID := imageFields[0].Descriptor()
	// image.DefaultID holds the default value on creation for the id field.
	image.DefaultID = imageDescID.Default.(func() uuid.UUID)
	imageofuserMixin := schema.ImageOfUser{}.Mixin()
	imageofuserMixinFields0 := imageofuserMixin[0].Fields()
	_ = imageofuserMixinFields0
	imageofuserFields := schema.ImageOfUser{}.Fields()
	_ = imageofuserFields
	// imageofuserDescCreateTime is the schema descriptor for create_time field.
	imageofuserDescCreateTime := imageofuserMixinFields0[0].Descriptor()
	// imageofuser.DefaultCreateTime holds the default value on creation for the create_time field.
	imageofuser.DefaultCreateTime = imageofuserDescCreateTime.Default.(func() time.Time)
	// imageofuserDescUpdateTime is the schema descriptor for update_time field.
	imageofuserDescUpdateTime := imageofuserMixinFields0[1].Descriptor()
	// imageofuser.DefaultUpdateTime holds the default value on creation for the update_time field.
	imageofuser.DefaultUpdateTime = imageofuserDescUpdateTime.Default.(func() time.Time)
	// imageofuser.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	imageofuser.UpdateDefaultUpdateTime = imageofuserDescUpdateTime.UpdateDefault.(func() time.Time)
	// imageofuserDescIsImagePublished is the schema descriptor for is_image_published field.
	imageofuserDescIsImagePublished := imageofuserFields[2].Descriptor()
	// imageofuser.DefaultIsImagePublished holds the default value on creation for the is_image_published field.
	imageofuser.DefaultIsImagePublished = imageofuserDescIsImagePublished.Default.(bool)
	// imageofuserDescID is the schema descriptor for id field.
	imageofuserDescID := imageofuserFields[0].Descriptor()
	// imageofuser.DefaultID holds the default value on creation for the id field.
	imageofuser.DefaultID = imageofuserDescID.Default.(func() uuid.UUID)
	itemhasownerMixin := schema.ItemHasOwner{}.Mixin()
	itemhasownerMixinFields0 := itemhasownerMixin[0].Fields()
	_ = itemhasownerMixinFields0
	itemhasownerFields := schema.ItemHasOwner{}.Fields()
	_ = itemhasownerFields
	// itemhasownerDescCreateTime is the schema descriptor for create_time field.
	itemhasownerDescCreateTime := itemhasownerMixinFields0[0].Descriptor()
	// itemhasowner.DefaultCreateTime holds the default value on creation for the create_time field.
	itemhasowner.DefaultCreateTime = itemhasownerDescCreateTime.Default.(func() time.Time)
	// itemhasownerDescUpdateTime is the schema descriptor for update_time field.
	itemhasownerDescUpdateTime := itemhasownerMixinFields0[1].Descriptor()
	// itemhasowner.DefaultUpdateTime holds the default value on creation for the update_time field.
	itemhasowner.DefaultUpdateTime = itemhasownerDescUpdateTime.Default.(func() time.Time)
	// itemhasowner.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	itemhasowner.UpdateDefaultUpdateTime = itemhasownerDescUpdateTime.UpdateDefault.(func() time.Time)
	// itemhasownerDescID is the schema descriptor for id field.
	itemhasownerDescID := itemhasownerFields[0].Descriptor()
	// itemhasowner.DefaultID holds the default value on creation for the id field.
	itemhasowner.DefaultID = itemhasownerDescID.Default.(func() uuid.UUID)
	orderMixin := schema.Order{}.Mixin()
	orderMixinFields0 := orderMixin[0].Fields()
	_ = orderMixinFields0
	orderFields := schema.Order{}.Fields()
	_ = orderFields
	// orderDescCreateTime is the schema descriptor for create_time field.
	orderDescCreateTime := orderMixinFields0[0].Descriptor()
	// order.DefaultCreateTime holds the default value on creation for the create_time field.
	order.DefaultCreateTime = orderDescCreateTime.Default.(func() time.Time)
	// orderDescUpdateTime is the schema descriptor for update_time field.
	orderDescUpdateTime := orderMixinFields0[1].Descriptor()
	// order.DefaultUpdateTime holds the default value on creation for the update_time field.
	order.DefaultUpdateTime = orderDescUpdateTime.Default.(func() time.Time)
	// order.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	order.UpdateDefaultUpdateTime = orderDescUpdateTime.UpdateDefault.(func() time.Time)
	// orderDescName is the schema descriptor for name field.
	orderDescName := orderFields[1].Descriptor()
	// order.NameValidator is a validator for the "name" field. It is called by the builders before save.
	order.NameValidator = orderDescName.Validators[0].(func(string) error)
	// orderDescAddress is the schema descriptor for address field.
	orderDescAddress := orderFields[2].Descriptor()
	// order.AddressValidator is a validator for the "address" field. It is called by the builders before save.
	order.AddressValidator = orderDescAddress.Validators[0].(func(string) error)
	// orderDescPhone is the schema descriptor for phone field.
	orderDescPhone := orderFields[3].Descriptor()
	// order.PhoneValidator is a validator for the "phone" field. It is called by the builders before save.
	order.PhoneValidator = orderDescPhone.Validators[0].(func(string) error)
	// orderDescID is the schema descriptor for id field.
	orderDescID := orderFields[0].Descriptor()
	// order.DefaultID holds the default value on creation for the id field.
	order.DefaultID = orderDescID.Default.(func() uuid.UUID)
	productquoteMixin := schema.ProductQuote{}.Mixin()
	productquoteMixinFields0 := productquoteMixin[0].Fields()
	_ = productquoteMixinFields0
	productquoteFields := schema.ProductQuote{}.Fields()
	_ = productquoteFields
	// productquoteDescCreateTime is the schema descriptor for create_time field.
	productquoteDescCreateTime := productquoteMixinFields0[0].Descriptor()
	// productquote.DefaultCreateTime holds the default value on creation for the create_time field.
	productquote.DefaultCreateTime = productquoteDescCreateTime.Default.(func() time.Time)
	// productquoteDescUpdateTime is the schema descriptor for update_time field.
	productquoteDescUpdateTime := productquoteMixinFields0[1].Descriptor()
	// productquote.DefaultUpdateTime holds the default value on creation for the update_time field.
	productquote.DefaultUpdateTime = productquoteDescUpdateTime.Default.(func() time.Time)
	// productquote.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	productquote.UpdateDefaultUpdateTime = productquoteDescUpdateTime.UpdateDefault.(func() time.Time)
	// productquoteDescIsActive is the schema descriptor for is_active field.
	productquoteDescIsActive := productquoteFields[2].Descriptor()
	// productquote.DefaultIsActive holds the default value on creation for the is_active field.
	productquote.DefaultIsActive = productquoteDescIsActive.Default.(bool)
	// productquoteDescID is the schema descriptor for id field.
	productquoteDescID := productquoteFields[0].Descriptor()
	// productquote.DefaultID holds the default value on creation for the id field.
	productquote.DefaultID = productquoteDescID.Default.(func() uuid.UUID)
	typeofgoodsMixin := schema.TypeOfGoods{}.Mixin()
	typeofgoodsMixinFields0 := typeofgoodsMixin[0].Fields()
	_ = typeofgoodsMixinFields0
	typeofgoodsFields := schema.TypeOfGoods{}.Fields()
	_ = typeofgoodsFields
	// typeofgoodsDescCreateTime is the schema descriptor for create_time field.
	typeofgoodsDescCreateTime := typeofgoodsMixinFields0[0].Descriptor()
	// typeofgoods.DefaultCreateTime holds the default value on creation for the create_time field.
	typeofgoods.DefaultCreateTime = typeofgoodsDescCreateTime.Default.(func() time.Time)
	// typeofgoodsDescUpdateTime is the schema descriptor for update_time field.
	typeofgoodsDescUpdateTime := typeofgoodsMixinFields0[1].Descriptor()
	// typeofgoods.DefaultUpdateTime holds the default value on creation for the update_time field.
	typeofgoods.DefaultUpdateTime = typeofgoodsDescUpdateTime.Default.(func() time.Time)
	// typeofgoods.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	typeofgoods.UpdateDefaultUpdateTime = typeofgoodsDescUpdateTime.UpdateDefault.(func() time.Time)
	// typeofgoodsDescName is the schema descriptor for name field.
	typeofgoodsDescName := typeofgoodsFields[1].Descriptor()
	// typeofgoods.NameValidator is a validator for the "name" field. It is called by the builders before save.
	typeofgoods.NameValidator = typeofgoodsDescName.Validators[0].(func(string) error)
	// typeofgoodsDescType is the schema descriptor for type field.
	typeofgoodsDescType := typeofgoodsFields[2].Descriptor()
	// typeofgoods.TypeValidator is a validator for the "type" field. It is called by the builders before save.
	typeofgoods.TypeValidator = typeofgoodsDescType.Validators[0].(func(string) error)
	// typeofgoodsDescID is the schema descriptor for id field.
	typeofgoodsDescID := typeofgoodsFields[0].Descriptor()
	// typeofgoods.DefaultID holds the default value on creation for the id field.
	typeofgoods.DefaultID = typeofgoodsDescID.Default.(func() uuid.UUID)
	userMixin := schema.User{}.Mixin()
	userMixinFields0 := userMixin[0].Fields()
	_ = userMixinFields0
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescCreateTime is the schema descriptor for create_time field.
	userDescCreateTime := userMixinFields0[0].Descriptor()
	// user.DefaultCreateTime holds the default value on creation for the create_time field.
	user.DefaultCreateTime = userDescCreateTime.Default.(func() time.Time)
	// userDescUpdateTime is the schema descriptor for update_time field.
	userDescUpdateTime := userMixinFields0[1].Descriptor()
	// user.DefaultUpdateTime holds the default value on creation for the update_time field.
	user.DefaultUpdateTime = userDescUpdateTime.Default.(func() time.Time)
	// user.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	user.UpdateDefaultUpdateTime = userDescUpdateTime.UpdateDefault.(func() time.Time)
	// userDescUsername is the schema descriptor for username field.
	userDescUsername := userFields[1].Descriptor()
	// user.UsernameValidator is a validator for the "username" field. It is called by the builders before save.
	user.UsernameValidator = userDescUsername.Validators[0].(func(string) error)
	// userDescIsActive is the schema descriptor for is_active field.
	userDescIsActive := userFields[2].Descriptor()
	// user.DefaultIsActive holds the default value on creation for the is_active field.
	user.DefaultIsActive = userDescIsActive.Default.(bool)
	// userDescHashedPwd is the schema descriptor for hashed_pwd field.
	userDescHashedPwd := userFields[3].Descriptor()
	// user.HashedPwdValidator is a validator for the "hashed_pwd" field. It is called by the builders before save.
	user.HashedPwdValidator = userDescHashedPwd.Validators[0].(func(string) error)
	// userDescIsSuperuser is the schema descriptor for is_superuser field.
	userDescIsSuperuser := userFields[4].Descriptor()
	// user.DefaultIsSuperuser holds the default value on creation for the is_superuser field.
	user.DefaultIsSuperuser = userDescIsSuperuser.Default.(bool)
	// userDescName is the schema descriptor for name field.
	userDescName := userFields[5].Descriptor()
	// user.NameValidator is a validator for the "name" field. It is called by the builders before save.
	user.NameValidator = userDescName.Validators[0].(func(string) error)
	// userDescID is the schema descriptor for id field.
	userDescID := userFields[0].Descriptor()
	// user.DefaultID holds the default value on creation for the id field.
	user.DefaultID = userDescID.Default.(func() uuid.UUID)
}
