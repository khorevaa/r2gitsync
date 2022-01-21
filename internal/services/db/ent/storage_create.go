// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/khorevaa/r2gitsync/internal/services/db/ent/project"
	"github.com/khorevaa/r2gitsync/internal/services/db/ent/storage"
)

// StorageCreate is the builder for creating a Storage entity.
type StorageCreate struct {
	config
	mutation *StorageMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (sc *StorageCreate) SetCreatedAt(t time.Time) *StorageCreate {
	sc.mutation.SetCreatedAt(t)
	return sc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (sc *StorageCreate) SetNillableCreatedAt(t *time.Time) *StorageCreate {
	if t != nil {
		sc.SetCreatedAt(*t)
	}
	return sc
}

// SetUpdatedAt sets the "updated_at" field.
func (sc *StorageCreate) SetUpdatedAt(t time.Time) *StorageCreate {
	sc.mutation.SetUpdatedAt(t)
	return sc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (sc *StorageCreate) SetNillableUpdatedAt(t *time.Time) *StorageCreate {
	if t != nil {
		sc.SetUpdatedAt(*t)
	}
	return sc
}

// SetDeletedAt sets the "deleted_at" field.
func (sc *StorageCreate) SetDeletedAt(t time.Time) *StorageCreate {
	sc.mutation.SetDeletedAt(t)
	return sc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (sc *StorageCreate) SetNillableDeletedAt(t *time.Time) *StorageCreate {
	if t != nil {
		sc.SetDeletedAt(*t)
	}
	return sc
}

// SetConnectionString sets the "connection_string" field.
func (sc *StorageCreate) SetConnectionString(s string) *StorageCreate {
	sc.mutation.SetConnectionString(s)
	return sc
}

// SetDevelop sets the "develop" field.
func (sc *StorageCreate) SetDevelop(b bool) *StorageCreate {
	sc.mutation.SetDevelop(b)
	return sc
}

// SetExtension sets the "extension" field.
func (sc *StorageCreate) SetExtension(s string) *StorageCreate {
	sc.mutation.SetExtension(s)
	return sc
}

// SetNillableExtension sets the "extension" field if the given value is not nil.
func (sc *StorageCreate) SetNillableExtension(s *string) *StorageCreate {
	if s != nil {
		sc.SetExtension(*s)
	}
	return sc
}

// SetType sets the "type" field.
func (sc *StorageCreate) SetType(s storage.Type) *StorageCreate {
	sc.mutation.SetType(s)
	return sc
}

// SetID sets the "id" field.
func (sc *StorageCreate) SetID(u uuid.UUID) *StorageCreate {
	sc.mutation.SetID(u)
	return sc
}

// SetProjectID sets the "project" edge to the Project entity by ID.
func (sc *StorageCreate) SetProjectID(id uuid.UUID) *StorageCreate {
	sc.mutation.SetProjectID(id)
	return sc
}

// SetProject sets the "project" edge to the Project entity.
func (sc *StorageCreate) SetProject(p *Project) *StorageCreate {
	return sc.SetProjectID(p.ID)
}

// SetParentID sets the "parent" edge to the Storage entity by ID.
func (sc *StorageCreate) SetParentID(id uuid.UUID) *StorageCreate {
	sc.mutation.SetParentID(id)
	return sc
}

// SetNillableParentID sets the "parent" edge to the Storage entity by ID if the given value is not nil.
func (sc *StorageCreate) SetNillableParentID(id *uuid.UUID) *StorageCreate {
	if id != nil {
		sc = sc.SetParentID(*id)
	}
	return sc
}

// SetParent sets the "parent" edge to the Storage entity.
func (sc *StorageCreate) SetParent(s *Storage) *StorageCreate {
	return sc.SetParentID(s.ID)
}

// Mutation returns the StorageMutation object of the builder.
func (sc *StorageCreate) Mutation() *StorageMutation {
	return sc.mutation
}

// Save creates the Storage in the database.
func (sc *StorageCreate) Save(ctx context.Context) (*Storage, error) {
	var (
		err  error
		node *Storage
	)
	sc.defaults()
	if len(sc.hooks) == 0 {
		if err = sc.check(); err != nil {
			return nil, err
		}
		node, err = sc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*StorageMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = sc.check(); err != nil {
				return nil, err
			}
			sc.mutation = mutation
			if node, err = sc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(sc.hooks) - 1; i >= 0; i-- {
			if sc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = sc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, sc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (sc *StorageCreate) SaveX(ctx context.Context) *Storage {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sc *StorageCreate) Exec(ctx context.Context) error {
	_, err := sc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sc *StorageCreate) ExecX(ctx context.Context) {
	if err := sc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sc *StorageCreate) defaults() {
	if _, ok := sc.mutation.CreatedAt(); !ok {
		v := storage.DefaultCreatedAt()
		sc.mutation.SetCreatedAt(v)
	}
	if _, ok := sc.mutation.UpdatedAt(); !ok {
		v := storage.DefaultUpdatedAt()
		sc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sc *StorageCreate) check() error {
	if _, ok := sc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Storage.created_at"`)}
	}
	if _, ok := sc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Storage.updated_at"`)}
	}
	if _, ok := sc.mutation.ConnectionString(); !ok {
		return &ValidationError{Name: "connection_string", err: errors.New(`ent: missing required field "Storage.connection_string"`)}
	}
	if _, ok := sc.mutation.Develop(); !ok {
		return &ValidationError{Name: "develop", err: errors.New(`ent: missing required field "Storage.develop"`)}
	}
	if _, ok := sc.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`ent: missing required field "Storage.type"`)}
	}
	if v, ok := sc.mutation.GetType(); ok {
		if err := storage.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "Storage.type": %w`, err)}
		}
	}
	if _, ok := sc.mutation.ProjectID(); !ok {
		return &ValidationError{Name: "project", err: errors.New(`ent: missing required edge "Storage.project"`)}
	}
	return nil
}

func (sc *StorageCreate) sqlSave(ctx context.Context) (*Storage, error) {
	_node, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
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
	return _node, nil
}

func (sc *StorageCreate) createSpec() (*Storage, *sqlgraph.CreateSpec) {
	var (
		_node = &Storage{config: sc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: storage.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: storage.FieldID,
			},
		}
	)
	if id, ok := sc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := sc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: storage.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := sc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: storage.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := sc.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: storage.FieldDeletedAt,
		})
		_node.DeletedAt = &value
	}
	if value, ok := sc.mutation.ConnectionString(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: storage.FieldConnectionString,
		})
		_node.ConnectionString = value
	}
	if value, ok := sc.mutation.Develop(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: storage.FieldDevelop,
		})
		_node.Develop = value
	}
	if value, ok := sc.mutation.Extension(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: storage.FieldExtension,
		})
		_node.Extension = value
	}
	if value, ok := sc.mutation.GetType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: storage.FieldType,
		})
		_node.Type = value
	}
	if nodes := sc.mutation.ProjectIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   storage.ProjectTable,
			Columns: []string{storage.ProjectColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: project.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.project_storages = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sc.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   storage.ParentTable,
			Columns: []string{storage.ParentColumn},
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: storage.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.storage_parent = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// StorageCreateBulk is the builder for creating many Storage entities in bulk.
type StorageCreateBulk struct {
	config
	builders []*StorageCreate
}

// Save creates the Storage entities in the database.
func (scb *StorageCreateBulk) Save(ctx context.Context) ([]*Storage, error) {
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*Storage, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*StorageMutation)
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
					_, err = mutators[i+1].Mutate(root, scb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, scb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
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
		if _, err := mutators[0].Mutate(ctx, scb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (scb *StorageCreateBulk) SaveX(ctx context.Context) []*Storage {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scb *StorageCreateBulk) Exec(ctx context.Context) error {
	_, err := scb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scb *StorageCreateBulk) ExecX(ctx context.Context) {
	if err := scb.Exec(ctx); err != nil {
		panic(err)
	}
}