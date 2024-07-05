// Code generated by entc, DO NOT EDIT.

package ent

import (
	"comment/app/job/internal/data/ent/commentcontent"
	"comment/app/job/internal/data/ent/predicate"
	"context"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CommentContentUpdate is the builder for updating CommentContent entities.
type CommentContentUpdate struct {
	config
	hooks    []Hook
	mutation *CommentContentMutation
}

// Where appends a list predicates to the CommentContentUpdate builder.
func (ccu *CommentContentUpdate) Where(ps ...predicate.CommentContent) *CommentContentUpdate {
	ccu.mutation.Where(ps...)
	return ccu
}

// SetObjID sets the "obj_id" field.
func (ccu *CommentContentUpdate) SetObjID(i int64) *CommentContentUpdate {
	ccu.mutation.ResetObjID()
	ccu.mutation.SetObjID(i)
	return ccu
}

// AddObjID adds i to the "obj_id" field.
func (ccu *CommentContentUpdate) AddObjID(i int64) *CommentContentUpdate {
	ccu.mutation.AddObjID(i)
	return ccu
}

// SetOwnerID sets the "owner_id" field.
func (ccu *CommentContentUpdate) SetOwnerID(i int64) *CommentContentUpdate {
	ccu.mutation.ResetOwnerID()
	ccu.mutation.SetOwnerID(i)
	return ccu
}

// AddOwnerID adds i to the "owner_id" field.
func (ccu *CommentContentUpdate) AddOwnerID(i int64) *CommentContentUpdate {
	ccu.mutation.AddOwnerID(i)
	return ccu
}

// SetRoot sets the "root" field.
func (ccu *CommentContentUpdate) SetRoot(i int64) *CommentContentUpdate {
	ccu.mutation.ResetRoot()
	ccu.mutation.SetRoot(i)
	return ccu
}

// AddRoot adds i to the "root" field.
func (ccu *CommentContentUpdate) AddRoot(i int64) *CommentContentUpdate {
	ccu.mutation.AddRoot(i)
	return ccu
}

// SetFloor sets the "floor" field.
func (ccu *CommentContentUpdate) SetFloor(i int32) *CommentContentUpdate {
	ccu.mutation.ResetFloor()
	ccu.mutation.SetFloor(i)
	return ccu
}

// AddFloor adds i to the "floor" field.
func (ccu *CommentContentUpdate) AddFloor(i int32) *CommentContentUpdate {
	ccu.mutation.AddFloor(i)
	return ccu
}

// SetMessage sets the "message" field.
func (ccu *CommentContentUpdate) SetMessage(s string) *CommentContentUpdate {
	ccu.mutation.SetMessage(s)
	return ccu
}

// SetNillableMessage sets the "message" field if the given value is not nil.
func (ccu *CommentContentUpdate) SetNillableMessage(s *string) *CommentContentUpdate {
	if s != nil {
		ccu.SetMessage(*s)
	}
	return ccu
}

// SetCreatedAt sets the "created_at" field.
func (ccu *CommentContentUpdate) SetCreatedAt(t time.Time) *CommentContentUpdate {
	ccu.mutation.SetCreatedAt(t)
	return ccu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ccu *CommentContentUpdate) SetNillableCreatedAt(t *time.Time) *CommentContentUpdate {
	if t != nil {
		ccu.SetCreatedAt(*t)
	}
	return ccu
}

// SetUpdatedAt sets the "updated_at" field.
func (ccu *CommentContentUpdate) SetUpdatedAt(t time.Time) *CommentContentUpdate {
	ccu.mutation.SetUpdatedAt(t)
	return ccu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ccu *CommentContentUpdate) SetNillableUpdatedAt(t *time.Time) *CommentContentUpdate {
	if t != nil {
		ccu.SetUpdatedAt(*t)
	}
	return ccu
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (ccu *CommentContentUpdate) ClearUpdatedAt() *CommentContentUpdate {
	ccu.mutation.ClearUpdatedAt()
	return ccu
}

// Mutation returns the CommentContentMutation object of the builder.
func (ccu *CommentContentUpdate) Mutation() *CommentContentMutation {
	return ccu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ccu *CommentContentUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ccu.hooks) == 0 {
		affected, err = ccu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CommentContentMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ccu.mutation = mutation
			affected, err = ccu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ccu.hooks) - 1; i >= 0; i-- {
			if ccu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ccu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ccu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ccu *CommentContentUpdate) SaveX(ctx context.Context) int {
	affected, err := ccu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ccu *CommentContentUpdate) Exec(ctx context.Context) error {
	_, err := ccu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccu *CommentContentUpdate) ExecX(ctx context.Context) {
	if err := ccu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ccu *CommentContentUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   commentcontent.Table,
			Columns: commentcontent.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: commentcontent.FieldID,
			},
		},
	}
	if ps := ccu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ccu.mutation.ObjID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: commentcontent.FieldObjID,
		})
	}
	if value, ok := ccu.mutation.AddedObjID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: commentcontent.FieldObjID,
		})
	}
	if value, ok := ccu.mutation.OwnerID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: commentcontent.FieldOwnerID,
		})
	}
	if value, ok := ccu.mutation.AddedOwnerID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: commentcontent.FieldOwnerID,
		})
	}
	if value, ok := ccu.mutation.Root(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: commentcontent.FieldRoot,
		})
	}
	if value, ok := ccu.mutation.AddedRoot(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: commentcontent.FieldRoot,
		})
	}
	if value, ok := ccu.mutation.Floor(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: commentcontent.FieldFloor,
		})
	}
	if value, ok := ccu.mutation.AddedFloor(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: commentcontent.FieldFloor,
		})
	}
	if value, ok := ccu.mutation.Message(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: commentcontent.FieldMessage,
		})
	}
	if value, ok := ccu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: commentcontent.FieldCreatedAt,
		})
	}
	if value, ok := ccu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: commentcontent.FieldUpdatedAt,
		})
	}
	if ccu.mutation.UpdatedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: commentcontent.FieldUpdatedAt,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ccu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{commentcontent.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// CommentContentUpdateOne is the builder for updating a single CommentContent entity.
type CommentContentUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CommentContentMutation
}

// SetObjID sets the "obj_id" field.
func (ccuo *CommentContentUpdateOne) SetObjID(i int64) *CommentContentUpdateOne {
	ccuo.mutation.ResetObjID()
	ccuo.mutation.SetObjID(i)
	return ccuo
}

// AddObjID adds i to the "obj_id" field.
func (ccuo *CommentContentUpdateOne) AddObjID(i int64) *CommentContentUpdateOne {
	ccuo.mutation.AddObjID(i)
	return ccuo
}

// SetOwnerID sets the "owner_id" field.
func (ccuo *CommentContentUpdateOne) SetOwnerID(i int64) *CommentContentUpdateOne {
	ccuo.mutation.ResetOwnerID()
	ccuo.mutation.SetOwnerID(i)
	return ccuo
}

// AddOwnerID adds i to the "owner_id" field.
func (ccuo *CommentContentUpdateOne) AddOwnerID(i int64) *CommentContentUpdateOne {
	ccuo.mutation.AddOwnerID(i)
	return ccuo
}

// SetRoot sets the "root" field.
func (ccuo *CommentContentUpdateOne) SetRoot(i int64) *CommentContentUpdateOne {
	ccuo.mutation.ResetRoot()
	ccuo.mutation.SetRoot(i)
	return ccuo
}

// AddRoot adds i to the "root" field.
func (ccuo *CommentContentUpdateOne) AddRoot(i int64) *CommentContentUpdateOne {
	ccuo.mutation.AddRoot(i)
	return ccuo
}

// SetFloor sets the "floor" field.
func (ccuo *CommentContentUpdateOne) SetFloor(i int32) *CommentContentUpdateOne {
	ccuo.mutation.ResetFloor()
	ccuo.mutation.SetFloor(i)
	return ccuo
}

// AddFloor adds i to the "floor" field.
func (ccuo *CommentContentUpdateOne) AddFloor(i int32) *CommentContentUpdateOne {
	ccuo.mutation.AddFloor(i)
	return ccuo
}

// SetMessage sets the "message" field.
func (ccuo *CommentContentUpdateOne) SetMessage(s string) *CommentContentUpdateOne {
	ccuo.mutation.SetMessage(s)
	return ccuo
}

// SetNillableMessage sets the "message" field if the given value is not nil.
func (ccuo *CommentContentUpdateOne) SetNillableMessage(s *string) *CommentContentUpdateOne {
	if s != nil {
		ccuo.SetMessage(*s)
	}
	return ccuo
}

// SetCreatedAt sets the "created_at" field.
func (ccuo *CommentContentUpdateOne) SetCreatedAt(t time.Time) *CommentContentUpdateOne {
	ccuo.mutation.SetCreatedAt(t)
	return ccuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ccuo *CommentContentUpdateOne) SetNillableCreatedAt(t *time.Time) *CommentContentUpdateOne {
	if t != nil {
		ccuo.SetCreatedAt(*t)
	}
	return ccuo
}

// SetUpdatedAt sets the "updated_at" field.
func (ccuo *CommentContentUpdateOne) SetUpdatedAt(t time.Time) *CommentContentUpdateOne {
	ccuo.mutation.SetUpdatedAt(t)
	return ccuo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ccuo *CommentContentUpdateOne) SetNillableUpdatedAt(t *time.Time) *CommentContentUpdateOne {
	if t != nil {
		ccuo.SetUpdatedAt(*t)
	}
	return ccuo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (ccuo *CommentContentUpdateOne) ClearUpdatedAt() *CommentContentUpdateOne {
	ccuo.mutation.ClearUpdatedAt()
	return ccuo
}

// Mutation returns the CommentContentMutation object of the builder.
func (ccuo *CommentContentUpdateOne) Mutation() *CommentContentMutation {
	return ccuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ccuo *CommentContentUpdateOne) Select(field string, fields ...string) *CommentContentUpdateOne {
	ccuo.fields = append([]string{field}, fields...)
	return ccuo
}

// Save executes the query and returns the updated CommentContent entity.
func (ccuo *CommentContentUpdateOne) Save(ctx context.Context) (*CommentContent, error) {
	var (
		err  error
		node *CommentContent
	)
	if len(ccuo.hooks) == 0 {
		node, err = ccuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CommentContentMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ccuo.mutation = mutation
			node, err = ccuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ccuo.hooks) - 1; i >= 0; i-- {
			if ccuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ccuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ccuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ccuo *CommentContentUpdateOne) SaveX(ctx context.Context) *CommentContent {
	node, err := ccuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ccuo *CommentContentUpdateOne) Exec(ctx context.Context) error {
	_, err := ccuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccuo *CommentContentUpdateOne) ExecX(ctx context.Context) {
	if err := ccuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ccuo *CommentContentUpdateOne) sqlSave(ctx context.Context) (_node *CommentContent, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   commentcontent.Table,
			Columns: commentcontent.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: commentcontent.FieldID,
			},
		},
	}
	id, ok := ccuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing CommentContent.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := ccuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, commentcontent.FieldID)
		for _, f := range fields {
			if !commentcontent.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != commentcontent.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ccuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ccuo.mutation.ObjID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: commentcontent.FieldObjID,
		})
	}
	if value, ok := ccuo.mutation.AddedObjID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: commentcontent.FieldObjID,
		})
	}
	if value, ok := ccuo.mutation.OwnerID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: commentcontent.FieldOwnerID,
		})
	}
	if value, ok := ccuo.mutation.AddedOwnerID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: commentcontent.FieldOwnerID,
		})
	}
	if value, ok := ccuo.mutation.Root(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: commentcontent.FieldRoot,
		})
	}
	if value, ok := ccuo.mutation.AddedRoot(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: commentcontent.FieldRoot,
		})
	}
	if value, ok := ccuo.mutation.Floor(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: commentcontent.FieldFloor,
		})
	}
	if value, ok := ccuo.mutation.AddedFloor(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: commentcontent.FieldFloor,
		})
	}
	if value, ok := ccuo.mutation.Message(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: commentcontent.FieldMessage,
		})
	}
	if value, ok := ccuo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: commentcontent.FieldCreatedAt,
		})
	}
	if value, ok := ccuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: commentcontent.FieldUpdatedAt,
		})
	}
	if ccuo.mutation.UpdatedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: commentcontent.FieldUpdatedAt,
		})
	}
	_node = &CommentContent{config: ccuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ccuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{commentcontent.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}