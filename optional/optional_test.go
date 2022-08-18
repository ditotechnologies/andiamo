/*
 * Copyright (c) Dito Technologies LLC. All rights reserved.
 */

package optional

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOptional1(t *testing.T) {
	o := NewEmpty[int]()
	assert.False(t, o.HasValue())
	fn := func() {
		_ = o.Value()
	}
	assert.Panics(t, fn)
}

func TestOptional2(t *testing.T) {
	o := NewWithValue[int](10)
	assert.True(t, o.HasValue())
	assert.Equal(t, o.Value(), 10)
}
