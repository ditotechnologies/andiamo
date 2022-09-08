/*
 * Copyright (c) Dito Technologies LLC. All rights reserved.
 */

package slice

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDeduplicate1(t *testing.T) {
	data := []int{0, 1, 2, 1, 3}
	deduplicated := Deduplicate(data)
	assert.Equal(t, len(deduplicated), 4)
	assert.Equal(t, deduplicated, []int{0, 1, 2, 3})
}
