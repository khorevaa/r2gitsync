// Code generated by entc, DO NOT EDIT.

package db

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/khorevaa/r2gitsync/internal/services/db/plugin"
	"github.com/khorevaa/r2gitsync/internal/services/db/pluginversion"
	"github.com/khorevaa/r2gitsync/internal/services/db/pluginversionproperty"
)

// PluginVersionPropertyCreate is the builder for creating a PluginVersionProperty entity.
type PluginVersionPropertyCreate struct {
	config
	mutation *PluginVersionPropertyMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (pvpc *PluginVersionPropertyCreate) SetCreatedAt(t time.Time) *PluginVersionPropertyCreate {
	pvpc.mutation.SetCreatedAt(t)
	return pvpc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (pvpc *PluginVersionPropertyCreate) SetNillableCreatedAt(t *time.Time) *PluginVersionPropertyCreate {
	if t != nil {
		pvpc.SetCreatedAt(*t)
	}
	return pvpc
}

// SetUpdatedAt sets the "updated_at" field.
func (pvpc *PluginVersionPropertyCreate) SetUpdatedAt(t time.Time) *PluginVersionPropertyCreate {
	pvpc.mutation.SetUpdatedAt(t)
	return pvpc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (pvpc *PluginVersionPropertyCreate) SetNillableUpdatedAt(t *time.Time) *PluginVersionPropertyCreate {
	if t != nil {
		pvpc.SetUpdatedAt(*t)
	}
	return pvpc
}

// SetDeletedAt sets the "deleted_at" field.
func (pvpc *PluginVersionPropertyCreate) SetDeletedAt(t time.Time) *PluginVersionPropertyCreate {
	pvpc.mutation.SetDeletedAt(t)
	return pvpc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (pvpc *PluginVersionPropertyCreate) SetNillableDeletedAt(t *time.Time) *PluginVersionPropertyCreate {
	if t != nil {
		pvpc.SetDeletedAt(*t)
	}
	return pvpc
}

// SetName sets the "name" field.
func (pvpc *PluginVersionPropertyCreate) SetName(s string) *PluginVersionPropertyCreate {
	pvpc.mutation.SetName(s)
	return pvpc
}

// SetDefault sets the "default" field.
func (pvpc *PluginVersionPropertyCreate) SetDefault(s string) *PluginVersionPropertyCreate {
	pvpc.mutation.SetDefault(s)
	return pvpc
}

// SetRequired sets the "required" field.
func (pvpc *PluginVersionPropertyCreate) SetRequired(b bool) *PluginVersionPropertyCreate {
	pvpc.mutation.SetRequired(b)
	return pvpc
}

// SetType sets the "type" field.
func (pvpc *PluginVersionPropertyCreate) SetType(pl pluginversionproperty.Type) *PluginVersionPropertyCreate {
	pvpc.mutation.SetType(pl)
	return pvpc
}

// SetID sets the "id" field.
func (pvpc *PluginVersionPropertyCreate) SetID(u uuid.UUID) *PluginVersionPropertyCreate {
	pvpc.mutation.SetID(u)
	return pvpc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (pvpc *PluginVersionPropertyCreate) SetNillableID(u *uuid.UUID) *PluginVersionPropertyCreate {
	if u != nil {
		pvpc.SetID(*u)
	}
	return pvpc
}

// SetPluginID sets the "plugin" edge to the Plugin entity by ID.
func (pvpc *PluginVersionPropertyCreate) SetPluginID(id uuid.UUID) *PluginVersionPropertyCreate {
	pvpc.mutation.SetPluginID(id)
	return pvpc
}

// SetPlugin sets the "plugin" edge to the Plugin entity.
func (pvpc *PluginVersionPropertyCreate) SetPlugin(p *Plugin) *PluginVersionPropertyCreate {
	return pvpc.SetPluginID(p.ID)
}

// SetVersionID sets the "version" edge to the PluginVersion entity by ID.
func (pvpc *PluginVersionPropertyCreate) SetVersionID(id uuid.UUID) *PluginVersionPropertyCreate {
	pvpc.mutation.SetVersionID(id)
	return pvpc
}

// SetVersion sets the "version" edge to the PluginVersion entity.
func (pvpc *PluginVersionPropertyCreate) SetVersion(p *PluginVersion) *PluginVersionPropertyCreate {
	return pvpc.SetVersionID(p.ID)
}

// Mutation returns the PluginVersionPropertyMutation object of the builder.
func (pvpc *PluginVersionPropertyCreate) Mutation() *PluginVersionPropertyMutation {
	return pvpc.mutation
}

// Save creates the PluginVersionProperty in the database.
func (pvpc *PluginVersionPropertyCreate) Save(ctx context.Context) (*PluginVersionProperty, error) {
	var (
		err  error
		node *PluginVersionProperty
	)
	pvpc.defaults()
	if len(pvpc.hooks) == 0 {
		if err = pvpc.check(); err != nil {
			return nil, err
		}
		node, err = pvpc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PluginVersionPropertyMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = pvpc.check(); err != nil {
				return nil, err
			}
			pvpc.mutation = mutation
			if node, err = pvpc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(pvpc.hooks) - 1; i >= 0; i-- {
			if pvpc.hooks[i] == nil {
				return nil, fmt.Errorf("db: uninitialized hook (forgotten import db/runtime?)")
			}
			mut = pvpc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pvpc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (pvpc *PluginVersionPropertyCreate) SaveX(ctx context.Context) *PluginVersionProperty {
	v, err := pvpc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pvpc *PluginVersionPropertyCreate) Exec(ctx context.Context) error {
	_, err := pvpc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pvpc *PluginVersionPropertyCreate) ExecX(ctx context.Context) {
	if err := pvpc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pvpc *PluginVersionPropertyCreate) defaults() {
	if _, ok := pvpc.mutation.CreatedAt(); !ok {
		v := pluginversionproperty.DefaultCreatedAt()
		pvpc.mutation.SetCreatedAt(v)
	}
	if _, ok := pvpc.mutation.UpdatedAt(); !ok {
		v := pluginversionproperty.DefaultUpdatedAt()
		pvpc.mutation.SetUpdatedAt(v)
	}
	if _, ok := pvpc.mutation.ID(); !ok {
		v := pluginversionproperty.DefaultID()
		pvpc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pvpc *PluginVersionPropertyCreate) check() error {
	if _, ok := pvpc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`db: missing required field "PluginVersionProperty.created_at"`)}
	}
	if _, ok := pvpc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`db: missing required field "PluginVersionProperty.updated_at"`)}
	}
	if _, ok := pvpc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`db: missing required field "PluginVersionProperty.name"`)}
	}
	if v, ok := pvpc.mutation.Name(); ok {
		if err := pluginversionproperty.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`db: validator failed for field "PluginVersionProperty.name": %w`, err)}
		}
	}
	if _, ok := pvpc.mutation.Default(); !ok {
		return &ValidationError{Name: "default", err: errors.New(`db: missing required field "PluginVersionProperty.default"`)}
	}
	if v, ok := pvpc.mutation.Default(); ok {
		if err := pluginversionproperty.DefaultValidator(v); err != nil {
			return &ValidationError{Name: "default", err: fmt.Errorf(`db: validator failed for field "PluginVersionProperty.default": %w`, err)}
		}
	}
	if _, ok := pvpc.mutation.Required(); !ok {
		return &ValidationError{Name: "required", err: errors.New(`db: missing required field "PluginVersionProperty.required"`)}
	}
	if _, ok := pvpc.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`db: missing required field "PluginVersionProperty.type"`)}
	}
	if v, ok := pvpc.mutation.GetType(); ok {
		if err := pluginversionproperty.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`db: validator failed for field "PluginVersionProperty.type": %w`, err)}
		}
	}
	if _, ok := pvpc.mutation.PluginID(); !ok {
		return &ValidationError{Name: "plugin", err: errors.New(`db: missing required edge "PluginVersionProperty.plugin"`)}
	}
	if _, ok := pvpc.mutation.VersionID(); !ok {
		return &ValidationError{Name: "version", err: errors.New(`db: missing required edge "PluginVersionProperty.version"`)}
	}
	return nil
}

func (pvpc *PluginVersionPropertyCreate) sqlSave(ctx context.Context) (*PluginVersionProperty, error) {
	_node, _spec := pvpc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pvpc.driver, _spec); err != nil {
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

func (pvpc *PluginVersionPropertyCreate) createSpec() (*PluginVersionProperty, *sqlgraph.CreateSpec) {
	var (
		_node = &PluginVersionProperty{config: pvpc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: pluginversionproperty.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: pluginversionproperty.FieldID,
			},
		}
	)
	if id, ok := pvpc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := pvpc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: pluginversionproperty.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := pvpc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: pluginversionproperty.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := pvpc.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: pluginversionproperty.FieldDeletedAt,
		})
		_node.DeletedAt = &value
	}
	if value, ok := pvpc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: pluginversionproperty.FieldName,
		})
		_node.Name = value
	}
	if value, ok := pvpc.mutation.Default(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: pluginversionproperty.FieldDefault,
		})
		_node.Default = value
	}
	if value, ok := pvpc.mutation.Required(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: pluginversionproperty.FieldRequired,
		})
		_node.Required = value
	}
	if value, ok := pvpc.mutation.GetType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: pluginversionproperty.FieldType,
		})
		_node.Type = value
	}
	if nodes := pvpc.mutation.PluginIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   pluginversionproperty.PluginTable,
			Columns: []string{pluginversionproperty.PluginColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: plugin.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.plugin_version_property_plugin = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pvpc.mutation.VersionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   pluginversionproperty.VersionTable,
			Columns: []string{pluginversionproperty.VersionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: pluginversion.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.plugin_version_property_version = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// PluginVersionPropertyCreateBulk is the builder for creating many PluginVersionProperty entities in bulk.
type PluginVersionPropertyCreateBulk struct {
	config
	builders []*PluginVersionPropertyCreate
}

// Save creates the PluginVersionProperty entities in the database.
func (pvpcb *PluginVersionPropertyCreateBulk) Save(ctx context.Context) ([]*PluginVersionProperty, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pvpcb.builders))
	nodes := make([]*PluginVersionProperty, len(pvpcb.builders))
	mutators := make([]Mutator, len(pvpcb.builders))
	for i := range pvpcb.builders {
		func(i int, root context.Context) {
			builder := pvpcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PluginVersionPropertyMutation)
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
					_, err = mutators[i+1].Mutate(root, pvpcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pvpcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, pvpcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pvpcb *PluginVersionPropertyCreateBulk) SaveX(ctx context.Context) []*PluginVersionProperty {
	v, err := pvpcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pvpcb *PluginVersionPropertyCreateBulk) Exec(ctx context.Context) error {
	_, err := pvpcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pvpcb *PluginVersionPropertyCreateBulk) ExecX(ctx context.Context) {
	if err := pvpcb.Exec(ctx); err != nil {
		panic(err)
	}
}