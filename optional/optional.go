/*
 * Copyright (c) Dito Technologies LLC. All rights reserved.
 */

package optional

type Optional[T any] interface {
	HasValue() bool
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

func NewWithValue[T any](value T) Optional[T] {
	return &internalOptional[T]{
		hasValue: true,
		value:    value,
	}
}

func NewEmpty[T any]() Optional[T] {
	return &internalOptional[T]{
		hasValue: false,
	}
}
