// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Hyeu249/noname-ent-schema/ent/imageofuser"
	"github.com/Hyeu249/noname-ent-schema/ent/predicate"
)

// ImageOfUserDelete is the builder for deleting a ImageOfUser entity.
type ImageOfUserDelete struct {
	config
	hooks    []Hook
	mutation *ImageOfUserMutation
}

// Where appends a list predicates to the ImageOfUserDelete builder.
func (ioud *ImageOfUserDelete) Where(ps ...predicate.ImageOfUser) *ImageOfUserDelete {
	ioud.mutation.Where(ps...)
	return ioud
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ioud *ImageOfUserDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, ImageOfUserMutation](ctx, ioud.sqlExec, ioud.mutation, ioud.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (ioud *ImageOfUserDelete) ExecX(ctx context.Context) int {
	n, err := ioud.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ioud *ImageOfUserDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: imageofuser.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: imageofuser.FieldID,
			},
		},
	}
	if ps := ioud.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, ioud.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	ioud.mutation.done = true
	return affected, err
}

// ImageOfUserDeleteOne is the builder for deleting a single ImageOfUser entity.
type ImageOfUserDeleteOne struct {
	ioud *ImageOfUserDelete
}

// Where appends a list predicates to the ImageOfUserDelete builder.
func (ioudo *ImageOfUserDeleteOne) Where(ps ...predicate.ImageOfUser) *ImageOfUserDeleteOne {
	ioudo.ioud.mutation.Where(ps...)
	return ioudo
}

// Exec executes the deletion query.
func (ioudo *ImageOfUserDeleteOne) Exec(ctx context.Context) error {
	n, err := ioudo.ioud.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{imageofuser.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ioudo *ImageOfUserDeleteOne) ExecX(ctx context.Context) {
	if err := ioudo.Exec(ctx); err != nil {
		panic(err)
	}
}
