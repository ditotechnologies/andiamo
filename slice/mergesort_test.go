/*
 * Copyright (c) Dito Technologies LLC. All rights reserved.
 */

package slice

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMergeSort1(t *testing.T) {
	testArr := []int{0, 3, 2, 1, 4}
	outputSmallestToBiggest := MergeSort(testArr, func(c1 int, c2 int) bool {
		return c1 < c2
	})
	assert.Equal(t, len(outputSmallestToBiggest), len(testArr))
	assert.Equal(t, outputSmallestToBiggest[0], 0)
	assert.Equal(t, outputSmallestToBiggest[1], 1)
	assert.Equal(t, outputSmallestToBiggest[2], 2)
	assert.Equal(t, outputSmallestToBiggest[3], 3)
	assert.Equal(t, outputSmallestToBiggest[4], 4)
}

func TestMergeSort2(t *testing.T) {
	testArr := []int{}
	output := MergeSort(testArr, func(c1 int, c2 int) bool {
		return c1 < c2
	})
	assert.Equal(t, len(output), len(testArr))
}
