/*
 * Copyright (c) Dito Technologies LLC. All rights reserved.
 */

package slice

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIndexOf1(t *testing.T) {
	slice := []int{0, 1, 2, 3, 4, 5}
	output, ok := IndexOf(slice, 3)
	assert.True(t, ok)
	assert.Equal(t, output, 3)
}
