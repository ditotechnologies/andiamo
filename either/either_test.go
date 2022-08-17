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

func TestEitherLeft(t *testing.T) {
	e := NewLeft[string, int]("adfasdf")
	assert.True(t, e.IsLeft())
	assert.Equal(t, e.Left(), "adfasdf")
}
