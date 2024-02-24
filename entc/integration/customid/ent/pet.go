// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/entc/integration/customid/ent/pet"
	"entgo.io/ent/entc/integration/customid/ent/user"
)

// Pet is the model entity for the Pet schema.
type Pet struct {
	config
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PetQuery when eager-loading is set.
	Edges           PetEdges `json:"edges"`
	pet_best_friend *string
	user_pets       *int
	selectValues    sql.SelectValues
}

// PetEdges holds the relations/edges for other nodes in the graph.
type PetEdges struct {
	// Owner holds the value of the owner edge.
	Owner *User `json:"owner,omitempty"`
	// Cars holds the value of the cars edge.
	Cars []*Car `json:"cars,omitempty"`
	// Friends holds the value of the friends edge.
	Friends []*Pet `json:"friends,omitempty"`
	// BestFriend holds the value of the best_friend edge.
	BestFriend *Pet `json:"best_friend,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// OwnerOrErr returns the Owner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PetEdges) OwnerOrErr() (*User, error) {
	if e.Owner != nil {
		return e.Owner, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: user.Label}
	}
	return nil, &NotLoadedError{edge: "owner"}
}

// CarsOrErr returns the Cars value or an error if the edge
// was not loaded in eager-loading.
func (e PetEdges) CarsOrErr() ([]*Car, error) {
	if e.loadedTypes[1] {
		return e.Cars, nil
	}
	return nil, &NotLoadedError{edge: "cars"}
}

// FriendsOrErr returns the Friends value or an error if the edge
// was not loaded in eager-loading.
func (e PetEdges) FriendsOrErr() ([]*Pet, error) {
	if e.loadedTypes[2] {
		return e.Friends, nil
	}
	return nil, &NotLoadedError{edge: "friends"}
}

// BestFriendOrErr returns the BestFriend value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PetEdges) BestFriendOrErr() (*Pet, error) {
	if e.BestFriend != nil {
		return e.BestFriend, nil
	} else if e.loadedTypes[3] {
		return nil, &NotFoundError{label: pet.Label}
	}
	return nil, &NotLoadedError{edge: "best_friend"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Pet) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case pet.FieldID:
			values[i] = new(sql.NullString)
		case pet.ForeignKeys[0]: // pet_best_friend
			values[i] = new(sql.NullString)
		case pet.ForeignKeys[1]: // user_pets
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Pet fields.
func (pe *Pet) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case pet.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				pe.ID = value.String
			}
		case pet.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field pet_best_friend", values[i])
			} else if value.Valid {
				pe.pet_best_friend = new(string)
				*pe.pet_best_friend = value.String
			}
		case pet.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field user_pets", value)
			} else if value.Valid {
				pe.user_pets = new(int)
				*pe.user_pets = int(value.Int64)
			}
		default:
			pe.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Pet.
// This includes values selected through modifiers, order, etc.
func (pe *Pet) Value(name string) (ent.Value, error) {
	return pe.selectValues.Get(name)
}

// QueryOwner queries the "owner" edge of the Pet entity.
func (pe *Pet) QueryOwner() *UserQuery {
	return NewPetClient(pe.config).QueryOwner(pe)
}

// QueryCars queries the "cars" edge of the Pet entity.
func (pe *Pet) QueryCars() *CarQuery {
	return NewPetClient(pe.config).QueryCars(pe)
}

// QueryFriends queries the "friends" edge of the Pet entity.
func (pe *Pet) QueryFriends() *PetQuery {
	return NewPetClient(pe.config).QueryFriends(pe)
}

// QueryBestFriend queries the "best_friend" edge of the Pet entity.
func (pe *Pet) QueryBestFriend() *PetQuery {
	return NewPetClient(pe.config).QueryBestFriend(pe)
}

// Update returns a builder for updating this Pet.
// Note that you need to call Pet.Unwrap() before calling this method if this Pet
// was returned from a transaction, and the transaction was committed or rolled back.
func (pe *Pet) Update() *PetUpdateOne {
	return NewPetClient(pe.config).UpdateOne(pe)
}

// Unwrap unwraps the Pet entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pe *Pet) Unwrap() *Pet {
	_tx, ok := pe.config.driver.(*txDriver)
	if !ok {
		panic("ent: Pet is not a transactional entity")
	}
	pe.config.driver = _tx.drv
	return pe
}

// String implements the fmt.Stringer.
func (pe *Pet) String() string {
	var builder strings.Builder
	builder.WriteString("Pet(")
	builder.WriteString(fmt.Sprintf("id=%v", pe.ID))
	builder.WriteByte(')')
	return builder.String()
}

// Pets is a parsable slice of Pet.
type Pets []*Pet
