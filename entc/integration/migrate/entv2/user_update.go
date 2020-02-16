// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package entv2

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/entc/integration/migrate/entv2/car"
	"github.com/facebookincubator/ent/entc/integration/migrate/entv2/pet"
	"github.com/facebookincubator/ent/entc/integration/migrate/entv2/predicate"
	"github.com/facebookincubator/ent/entc/integration/migrate/entv2/user"
	"github.com/facebookincubator/ent/schema/field"
)

// UserUpdate is the builder for updating User entities.
type UserUpdate struct {
	config
	age           *int
	addage        *int
	name          *string
	nickname      *string
	phone         *string
	buffer        *[]byte
	clearbuffer   bool
	title         *string
	new_name      *string
	clearnew_name bool
	blob          *[]byte
	clearblob     bool
	state         *user.State
	clearstate    bool
	car           map[int]struct{}
	pets          map[int]struct{}
	removedCar    map[int]struct{}
	clearedPets   bool
	predicates    []predicate.User
}

// Where adds a new predicate for the builder.
func (uu *UserUpdate) Where(ps ...predicate.User) *UserUpdate {
	uu.predicates = append(uu.predicates, ps...)
	return uu
}

// SetAge sets the age field.
func (uu *UserUpdate) SetAge(i int) *UserUpdate {
	uu.age = &i
	uu.addage = nil
	return uu
}

// AddAge adds i to age.
func (uu *UserUpdate) AddAge(i int) *UserUpdate {
	if uu.addage == nil {
		uu.addage = &i
	} else {
		*uu.addage += i
	}
	return uu
}

// SetName sets the name field.
func (uu *UserUpdate) SetName(s string) *UserUpdate {
	uu.name = &s
	return uu
}

// SetNickname sets the nickname field.
func (uu *UserUpdate) SetNickname(s string) *UserUpdate {
	uu.nickname = &s
	return uu
}

// SetPhone sets the phone field.
func (uu *UserUpdate) SetPhone(s string) *UserUpdate {
	uu.phone = &s
	return uu
}

// SetNillablePhone sets the phone field if the given value is not nil.
func (uu *UserUpdate) SetNillablePhone(s *string) *UserUpdate {
	if s != nil {
		uu.SetPhone(*s)
	}
	return uu
}

// SetBuffer sets the buffer field.
func (uu *UserUpdate) SetBuffer(b []byte) *UserUpdate {
	uu.buffer = &b
	return uu
}

// ClearBuffer clears the value of buffer.
func (uu *UserUpdate) ClearBuffer() *UserUpdate {
	uu.buffer = nil
	uu.clearbuffer = true
	return uu
}

// SetTitle sets the title field.
func (uu *UserUpdate) SetTitle(s string) *UserUpdate {
	uu.title = &s
	return uu
}

// SetNillableTitle sets the title field if the given value is not nil.
func (uu *UserUpdate) SetNillableTitle(s *string) *UserUpdate {
	if s != nil {
		uu.SetTitle(*s)
	}
	return uu
}

// SetNewName sets the new_name field.
func (uu *UserUpdate) SetNewName(s string) *UserUpdate {
	uu.new_name = &s
	return uu
}

// SetNillableNewName sets the new_name field if the given value is not nil.
func (uu *UserUpdate) SetNillableNewName(s *string) *UserUpdate {
	if s != nil {
		uu.SetNewName(*s)
	}
	return uu
}

// ClearNewName clears the value of new_name.
func (uu *UserUpdate) ClearNewName() *UserUpdate {
	uu.new_name = nil
	uu.clearnew_name = true
	return uu
}

// SetBlob sets the blob field.
func (uu *UserUpdate) SetBlob(b []byte) *UserUpdate {
	uu.blob = &b
	return uu
}

// ClearBlob clears the value of blob.
func (uu *UserUpdate) ClearBlob() *UserUpdate {
	uu.blob = nil
	uu.clearblob = true
	return uu
}

// SetState sets the state field.
func (uu *UserUpdate) SetState(u user.State) *UserUpdate {
	uu.state = &u
	return uu
}

// SetNillableState sets the state field if the given value is not nil.
func (uu *UserUpdate) SetNillableState(u *user.State) *UserUpdate {
	if u != nil {
		uu.SetState(*u)
	}
	return uu
}

// ClearState clears the value of state.
func (uu *UserUpdate) ClearState() *UserUpdate {
	uu.state = nil
	uu.clearstate = true
	return uu
}

// AddCarIDs adds the car edge to Car by ids.
func (uu *UserUpdate) AddCarIDs(ids ...int) *UserUpdate {
	if uu.car == nil {
		uu.car = make(map[int]struct{})
	}
	for i := range ids {
		uu.car[ids[i]] = struct{}{}
	}
	return uu
}

// AddCar adds the car edges to Car.
func (uu *UserUpdate) AddCar(c ...*Car) *UserUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return uu.AddCarIDs(ids...)
}

// SetPetsID sets the pets edge to Pet by id.
func (uu *UserUpdate) SetPetsID(id int) *UserUpdate {
	if uu.pets == nil {
		uu.pets = make(map[int]struct{})
	}
	uu.pets[id] = struct{}{}
	return uu
}

// SetNillablePetsID sets the pets edge to Pet by id if the given value is not nil.
func (uu *UserUpdate) SetNillablePetsID(id *int) *UserUpdate {
	if id != nil {
		uu = uu.SetPetsID(*id)
	}
	return uu
}

// SetPets sets the pets edge to Pet.
func (uu *UserUpdate) SetPets(p *Pet) *UserUpdate {
	return uu.SetPetsID(p.ID)
}

// RemoveCarIDs removes the car edge to Car by ids.
func (uu *UserUpdate) RemoveCarIDs(ids ...int) *UserUpdate {
	if uu.removedCar == nil {
		uu.removedCar = make(map[int]struct{})
	}
	for i := range ids {
		uu.removedCar[ids[i]] = struct{}{}
	}
	return uu
}

// RemoveCar removes car edges to Car.
func (uu *UserUpdate) RemoveCar(c ...*Car) *UserUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return uu.RemoveCarIDs(ids...)
}

// ClearPets clears the pets edge to Pet.
func (uu *UserUpdate) ClearPets() *UserUpdate {
	uu.clearedPets = true
	return uu
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (uu *UserUpdate) Save(ctx context.Context) (int, error) {
	if uu.state != nil {
		if err := user.StateValidator(*uu.state); err != nil {
			return 0, fmt.Errorf("entv2: validator failed for field \"state\": %v", err)
		}
	}
	if len(uu.pets) > 1 {
		return 0, errors.New("entv2: multiple assignments on a unique edge \"pets\"")
	}
	return uu.sqlSave(ctx)
}

// SaveX is like Save, but panics if an error occurs.
func (uu *UserUpdate) SaveX(ctx context.Context) int {
	affected, err := uu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uu *UserUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uu *UserUpdate) ExecX(ctx context.Context) {
	if err := uu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (uu *UserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   user.Table,
			Columns: user.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: user.FieldID,
			},
		},
	}
	if ps := uu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value := uu.age; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  *value,
			Column: user.FieldAge,
		})
	}
	if value := uu.addage; value != nil {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  *value,
			Column: user.FieldAge,
		})
	}
	if value := uu.name; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: user.FieldName,
		})
	}
	if value := uu.nickname; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: user.FieldNickname,
		})
	}
	if value := uu.phone; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: user.FieldPhone,
		})
	}
	if value := uu.buffer; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBytes,
			Value:  *value,
			Column: user.FieldBuffer,
		})
	}
	if uu.clearbuffer {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeBytes,
			Column: user.FieldBuffer,
		})
	}
	if value := uu.title; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: user.FieldTitle,
		})
	}
	if value := uu.new_name; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: user.FieldNewName,
		})
	}
	if uu.clearnew_name {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: user.FieldNewName,
		})
	}
	if value := uu.blob; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBytes,
			Value:  *value,
			Column: user.FieldBlob,
		})
	}
	if uu.clearblob {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeBytes,
			Column: user.FieldBlob,
		})
	}
	if value := uu.state; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  *value,
			Column: user.FieldState,
		})
	}
	if uu.clearstate {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Column: user.FieldState,
		})
	}
	if nodes := uu.removedCar; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.CarTable,
			Columns: []string{user.CarColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: car.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.car; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.CarTable,
			Columns: []string{user.CarColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: car.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.clearedPets {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   user.PetsTable,
			Columns: []string{user.PetsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: pet.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.pets; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   user.PetsTable,
			Columns: []string{user.PetsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: pet.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// UserUpdateOne is the builder for updating a single User entity.
type UserUpdateOne struct {
	config
	id            int
	age           *int
	addage        *int
	name          *string
	nickname      *string
	phone         *string
	buffer        *[]byte
	clearbuffer   bool
	title         *string
	new_name      *string
	clearnew_name bool
	blob          *[]byte
	clearblob     bool
	state         *user.State
	clearstate    bool
	car           map[int]struct{}
	pets          map[int]struct{}
	removedCar    map[int]struct{}
	clearedPets   bool
}

// SetAge sets the age field.
func (uuo *UserUpdateOne) SetAge(i int) *UserUpdateOne {
	uuo.age = &i
	uuo.addage = nil
	return uuo
}

// AddAge adds i to age.
func (uuo *UserUpdateOne) AddAge(i int) *UserUpdateOne {
	if uuo.addage == nil {
		uuo.addage = &i
	} else {
		*uuo.addage += i
	}
	return uuo
}

// SetName sets the name field.
func (uuo *UserUpdateOne) SetName(s string) *UserUpdateOne {
	uuo.name = &s
	return uuo
}

// SetNickname sets the nickname field.
func (uuo *UserUpdateOne) SetNickname(s string) *UserUpdateOne {
	uuo.nickname = &s
	return uuo
}

// SetPhone sets the phone field.
func (uuo *UserUpdateOne) SetPhone(s string) *UserUpdateOne {
	uuo.phone = &s
	return uuo
}

// SetNillablePhone sets the phone field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillablePhone(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetPhone(*s)
	}
	return uuo
}

// SetBuffer sets the buffer field.
func (uuo *UserUpdateOne) SetBuffer(b []byte) *UserUpdateOne {
	uuo.buffer = &b
	return uuo
}

// ClearBuffer clears the value of buffer.
func (uuo *UserUpdateOne) ClearBuffer() *UserUpdateOne {
	uuo.buffer = nil
	uuo.clearbuffer = true
	return uuo
}

// SetTitle sets the title field.
func (uuo *UserUpdateOne) SetTitle(s string) *UserUpdateOne {
	uuo.title = &s
	return uuo
}

// SetNillableTitle sets the title field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableTitle(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetTitle(*s)
	}
	return uuo
}

// SetNewName sets the new_name field.
func (uuo *UserUpdateOne) SetNewName(s string) *UserUpdateOne {
	uuo.new_name = &s
	return uuo
}

// SetNillableNewName sets the new_name field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableNewName(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetNewName(*s)
	}
	return uuo
}

// ClearNewName clears the value of new_name.
func (uuo *UserUpdateOne) ClearNewName() *UserUpdateOne {
	uuo.new_name = nil
	uuo.clearnew_name = true
	return uuo
}

// SetBlob sets the blob field.
func (uuo *UserUpdateOne) SetBlob(b []byte) *UserUpdateOne {
	uuo.blob = &b
	return uuo
}

// ClearBlob clears the value of blob.
func (uuo *UserUpdateOne) ClearBlob() *UserUpdateOne {
	uuo.blob = nil
	uuo.clearblob = true
	return uuo
}

// SetState sets the state field.
func (uuo *UserUpdateOne) SetState(u user.State) *UserUpdateOne {
	uuo.state = &u
	return uuo
}

// SetNillableState sets the state field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableState(u *user.State) *UserUpdateOne {
	if u != nil {
		uuo.SetState(*u)
	}
	return uuo
}

// ClearState clears the value of state.
func (uuo *UserUpdateOne) ClearState() *UserUpdateOne {
	uuo.state = nil
	uuo.clearstate = true
	return uuo
}

// AddCarIDs adds the car edge to Car by ids.
func (uuo *UserUpdateOne) AddCarIDs(ids ...int) *UserUpdateOne {
	if uuo.car == nil {
		uuo.car = make(map[int]struct{})
	}
	for i := range ids {
		uuo.car[ids[i]] = struct{}{}
	}
	return uuo
}

// AddCar adds the car edges to Car.
func (uuo *UserUpdateOne) AddCar(c ...*Car) *UserUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return uuo.AddCarIDs(ids...)
}

// SetPetsID sets the pets edge to Pet by id.
func (uuo *UserUpdateOne) SetPetsID(id int) *UserUpdateOne {
	if uuo.pets == nil {
		uuo.pets = make(map[int]struct{})
	}
	uuo.pets[id] = struct{}{}
	return uuo
}

// SetNillablePetsID sets the pets edge to Pet by id if the given value is not nil.
func (uuo *UserUpdateOne) SetNillablePetsID(id *int) *UserUpdateOne {
	if id != nil {
		uuo = uuo.SetPetsID(*id)
	}
	return uuo
}

// SetPets sets the pets edge to Pet.
func (uuo *UserUpdateOne) SetPets(p *Pet) *UserUpdateOne {
	return uuo.SetPetsID(p.ID)
}

// RemoveCarIDs removes the car edge to Car by ids.
func (uuo *UserUpdateOne) RemoveCarIDs(ids ...int) *UserUpdateOne {
	if uuo.removedCar == nil {
		uuo.removedCar = make(map[int]struct{})
	}
	for i := range ids {
		uuo.removedCar[ids[i]] = struct{}{}
	}
	return uuo
}

// RemoveCar removes car edges to Car.
func (uuo *UserUpdateOne) RemoveCar(c ...*Car) *UserUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return uuo.RemoveCarIDs(ids...)
}

// ClearPets clears the pets edge to Pet.
func (uuo *UserUpdateOne) ClearPets() *UserUpdateOne {
	uuo.clearedPets = true
	return uuo
}

// Save executes the query and returns the updated entity.
func (uuo *UserUpdateOne) Save(ctx context.Context) (*User, error) {
	if uuo.state != nil {
		if err := user.StateValidator(*uuo.state); err != nil {
			return nil, fmt.Errorf("entv2: validator failed for field \"state\": %v", err)
		}
	}
	if len(uuo.pets) > 1 {
		return nil, errors.New("entv2: multiple assignments on a unique edge \"pets\"")
	}
	return uuo.sqlSave(ctx)
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *UserUpdateOne) SaveX(ctx context.Context) *User {
	u, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return u
}

// Exec executes the query on the entity.
func (uuo *UserUpdateOne) Exec(ctx context.Context) error {
	_, err := uuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uuo *UserUpdateOne) ExecX(ctx context.Context) {
	if err := uuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (uuo *UserUpdateOne) sqlSave(ctx context.Context) (u *User, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   user.Table,
			Columns: user.Columns,
			ID: &sqlgraph.FieldSpec{
				Value:  uuo.id,
				Type:   field.TypeInt,
				Column: user.FieldID,
			},
		},
	}
	if value := uuo.age; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  *value,
			Column: user.FieldAge,
		})
	}
	if value := uuo.addage; value != nil {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  *value,
			Column: user.FieldAge,
		})
	}
	if value := uuo.name; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: user.FieldName,
		})
	}
	if value := uuo.nickname; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: user.FieldNickname,
		})
	}
	if value := uuo.phone; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: user.FieldPhone,
		})
	}
	if value := uuo.buffer; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBytes,
			Value:  *value,
			Column: user.FieldBuffer,
		})
	}
	if uuo.clearbuffer {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeBytes,
			Column: user.FieldBuffer,
		})
	}
	if value := uuo.title; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: user.FieldTitle,
		})
	}
	if value := uuo.new_name; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: user.FieldNewName,
		})
	}
	if uuo.clearnew_name {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: user.FieldNewName,
		})
	}
	if value := uuo.blob; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBytes,
			Value:  *value,
			Column: user.FieldBlob,
		})
	}
	if uuo.clearblob {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeBytes,
			Column: user.FieldBlob,
		})
	}
	if value := uuo.state; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  *value,
			Column: user.FieldState,
		})
	}
	if uuo.clearstate {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Column: user.FieldState,
		})
	}
	if nodes := uuo.removedCar; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.CarTable,
			Columns: []string{user.CarColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: car.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.car; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.CarTable,
			Columns: []string{user.CarColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: car.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.clearedPets {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   user.PetsTable,
			Columns: []string{user.PetsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: pet.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.pets; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   user.PetsTable,
			Columns: []string{user.PetsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: pet.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	u = &User{config: uuo.config}
	_spec.Assign = u.assignValues
	_spec.ScanValues = u.scanValues()
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return u, nil
}
