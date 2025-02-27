// Code generated by entc, DO NOT EDIT.

package ent

import (
	"comment/app/service/internal/data/ent/commentsubject"
	"comment/app/service/internal/data/ent/predicate"
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CommentSubjectDelete is the builder for deleting a CommentSubject entity.
type CommentSubjectDelete struct {
	config
	hooks    []Hook
	mutation *CommentSubjectMutation
}

// Where appends a list predicates to the CommentSubjectDelete builder.
func (csd *CommentSubjectDelete) Where(ps ...predicate.CommentSubject) *CommentSubjectDelete {
	csd.mutation.Where(ps...)
	return csd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (csd *CommentSubjectDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(csd.hooks) == 0 {
		affected, err = csd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CommentSubjectMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			csd.mutation = mutation
			affected, err = csd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(csd.hooks) - 1; i >= 0; i-- {
			if csd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = csd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, csd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (csd *CommentSubjectDelete) ExecX(ctx context.Context) int {
	n, err := csd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (csd *CommentSubjectDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: commentsubject.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: commentsubject.FieldID,
			},
		},
	}
	if ps := csd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, csd.driver, _spec)
}

// CommentSubjectDeleteOne is the builder for deleting a single CommentSubject entity.
type CommentSubjectDeleteOne struct {
	csd *CommentSubjectDelete
}

// Exec executes the deletion query.
func (csdo *CommentSubjectDeleteOne) Exec(ctx context.Context) error {
	n, err := csdo.csd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{commentsubject.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (csdo *CommentSubjectDeleteOne) ExecX(ctx context.Context) {
	csdo.csd.ExecX(ctx)
}
