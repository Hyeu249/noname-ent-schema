// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Hyeu249/noname-ent-schema/ent/itemhasowner"
	"github.com/Hyeu249/noname-ent-schema/ent/order"
	"github.com/Hyeu249/noname-ent-schema/ent/predicate"
	"github.com/Hyeu249/noname-ent-schema/ent/productquote"
	"github.com/Hyeu249/noname-ent-schema/ent/user"
	"github.com/google/uuid"
)

// ItemHasOwnerQuery is the builder for querying ItemHasOwner entities.
type ItemHasOwnerQuery struct {
	config
	ctx              *QueryContext
	order            []OrderFunc
	inters           []Interceptor
	predicates       []predicate.ItemHasOwner
	withProductQuote *ProductQuoteQuery
	withUser         *UserQuery
	withOrder        *OrderQuery
	withFKs          bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ItemHasOwnerQuery builder.
func (ihoq *ItemHasOwnerQuery) Where(ps ...predicate.ItemHasOwner) *ItemHasOwnerQuery {
	ihoq.predicates = append(ihoq.predicates, ps...)
	return ihoq
}

// Limit the number of records to be returned by this query.
func (ihoq *ItemHasOwnerQuery) Limit(limit int) *ItemHasOwnerQuery {
	ihoq.ctx.Limit = &limit
	return ihoq
}

// Offset to start from.
func (ihoq *ItemHasOwnerQuery) Offset(offset int) *ItemHasOwnerQuery {
	ihoq.ctx.Offset = &offset
	return ihoq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ihoq *ItemHasOwnerQuery) Unique(unique bool) *ItemHasOwnerQuery {
	ihoq.ctx.Unique = &unique
	return ihoq
}

// Order specifies how the records should be ordered.
func (ihoq *ItemHasOwnerQuery) Order(o ...OrderFunc) *ItemHasOwnerQuery {
	ihoq.order = append(ihoq.order, o...)
	return ihoq
}

// QueryProductQuote chains the current query on the "product_quote" edge.
func (ihoq *ItemHasOwnerQuery) QueryProductQuote() *ProductQuoteQuery {
	query := (&ProductQuoteClient{config: ihoq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ihoq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ihoq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(itemhasowner.Table, itemhasowner.FieldID, selector),
			sqlgraph.To(productquote.Table, productquote.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, itemhasowner.ProductQuoteTable, itemhasowner.ProductQuoteColumn),
		)
		fromU = sqlgraph.SetNeighbors(ihoq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryUser chains the current query on the "user" edge.
func (ihoq *ItemHasOwnerQuery) QueryUser() *UserQuery {
	query := (&UserClient{config: ihoq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ihoq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ihoq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(itemhasowner.Table, itemhasowner.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, itemhasowner.UserTable, itemhasowner.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(ihoq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryOrder chains the current query on the "order" edge.
func (ihoq *ItemHasOwnerQuery) QueryOrder() *OrderQuery {
	query := (&OrderClient{config: ihoq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ihoq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ihoq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(itemhasowner.Table, itemhasowner.FieldID, selector),
			sqlgraph.To(order.Table, order.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, itemhasowner.OrderTable, itemhasowner.OrderColumn),
		)
		fromU = sqlgraph.SetNeighbors(ihoq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first ItemHasOwner entity from the query.
// Returns a *NotFoundError when no ItemHasOwner was found.
func (ihoq *ItemHasOwnerQuery) First(ctx context.Context) (*ItemHasOwner, error) {
	nodes, err := ihoq.Limit(1).All(setContextOp(ctx, ihoq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{itemhasowner.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ihoq *ItemHasOwnerQuery) FirstX(ctx context.Context) *ItemHasOwner {
	node, err := ihoq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ItemHasOwner ID from the query.
// Returns a *NotFoundError when no ItemHasOwner ID was found.
func (ihoq *ItemHasOwnerQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = ihoq.Limit(1).IDs(setContextOp(ctx, ihoq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{itemhasowner.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ihoq *ItemHasOwnerQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := ihoq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ItemHasOwner entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one ItemHasOwner entity is found.
// Returns a *NotFoundError when no ItemHasOwner entities are found.
func (ihoq *ItemHasOwnerQuery) Only(ctx context.Context) (*ItemHasOwner, error) {
	nodes, err := ihoq.Limit(2).All(setContextOp(ctx, ihoq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{itemhasowner.Label}
	default:
		return nil, &NotSingularError{itemhasowner.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ihoq *ItemHasOwnerQuery) OnlyX(ctx context.Context) *ItemHasOwner {
	node, err := ihoq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ItemHasOwner ID in the query.
// Returns a *NotSingularError when more than one ItemHasOwner ID is found.
// Returns a *NotFoundError when no entities are found.
func (ihoq *ItemHasOwnerQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = ihoq.Limit(2).IDs(setContextOp(ctx, ihoq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{itemhasowner.Label}
	default:
		err = &NotSingularError{itemhasowner.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ihoq *ItemHasOwnerQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := ihoq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ItemHasOwners.
func (ihoq *ItemHasOwnerQuery) All(ctx context.Context) ([]*ItemHasOwner, error) {
	ctx = setContextOp(ctx, ihoq.ctx, "All")
	if err := ihoq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*ItemHasOwner, *ItemHasOwnerQuery]()
	return withInterceptors[[]*ItemHasOwner](ctx, ihoq, qr, ihoq.inters)
}

// AllX is like All, but panics if an error occurs.
func (ihoq *ItemHasOwnerQuery) AllX(ctx context.Context) []*ItemHasOwner {
	nodes, err := ihoq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ItemHasOwner IDs.
func (ihoq *ItemHasOwnerQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	ctx = setContextOp(ctx, ihoq.ctx, "IDs")
	if err := ihoq.Select(itemhasowner.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ihoq *ItemHasOwnerQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := ihoq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ihoq *ItemHasOwnerQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, ihoq.ctx, "Count")
	if err := ihoq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, ihoq, querierCount[*ItemHasOwnerQuery](), ihoq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (ihoq *ItemHasOwnerQuery) CountX(ctx context.Context) int {
	count, err := ihoq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ihoq *ItemHasOwnerQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, ihoq.ctx, "Exist")
	switch _, err := ihoq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (ihoq *ItemHasOwnerQuery) ExistX(ctx context.Context) bool {
	exist, err := ihoq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ItemHasOwnerQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ihoq *ItemHasOwnerQuery) Clone() *ItemHasOwnerQuery {
	if ihoq == nil {
		return nil
	}
	return &ItemHasOwnerQuery{
		config:           ihoq.config,
		ctx:              ihoq.ctx.Clone(),
		order:            append([]OrderFunc{}, ihoq.order...),
		inters:           append([]Interceptor{}, ihoq.inters...),
		predicates:       append([]predicate.ItemHasOwner{}, ihoq.predicates...),
		withProductQuote: ihoq.withProductQuote.Clone(),
		withUser:         ihoq.withUser.Clone(),
		withOrder:        ihoq.withOrder.Clone(),
		// clone intermediate query.
		sql:  ihoq.sql.Clone(),
		path: ihoq.path,
	}
}

// WithProductQuote tells the query-builder to eager-load the nodes that are connected to
// the "product_quote" edge. The optional arguments are used to configure the query builder of the edge.
func (ihoq *ItemHasOwnerQuery) WithProductQuote(opts ...func(*ProductQuoteQuery)) *ItemHasOwnerQuery {
	query := (&ProductQuoteClient{config: ihoq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	ihoq.withProductQuote = query
	return ihoq
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (ihoq *ItemHasOwnerQuery) WithUser(opts ...func(*UserQuery)) *ItemHasOwnerQuery {
	query := (&UserClient{config: ihoq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	ihoq.withUser = query
	return ihoq
}

// WithOrder tells the query-builder to eager-load the nodes that are connected to
// the "order" edge. The optional arguments are used to configure the query builder of the edge.
func (ihoq *ItemHasOwnerQuery) WithOrder(opts ...func(*OrderQuery)) *ItemHasOwnerQuery {
	query := (&OrderClient{config: ihoq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	ihoq.withOrder = query
	return ihoq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreateTime time.Time `json:"create_time,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.ItemHasOwner.Query().
//		GroupBy(itemhasowner.FieldCreateTime).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (ihoq *ItemHasOwnerQuery) GroupBy(field string, fields ...string) *ItemHasOwnerGroupBy {
	ihoq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ItemHasOwnerGroupBy{build: ihoq}
	grbuild.flds = &ihoq.ctx.Fields
	grbuild.label = itemhasowner.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreateTime time.Time `json:"create_time,omitempty"`
//	}
//
//	client.ItemHasOwner.Query().
//		Select(itemhasowner.FieldCreateTime).
//		Scan(ctx, &v)
func (ihoq *ItemHasOwnerQuery) Select(fields ...string) *ItemHasOwnerSelect {
	ihoq.ctx.Fields = append(ihoq.ctx.Fields, fields...)
	sbuild := &ItemHasOwnerSelect{ItemHasOwnerQuery: ihoq}
	sbuild.label = itemhasowner.Label
	sbuild.flds, sbuild.scan = &ihoq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ItemHasOwnerSelect configured with the given aggregations.
func (ihoq *ItemHasOwnerQuery) Aggregate(fns ...AggregateFunc) *ItemHasOwnerSelect {
	return ihoq.Select().Aggregate(fns...)
}

func (ihoq *ItemHasOwnerQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range ihoq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, ihoq); err != nil {
				return err
			}
		}
	}
	for _, f := range ihoq.ctx.Fields {
		if !itemhasowner.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if ihoq.path != nil {
		prev, err := ihoq.path(ctx)
		if err != nil {
			return err
		}
		ihoq.sql = prev
	}
	return nil
}

func (ihoq *ItemHasOwnerQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*ItemHasOwner, error) {
	var (
		nodes       = []*ItemHasOwner{}
		withFKs     = ihoq.withFKs
		_spec       = ihoq.querySpec()
		loadedTypes = [3]bool{
			ihoq.withProductQuote != nil,
			ihoq.withUser != nil,
			ihoq.withOrder != nil,
		}
	)
	if ihoq.withProductQuote != nil || ihoq.withUser != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, itemhasowner.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*ItemHasOwner).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &ItemHasOwner{config: ihoq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, ihoq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := ihoq.withProductQuote; query != nil {
		if err := ihoq.loadProductQuote(ctx, query, nodes, nil,
			func(n *ItemHasOwner, e *ProductQuote) { n.Edges.ProductQuote = e }); err != nil {
			return nil, err
		}
	}
	if query := ihoq.withUser; query != nil {
		if err := ihoq.loadUser(ctx, query, nodes, nil,
			func(n *ItemHasOwner, e *User) { n.Edges.User = e }); err != nil {
			return nil, err
		}
	}
	if query := ihoq.withOrder; query != nil {
		if err := ihoq.loadOrder(ctx, query, nodes,
			func(n *ItemHasOwner) { n.Edges.Order = []*Order{} },
			func(n *ItemHasOwner, e *Order) { n.Edges.Order = append(n.Edges.Order, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (ihoq *ItemHasOwnerQuery) loadProductQuote(ctx context.Context, query *ProductQuoteQuery, nodes []*ItemHasOwner, init func(*ItemHasOwner), assign func(*ItemHasOwner, *ProductQuote)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*ItemHasOwner)
	for i := range nodes {
		if nodes[i].item_has_owner_product_quote == nil {
			continue
		}
		fk := *nodes[i].item_has_owner_product_quote
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(productquote.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "item_has_owner_product_quote" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (ihoq *ItemHasOwnerQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*ItemHasOwner, init func(*ItemHasOwner), assign func(*ItemHasOwner, *User)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*ItemHasOwner)
	for i := range nodes {
		if nodes[i].item_has_owner_user == nil {
			continue
		}
		fk := *nodes[i].item_has_owner_user
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(user.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "item_has_owner_user" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (ihoq *ItemHasOwnerQuery) loadOrder(ctx context.Context, query *OrderQuery, nodes []*ItemHasOwner, init func(*ItemHasOwner), assign func(*ItemHasOwner, *Order)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*ItemHasOwner)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Order(func(s *sql.Selector) {
		s.Where(sql.InValues(itemhasowner.OrderColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.order_item_has_owner
		if fk == nil {
			return fmt.Errorf(`foreign-key "order_item_has_owner" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "order_item_has_owner" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (ihoq *ItemHasOwnerQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ihoq.querySpec()
	_spec.Node.Columns = ihoq.ctx.Fields
	if len(ihoq.ctx.Fields) > 0 {
		_spec.Unique = ihoq.ctx.Unique != nil && *ihoq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, ihoq.driver, _spec)
}

func (ihoq *ItemHasOwnerQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   itemhasowner.Table,
			Columns: itemhasowner.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: itemhasowner.FieldID,
			},
		},
		From:   ihoq.sql,
		Unique: true,
	}
	if unique := ihoq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := ihoq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, itemhasowner.FieldID)
		for i := range fields {
			if fields[i] != itemhasowner.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := ihoq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ihoq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ihoq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ihoq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ihoq *ItemHasOwnerQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ihoq.driver.Dialect())
	t1 := builder.Table(itemhasowner.Table)
	columns := ihoq.ctx.Fields
	if len(columns) == 0 {
		columns = itemhasowner.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ihoq.sql != nil {
		selector = ihoq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if ihoq.ctx.Unique != nil && *ihoq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range ihoq.predicates {
		p(selector)
	}
	for _, p := range ihoq.order {
		p(selector)
	}
	if offset := ihoq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ihoq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ItemHasOwnerGroupBy is the group-by builder for ItemHasOwner entities.
type ItemHasOwnerGroupBy struct {
	selector
	build *ItemHasOwnerQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ihogb *ItemHasOwnerGroupBy) Aggregate(fns ...AggregateFunc) *ItemHasOwnerGroupBy {
	ihogb.fns = append(ihogb.fns, fns...)
	return ihogb
}

// Scan applies the selector query and scans the result into the given value.
func (ihogb *ItemHasOwnerGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ihogb.build.ctx, "GroupBy")
	if err := ihogb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ItemHasOwnerQuery, *ItemHasOwnerGroupBy](ctx, ihogb.build, ihogb, ihogb.build.inters, v)
}

func (ihogb *ItemHasOwnerGroupBy) sqlScan(ctx context.Context, root *ItemHasOwnerQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(ihogb.fns))
	for _, fn := range ihogb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*ihogb.flds)+len(ihogb.fns))
		for _, f := range *ihogb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*ihogb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ihogb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ItemHasOwnerSelect is the builder for selecting fields of ItemHasOwner entities.
type ItemHasOwnerSelect struct {
	*ItemHasOwnerQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ihos *ItemHasOwnerSelect) Aggregate(fns ...AggregateFunc) *ItemHasOwnerSelect {
	ihos.fns = append(ihos.fns, fns...)
	return ihos
}

// Scan applies the selector query and scans the result into the given value.
func (ihos *ItemHasOwnerSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ihos.ctx, "Select")
	if err := ihos.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ItemHasOwnerQuery, *ItemHasOwnerSelect](ctx, ihos.ItemHasOwnerQuery, ihos, ihos.inters, v)
}

func (ihos *ItemHasOwnerSelect) sqlScan(ctx context.Context, root *ItemHasOwnerQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ihos.fns))
	for _, fn := range ihos.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ihos.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ihos.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
