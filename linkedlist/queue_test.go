/*
 * Copyright (c) Dito Technologies LLC. All rights reserved.
 */

package linkedlist

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewQueue(t *testing.T) {
	q := NewQueue[int]()
	assert.Equal(t, q.Len(), 0)
}
