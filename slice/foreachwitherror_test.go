/*
 * Copyright (c) Dito Technologies LLC. All rights reserved.
 */

package slice

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"sync/atomic"
	"testing"
)

func TestForEachWithError1(t *testing.T) {
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	var totalCount uint32 = 0
	err := ForEachWithError(slice, func(i int) error {
		atomic.AddUint32(&totalCount, 1)
		return nil
	})
	assert.Nil(t, err)
	assert.Equal(t, uint32(len(slice)), totalCount)
}

func TestForEachWithError2(t *testing.T) {
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	err := ForEachWithError(slice, func(i int) error {
		if i == 3 {
			return errors.New("test error")
		}
		return nil
	})
	assert.NotNil(t, err)
}
