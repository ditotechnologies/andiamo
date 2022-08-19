/*
 * Copyright (c) Dito Technologies LLC. All rights reserved.
 */

package slice

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFirstOr1(t *testing.T) {
	arr := []int{1, 2}
	output := FirstOr(arr, 3)
	assert.Equal(t, output, 1)
}

func TestFirstOr2(t *testing.T) {
	var arr []int
	output := FirstOr(arr, 3)
	assert.Equal(t, output, 3)
}

func TestFirst1(t *testing.T) {
	arr := []int{1, 2}
	output := First(arr)
	assert.Equal(t, output.Value(), 1)
	assert.True(t, output.HasValue())
}

func TestFirst2(t *testing.T) {
	var arr []int
	output := First(arr)
	assert.False(t, output.HasValue())
}
