/*
 * Copyright (c) Dito Technologies LLC. All rights reserved.
 */

package set

import (
	"context"
	"github.com/ditotechnologies/andiamo/channel"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSetInit1(t *testing.T) {
	s := New[int]()
	assert.Equal(t, s.Len(), 0)
	assert.Equal(t, s.Items(), []int{})
}

func TestSetInit2(t *testing.T) {
	s := New[string]()
	assert.Equal(t, s.Len(), 0)
	assert.Equal(t, s.Items(), []string{})
}

func TestAddAndRemove1(t *testing.T) {
	s := New[int]()
	assert.Equal(t, s.Len(), 0)
	s.Add(1)
	assert.Equal(t, s.Len(), 1)
	s.Add(2)
	assert.Equal(t, s.Len(), 2)
	s.Add(1)
	assert.Equal(t, s.Len(), 2)
	assert.True(t, s.Contains(1))
	assert.True(t, s.Contains(2))
}

func TestAddAndRemove2(t *testing.T) {
	s := New[int]()
	s.Add(1)
	assert.Equal(t, s.Len(), 1)
	assert.True(t, s.Contains(1))
	s.Remove(1)
	assert.Equal(t, s.Len(), 0)
	assert.False(t, s.Contains(1))
}

func TestItemsCh(t *testing.T) {
	s := New[int]()
	ch1 := s.ItemsCh(context.Background())
	assert.Equal(t, channel.ReadAndCount(ch1), 0)
	s.Add(1)
	ch2 := s.ItemsCh(context.Background())
	slice := channel.ToSlice(ch2)
	assert.Equal(t, slice, []int{1})
	items := s.Items()
	assert.Equal(t, items, []int{1})
}
