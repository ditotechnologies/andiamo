/*
 * Copyright (c) Dito Technologies LLC. All rights reserved.
 */

// Package optional implements a data structure that either has a value or does not.
package optional

// Optional The interface that implements the optional data structure
type Optional[T any] interface {
	// HasValue Returns true if the optional has a value (AKA it was created with [NewWithValue]) or false otherwise.
	HasValue() bool
	// Value Returns the value wrapped by the optional. Or, panics if [HasValue] returns false.
	Value() T
}

type internalOptional[T any] struct {
	hasValue bool
	value    T
}

func (o *internalOptional[T]) HasValue() bool {
	return o.hasValue
}

func (o *internalOptional[T]) Value() T {
	if !o.hasValue {
		panic("cannot access the value of an optional without a value")
	}
	return o.value
}

// NewWithValue Creates a new optional where [HasValue()] will return true and [Value()] will return the value passed in.
func NewWithValue[T any](value T) Optional[T] {
	return &internalOptional[T]{
		hasValue: true,
		value:    value,
	}
}

// NewEmpty Creates a new optional where [HasValue()] will return false
func NewEmpty[T any]() Optional[T] {
	return &internalOptional[T]{
		hasValue: false,
	}
}
