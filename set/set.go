/*
 * Copyright (c) Dito Technologies LLC. All rights reserved.
 */

package set

import (
	"context"
	"sync"
)

type Set[T comparable] interface {
	Items() []T
	ItemsCh(ctx context.Context) <-chan T
	Contains(T) bool
	Add(T)
	Remove(T)
	Len() int
}

type internalSet[T comparable] struct {
	underlyingData     map[T]bool
	underlyingDataLock sync.RWMutex
}

func (set *internalSet[T]) Items() []T {
	set.underlyingDataLock.RLock()
	defer set.underlyingDataLock.RUnlock()
	output := make([]T, 0)
	for k := range set.underlyingData {
		output = append(output, k)
	}
	return output
}

func (set *internalSet[T]) ItemsCh(ctx context.Context) <-chan T {

	output := make(chan T)

	go func() {

		defer close(output)

		set.underlyingDataLock.RLock()
		defer set.underlyingDataLock.RUnlock()
		for k := range set.underlyingData {
			select {
			case output <- k:
			case <-ctx.Done():
			}
		}
	}()

	return output

}

func (set *internalSet[T]) Contains(elem T) bool {
	set.underlyingDataLock.RLock()
	defer set.underlyingDataLock.RUnlock()
	_, exists := set.underlyingData[elem]
	return exists
}

func (set *internalSet[T]) Len() int {
	set.underlyingDataLock.RLock()
	defer set.underlyingDataLock.RUnlock()
	return len(set.underlyingData)
}

func (set *internalSet[T]) Add(elem T) {
	set.underlyingDataLock.Lock()
	defer set.underlyingDataLock.Unlock()
	set.underlyingData[elem] = true
}

func (set *internalSet[T]) Remove(elem T) {
	set.underlyingDataLock.Lock()
	defer set.underlyingDataLock.Unlock()
	delete(set.underlyingData, elem)
}

func New[T comparable]() Set[T] {
	s := internalSet[T]{
		underlyingData: make(map[T]bool),
	}
	return &s
}
