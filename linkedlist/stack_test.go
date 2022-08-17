/*
 * Copyright (c) Dito Technologies LLC. All rights reserved.
 */

package linkedlist

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewStack(t *testing.T) {
	s := NewStack[int]()
	assert.Equal(t, s.Len(), 0)
}
