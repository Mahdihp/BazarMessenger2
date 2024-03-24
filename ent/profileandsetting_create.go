// Code generated by ent, DO NOT EDIT.

package ent

import (
	"BazarMessenger/ent/profileandsetting"
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ProfileAndSettingCreate is the builder for creating a ProfileAndSetting entity.
type ProfileAndSettingCreate struct {
	config
	mutation *ProfileAndSettingMutation
	hooks    []Hook
}

// Mutation returns the ProfileAndSettingMutation object of the builder.
func (pasc *ProfileAndSettingCreate) Mutation() *ProfileAndSettingMutation {
	return pasc.mutation
}

// Save creates the ProfileAndSetting in the database.
func (pasc *ProfileAndSettingCreate) Save(ctx context.Context) (*ProfileAndSetting, error) {
	return withHooks(ctx, pasc.sqlSave, pasc.mutation, pasc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (pasc *ProfileAndSettingCreate) SaveX(ctx context.Context) *ProfileAndSetting {
	v, err := pasc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pasc *ProfileAndSettingCreate) Exec(ctx context.Context) error {
	_, err := pasc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pasc *ProfileAndSettingCreate) ExecX(ctx context.Context) {
	if err := pasc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pasc *ProfileAndSettingCreate) check() error {
	return nil
}

func (pasc *ProfileAndSettingCreate) sqlSave(ctx context.Context) (*ProfileAndSetting, error) {
	if err := pasc.check(); err != nil {
		return nil, err
	}
	_node, _spec := pasc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pasc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	pasc.mutation.id = &_node.ID
	pasc.mutation.done = true
	return _node, nil
}

func (pasc *ProfileAndSettingCreate) createSpec() (*ProfileAndSetting, *sqlgraph.CreateSpec) {
	var (
		_node = &ProfileAndSetting{config: pasc.config}
		_spec = sqlgraph.NewCreateSpec(profileandsetting.Table, sqlgraph.NewFieldSpec(profileandsetting.FieldID, field.TypeInt))
	)
	return _node, _spec
}

// ProfileAndSettingCreateBulk is the builder for creating many ProfileAndSetting entities in bulk.
type ProfileAndSettingCreateBulk struct {
	config
	err      error
	builders []*ProfileAndSettingCreate
}

// Save creates the ProfileAndSetting entities in the database.
func (pascb *ProfileAndSettingCreateBulk) Save(ctx context.Context) ([]*ProfileAndSetting, error) {
	if pascb.err != nil {
		return nil, pascb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(pascb.builders))
	nodes := make([]*ProfileAndSetting, len(pascb.builders))
	mutators := make([]Mutator, len(pascb.builders))
	for i := range pascb.builders {
		func(i int, root context.Context) {
			builder := pascb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ProfileAndSettingMutation)
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
					_, err = mutators[i+1].Mutate(root, pascb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pascb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
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
		if _, err := mutators[0].Mutate(ctx, pascb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pascb *ProfileAndSettingCreateBulk) SaveX(ctx context.Context) []*ProfileAndSetting {
	v, err := pascb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pascb *ProfileAndSettingCreateBulk) Exec(ctx context.Context) error {
	_, err := pascb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pascb *ProfileAndSettingCreateBulk) ExecX(ctx context.Context) {
	if err := pascb.Exec(ctx); err != nil {
		panic(err)
	}
}
