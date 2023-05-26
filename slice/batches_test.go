/*
 * Copyright (c) Dito Technologies LLC. All rights reserved.
 */

package slice

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBatches1(t *testing.T) {
	input := []int{0, 1, 2, 3}
	batches := Batches(input, 2)
	assert.Equal(t, len(batches), 2)
	assert.Equal(t, len(batches[0]), 2)
	assert.Equal(t, len(batches[1]), 2)
	assert.Equal(t, batches[0][0], 0)
	assert.Equal(t, batches[0][1], 1)
	assert.Equal(t, batches[1][0], 2)
	assert.Equal(t, batches[1][1], 3)
}

func TestBatches2(t *testing.T) {
	input := []int{0, 1, 2, 3, 4}
	batches := Batches(input, 2)
	assert.Equal(t, len(batches), 3)
	assert.Equal(t, len(batches[0]), 2)
	assert.Equal(t, len(batches[1]), 2)
	assert.Equal(t, len(batches[2]), 1)
}
