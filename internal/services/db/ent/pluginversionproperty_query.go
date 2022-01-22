// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/khorevaa/r2gitsync/internal/services/db/ent/plugin"
	"github.com/khorevaa/r2gitsync/internal/services/db/ent/pluginversion"
	"github.com/khorevaa/r2gitsync/internal/services/db/ent/pluginversionproperty"
	"github.com/khorevaa/r2gitsync/internal/services/db/ent/predicate"
)

// PluginVersionPropertyQuery is the builder for querying PluginVersionProperty entities.
type PluginVersionPropertyQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.PluginVersionProperty
	// eager-loading edges.
	withPlugin  *PluginQuery
	withVersion *PluginVersionQuery
	withFKs     bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the PluginVersionPropertyQuery builder.
func (pvpq *PluginVersionPropertyQuery) Where(ps ...predicate.PluginVersionProperty) *PluginVersionPropertyQuery {
	pvpq.predicates = append(pvpq.predicates, ps...)
	return pvpq
}

// Limit adds a limit step to the query.
func (pvpq *PluginVersionPropertyQuery) Limit(limit int) *PluginVersionPropertyQuery {
	pvpq.limit = &limit
	return pvpq
}

// Offset adds an offset step to the query.
func (pvpq *PluginVersionPropertyQuery) Offset(offset int) *PluginVersionPropertyQuery {
	pvpq.offset = &offset
	return pvpq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (pvpq *PluginVersionPropertyQuery) Unique(unique bool) *PluginVersionPropertyQuery {
	pvpq.unique = &unique
	return pvpq
}

// Order adds an order step to the query.
func (pvpq *PluginVersionPropertyQuery) Order(o ...OrderFunc) *PluginVersionPropertyQuery {
	pvpq.order = append(pvpq.order, o...)
	return pvpq
}

// QueryPlugin chains the current query on the "plugin" edge.
func (pvpq *PluginVersionPropertyQuery) QueryPlugin() *PluginQuery {
	query := &PluginQuery{config: pvpq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pvpq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pvpq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(pluginversionproperty.Table, pluginversionproperty.FieldID, selector),
			sqlgraph.To(plugin.Table, plugin.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, pluginversionproperty.PluginTable, pluginversionproperty.PluginColumn),
		)
		fromU = sqlgraph.SetNeighbors(pvpq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryVersion chains the current query on the "version" edge.
func (pvpq *PluginVersionPropertyQuery) QueryVersion() *PluginVersionQuery {
	query := &PluginVersionQuery{config: pvpq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pvpq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pvpq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(pluginversionproperty.Table, pluginversionproperty.FieldID, selector),
			sqlgraph.To(pluginversion.Table, pluginversion.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, pluginversionproperty.VersionTable, pluginversionproperty.VersionColumn),
		)
		fromU = sqlgraph.SetNeighbors(pvpq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first PluginVersionProperty entity from the query.
// Returns a *NotFoundError when no PluginVersionProperty was found.
func (pvpq *PluginVersionPropertyQuery) First(ctx context.Context) (*PluginVersionProperty, error) {
	nodes, err := pvpq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{pluginversionproperty.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (pvpq *PluginVersionPropertyQuery) FirstX(ctx context.Context) *PluginVersionProperty {
	node, err := pvpq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first PluginVersionProperty ID from the query.
// Returns a *NotFoundError when no PluginVersionProperty ID was found.
func (pvpq *PluginVersionPropertyQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = pvpq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{pluginversionproperty.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (pvpq *PluginVersionPropertyQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := pvpq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single PluginVersionProperty entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one PluginVersionProperty entity is not found.
// Returns a *NotFoundError when no PluginVersionProperty entities are found.
func (pvpq *PluginVersionPropertyQuery) Only(ctx context.Context) (*PluginVersionProperty, error) {
	nodes, err := pvpq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{pluginversionproperty.Label}
	default:
		return nil, &NotSingularError{pluginversionproperty.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (pvpq *PluginVersionPropertyQuery) OnlyX(ctx context.Context) *PluginVersionProperty {
	node, err := pvpq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only PluginVersionProperty ID in the query.
// Returns a *NotSingularError when exactly one PluginVersionProperty ID is not found.
// Returns a *NotFoundError when no entities are found.
func (pvpq *PluginVersionPropertyQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = pvpq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{pluginversionproperty.Label}
	default:
		err = &NotSingularError{pluginversionproperty.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (pvpq *PluginVersionPropertyQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := pvpq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of PluginVersionProperties.
func (pvpq *PluginVersionPropertyQuery) All(ctx context.Context) ([]*PluginVersionProperty, error) {
	if err := pvpq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return pvpq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (pvpq *PluginVersionPropertyQuery) AllX(ctx context.Context) []*PluginVersionProperty {
	nodes, err := pvpq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of PluginVersionProperty IDs.
func (pvpq *PluginVersionPropertyQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := pvpq.Select(pluginversionproperty.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (pvpq *PluginVersionPropertyQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := pvpq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (pvpq *PluginVersionPropertyQuery) Count(ctx context.Context) (int, error) {
	if err := pvpq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return pvpq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (pvpq *PluginVersionPropertyQuery) CountX(ctx context.Context) int {
	count, err := pvpq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (pvpq *PluginVersionPropertyQuery) Exist(ctx context.Context) (bool, error) {
	if err := pvpq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return pvpq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (pvpq *PluginVersionPropertyQuery) ExistX(ctx context.Context) bool {
	exist, err := pvpq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the PluginVersionPropertyQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (pvpq *PluginVersionPropertyQuery) Clone() *PluginVersionPropertyQuery {
	if pvpq == nil {
		return nil
	}
	return &PluginVersionPropertyQuery{
		config:      pvpq.config,
		limit:       pvpq.limit,
		offset:      pvpq.offset,
		order:       append([]OrderFunc{}, pvpq.order...),
		predicates:  append([]predicate.PluginVersionProperty{}, pvpq.predicates...),
		withPlugin:  pvpq.withPlugin.Clone(),
		withVersion: pvpq.withVersion.Clone(),
		// clone intermediate query.
		sql:  pvpq.sql.Clone(),
		path: pvpq.path,
	}
}

// WithPlugin tells the query-builder to eager-load the nodes that are connected to
// the "plugin" edge. The optional arguments are used to configure the query builder of the edge.
func (pvpq *PluginVersionPropertyQuery) WithPlugin(opts ...func(*PluginQuery)) *PluginVersionPropertyQuery {
	query := &PluginQuery{config: pvpq.config}
	for _, opt := range opts {
		opt(query)
	}
	pvpq.withPlugin = query
	return pvpq
}

// WithVersion tells the query-builder to eager-load the nodes that are connected to
// the "version" edge. The optional arguments are used to configure the query builder of the edge.
func (pvpq *PluginVersionPropertyQuery) WithVersion(opts ...func(*PluginVersionQuery)) *PluginVersionPropertyQuery {
	query := &PluginVersionQuery{config: pvpq.config}
	for _, opt := range opts {
		opt(query)
	}
	pvpq.withVersion = query
	return pvpq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.PluginVersionProperty.Query().
//		GroupBy(pluginversionproperty.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (pvpq *PluginVersionPropertyQuery) GroupBy(field string, fields ...string) *PluginVersionPropertyGroupBy {
	group := &PluginVersionPropertyGroupBy{config: pvpq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := pvpq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return pvpq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//	}
//
//	client.PluginVersionProperty.Query().
//		Select(pluginversionproperty.FieldCreatedAt).
//		Scan(ctx, &v)
//
func (pvpq *PluginVersionPropertyQuery) Select(fields ...string) *PluginVersionPropertySelect {
	pvpq.fields = append(pvpq.fields, fields...)
	return &PluginVersionPropertySelect{PluginVersionPropertyQuery: pvpq}
}

func (pvpq *PluginVersionPropertyQuery) prepareQuery(ctx context.Context) error {
	for _, f := range pvpq.fields {
		if !pluginversionproperty.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if pvpq.path != nil {
		prev, err := pvpq.path(ctx)
		if err != nil {
			return err
		}
		pvpq.sql = prev
	}
	return nil
}

func (pvpq *PluginVersionPropertyQuery) sqlAll(ctx context.Context) ([]*PluginVersionProperty, error) {
	var (
		nodes       = []*PluginVersionProperty{}
		withFKs     = pvpq.withFKs
		_spec       = pvpq.querySpec()
		loadedTypes = [2]bool{
			pvpq.withPlugin != nil,
			pvpq.withVersion != nil,
		}
	)
	if pvpq.withPlugin != nil || pvpq.withVersion != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, pluginversionproperty.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &PluginVersionProperty{config: pvpq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, pvpq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := pvpq.withPlugin; query != nil {
		ids := make([]uuid.UUID, 0, len(nodes))
		nodeids := make(map[uuid.UUID][]*PluginVersionProperty)
		for i := range nodes {
			if nodes[i].plugin_version_property_plugin == nil {
				continue
			}
			fk := *nodes[i].plugin_version_property_plugin
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(plugin.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "plugin_version_property_plugin" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Plugin = n
			}
		}
	}

	if query := pvpq.withVersion; query != nil {
		ids := make([]uuid.UUID, 0, len(nodes))
		nodeids := make(map[uuid.UUID][]*PluginVersionProperty)
		for i := range nodes {
			if nodes[i].plugin_version_property_version == nil {
				continue
			}
			fk := *nodes[i].plugin_version_property_version
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(pluginversion.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "plugin_version_property_version" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Version = n
			}
		}
	}

	return nodes, nil
}

func (pvpq *PluginVersionPropertyQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := pvpq.querySpec()
	_spec.Node.Columns = pvpq.fields
	if len(pvpq.fields) > 0 {
		_spec.Unique = pvpq.unique != nil && *pvpq.unique
	}
	return sqlgraph.CountNodes(ctx, pvpq.driver, _spec)
}

func (pvpq *PluginVersionPropertyQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := pvpq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (pvpq *PluginVersionPropertyQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   pluginversionproperty.Table,
			Columns: pluginversionproperty.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: pluginversionproperty.FieldID,
			},
		},
		From:   pvpq.sql,
		Unique: true,
	}
	if unique := pvpq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := pvpq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, pluginversionproperty.FieldID)
		for i := range fields {
			if fields[i] != pluginversionproperty.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := pvpq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := pvpq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := pvpq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := pvpq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (pvpq *PluginVersionPropertyQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(pvpq.driver.Dialect())
	t1 := builder.Table(pluginversionproperty.Table)
	columns := pvpq.fields
	if len(columns) == 0 {
		columns = pluginversionproperty.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if pvpq.sql != nil {
		selector = pvpq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if pvpq.unique != nil && *pvpq.unique {
		selector.Distinct()
	}
	for _, p := range pvpq.predicates {
		p(selector)
	}
	for _, p := range pvpq.order {
		p(selector)
	}
	if offset := pvpq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := pvpq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// PluginVersionPropertyGroupBy is the group-by builder for PluginVersionProperty entities.
type PluginVersionPropertyGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (pvpgb *PluginVersionPropertyGroupBy) Aggregate(fns ...AggregateFunc) *PluginVersionPropertyGroupBy {
	pvpgb.fns = append(pvpgb.fns, fns...)
	return pvpgb
}

// Scan applies the group-by query and scans the result into the given value.
func (pvpgb *PluginVersionPropertyGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := pvpgb.path(ctx)
	if err != nil {
		return err
	}
	pvpgb.sql = query
	return pvpgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (pvpgb *PluginVersionPropertyGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := pvpgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (pvpgb *PluginVersionPropertyGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(pvpgb.fields) > 1 {
		return nil, errors.New("ent: PluginVersionPropertyGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := pvpgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (pvpgb *PluginVersionPropertyGroupBy) StringsX(ctx context.Context) []string {
	v, err := pvpgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (pvpgb *PluginVersionPropertyGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = pvpgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{pluginversionproperty.Label}
	default:
		err = fmt.Errorf("ent: PluginVersionPropertyGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (pvpgb *PluginVersionPropertyGroupBy) StringX(ctx context.Context) string {
	v, err := pvpgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (pvpgb *PluginVersionPropertyGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(pvpgb.fields) > 1 {
		return nil, errors.New("ent: PluginVersionPropertyGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := pvpgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (pvpgb *PluginVersionPropertyGroupBy) IntsX(ctx context.Context) []int {
	v, err := pvpgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (pvpgb *PluginVersionPropertyGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = pvpgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{pluginversionproperty.Label}
	default:
		err = fmt.Errorf("ent: PluginVersionPropertyGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (pvpgb *PluginVersionPropertyGroupBy) IntX(ctx context.Context) int {
	v, err := pvpgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (pvpgb *PluginVersionPropertyGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(pvpgb.fields) > 1 {
		return nil, errors.New("ent: PluginVersionPropertyGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := pvpgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (pvpgb *PluginVersionPropertyGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := pvpgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (pvpgb *PluginVersionPropertyGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = pvpgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{pluginversionproperty.Label}
	default:
		err = fmt.Errorf("ent: PluginVersionPropertyGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (pvpgb *PluginVersionPropertyGroupBy) Float64X(ctx context.Context) float64 {
	v, err := pvpgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (pvpgb *PluginVersionPropertyGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(pvpgb.fields) > 1 {
		return nil, errors.New("ent: PluginVersionPropertyGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := pvpgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (pvpgb *PluginVersionPropertyGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := pvpgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (pvpgb *PluginVersionPropertyGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = pvpgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{pluginversionproperty.Label}
	default:
		err = fmt.Errorf("ent: PluginVersionPropertyGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (pvpgb *PluginVersionPropertyGroupBy) BoolX(ctx context.Context) bool {
	v, err := pvpgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (pvpgb *PluginVersionPropertyGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range pvpgb.fields {
		if !pluginversionproperty.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := pvpgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pvpgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (pvpgb *PluginVersionPropertyGroupBy) sqlQuery() *sql.Selector {
	selector := pvpgb.sql.Select()
	aggregation := make([]string, 0, len(pvpgb.fns))
	for _, fn := range pvpgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(pvpgb.fields)+len(pvpgb.fns))
		for _, f := range pvpgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(pvpgb.fields...)...)
}

// PluginVersionPropertySelect is the builder for selecting fields of PluginVersionProperty entities.
type PluginVersionPropertySelect struct {
	*PluginVersionPropertyQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (pvps *PluginVersionPropertySelect) Scan(ctx context.Context, v interface{}) error {
	if err := pvps.prepareQuery(ctx); err != nil {
		return err
	}
	pvps.sql = pvps.PluginVersionPropertyQuery.sqlQuery(ctx)
	return pvps.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (pvps *PluginVersionPropertySelect) ScanX(ctx context.Context, v interface{}) {
	if err := pvps.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (pvps *PluginVersionPropertySelect) Strings(ctx context.Context) ([]string, error) {
	if len(pvps.fields) > 1 {
		return nil, errors.New("ent: PluginVersionPropertySelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := pvps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (pvps *PluginVersionPropertySelect) StringsX(ctx context.Context) []string {
	v, err := pvps.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (pvps *PluginVersionPropertySelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = pvps.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{pluginversionproperty.Label}
	default:
		err = fmt.Errorf("ent: PluginVersionPropertySelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (pvps *PluginVersionPropertySelect) StringX(ctx context.Context) string {
	v, err := pvps.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (pvps *PluginVersionPropertySelect) Ints(ctx context.Context) ([]int, error) {
	if len(pvps.fields) > 1 {
		return nil, errors.New("ent: PluginVersionPropertySelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := pvps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (pvps *PluginVersionPropertySelect) IntsX(ctx context.Context) []int {
	v, err := pvps.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (pvps *PluginVersionPropertySelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = pvps.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{pluginversionproperty.Label}
	default:
		err = fmt.Errorf("ent: PluginVersionPropertySelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (pvps *PluginVersionPropertySelect) IntX(ctx context.Context) int {
	v, err := pvps.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (pvps *PluginVersionPropertySelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(pvps.fields) > 1 {
		return nil, errors.New("ent: PluginVersionPropertySelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := pvps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (pvps *PluginVersionPropertySelect) Float64sX(ctx context.Context) []float64 {
	v, err := pvps.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (pvps *PluginVersionPropertySelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = pvps.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{pluginversionproperty.Label}
	default:
		err = fmt.Errorf("ent: PluginVersionPropertySelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (pvps *PluginVersionPropertySelect) Float64X(ctx context.Context) float64 {
	v, err := pvps.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (pvps *PluginVersionPropertySelect) Bools(ctx context.Context) ([]bool, error) {
	if len(pvps.fields) > 1 {
		return nil, errors.New("ent: PluginVersionPropertySelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := pvps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (pvps *PluginVersionPropertySelect) BoolsX(ctx context.Context) []bool {
	v, err := pvps.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (pvps *PluginVersionPropertySelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = pvps.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{pluginversionproperty.Label}
	default:
		err = fmt.Errorf("ent: PluginVersionPropertySelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (pvps *PluginVersionPropertySelect) BoolX(ctx context.Context) bool {
	v, err := pvps.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (pvps *PluginVersionPropertySelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := pvps.sql.Query()
	if err := pvps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
