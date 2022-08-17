/*
 * Copyright (c) Dito Technologies LLC. All rights reserved.
 */

package slice

import (
	"github.com/stretchr/testify/assert"
	"sync/atomic"
	"testing"
)

func TestForEach(t *testing.T) {
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	var totalCount uint32 = 0
	ForEach(slice, func(_ int) {
		atomic.AddUint32(&totalCount, 1)
	})
	assert.Equal(t, uint32(len(slice)), totalCount)
}
