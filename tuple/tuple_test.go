/*
 * Copyright (c) Dito Technologies LLC. All rights reserved.
 */

package tuple

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTuple1(t *testing.T) {
	tuple := NewTuple[int, string](10, "ten")
	assert.Equal(t, tuple.First(), 10)
	assert.Equal(t, tuple.Second(), "ten")
}
