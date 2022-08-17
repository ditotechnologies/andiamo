/*
 * Copyright (c) Dito Technologies LLC. All rights reserved.
 */

package slice

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMap1(t *testing.T) {
	input := []int{0, 1, 2}
	output := Map(input, func(e int) int {
		return e + 2
	})
	assert.Equal(t, len(input), len(output))
	assert.Equal(t, output[0], 2)
	assert.Equal(t, output[1], 3)
	assert.Equal(t, output[2], 4)
}
