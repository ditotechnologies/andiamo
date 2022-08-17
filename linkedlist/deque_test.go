/*
 * Copyright (c) Dito Technologies LLC. All rights reserved.
 */

package linkedlist

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewDeque(t *testing.T) {
	d := NewDeque[int]()
	assert.Equal(t, d.Len(), 0)
}
