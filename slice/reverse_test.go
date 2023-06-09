/*
 * Copyright (c) Dito Technologies LLC. All rights reserved.
 */

package slice

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverse1(t *testing.T) {
	var test []string
	Reverse(test)
	assert.Equal(t, len(test), 0)
}

func TestReverse2(t *testing.T) {
	test := []string{"a"}
	Reverse(test)
	assert.Equal(t, len(test), 1)
	assert.Equal(t, test[0], "a")
}

func TestReverse3(t *testing.T) {
	test := []string{"a", "b"}
	Reverse(test)
	assert.Equal(t, len(test), 2)
	assert.Equal(t, test[0], "b")
	assert.Equal(t, test[1], "a")
}

func TestReverse4(t *testing.T) {
	test := []string{"a", "b", "c"}
	Reverse(test)
	assert.Equal(t, len(test), 3)
	assert.Equal(t, test[0], "c")
	assert.Equal(t, test[1], "b")
	assert.Equal(t, test[2], "a")
}

func TestReverse5(t *testing.T) {
	test := []string{"a", "b", "c", "d"}
	Reverse(test)
	assert.Equal(t, len(test), 4)
	assert.Equal(t, test[0], "d")
	assert.Equal(t, test[1], "c")
	assert.Equal(t, test[2], "b")
	assert.Equal(t, test[3], "a")
}

func TestReverse6(t *testing.T) {
	test := make([]int, 0)
	for len(test) < 1000 {
		test = append(test, len(test))
	}
	Reverse(test)
	assert.Equal(t, test[0], 999)
	assert.Equal(t, test[1], 998)
	assert.Equal(t, test[len(test)-1], 0)
}
