/*
 * Copyright (c) Dito Technologies LLC. All rights reserved.
 */

package channel

import (
	"context"
	"fmt"
	"github.com/ditotechnologies/andiamo/slice"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParallelizeFunctionsToResultAndError1(t *testing.T) {
	f1 := func() (int, error) {
		return 10, nil
	}
	f2 := func() (int, error) {
		return 20, nil
	}
	arr := make([]func() (int, error), 0)
	arr = append(arr, f1)
	arr = append(arr, f2)
	result, err := ParallelizeFunctionsToResultAndError(context.Background(), arr)
	assert.Nil(t, err)
	result = slice.MergeSort(result)
	assert.Equal(t, len(result), 2)
	assert.Equal(t, result[0], 10)
	assert.Equal(t, result[1], 20)
}

func TestParallelizeFunctionsToError1(t *testing.T) {
	f1 := func() error {
		return nil
	}
	f2 := func() error {
		return fmt.Errorf("errored")
	}
	arr := make([]func() error, 0)
	arr = append(arr, f1)
	arr = append(arr, f2)
	err := ParallelizeFunctionsToError(context.Background(), arr)
	assert.NotNil(t, err)
}
