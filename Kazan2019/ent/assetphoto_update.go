// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/qin-guan/WorldSkills/Kazan2019/ent/asset"
	"github.com/qin-guan/WorldSkills/Kazan2019/ent/assetphoto"
	"github.com/qin-guan/WorldSkills/Kazan2019/ent/predicate"
)

// AssetPhotoUpdate is the builder for updating AssetPhoto entities.
type AssetPhotoUpdate struct {
	config
	hooks    []Hook
	mutation *AssetPhotoMutation
}

// Where appends a list predicates to the AssetPhotoUpdate builder.
func (apu *AssetPhotoUpdate) Where(ps ...predicate.AssetPhoto) *AssetPhotoUpdate {
	apu.mutation.Where(ps...)
	return apu
}

// SetAssetPhoto sets the "AssetPhoto" field.
func (apu *AssetPhotoUpdate) SetAssetPhoto(s string) *AssetPhotoUpdate {
	apu.mutation.SetAssetPhoto(s)
	return apu
}

// SetAssetID sets the "Asset" edge to the Asset entity by ID.
func (apu *AssetPhotoUpdate) SetAssetID(id int) *AssetPhotoUpdate {
	apu.mutation.SetAssetID(id)
	return apu
}

// SetAsset sets the "Asset" edge to the Asset entity.
func (apu *AssetPhotoUpdate) SetAsset(a *Asset) *AssetPhotoUpdate {
	return apu.SetAssetID(a.ID)
}

// Mutation returns the AssetPhotoMutation object of the builder.
func (apu *AssetPhotoUpdate) Mutation() *AssetPhotoMutation {
	return apu.mutation
}

// ClearAsset clears the "Asset" edge to the Asset entity.
func (apu *AssetPhotoUpdate) ClearAsset() *AssetPhotoUpdate {
	apu.mutation.ClearAsset()
	return apu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (apu *AssetPhotoUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, AssetPhotoMutation](ctx, apu.sqlSave, apu.mutation, apu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (apu *AssetPhotoUpdate) SaveX(ctx context.Context) int {
	affected, err := apu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (apu *AssetPhotoUpdate) Exec(ctx context.Context) error {
	_, err := apu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (apu *AssetPhotoUpdate) ExecX(ctx context.Context) {
	if err := apu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (apu *AssetPhotoUpdate) check() error {
	if _, ok := apu.mutation.AssetID(); apu.mutation.AssetCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "AssetPhoto.Asset"`)
	}
	return nil
}

func (apu *AssetPhotoUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := apu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(assetphoto.Table, assetphoto.Columns, sqlgraph.NewFieldSpec(assetphoto.FieldID, field.TypeInt))
	if ps := apu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := apu.mutation.AssetPhoto(); ok {
		_spec.SetField(assetphoto.FieldAssetPhoto, field.TypeString, value)
	}
	if apu.mutation.AssetCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   assetphoto.AssetTable,
			Columns: []string{assetphoto.AssetColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(asset.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := apu.mutation.AssetIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   assetphoto.AssetTable,
			Columns: []string{assetphoto.AssetColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(asset.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, apu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{assetphoto.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	apu.mutation.done = true
	return n, nil
}

// AssetPhotoUpdateOne is the builder for updating a single AssetPhoto entity.
type AssetPhotoUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AssetPhotoMutation
}

// SetAssetPhoto sets the "AssetPhoto" field.
func (apuo *AssetPhotoUpdateOne) SetAssetPhoto(s string) *AssetPhotoUpdateOne {
	apuo.mutation.SetAssetPhoto(s)
	return apuo
}

// SetAssetID sets the "Asset" edge to the Asset entity by ID.
func (apuo *AssetPhotoUpdateOne) SetAssetID(id int) *AssetPhotoUpdateOne {
	apuo.mutation.SetAssetID(id)
	return apuo
}

// SetAsset sets the "Asset" edge to the Asset entity.
func (apuo *AssetPhotoUpdateOne) SetAsset(a *Asset) *AssetPhotoUpdateOne {
	return apuo.SetAssetID(a.ID)
}

// Mutation returns the AssetPhotoMutation object of the builder.
func (apuo *AssetPhotoUpdateOne) Mutation() *AssetPhotoMutation {
	return apuo.mutation
}

// ClearAsset clears the "Asset" edge to the Asset entity.
func (apuo *AssetPhotoUpdateOne) ClearAsset() *AssetPhotoUpdateOne {
	apuo.mutation.ClearAsset()
	return apuo
}

// Where appends a list predicates to the AssetPhotoUpdate builder.
func (apuo *AssetPhotoUpdateOne) Where(ps ...predicate.AssetPhoto) *AssetPhotoUpdateOne {
	apuo.mutation.Where(ps...)
	return apuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (apuo *AssetPhotoUpdateOne) Select(field string, fields ...string) *AssetPhotoUpdateOne {
	apuo.fields = append([]string{field}, fields...)
	return apuo
}

// Save executes the query and returns the updated AssetPhoto entity.
func (apuo *AssetPhotoUpdateOne) Save(ctx context.Context) (*AssetPhoto, error) {
	return withHooks[*AssetPhoto, AssetPhotoMutation](ctx, apuo.sqlSave, apuo.mutation, apuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (apuo *AssetPhotoUpdateOne) SaveX(ctx context.Context) *AssetPhoto {
	node, err := apuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (apuo *AssetPhotoUpdateOne) Exec(ctx context.Context) error {
	_, err := apuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (apuo *AssetPhotoUpdateOne) ExecX(ctx context.Context) {
	if err := apuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (apuo *AssetPhotoUpdateOne) check() error {
	if _, ok := apuo.mutation.AssetID(); apuo.mutation.AssetCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "AssetPhoto.Asset"`)
	}
	return nil
}

func (apuo *AssetPhotoUpdateOne) sqlSave(ctx context.Context) (_node *AssetPhoto, err error) {
	if err := apuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(assetphoto.Table, assetphoto.Columns, sqlgraph.NewFieldSpec(assetphoto.FieldID, field.TypeInt))
	id, ok := apuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "AssetPhoto.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := apuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, assetphoto.FieldID)
		for _, f := range fields {
			if !assetphoto.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != assetphoto.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := apuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := apuo.mutation.AssetPhoto(); ok {
		_spec.SetField(assetphoto.FieldAssetPhoto, field.TypeString, value)
	}
	if apuo.mutation.AssetCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   assetphoto.AssetTable,
			Columns: []string{assetphoto.AssetColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(asset.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := apuo.mutation.AssetIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   assetphoto.AssetTable,
			Columns: []string{assetphoto.AssetColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(asset.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &AssetPhoto{config: apuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, apuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{assetphoto.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	apuo.mutation.done = true
	return _node, nil
}