// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Hyeu249/noname-ent-schema/ent/goods"
	"github.com/Hyeu249/noname-ent-schema/ent/image"
	"github.com/Hyeu249/noname-ent-schema/ent/imageofuser"
	"github.com/Hyeu249/noname-ent-schema/ent/productquote"
	"github.com/Hyeu249/noname-ent-schema/ent/user"
	"github.com/google/uuid"
)

// ImageOfUserCreate is the builder for creating a ImageOfUser entity.
type ImageOfUserCreate struct {
	config
	mutation *ImageOfUserMutation
	hooks    []Hook
}

// SetCreateTime sets the "create_time" field.
func (iouc *ImageOfUserCreate) SetCreateTime(t time.Time) *ImageOfUserCreate {
	iouc.mutation.SetCreateTime(t)
	return iouc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (iouc *ImageOfUserCreate) SetNillableCreateTime(t *time.Time) *ImageOfUserCreate {
	if t != nil {
		iouc.SetCreateTime(*t)
	}
	return iouc
}

// SetUpdateTime sets the "update_time" field.
func (iouc *ImageOfUserCreate) SetUpdateTime(t time.Time) *ImageOfUserCreate {
	iouc.mutation.SetUpdateTime(t)
	return iouc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (iouc *ImageOfUserCreate) SetNillableUpdateTime(t *time.Time) *ImageOfUserCreate {
	if t != nil {
		iouc.SetUpdateTime(*t)
	}
	return iouc
}

// SetDeletedAt sets the "deleted_at" field.
func (iouc *ImageOfUserCreate) SetDeletedAt(t time.Time) *ImageOfUserCreate {
	iouc.mutation.SetDeletedAt(t)
	return iouc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (iouc *ImageOfUserCreate) SetNillableDeletedAt(t *time.Time) *ImageOfUserCreate {
	if t != nil {
		iouc.SetDeletedAt(*t)
	}
	return iouc
}

// SetRoyalty sets the "royalty" field.
func (iouc *ImageOfUserCreate) SetRoyalty(f float32) *ImageOfUserCreate {
	iouc.mutation.SetRoyalty(f)
	return iouc
}

// SetNillableRoyalty sets the "royalty" field if the given value is not nil.
func (iouc *ImageOfUserCreate) SetNillableRoyalty(f *float32) *ImageOfUserCreate {
	if f != nil {
		iouc.SetRoyalty(*f)
	}
	return iouc
}

// SetIsImagePublished sets the "is_image_published" field.
func (iouc *ImageOfUserCreate) SetIsImagePublished(b bool) *ImageOfUserCreate {
	iouc.mutation.SetIsImagePublished(b)
	return iouc
}

// SetNillableIsImagePublished sets the "is_image_published" field if the given value is not nil.
func (iouc *ImageOfUserCreate) SetNillableIsImagePublished(b *bool) *ImageOfUserCreate {
	if b != nil {
		iouc.SetIsImagePublished(*b)
	}
	return iouc
}

// SetID sets the "id" field.
func (iouc *ImageOfUserCreate) SetID(u uuid.UUID) *ImageOfUserCreate {
	iouc.mutation.SetID(u)
	return iouc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (iouc *ImageOfUserCreate) SetNillableID(u *uuid.UUID) *ImageOfUserCreate {
	if u != nil {
		iouc.SetID(*u)
	}
	return iouc
}

// SetImageID sets the "image" edge to the Image entity by ID.
func (iouc *ImageOfUserCreate) SetImageID(id uuid.UUID) *ImageOfUserCreate {
	iouc.mutation.SetImageID(id)
	return iouc
}

// SetImage sets the "image" edge to the Image entity.
func (iouc *ImageOfUserCreate) SetImage(i *Image) *ImageOfUserCreate {
	return iouc.SetImageID(i.ID)
}

// SetUserID sets the "user" edge to the User entity by ID.
func (iouc *ImageOfUserCreate) SetUserID(id uuid.UUID) *ImageOfUserCreate {
	iouc.mutation.SetUserID(id)
	return iouc
}

// SetUser sets the "user" edge to the User entity.
func (iouc *ImageOfUserCreate) SetUser(u *User) *ImageOfUserCreate {
	return iouc.SetUserID(u.ID)
}

// SetGoodsID sets the "goods" edge to the Goods entity by ID.
func (iouc *ImageOfUserCreate) SetGoodsID(id uuid.UUID) *ImageOfUserCreate {
	iouc.mutation.SetGoodsID(id)
	return iouc
}

// SetNillableGoodsID sets the "goods" edge to the Goods entity by ID if the given value is not nil.
func (iouc *ImageOfUserCreate) SetNillableGoodsID(id *uuid.UUID) *ImageOfUserCreate {
	if id != nil {
		iouc = iouc.SetGoodsID(*id)
	}
	return iouc
}

// SetGoods sets the "goods" edge to the Goods entity.
func (iouc *ImageOfUserCreate) SetGoods(g *Goods) *ImageOfUserCreate {
	return iouc.SetGoodsID(g.ID)
}

// AddProductQuoteIDs adds the "product_quote" edge to the ProductQuote entity by IDs.
func (iouc *ImageOfUserCreate) AddProductQuoteIDs(ids ...uuid.UUID) *ImageOfUserCreate {
	iouc.mutation.AddProductQuoteIDs(ids...)
	return iouc
}

// AddProductQuote adds the "product_quote" edges to the ProductQuote entity.
func (iouc *ImageOfUserCreate) AddProductQuote(p ...*ProductQuote) *ImageOfUserCreate {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return iouc.AddProductQuoteIDs(ids...)
}

// Mutation returns the ImageOfUserMutation object of the builder.
func (iouc *ImageOfUserCreate) Mutation() *ImageOfUserMutation {
	return iouc.mutation
}

// Save creates the ImageOfUser in the database.
func (iouc *ImageOfUserCreate) Save(ctx context.Context) (*ImageOfUser, error) {
	iouc.defaults()
	return withHooks[*ImageOfUser, ImageOfUserMutation](ctx, iouc.sqlSave, iouc.mutation, iouc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (iouc *ImageOfUserCreate) SaveX(ctx context.Context) *ImageOfUser {
	v, err := iouc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (iouc *ImageOfUserCreate) Exec(ctx context.Context) error {
	_, err := iouc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iouc *ImageOfUserCreate) ExecX(ctx context.Context) {
	if err := iouc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (iouc *ImageOfUserCreate) defaults() {
	if _, ok := iouc.mutation.CreateTime(); !ok {
		v := imageofuser.DefaultCreateTime()
		iouc.mutation.SetCreateTime(v)
	}
	if _, ok := iouc.mutation.UpdateTime(); !ok {
		v := imageofuser.DefaultUpdateTime()
		iouc.mutation.SetUpdateTime(v)
	}
	if _, ok := iouc.mutation.IsImagePublished(); !ok {
		v := imageofuser.DefaultIsImagePublished
		iouc.mutation.SetIsImagePublished(v)
	}
	if _, ok := iouc.mutation.ID(); !ok {
		v := imageofuser.DefaultID()
		iouc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (iouc *ImageOfUserCreate) check() error {
	if _, ok := iouc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`ent: missing required field "ImageOfUser.create_time"`)}
	}
	if _, ok := iouc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`ent: missing required field "ImageOfUser.update_time"`)}
	}
	if _, ok := iouc.mutation.IsImagePublished(); !ok {
		return &ValidationError{Name: "is_image_published", err: errors.New(`ent: missing required field "ImageOfUser.is_image_published"`)}
	}
	if _, ok := iouc.mutation.ImageID(); !ok {
		return &ValidationError{Name: "image", err: errors.New(`ent: missing required edge "ImageOfUser.image"`)}
	}
	if _, ok := iouc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user", err: errors.New(`ent: missing required edge "ImageOfUser.user"`)}
	}
	return nil
}

func (iouc *ImageOfUserCreate) sqlSave(ctx context.Context) (*ImageOfUser, error) {
	if err := iouc.check(); err != nil {
		return nil, err
	}
	_node, _spec := iouc.createSpec()
	if err := sqlgraph.CreateNode(ctx, iouc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	iouc.mutation.id = &_node.ID
	iouc.mutation.done = true
	return _node, nil
}

func (iouc *ImageOfUserCreate) createSpec() (*ImageOfUser, *sqlgraph.CreateSpec) {
	var (
		_node = &ImageOfUser{config: iouc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: imageofuser.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: imageofuser.FieldID,
			},
		}
	)
	if id, ok := iouc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := iouc.mutation.CreateTime(); ok {
		_spec.SetField(imageofuser.FieldCreateTime, field.TypeTime, value)
		_node.CreateTime = value
	}
	if value, ok := iouc.mutation.UpdateTime(); ok {
		_spec.SetField(imageofuser.FieldUpdateTime, field.TypeTime, value)
		_node.UpdateTime = value
	}
	if value, ok := iouc.mutation.DeletedAt(); ok {
		_spec.SetField(imageofuser.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = &value
	}
	if value, ok := iouc.mutation.Royalty(); ok {
		_spec.SetField(imageofuser.FieldRoyalty, field.TypeFloat32, value)
		_node.Royalty = value
	}
	if value, ok := iouc.mutation.IsImagePublished(); ok {
		_spec.SetField(imageofuser.FieldIsImagePublished, field.TypeBool, value)
		_node.IsImagePublished = value
	}
	if nodes := iouc.mutation.ImageIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   imageofuser.ImageTable,
			Columns: []string{imageofuser.ImageColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: image.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := iouc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   imageofuser.UserTable,
			Columns: []string{imageofuser.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.image_of_user_user = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := iouc.mutation.GoodsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   imageofuser.GoodsTable,
			Columns: []string{imageofuser.GoodsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: goods.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.goods_image_of_user = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := iouc.mutation.ProductQuoteIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   imageofuser.ProductQuoteTable,
			Columns: []string{imageofuser.ProductQuoteColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: productquote.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ImageOfUserCreateBulk is the builder for creating many ImageOfUser entities in bulk.
type ImageOfUserCreateBulk struct {
	config
	builders []*ImageOfUserCreate
}

// Save creates the ImageOfUser entities in the database.
func (ioucb *ImageOfUserCreateBulk) Save(ctx context.Context) ([]*ImageOfUser, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ioucb.builders))
	nodes := make([]*ImageOfUser, len(ioucb.builders))
	mutators := make([]Mutator, len(ioucb.builders))
	for i := range ioucb.builders {
		func(i int, root context.Context) {
			builder := ioucb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ImageOfUserMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ioucb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ioucb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ioucb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ioucb *ImageOfUserCreateBulk) SaveX(ctx context.Context) []*ImageOfUser {
	v, err := ioucb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ioucb *ImageOfUserCreateBulk) Exec(ctx context.Context) error {
	_, err := ioucb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ioucb *ImageOfUserCreateBulk) ExecX(ctx context.Context) {
	if err := ioucb.Exec(ctx); err != nil {
		panic(err)
	}
}
