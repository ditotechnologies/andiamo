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
	Union(Set[T]) Set[T]
	Intersection(Set[T]) Set[T]
	Difference(Set[T]) Set[T]
	ForEach(func(T))
}

type internalSet[T comparable] struct {
	underlyingData     map[T]bool
	underlyingDataLock sync.RWMutex
}

func (s *internalSet[T]) Items() []T {
	s.underlyingDataLock.RLock()
	defer s.underlyingDataLock.RUnlock()
	output := make([]T, 0)
	for k := range s.underlyingData {
		output = append(output, k)
	}
	return output
}

func (s *internalSet[T]) ItemsCh(ctx context.Context) <-chan T {

	output := make(chan T)

	go func() {

		defer close(output)

		s.underlyingDataLock.RLock()
		defer s.underlyingDataLock.RUnlock()
		for k := range s.underlyingData {
			select {
			case output <- k:
			case <-ctx.Done():
			}
		}
	}()

	return output

}

func (s *internalSet[T]) Contains(elem T) bool {
	s.underlyingDataLock.RLock()
	defer s.underlyingDataLock.RUnlock()
	_, exists := s.underlyingData[elem]
	return exists
}

func (s *internalSet[T]) Len() int {
	s.underlyingDataLock.RLock()
	defer s.underlyingDataLock.RUnlock()
	return len(s.underlyingData)
}

func (s *internalSet[T]) Add(elem T) {
	s.underlyingDataLock.Lock()
	defer s.underlyingDataLock.Unlock()
	s.underlyingData[elem] = true
}

func (s *internalSet[T]) Remove(elem T) {
	s.underlyingDataLock.Lock()
	defer s.underlyingDataLock.Unlock()
	delete(s.underlyingData, elem)
}

func (s *internalSet[T]) Union(other Set[T]) Set[T] {
	output := New[T]()
	s.ForEach(func(elem T) {
		output.Add(elem)
	})
	other.ForEach(func(elem T) {
		output.Add(elem)
	})
	return output
}

func (s *internalSet[T]) Difference(other Set[T]) Set[T] {
	output := New[T]()
	s.ForEach(func(elem T) {
		if !other.Contains(elem) {
			output.Add(elem)
		}
	})
	return output
}

func (s *internalSet[T]) Intersection(other Set[T]) Set[T] {
	// have the sorter set be set 1 to make this faster
	var set1 Set[T] = s
	var set2 = other
	if s.Len() > other.Len() {
		set1 = other
		set2 = s
	}

	// do the intersection
	output := New[T]()
	set1.ForEach(func(elem T) {
		if set2.Contains(elem) {
			output.Add(elem)
		}
	})
	return output
}

func (s *internalSet[T]) ForEach(fn func(elem T)) {
	var wg sync.WaitGroup
	s.underlyingDataLock.RLock()
	defer s.underlyingDataLock.RUnlock()
	for _elem := range s.underlyingData {
		wg.Add(1)
		go func(elem T) {
			defer wg.Done()
			fn(elem)
		}(_elem)
	}
	wg.Wait()
}

func NewWithSlice[T comparable](items []T) Set[T] {
	newSet := New[T]()
	for _, item := range items {
		newSet.Add(item)
	}
	return newSet
}

func New[T comparable]() Set[T] {
	s := internalSet[T]{
		underlyingData: make(map[T]bool),
	}
	return &s
}
