// Code generated by entc, DO NOT EDIT.

package db

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/khorevaa/r2gitsync/internal/services/db/predicate"
	"github.com/khorevaa/r2gitsync/internal/services/db/project"
	"github.com/khorevaa/r2gitsync/internal/services/db/storage"
)

// ProjectUpdate is the builder for updating Project entities.
type ProjectUpdate struct {
	config
	hooks    []Hook
	mutation *ProjectMutation
}

// Where appends a list predicates to the ProjectUpdate builder.
func (pu *ProjectUpdate) Where(ps ...predicate.Project) *ProjectUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetUpdatedAt sets the "updated_at" field.
func (pu *ProjectUpdate) SetUpdatedAt(t time.Time) *ProjectUpdate {
	pu.mutation.SetUpdatedAt(t)
	return pu
}

// SetDeletedAt sets the "deleted_at" field.
func (pu *ProjectUpdate) SetDeletedAt(t time.Time) *ProjectUpdate {
	pu.mutation.SetDeletedAt(t)
	return pu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (pu *ProjectUpdate) SetNillableDeletedAt(t *time.Time) *ProjectUpdate {
	if t != nil {
		pu.SetDeletedAt(*t)
	}
	return pu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (pu *ProjectUpdate) ClearDeletedAt() *ProjectUpdate {
	pu.mutation.ClearDeletedAt()
	return pu
}

// SetCode sets the "code" field.
func (pu *ProjectUpdate) SetCode(s string) *ProjectUpdate {
	pu.mutation.SetCode(s)
	return pu
}

// SetName sets the "name" field.
func (pu *ProjectUpdate) SetName(s string) *ProjectUpdate {
	pu.mutation.SetName(s)
	return pu
}

// SetDescription sets the "Description" field.
func (pu *ProjectUpdate) SetDescription(s string) *ProjectUpdate {
	pu.mutation.SetDescription(s)
	return pu
}

// SetType sets the "type" field.
func (pu *ProjectUpdate) SetType(pr project.Type) *ProjectUpdate {
	pu.mutation.SetType(pr)
	return pu
}

// AddStorageIDs adds the "storages" edge to the Storage entity by IDs.
func (pu *ProjectUpdate) AddStorageIDs(ids ...uuid.UUID) *ProjectUpdate {
	pu.mutation.AddStorageIDs(ids...)
	return pu
}

// AddStorages adds the "storages" edges to the Storage entity.
func (pu *ProjectUpdate) AddStorages(s ...*Storage) *ProjectUpdate {
	ids := make([]uuid.UUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return pu.AddStorageIDs(ids...)
}

// SetMasterStorageID sets the "master_storage" edge to the Storage entity by ID.
func (pu *ProjectUpdate) SetMasterStorageID(id uuid.UUID) *ProjectUpdate {
	pu.mutation.SetMasterStorageID(id)
	return pu
}

// SetNillableMasterStorageID sets the "master_storage" edge to the Storage entity by ID if the given value is not nil.
func (pu *ProjectUpdate) SetNillableMasterStorageID(id *uuid.UUID) *ProjectUpdate {
	if id != nil {
		pu = pu.SetMasterStorageID(*id)
	}
	return pu
}

// SetMasterStorage sets the "master_storage" edge to the Storage entity.
func (pu *ProjectUpdate) SetMasterStorage(s *Storage) *ProjectUpdate {
	return pu.SetMasterStorageID(s.ID)
}

// SetDevelopStorageID sets the "develop_storage" edge to the Storage entity by ID.
func (pu *ProjectUpdate) SetDevelopStorageID(id uuid.UUID) *ProjectUpdate {
	pu.mutation.SetDevelopStorageID(id)
	return pu
}

// SetNillableDevelopStorageID sets the "develop_storage" edge to the Storage entity by ID if the given value is not nil.
func (pu *ProjectUpdate) SetNillableDevelopStorageID(id *uuid.UUID) *ProjectUpdate {
	if id != nil {
		pu = pu.SetDevelopStorageID(*id)
	}
	return pu
}

// SetDevelopStorage sets the "develop_storage" edge to the Storage entity.
func (pu *ProjectUpdate) SetDevelopStorage(s *Storage) *ProjectUpdate {
	return pu.SetDevelopStorageID(s.ID)
}

// Mutation returns the ProjectMutation object of the builder.
func (pu *ProjectUpdate) Mutation() *ProjectMutation {
	return pu.mutation
}

// ClearStorages clears all "storages" edges to the Storage entity.
func (pu *ProjectUpdate) ClearStorages() *ProjectUpdate {
	pu.mutation.ClearStorages()
	return pu
}

// RemoveStorageIDs removes the "storages" edge to Storage entities by IDs.
func (pu *ProjectUpdate) RemoveStorageIDs(ids ...uuid.UUID) *ProjectUpdate {
	pu.mutation.RemoveStorageIDs(ids...)
	return pu
}

// RemoveStorages removes "storages" edges to Storage entities.
func (pu *ProjectUpdate) RemoveStorages(s ...*Storage) *ProjectUpdate {
	ids := make([]uuid.UUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return pu.RemoveStorageIDs(ids...)
}

// ClearMasterStorage clears the "master_storage" edge to the Storage entity.
func (pu *ProjectUpdate) ClearMasterStorage() *ProjectUpdate {
	pu.mutation.ClearMasterStorage()
	return pu
}

// ClearDevelopStorage clears the "develop_storage" edge to the Storage entity.
func (pu *ProjectUpdate) ClearDevelopStorage() *ProjectUpdate {
	pu.mutation.ClearDevelopStorage()
	return pu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *ProjectUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	pu.defaults()
	if len(pu.hooks) == 0 {
		if err = pu.check(); err != nil {
			return 0, err
		}
		affected, err = pu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ProjectMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = pu.check(); err != nil {
				return 0, err
			}
			pu.mutation = mutation
			affected, err = pu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(pu.hooks) - 1; i >= 0; i-- {
			if pu.hooks[i] == nil {
				return 0, fmt.Errorf("db: uninitialized hook (forgotten import db/runtime?)")
			}
			mut = pu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (pu *ProjectUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *ProjectUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *ProjectUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pu *ProjectUpdate) defaults() {
	if _, ok := pu.mutation.UpdatedAt(); !ok {
		v := project.UpdateDefaultUpdatedAt()
		pu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pu *ProjectUpdate) check() error {
	if v, ok := pu.mutation.Code(); ok {
		if err := project.CodeValidator(v); err != nil {
			return &ValidationError{Name: "code", err: fmt.Errorf(`db: validator failed for field "Project.code": %w`, err)}
		}
	}
	if v, ok := pu.mutation.GetType(); ok {
		if err := project.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`db: validator failed for field "Project.type": %w`, err)}
		}
	}
	return nil
}

func (pu *ProjectUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   project.Table,
			Columns: project.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: project.FieldID,
			},
		},
	}
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: project.FieldUpdatedAt,
		})
	}
	if value, ok := pu.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: project.FieldDeletedAt,
		})
	}
	if pu.mutation.DeletedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: project.FieldDeletedAt,
		})
	}
	if value, ok := pu.mutation.Code(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: project.FieldCode,
		})
	}
	if value, ok := pu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: project.FieldName,
		})
	}
	if value, ok := pu.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: project.FieldDescription,
		})
	}
	if value, ok := pu.mutation.GetType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: project.FieldType,
		})
	}
	if pu.mutation.StoragesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   project.StoragesTable,
			Columns: []string{project.StoragesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: storage.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.RemovedStoragesIDs(); len(nodes) > 0 && !pu.mutation.StoragesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   project.StoragesTable,
			Columns: []string{project.StoragesColumn},
			Bidi:    false,
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.StoragesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   project.StoragesTable,
			Columns: []string{project.StoragesColumn},
			Bidi:    false,
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pu.mutation.MasterStorageCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   project.MasterStorageTable,
			Columns: []string{project.MasterStorageColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: storage.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.MasterStorageIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   project.MasterStorageTable,
			Columns: []string{project.MasterStorageColumn},
			Bidi:    false,
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pu.mutation.DevelopStorageCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   project.DevelopStorageTable,
			Columns: []string{project.DevelopStorageColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: storage.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.DevelopStorageIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   project.DevelopStorageTable,
			Columns: []string{project.DevelopStorageColumn},
			Bidi:    false,
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{project.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// ProjectUpdateOne is the builder for updating a single Project entity.
type ProjectUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ProjectMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (puo *ProjectUpdateOne) SetUpdatedAt(t time.Time) *ProjectUpdateOne {
	puo.mutation.SetUpdatedAt(t)
	return puo
}

// SetDeletedAt sets the "deleted_at" field.
func (puo *ProjectUpdateOne) SetDeletedAt(t time.Time) *ProjectUpdateOne {
	puo.mutation.SetDeletedAt(t)
	return puo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (puo *ProjectUpdateOne) SetNillableDeletedAt(t *time.Time) *ProjectUpdateOne {
	if t != nil {
		puo.SetDeletedAt(*t)
	}
	return puo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (puo *ProjectUpdateOne) ClearDeletedAt() *ProjectUpdateOne {
	puo.mutation.ClearDeletedAt()
	return puo
}

// SetCode sets the "code" field.
func (puo *ProjectUpdateOne) SetCode(s string) *ProjectUpdateOne {
	puo.mutation.SetCode(s)
	return puo
}

// SetName sets the "name" field.
func (puo *ProjectUpdateOne) SetName(s string) *ProjectUpdateOne {
	puo.mutation.SetName(s)
	return puo
}

// SetDescription sets the "Description" field.
func (puo *ProjectUpdateOne) SetDescription(s string) *ProjectUpdateOne {
	puo.mutation.SetDescription(s)
	return puo
}

// SetType sets the "type" field.
func (puo *ProjectUpdateOne) SetType(pr project.Type) *ProjectUpdateOne {
	puo.mutation.SetType(pr)
	return puo
}

// AddStorageIDs adds the "storages" edge to the Storage entity by IDs.
func (puo *ProjectUpdateOne) AddStorageIDs(ids ...uuid.UUID) *ProjectUpdateOne {
	puo.mutation.AddStorageIDs(ids...)
	return puo
}

// AddStorages adds the "storages" edges to the Storage entity.
func (puo *ProjectUpdateOne) AddStorages(s ...*Storage) *ProjectUpdateOne {
	ids := make([]uuid.UUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return puo.AddStorageIDs(ids...)
}

// SetMasterStorageID sets the "master_storage" edge to the Storage entity by ID.
func (puo *ProjectUpdateOne) SetMasterStorageID(id uuid.UUID) *ProjectUpdateOne {
	puo.mutation.SetMasterStorageID(id)
	return puo
}

// SetNillableMasterStorageID sets the "master_storage" edge to the Storage entity by ID if the given value is not nil.
func (puo *ProjectUpdateOne) SetNillableMasterStorageID(id *uuid.UUID) *ProjectUpdateOne {
	if id != nil {
		puo = puo.SetMasterStorageID(*id)
	}
	return puo
}

// SetMasterStorage sets the "master_storage" edge to the Storage entity.
func (puo *ProjectUpdateOne) SetMasterStorage(s *Storage) *ProjectUpdateOne {
	return puo.SetMasterStorageID(s.ID)
}

// SetDevelopStorageID sets the "develop_storage" edge to the Storage entity by ID.
func (puo *ProjectUpdateOne) SetDevelopStorageID(id uuid.UUID) *ProjectUpdateOne {
	puo.mutation.SetDevelopStorageID(id)
	return puo
}

// SetNillableDevelopStorageID sets the "develop_storage" edge to the Storage entity by ID if the given value is not nil.
func (puo *ProjectUpdateOne) SetNillableDevelopStorageID(id *uuid.UUID) *ProjectUpdateOne {
	if id != nil {
		puo = puo.SetDevelopStorageID(*id)
	}
	return puo
}

// SetDevelopStorage sets the "develop_storage" edge to the Storage entity.
func (puo *ProjectUpdateOne) SetDevelopStorage(s *Storage) *ProjectUpdateOne {
	return puo.SetDevelopStorageID(s.ID)
}

// Mutation returns the ProjectMutation object of the builder.
func (puo *ProjectUpdateOne) Mutation() *ProjectMutation {
	return puo.mutation
}

// ClearStorages clears all "storages" edges to the Storage entity.
func (puo *ProjectUpdateOne) ClearStorages() *ProjectUpdateOne {
	puo.mutation.ClearStorages()
	return puo
}

// RemoveStorageIDs removes the "storages" edge to Storage entities by IDs.
func (puo *ProjectUpdateOne) RemoveStorageIDs(ids ...uuid.UUID) *ProjectUpdateOne {
	puo.mutation.RemoveStorageIDs(ids...)
	return puo
}

// RemoveStorages removes "storages" edges to Storage entities.
func (puo *ProjectUpdateOne) RemoveStorages(s ...*Storage) *ProjectUpdateOne {
	ids := make([]uuid.UUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return puo.RemoveStorageIDs(ids...)
}

// ClearMasterStorage clears the "master_storage" edge to the Storage entity.
func (puo *ProjectUpdateOne) ClearMasterStorage() *ProjectUpdateOne {
	puo.mutation.ClearMasterStorage()
	return puo
}

// ClearDevelopStorage clears the "develop_storage" edge to the Storage entity.
func (puo *ProjectUpdateOne) ClearDevelopStorage() *ProjectUpdateOne {
	puo.mutation.ClearDevelopStorage()
	return puo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *ProjectUpdateOne) Select(field string, fields ...string) *ProjectUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Project entity.
func (puo *ProjectUpdateOne) Save(ctx context.Context) (*Project, error) {
	var (
		err  error
		node *Project
	)
	puo.defaults()
	if len(puo.hooks) == 0 {
		if err = puo.check(); err != nil {
			return nil, err
		}
		node, err = puo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ProjectMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = puo.check(); err != nil {
				return nil, err
			}
			puo.mutation = mutation
			node, err = puo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(puo.hooks) - 1; i >= 0; i-- {
			if puo.hooks[i] == nil {
				return nil, fmt.Errorf("db: uninitialized hook (forgotten import db/runtime?)")
			}
			mut = puo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, puo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (puo *ProjectUpdateOne) SaveX(ctx context.Context) *Project {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *ProjectUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *ProjectUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (puo *ProjectUpdateOne) defaults() {
	if _, ok := puo.mutation.UpdatedAt(); !ok {
		v := project.UpdateDefaultUpdatedAt()
		puo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (puo *ProjectUpdateOne) check() error {
	if v, ok := puo.mutation.Code(); ok {
		if err := project.CodeValidator(v); err != nil {
			return &ValidationError{Name: "code", err: fmt.Errorf(`db: validator failed for field "Project.code": %w`, err)}
		}
	}
	if v, ok := puo.mutation.GetType(); ok {
		if err := project.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`db: validator failed for field "Project.type": %w`, err)}
		}
	}
	return nil
}

func (puo *ProjectUpdateOne) sqlSave(ctx context.Context) (_node *Project, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   project.Table,
			Columns: project.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: project.FieldID,
			},
		},
	}
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`db: missing "Project.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, project.FieldID)
		for _, f := range fields {
			if !project.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("db: invalid field %q for query", f)}
			}
			if f != project.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := puo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: project.FieldUpdatedAt,
		})
	}
	if value, ok := puo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: project.FieldDeletedAt,
		})
	}
	if puo.mutation.DeletedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: project.FieldDeletedAt,
		})
	}
	if value, ok := puo.mutation.Code(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: project.FieldCode,
		})
	}
	if value, ok := puo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: project.FieldName,
		})
	}
	if value, ok := puo.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: project.FieldDescription,
		})
	}
	if value, ok := puo.mutation.GetType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: project.FieldType,
		})
	}
	if puo.mutation.StoragesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   project.StoragesTable,
			Columns: []string{project.StoragesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: storage.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.RemovedStoragesIDs(); len(nodes) > 0 && !puo.mutation.StoragesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   project.StoragesTable,
			Columns: []string{project.StoragesColumn},
			Bidi:    false,
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.StoragesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   project.StoragesTable,
			Columns: []string{project.StoragesColumn},
			Bidi:    false,
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if puo.mutation.MasterStorageCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   project.MasterStorageTable,
			Columns: []string{project.MasterStorageColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: storage.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.MasterStorageIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   project.MasterStorageTable,
			Columns: []string{project.MasterStorageColumn},
			Bidi:    false,
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if puo.mutation.DevelopStorageCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   project.DevelopStorageTable,
			Columns: []string{project.DevelopStorageColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: storage.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.DevelopStorageIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   project.DevelopStorageTable,
			Columns: []string{project.DevelopStorageColumn},
			Bidi:    false,
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Project{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{project.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}