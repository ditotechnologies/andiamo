/*
 * Copyright (c) Dito Technologies LLC. All rights reserved.
 */

package slice

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestContains1(t *testing.T) {
	data := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	contains1 := Contains(data, 3)
	assert.True(t, contains1)
	contains2 := Contains(data, -1)
	assert.False(t, contains2)
}
