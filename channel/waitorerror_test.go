/*
 * Copyright (c) Dito Technologies LLC. All rights reserved.
 */

package channel

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWaitOrError1(t *testing.T) {
	ctx := context.Background()
	c1 := make(chan error)
	go func() {
		defer close(c1)
		c1 <- fmt.Errorf("whatever")
	}()
	c2 := make(chan error)
	go func() {
		defer close(c2)
		c2 <- fmt.Errorf("nothing")
	}()
	arr := []<-chan error{c1, c2}
	err := WaitOrError(ctx, arr)
	assert.NotNil(t, err)
}

func TestWaitOrError2(t *testing.T) {
	ctx := context.Background()
	c1 := make(chan error)
	go func() {
		defer close(c1)
	}()
	c2 := make(chan error)
	go func() {
		defer close(c2)
	}()
	arr := []<-chan error{c1, c2}
	err := WaitOrError(ctx, arr)
	assert.Nil(t, err)
}
