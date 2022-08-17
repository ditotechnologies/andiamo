/*
 * Copyright (c) Dito Technologies LLC. All rights reserved.
 */

package either

type Either[T1 any, T2 any] interface {
	IsLeft() bool
	Left() T1
	Right() T2
}

type internalEither[T1 any, T2 any] struct {
	isLeft bool
	left   *T1
	right  *T2
}

func (e internalEither[T1, T2]) IsLeft() bool {
	return e.isLeft
}

func (e internalEither[T1, T2]) Left() T1 {
	if !e.isLeft {
		panic("either is right, but tried to take left")
	}
	return *e.left
}

func (e internalEither[T1, T2]) Right() T2 {
	if e.isLeft {
		panic("either is left, but tried to take right")
	}
	return *e.right
}

func NewLeft[T1 any, T2 any](val T1) Either[T1, T2] {
	return internalEither[T1, T2]{
		isLeft: true,
		left:   &val,
	}
}
func NewRight[T1 any, T2 any](val T2) Either[T1, T2] {
	return internalEither[T1, T2]{
		isLeft: false,
		right:  &val,
	}
}
