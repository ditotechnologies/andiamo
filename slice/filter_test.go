/*
 * Copyright (c) Dito Technologies LLC. All rights reserved.
 */

package slice

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFilter1(t *testing.T) {
	slice := make([]int, 0)
	output := Filter(slice, func(e int) bool {
		return true
	})
	assert.Equal(t, output, slice)
}

func TestFilter2(t *testing.T) {
	slice := []int{0, 1, 2, 3}
	output := Filter(slice, func(e int) bool {
		return e != 1
	})
	assert.Equal(t, output, []int{0, 2, 3})
}
