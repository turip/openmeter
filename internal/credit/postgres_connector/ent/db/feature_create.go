// Code generated by ent, DO NOT EDIT.

package db

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/openmeterio/openmeter/internal/credit/postgres_connector/ent/db/creditentry"
	"github.com/openmeterio/openmeter/internal/credit/postgres_connector/ent/db/feature"
)

// FeatureCreate is the builder for creating a Feature entity.
type FeatureCreate struct {
	config
	mutation *FeatureMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (fc *FeatureCreate) SetCreatedAt(t time.Time) *FeatureCreate {
	fc.mutation.SetCreatedAt(t)
	return fc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (fc *FeatureCreate) SetNillableCreatedAt(t *time.Time) *FeatureCreate {
	if t != nil {
		fc.SetCreatedAt(*t)
	}
	return fc
}

// SetUpdatedAt sets the "updated_at" field.
func (fc *FeatureCreate) SetUpdatedAt(t time.Time) *FeatureCreate {
	fc.mutation.SetUpdatedAt(t)
	return fc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (fc *FeatureCreate) SetNillableUpdatedAt(t *time.Time) *FeatureCreate {
	if t != nil {
		fc.SetUpdatedAt(*t)
	}
	return fc
}

// SetNamespace sets the "namespace" field.
func (fc *FeatureCreate) SetNamespace(s string) *FeatureCreate {
	fc.mutation.SetNamespace(s)
	return fc
}

// SetName sets the "name" field.
func (fc *FeatureCreate) SetName(s string) *FeatureCreate {
	fc.mutation.SetName(s)
	return fc
}

// SetMeterSlug sets the "meter_slug" field.
func (fc *FeatureCreate) SetMeterSlug(s string) *FeatureCreate {
	fc.mutation.SetMeterSlug(s)
	return fc
}

// SetMeterGroupByFilters sets the "meter_group_by_filters" field.
func (fc *FeatureCreate) SetMeterGroupByFilters(m map[string]string) *FeatureCreate {
	fc.mutation.SetMeterGroupByFilters(m)
	return fc
}

// SetArchived sets the "archived" field.
func (fc *FeatureCreate) SetArchived(b bool) *FeatureCreate {
	fc.mutation.SetArchived(b)
	return fc
}

// SetNillableArchived sets the "archived" field if the given value is not nil.
func (fc *FeatureCreate) SetNillableArchived(b *bool) *FeatureCreate {
	if b != nil {
		fc.SetArchived(*b)
	}
	return fc
}

// SetID sets the "id" field.
func (fc *FeatureCreate) SetID(s string) *FeatureCreate {
	fc.mutation.SetID(s)
	return fc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (fc *FeatureCreate) SetNillableID(s *string) *FeatureCreate {
	if s != nil {
		fc.SetID(*s)
	}
	return fc
}

// AddCreditGrantIDs adds the "credit_grants" edge to the CreditEntry entity by IDs.
func (fc *FeatureCreate) AddCreditGrantIDs(ids ...string) *FeatureCreate {
	fc.mutation.AddCreditGrantIDs(ids...)
	return fc
}

// AddCreditGrants adds the "credit_grants" edges to the CreditEntry entity.
func (fc *FeatureCreate) AddCreditGrants(c ...*CreditEntry) *FeatureCreate {
	ids := make([]string, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return fc.AddCreditGrantIDs(ids...)
}

// Mutation returns the FeatureMutation object of the builder.
func (fc *FeatureCreate) Mutation() *FeatureMutation {
	return fc.mutation
}

// Save creates the Feature in the database.
func (fc *FeatureCreate) Save(ctx context.Context) (*Feature, error) {
	fc.defaults()
	return withHooks(ctx, fc.sqlSave, fc.mutation, fc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (fc *FeatureCreate) SaveX(ctx context.Context) *Feature {
	v, err := fc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (fc *FeatureCreate) Exec(ctx context.Context) error {
	_, err := fc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fc *FeatureCreate) ExecX(ctx context.Context) {
	if err := fc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (fc *FeatureCreate) defaults() {
	if _, ok := fc.mutation.CreatedAt(); !ok {
		v := feature.DefaultCreatedAt()
		fc.mutation.SetCreatedAt(v)
	}
	if _, ok := fc.mutation.UpdatedAt(); !ok {
		v := feature.DefaultUpdatedAt()
		fc.mutation.SetUpdatedAt(v)
	}
	if _, ok := fc.mutation.Archived(); !ok {
		v := feature.DefaultArchived
		fc.mutation.SetArchived(v)
	}
	if _, ok := fc.mutation.ID(); !ok {
		v := feature.DefaultID()
		fc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (fc *FeatureCreate) check() error {
	if _, ok := fc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`db: missing required field "Feature.created_at"`)}
	}
	if _, ok := fc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`db: missing required field "Feature.updated_at"`)}
	}
	if _, ok := fc.mutation.Namespace(); !ok {
		return &ValidationError{Name: "namespace", err: errors.New(`db: missing required field "Feature.namespace"`)}
	}
	if v, ok := fc.mutation.Namespace(); ok {
		if err := feature.NamespaceValidator(v); err != nil {
			return &ValidationError{Name: "namespace", err: fmt.Errorf(`db: validator failed for field "Feature.namespace": %w`, err)}
		}
	}
	if _, ok := fc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`db: missing required field "Feature.name"`)}
	}
	if v, ok := fc.mutation.Name(); ok {
		if err := feature.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`db: validator failed for field "Feature.name": %w`, err)}
		}
	}
	if _, ok := fc.mutation.MeterSlug(); !ok {
		return &ValidationError{Name: "meter_slug", err: errors.New(`db: missing required field "Feature.meter_slug"`)}
	}
	if v, ok := fc.mutation.MeterSlug(); ok {
		if err := feature.MeterSlugValidator(v); err != nil {
			return &ValidationError{Name: "meter_slug", err: fmt.Errorf(`db: validator failed for field "Feature.meter_slug": %w`, err)}
		}
	}
	if _, ok := fc.mutation.Archived(); !ok {
		return &ValidationError{Name: "archived", err: errors.New(`db: missing required field "Feature.archived"`)}
	}
	return nil
}

func (fc *FeatureCreate) sqlSave(ctx context.Context) (*Feature, error) {
	if err := fc.check(); err != nil {
		return nil, err
	}
	_node, _spec := fc.createSpec()
	if err := sqlgraph.CreateNode(ctx, fc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected Feature.ID type: %T", _spec.ID.Value)
		}
	}
	fc.mutation.id = &_node.ID
	fc.mutation.done = true
	return _node, nil
}

func (fc *FeatureCreate) createSpec() (*Feature, *sqlgraph.CreateSpec) {
	var (
		_node = &Feature{config: fc.config}
		_spec = sqlgraph.NewCreateSpec(feature.Table, sqlgraph.NewFieldSpec(feature.FieldID, field.TypeString))
	)
	_spec.OnConflict = fc.conflict
	if id, ok := fc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := fc.mutation.CreatedAt(); ok {
		_spec.SetField(feature.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := fc.mutation.UpdatedAt(); ok {
		_spec.SetField(feature.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := fc.mutation.Namespace(); ok {
		_spec.SetField(feature.FieldNamespace, field.TypeString, value)
		_node.Namespace = value
	}
	if value, ok := fc.mutation.Name(); ok {
		_spec.SetField(feature.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := fc.mutation.MeterSlug(); ok {
		_spec.SetField(feature.FieldMeterSlug, field.TypeString, value)
		_node.MeterSlug = value
	}
	if value, ok := fc.mutation.MeterGroupByFilters(); ok {
		_spec.SetField(feature.FieldMeterGroupByFilters, field.TypeJSON, value)
		_node.MeterGroupByFilters = value
	}
	if value, ok := fc.mutation.Archived(); ok {
		_spec.SetField(feature.FieldArchived, field.TypeBool, value)
		_node.Archived = value
	}
	if nodes := fc.mutation.CreditGrantsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   feature.CreditGrantsTable,
			Columns: []string{feature.CreditGrantsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(creditentry.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Feature.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.FeatureUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (fc *FeatureCreate) OnConflict(opts ...sql.ConflictOption) *FeatureUpsertOne {
	fc.conflict = opts
	return &FeatureUpsertOne{
		create: fc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Feature.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (fc *FeatureCreate) OnConflictColumns(columns ...string) *FeatureUpsertOne {
	fc.conflict = append(fc.conflict, sql.ConflictColumns(columns...))
	return &FeatureUpsertOne{
		create: fc,
	}
}

type (
	// FeatureUpsertOne is the builder for "upsert"-ing
	//  one Feature node.
	FeatureUpsertOne struct {
		create *FeatureCreate
	}

	// FeatureUpsert is the "OnConflict" setter.
	FeatureUpsert struct {
		*sql.UpdateSet
	}
)

// SetUpdatedAt sets the "updated_at" field.
func (u *FeatureUpsert) SetUpdatedAt(v time.Time) *FeatureUpsert {
	u.Set(feature.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *FeatureUpsert) UpdateUpdatedAt() *FeatureUpsert {
	u.SetExcluded(feature.FieldUpdatedAt)
	return u
}

// SetName sets the "name" field.
func (u *FeatureUpsert) SetName(v string) *FeatureUpsert {
	u.Set(feature.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *FeatureUpsert) UpdateName() *FeatureUpsert {
	u.SetExcluded(feature.FieldName)
	return u
}

// SetMeterGroupByFilters sets the "meter_group_by_filters" field.
func (u *FeatureUpsert) SetMeterGroupByFilters(v map[string]string) *FeatureUpsert {
	u.Set(feature.FieldMeterGroupByFilters, v)
	return u
}

// UpdateMeterGroupByFilters sets the "meter_group_by_filters" field to the value that was provided on create.
func (u *FeatureUpsert) UpdateMeterGroupByFilters() *FeatureUpsert {
	u.SetExcluded(feature.FieldMeterGroupByFilters)
	return u
}

// ClearMeterGroupByFilters clears the value of the "meter_group_by_filters" field.
func (u *FeatureUpsert) ClearMeterGroupByFilters() *FeatureUpsert {
	u.SetNull(feature.FieldMeterGroupByFilters)
	return u
}

// SetArchived sets the "archived" field.
func (u *FeatureUpsert) SetArchived(v bool) *FeatureUpsert {
	u.Set(feature.FieldArchived, v)
	return u
}

// UpdateArchived sets the "archived" field to the value that was provided on create.
func (u *FeatureUpsert) UpdateArchived() *FeatureUpsert {
	u.SetExcluded(feature.FieldArchived)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Feature.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(feature.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *FeatureUpsertOne) UpdateNewValues() *FeatureUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(feature.FieldID)
		}
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(feature.FieldCreatedAt)
		}
		if _, exists := u.create.mutation.Namespace(); exists {
			s.SetIgnore(feature.FieldNamespace)
		}
		if _, exists := u.create.mutation.MeterSlug(); exists {
			s.SetIgnore(feature.FieldMeterSlug)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Feature.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *FeatureUpsertOne) Ignore() *FeatureUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *FeatureUpsertOne) DoNothing() *FeatureUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the FeatureCreate.OnConflict
// documentation for more info.
func (u *FeatureUpsertOne) Update(set func(*FeatureUpsert)) *FeatureUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&FeatureUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *FeatureUpsertOne) SetUpdatedAt(v time.Time) *FeatureUpsertOne {
	return u.Update(func(s *FeatureUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *FeatureUpsertOne) UpdateUpdatedAt() *FeatureUpsertOne {
	return u.Update(func(s *FeatureUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetName sets the "name" field.
func (u *FeatureUpsertOne) SetName(v string) *FeatureUpsertOne {
	return u.Update(func(s *FeatureUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *FeatureUpsertOne) UpdateName() *FeatureUpsertOne {
	return u.Update(func(s *FeatureUpsert) {
		s.UpdateName()
	})
}

// SetMeterGroupByFilters sets the "meter_group_by_filters" field.
func (u *FeatureUpsertOne) SetMeterGroupByFilters(v map[string]string) *FeatureUpsertOne {
	return u.Update(func(s *FeatureUpsert) {
		s.SetMeterGroupByFilters(v)
	})
}

// UpdateMeterGroupByFilters sets the "meter_group_by_filters" field to the value that was provided on create.
func (u *FeatureUpsertOne) UpdateMeterGroupByFilters() *FeatureUpsertOne {
	return u.Update(func(s *FeatureUpsert) {
		s.UpdateMeterGroupByFilters()
	})
}

// ClearMeterGroupByFilters clears the value of the "meter_group_by_filters" field.
func (u *FeatureUpsertOne) ClearMeterGroupByFilters() *FeatureUpsertOne {
	return u.Update(func(s *FeatureUpsert) {
		s.ClearMeterGroupByFilters()
	})
}

// SetArchived sets the "archived" field.
func (u *FeatureUpsertOne) SetArchived(v bool) *FeatureUpsertOne {
	return u.Update(func(s *FeatureUpsert) {
		s.SetArchived(v)
	})
}

// UpdateArchived sets the "archived" field to the value that was provided on create.
func (u *FeatureUpsertOne) UpdateArchived() *FeatureUpsertOne {
	return u.Update(func(s *FeatureUpsert) {
		s.UpdateArchived()
	})
}

// Exec executes the query.
func (u *FeatureUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("db: missing options for FeatureCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *FeatureUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *FeatureUpsertOne) ID(ctx context.Context) (id string, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("db: FeatureUpsertOne.ID is not supported by MySQL driver. Use FeatureUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *FeatureUpsertOne) IDX(ctx context.Context) string {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// FeatureCreateBulk is the builder for creating many Feature entities in bulk.
type FeatureCreateBulk struct {
	config
	err      error
	builders []*FeatureCreate
	conflict []sql.ConflictOption
}

// Save creates the Feature entities in the database.
func (fcb *FeatureCreateBulk) Save(ctx context.Context) ([]*Feature, error) {
	if fcb.err != nil {
		return nil, fcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(fcb.builders))
	nodes := make([]*Feature, len(fcb.builders))
	mutators := make([]Mutator, len(fcb.builders))
	for i := range fcb.builders {
		func(i int, root context.Context) {
			builder := fcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*FeatureMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, fcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = fcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, fcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, fcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (fcb *FeatureCreateBulk) SaveX(ctx context.Context) []*Feature {
	v, err := fcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (fcb *FeatureCreateBulk) Exec(ctx context.Context) error {
	_, err := fcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fcb *FeatureCreateBulk) ExecX(ctx context.Context) {
	if err := fcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Feature.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.FeatureUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (fcb *FeatureCreateBulk) OnConflict(opts ...sql.ConflictOption) *FeatureUpsertBulk {
	fcb.conflict = opts
	return &FeatureUpsertBulk{
		create: fcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Feature.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (fcb *FeatureCreateBulk) OnConflictColumns(columns ...string) *FeatureUpsertBulk {
	fcb.conflict = append(fcb.conflict, sql.ConflictColumns(columns...))
	return &FeatureUpsertBulk{
		create: fcb,
	}
}

// FeatureUpsertBulk is the builder for "upsert"-ing
// a bulk of Feature nodes.
type FeatureUpsertBulk struct {
	create *FeatureCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Feature.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(feature.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *FeatureUpsertBulk) UpdateNewValues() *FeatureUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(feature.FieldID)
			}
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(feature.FieldCreatedAt)
			}
			if _, exists := b.mutation.Namespace(); exists {
				s.SetIgnore(feature.FieldNamespace)
			}
			if _, exists := b.mutation.MeterSlug(); exists {
				s.SetIgnore(feature.FieldMeterSlug)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Feature.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *FeatureUpsertBulk) Ignore() *FeatureUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *FeatureUpsertBulk) DoNothing() *FeatureUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the FeatureCreateBulk.OnConflict
// documentation for more info.
func (u *FeatureUpsertBulk) Update(set func(*FeatureUpsert)) *FeatureUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&FeatureUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *FeatureUpsertBulk) SetUpdatedAt(v time.Time) *FeatureUpsertBulk {
	return u.Update(func(s *FeatureUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *FeatureUpsertBulk) UpdateUpdatedAt() *FeatureUpsertBulk {
	return u.Update(func(s *FeatureUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetName sets the "name" field.
func (u *FeatureUpsertBulk) SetName(v string) *FeatureUpsertBulk {
	return u.Update(func(s *FeatureUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *FeatureUpsertBulk) UpdateName() *FeatureUpsertBulk {
	return u.Update(func(s *FeatureUpsert) {
		s.UpdateName()
	})
}

// SetMeterGroupByFilters sets the "meter_group_by_filters" field.
func (u *FeatureUpsertBulk) SetMeterGroupByFilters(v map[string]string) *FeatureUpsertBulk {
	return u.Update(func(s *FeatureUpsert) {
		s.SetMeterGroupByFilters(v)
	})
}

// UpdateMeterGroupByFilters sets the "meter_group_by_filters" field to the value that was provided on create.
func (u *FeatureUpsertBulk) UpdateMeterGroupByFilters() *FeatureUpsertBulk {
	return u.Update(func(s *FeatureUpsert) {
		s.UpdateMeterGroupByFilters()
	})
}

// ClearMeterGroupByFilters clears the value of the "meter_group_by_filters" field.
func (u *FeatureUpsertBulk) ClearMeterGroupByFilters() *FeatureUpsertBulk {
	return u.Update(func(s *FeatureUpsert) {
		s.ClearMeterGroupByFilters()
	})
}

// SetArchived sets the "archived" field.
func (u *FeatureUpsertBulk) SetArchived(v bool) *FeatureUpsertBulk {
	return u.Update(func(s *FeatureUpsert) {
		s.SetArchived(v)
	})
}

// UpdateArchived sets the "archived" field to the value that was provided on create.
func (u *FeatureUpsertBulk) UpdateArchived() *FeatureUpsertBulk {
	return u.Update(func(s *FeatureUpsert) {
		s.UpdateArchived()
	})
}

// Exec executes the query.
func (u *FeatureUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("db: OnConflict was set for builder %d. Set it on the FeatureCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("db: missing options for FeatureCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *FeatureUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}