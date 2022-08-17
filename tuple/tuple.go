/*
 * Copyright (c) Dito Technologies LLC. All rights reserved.
 */

package tuple

type Tuple[T1 any, T2 any] interface {
	First() T1
	Second() T2
}

type internalTuple[T1 any, T2 any] struct {
	first  T1
	second T2
}

func (t *internalTuple[T1, T2]) First() T1 {
	return t.first
}

func (t *internalTuple[T1, T2]) Second() T2 {
	return t.second
}

func NewTuple[T1 any, T2 any](first T1, second T2) Tuple[T1, T2] {
	return &internalTuple[T1, T2]{first: first, second: second}
}
