/*
 * Copyright (c) Dito Technologies LLC. All rights reserved.
 */

package slice

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSubsliceAfter1(t *testing.T) {
	test := []string{"a", "b", "c", "d", "e"}
	output := SubsliceAfter(test, 2)
	assert.Equal(t, len(output), 3)
	assert.Equal(t, output[0], "c")
	assert.Equal(t, output[1], "d")
	assert.Equal(t, output[2], "e")
}
