/*
 * Copyright (c) Dito Technologies LLC. All rights reserved.
 */

package linkedlist

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewDeque(t *testing.T) {
	d := NewDeque[int]()
	assert.Equal(t, d.Len(), 0)
}

func TestAddFirst1(t *testing.T) {
	d := NewDeque[int]()
	d.AddFirst(1)
	assert.Equal(t, d.Len(), 1)
	d.AddFirst(2)
	assert.Equal(t, d.Len(), 2)
	assert.Equal(t, d.PeekFirst(), 2)
	assert.Equal(t, d.PeekLast(), 1)
}

func TestRemoveFirstNoItems(t *testing.T) {
	d := NewDeque[int]()
	f := func() {
		d.RemoveFirst()
	}
	assert.Panics(t, f)
}

func TestRemoveLastNoItems(t *testing.T) {
	d := NewDeque[int]()
	f := func() {
		d.RemoveLast()
	}
	assert.Panics(t, f)
}

func TestRemoveFirst1(t *testing.T) {
	d := NewDeque[int]()
	d.AddFirst(1)
	output := d.RemoveFirst()
	assert.Equal(t, output, 1)
	assert.Equal(t, d.Len(), 0)
}

func TestRemoveFirst2(t *testing.T) {
	d := NewDeque[int]()
	d.AddFirst(1)
	d.AddLast(2)
	output := d.RemoveFirst()
	assert.Equal(t, output, 1)
}
