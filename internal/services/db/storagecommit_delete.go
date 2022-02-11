// Code generated by entc, DO NOT EDIT.

package db

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/khorevaa/r2gitsync/internal/services/db/predicate"
	"github.com/khorevaa/r2gitsync/internal/services/db/storagecommit"
)

// StorageCommitDelete is the builder for deleting a StorageCommit entity.
type StorageCommitDelete struct {
	config
	hooks    []Hook
	mutation *StorageCommitMutation
}

// Where appends a list predicates to the StorageCommitDelete builder.
func (scd *StorageCommitDelete) Where(ps ...predicate.StorageCommit) *StorageCommitDelete {
	scd.mutation.Where(ps...)
	return scd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (scd *StorageCommitDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(scd.hooks) == 0 {
		affected, err = scd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*StorageCommitMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			scd.mutation = mutation
			affected, err = scd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(scd.hooks) - 1; i >= 0; i-- {
			if scd.hooks[i] == nil {
				return 0, fmt.Errorf("db: uninitialized hook (forgotten import db/runtime?)")
			}
			mut = scd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, scd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (scd *StorageCommitDelete) ExecX(ctx context.Context) int {
	n, err := scd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (scd *StorageCommitDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: storagecommit.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: storagecommit.FieldID,
			},
		},
	}
	if ps := scd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, scd.driver, _spec)
}

// StorageCommitDeleteOne is the builder for deleting a single StorageCommit entity.
type StorageCommitDeleteOne struct {
	scd *StorageCommitDelete
}

// Exec executes the deletion query.
func (scdo *StorageCommitDeleteOne) Exec(ctx context.Context) error {
	n, err := scdo.scd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{storagecommit.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (scdo *StorageCommitDeleteOne) ExecX(ctx context.Context) {
	scdo.scd.ExecX(ctx)
}