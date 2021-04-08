// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"github.com/dreamvo/gilfoyle/ent/media"
	"github.com/dreamvo/gilfoyle/ent/mediafile"
	"github.com/dreamvo/gilfoyle/ent/predicate"
	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/google/uuid"
)

// MediaUpdate is the builder for updating Media entities.
type MediaUpdate struct {
	config
	hooks    []Hook
	mutation *MediaMutation
}

// Where adds a new predicate for the MediaUpdate builder.
func (mu *MediaUpdate) Where(ps ...predicate.Media) *MediaUpdate {
	mu.mutation.predicates = append(mu.mutation.predicates, ps...)
	return mu
}

// SetTitle sets the "title" field.
func (mu *MediaUpdate) SetTitle(s string) *MediaUpdate {
	mu.mutation.SetTitle(s)
	return mu
}

// SetOriginalFilename sets the "original_filename" field.
func (mu *MediaUpdate) SetOriginalFilename(s string) *MediaUpdate {
	mu.mutation.SetOriginalFilename(s)
	return mu
}

// SetNillableOriginalFilename sets the "original_filename" field if the given value is not nil.
func (mu *MediaUpdate) SetNillableOriginalFilename(s *string) *MediaUpdate {
	if s != nil {
		mu.SetOriginalFilename(*s)
	}
	return mu
}

// ClearOriginalFilename clears the value of the "original_filename" field.
func (mu *MediaUpdate) ClearOriginalFilename() *MediaUpdate {
	mu.mutation.ClearOriginalFilename()
	return mu
}

// SetStatus sets the "status" field.
func (mu *MediaUpdate) SetStatus(m media.Status) *MediaUpdate {
	mu.mutation.SetStatus(m)
	return mu
}

// SetCreatedAt sets the "created_at" field.
func (mu *MediaUpdate) SetCreatedAt(t time.Time) *MediaUpdate {
	mu.mutation.SetCreatedAt(t)
	return mu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (mu *MediaUpdate) SetNillableCreatedAt(t *time.Time) *MediaUpdate {
	if t != nil {
		mu.SetCreatedAt(*t)
	}
	return mu
}

// SetUpdatedAt sets the "updated_at" field.
func (mu *MediaUpdate) SetUpdatedAt(t time.Time) *MediaUpdate {
	mu.mutation.SetUpdatedAt(t)
	return mu
}

// AddMediaFileIDs adds the "media_files" edge to the MediaFile entity by IDs.
func (mu *MediaUpdate) AddMediaFileIDs(ids ...uuid.UUID) *MediaUpdate {
	mu.mutation.AddMediaFileIDs(ids...)
	return mu
}

// AddMediaFiles adds the "media_files" edges to the MediaFile entity.
func (mu *MediaUpdate) AddMediaFiles(m ...*MediaFile) *MediaUpdate {
	ids := make([]uuid.UUID, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return mu.AddMediaFileIDs(ids...)
}

// Mutation returns the MediaMutation object of the builder.
func (mu *MediaUpdate) Mutation() *MediaMutation {
	return mu.mutation
}

// ClearMediaFiles clears all "media_files" edges to the MediaFile entity.
func (mu *MediaUpdate) ClearMediaFiles() *MediaUpdate {
	mu.mutation.ClearMediaFiles()
	return mu
}

// RemoveMediaFileIDs removes the "media_files" edge to MediaFile entities by IDs.
func (mu *MediaUpdate) RemoveMediaFileIDs(ids ...uuid.UUID) *MediaUpdate {
	mu.mutation.RemoveMediaFileIDs(ids...)
	return mu
}

// RemoveMediaFiles removes "media_files" edges to MediaFile entities.
func (mu *MediaUpdate) RemoveMediaFiles(m ...*MediaFile) *MediaUpdate {
	ids := make([]uuid.UUID, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return mu.RemoveMediaFileIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (mu *MediaUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	mu.defaults()
	if len(mu.hooks) == 0 {
		if err = mu.check(); err != nil {
			return 0, err
		}
		affected, err = mu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MediaMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = mu.check(); err != nil {
				return 0, err
			}
			mu.mutation = mutation
			affected, err = mu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(mu.hooks) - 1; i >= 0; i-- {
			mut = mu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, mu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (mu *MediaUpdate) SaveX(ctx context.Context) int {
	affected, err := mu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mu *MediaUpdate) Exec(ctx context.Context) error {
	_, err := mu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mu *MediaUpdate) ExecX(ctx context.Context) {
	if err := mu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mu *MediaUpdate) defaults() {
	if _, ok := mu.mutation.UpdatedAt(); !ok {
		v := media.UpdateDefaultUpdatedAt()
		mu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mu *MediaUpdate) check() error {
	if v, ok := mu.mutation.Title(); ok {
		if err := media.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf("ent: validator failed for field \"title\": %w", err)}
		}
	}
	if v, ok := mu.mutation.OriginalFilename(); ok {
		if err := media.OriginalFilenameValidator(v); err != nil {
			return &ValidationError{Name: "original_filename", err: fmt.Errorf("ent: validator failed for field \"original_filename\": %w", err)}
		}
	}
	if v, ok := mu.mutation.Status(); ok {
		if err := media.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf("ent: validator failed for field \"status\": %w", err)}
		}
	}
	return nil
}

func (mu *MediaUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   media.Table,
			Columns: media.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: media.FieldID,
			},
		},
	}
	if ps := mu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mu.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: media.FieldTitle,
		})
	}
	if value, ok := mu.mutation.OriginalFilename(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: media.FieldOriginalFilename,
		})
	}
	if mu.mutation.OriginalFilenameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: media.FieldOriginalFilename,
		})
	}
	if value, ok := mu.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: media.FieldStatus,
		})
	}
	if value, ok := mu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: media.FieldCreatedAt,
		})
	}
	if value, ok := mu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: media.FieldUpdatedAt,
		})
	}
	if mu.mutation.MediaFilesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   media.MediaFilesTable,
			Columns: []string{media.MediaFilesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: mediafile.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.RemovedMediaFilesIDs(); len(nodes) > 0 && !mu.mutation.MediaFilesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   media.MediaFilesTable,
			Columns: []string{media.MediaFilesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: mediafile.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.MediaFilesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   media.MediaFilesTable,
			Columns: []string{media.MediaFilesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: mediafile.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, mu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{media.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// MediaUpdateOne is the builder for updating a single Media entity.
type MediaUpdateOne struct {
	config
	hooks    []Hook
	mutation *MediaMutation
}

// SetTitle sets the "title" field.
func (muo *MediaUpdateOne) SetTitle(s string) *MediaUpdateOne {
	muo.mutation.SetTitle(s)
	return muo
}

// SetOriginalFilename sets the "original_filename" field.
func (muo *MediaUpdateOne) SetOriginalFilename(s string) *MediaUpdateOne {
	muo.mutation.SetOriginalFilename(s)
	return muo
}

// SetNillableOriginalFilename sets the "original_filename" field if the given value is not nil.
func (muo *MediaUpdateOne) SetNillableOriginalFilename(s *string) *MediaUpdateOne {
	if s != nil {
		muo.SetOriginalFilename(*s)
	}
	return muo
}

// ClearOriginalFilename clears the value of the "original_filename" field.
func (muo *MediaUpdateOne) ClearOriginalFilename() *MediaUpdateOne {
	muo.mutation.ClearOriginalFilename()
	return muo
}

// SetStatus sets the "status" field.
func (muo *MediaUpdateOne) SetStatus(m media.Status) *MediaUpdateOne {
	muo.mutation.SetStatus(m)
	return muo
}

// SetCreatedAt sets the "created_at" field.
func (muo *MediaUpdateOne) SetCreatedAt(t time.Time) *MediaUpdateOne {
	muo.mutation.SetCreatedAt(t)
	return muo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (muo *MediaUpdateOne) SetNillableCreatedAt(t *time.Time) *MediaUpdateOne {
	if t != nil {
		muo.SetCreatedAt(*t)
	}
	return muo
}

// SetUpdatedAt sets the "updated_at" field.
func (muo *MediaUpdateOne) SetUpdatedAt(t time.Time) *MediaUpdateOne {
	muo.mutation.SetUpdatedAt(t)
	return muo
}

// AddMediaFileIDs adds the "media_files" edge to the MediaFile entity by IDs.
func (muo *MediaUpdateOne) AddMediaFileIDs(ids ...uuid.UUID) *MediaUpdateOne {
	muo.mutation.AddMediaFileIDs(ids...)
	return muo
}

// AddMediaFiles adds the "media_files" edges to the MediaFile entity.
func (muo *MediaUpdateOne) AddMediaFiles(m ...*MediaFile) *MediaUpdateOne {
	ids := make([]uuid.UUID, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return muo.AddMediaFileIDs(ids...)
}

// Mutation returns the MediaMutation object of the builder.
func (muo *MediaUpdateOne) Mutation() *MediaMutation {
	return muo.mutation
}

// ClearMediaFiles clears all "media_files" edges to the MediaFile entity.
func (muo *MediaUpdateOne) ClearMediaFiles() *MediaUpdateOne {
	muo.mutation.ClearMediaFiles()
	return muo
}

// RemoveMediaFileIDs removes the "media_files" edge to MediaFile entities by IDs.
func (muo *MediaUpdateOne) RemoveMediaFileIDs(ids ...uuid.UUID) *MediaUpdateOne {
	muo.mutation.RemoveMediaFileIDs(ids...)
	return muo
}

// RemoveMediaFiles removes "media_files" edges to MediaFile entities.
func (muo *MediaUpdateOne) RemoveMediaFiles(m ...*MediaFile) *MediaUpdateOne {
	ids := make([]uuid.UUID, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return muo.RemoveMediaFileIDs(ids...)
}

// Save executes the query and returns the updated Media entity.
func (muo *MediaUpdateOne) Save(ctx context.Context) (*Media, error) {
	var (
		err  error
		node *Media
	)
	muo.defaults()
	if len(muo.hooks) == 0 {
		if err = muo.check(); err != nil {
			return nil, err
		}
		node, err = muo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MediaMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = muo.check(); err != nil {
				return nil, err
			}
			muo.mutation = mutation
			node, err = muo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(muo.hooks) - 1; i >= 0; i-- {
			mut = muo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, muo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (muo *MediaUpdateOne) SaveX(ctx context.Context) *Media {
	node, err := muo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (muo *MediaUpdateOne) Exec(ctx context.Context) error {
	_, err := muo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (muo *MediaUpdateOne) ExecX(ctx context.Context) {
	if err := muo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (muo *MediaUpdateOne) defaults() {
	if _, ok := muo.mutation.UpdatedAt(); !ok {
		v := media.UpdateDefaultUpdatedAt()
		muo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (muo *MediaUpdateOne) check() error {
	if v, ok := muo.mutation.Title(); ok {
		if err := media.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf("ent: validator failed for field \"title\": %w", err)}
		}
	}
	if v, ok := muo.mutation.OriginalFilename(); ok {
		if err := media.OriginalFilenameValidator(v); err != nil {
			return &ValidationError{Name: "original_filename", err: fmt.Errorf("ent: validator failed for field \"original_filename\": %w", err)}
		}
	}
	if v, ok := muo.mutation.Status(); ok {
		if err := media.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf("ent: validator failed for field \"status\": %w", err)}
		}
	}
	return nil
}

func (muo *MediaUpdateOne) sqlSave(ctx context.Context) (_node *Media, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   media.Table,
			Columns: media.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: media.FieldID,
			},
		},
	}
	id, ok := muo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Media.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := muo.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: media.FieldTitle,
		})
	}
	if value, ok := muo.mutation.OriginalFilename(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: media.FieldOriginalFilename,
		})
	}
	if muo.mutation.OriginalFilenameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: media.FieldOriginalFilename,
		})
	}
	if value, ok := muo.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: media.FieldStatus,
		})
	}
	if value, ok := muo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: media.FieldCreatedAt,
		})
	}
	if value, ok := muo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: media.FieldUpdatedAt,
		})
	}
	if muo.mutation.MediaFilesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   media.MediaFilesTable,
			Columns: []string{media.MediaFilesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: mediafile.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.RemovedMediaFilesIDs(); len(nodes) > 0 && !muo.mutation.MediaFilesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   media.MediaFilesTable,
			Columns: []string{media.MediaFilesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: mediafile.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.MediaFilesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   media.MediaFilesTable,
			Columns: []string{media.MediaFilesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: mediafile.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Media{config: muo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, muo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{media.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
