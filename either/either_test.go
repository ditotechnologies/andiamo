/*
 * Copyright (c) Dito Technologies LLC. All rights reserved.
 */

package either

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEitherRight(t *testing.T) {
	e := NewRight[string, int](10)
	assert.False(t, e.IsLeft())
	assert.Equal(t, e.Right(), 10)
}

func TestEitherRightWhenActuallyLeft(t *testing.T) {
	e := NewLeft[string, int]("fsdfs")
	fn := func() {
		_ = e.Right()
	}
	assert.Panics(t, fn)
}

func TestEitherLeft(t *testing.T) {
	e := NewLeft[string, int]("adfasdf")
	assert.True(t, e.IsLeft())
	assert.Equal(t, e.Left(), "adfasdf")
}

func TestEitherLeftWhenActuallyRight(t *testing.T) {
	e := NewRight[string, int](2394)
	fn := func() {
		_ = e.Left()
	}
	assert.Panics(t, fn)
}
