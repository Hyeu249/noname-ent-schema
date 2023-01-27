// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Hyeu249/noname-ent-schema/ent/itemhasowner"
	"github.com/Hyeu249/noname-ent-schema/ent/predicate"
)

// ItemHasOwnerDelete is the builder for deleting a ItemHasOwner entity.
type ItemHasOwnerDelete struct {
	config
	hooks    []Hook
	mutation *ItemHasOwnerMutation
}

// Where appends a list predicates to the ItemHasOwnerDelete builder.
func (ihod *ItemHasOwnerDelete) Where(ps ...predicate.ItemHasOwner) *ItemHasOwnerDelete {
	ihod.mutation.Where(ps...)
	return ihod
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ihod *ItemHasOwnerDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, ItemHasOwnerMutation](ctx, ihod.sqlExec, ihod.mutation, ihod.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (ihod *ItemHasOwnerDelete) ExecX(ctx context.Context) int {
	n, err := ihod.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ihod *ItemHasOwnerDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: itemhasowner.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: itemhasowner.FieldID,
			},
		},
	}
	if ps := ihod.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, ihod.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	ihod.mutation.done = true
	return affected, err
}

// ItemHasOwnerDeleteOne is the builder for deleting a single ItemHasOwner entity.
type ItemHasOwnerDeleteOne struct {
	ihod *ItemHasOwnerDelete
}

// Where appends a list predicates to the ItemHasOwnerDelete builder.
func (ihodo *ItemHasOwnerDeleteOne) Where(ps ...predicate.ItemHasOwner) *ItemHasOwnerDeleteOne {
	ihodo.ihod.mutation.Where(ps...)
	return ihodo
}

// Exec executes the deletion query.
func (ihodo *ItemHasOwnerDeleteOne) Exec(ctx context.Context) error {
	n, err := ihodo.ihod.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{itemhasowner.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ihodo *ItemHasOwnerDeleteOne) ExecX(ctx context.Context) {
	if err := ihodo.Exec(ctx); err != nil {
		panic(err)
	}
}