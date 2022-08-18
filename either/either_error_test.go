/*
 * Copyright (c) Dito Technologies LLC. All rights reserved.
 */

package either

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCollectOrError1(t *testing.T) {
	arr := []Either[int, error]{NewLeft[int, error](10), NewLeft[int, error](40)}
	data, err := CollectOrError(arr)
	assert.Nil(t, err)
	assert.Equal(t, len(data), 2)
}

func TestCollectOrError2(t *testing.T) {
	arr := []Either[int, error]{NewLeft[int, error](10), NewRight[int, error](fmt.Errorf("whatever"))}
	data, err := CollectOrError(arr)
	assert.Nil(t, data)
	assert.NotNil(t, err)
}
