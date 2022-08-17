/*
 * Copyright (c) Dito Technologies LLC. All rights reserved.
 */

package box

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBoxConstrain1(t *testing.T) {
	b1, _ := NewWithXYWidthHeight(10, 10, 10, 10)
	b2, _ := NewWithXYWidthHeight(0, 0, 100, 100)
	output, _ := b1.ConstrainedTo(b2)
	assert.Equal(t, output, b1)
}

func TestBoxConstrain2(t *testing.T) {
	b1, _ := NewWithXYWidthHeight(-10, 0, 20, 20)
	b2, _ := NewWithXYWidthHeight(0, 0, 100, 100)
	output, _ := b1.ConstrainedTo(b2)
	expectedOutput, _ := NewWithXYWidthHeight(0, 0, 20, 20)
	assert.Equal(t, output, expectedOutput)
}

func TestBoxConstrain3(t *testing.T) {
	b1, _ := NewWithXYWidthHeight(-10, 0, 120, 20)
	b2, _ := NewWithXYWidthHeight(0, 0, 100, 100)
	output, ok := b1.ConstrainedTo(b2)
	assert.False(t, ok)
	assert.Nil(t, output)
}

func TestBoxConstrain4(t *testing.T) {
	b1, _ := NewWithXYWidthHeight(10, 130, 20, 20)
	b2, _ := NewWithXYWidthHeight(0, 0, 100, 100)
	assert.False(t, b1 == b2)
	output, _ := b1.ConstrainedTo(b2)
	expectedOutput, _ := NewWithXYWidthHeight(10, 80, 20, 20)
	assert.Equal(t, output, expectedOutput)
}

func TestBoxConstrain5(t *testing.T) {
	b1, _ := NewWithXYWidthHeight(2432, -76, 756, 756)
	b2, _ := NewWithXYWidthHeight(0, 0, 3024, 4032)
	output, err := b1.ConstrainedTo(b2)
	assert.True(t, err)
	assert.NotNil(t, output)
}

func TestIntersects1(t *testing.T) {
	b1, _ := NewWithXYWidthHeight(0, 0, 20, 20)
	intersects := b1.Intersects(b1)
	assert.True(t, intersects)
}

func TestIntersects2(t *testing.T) {
	b1, _ := NewWithXYWidthHeight(0, 0, 20, 20)
	b2, _ := NewWithXYWidthHeight(30, 30, 5, 5)
	intersects := b1.Intersects(b2)
	assert.False(t, intersects)
}

func TestIntersects3(t *testing.T) {
	b1, _ := NewWithXYWidthHeight(0, 0, 20, 20)
	b2, _ := NewWithXYWidthHeight(10, 10, 20, 20)
	intersects := b1.Intersects(b2)
	assert.True(t, intersects)
}

func TestIntersects4(t *testing.T) {
	b1, _ := NewWithXYWidthHeight(0, 0, 20, 20)
	b2, _ := NewWithXYWidthHeight(10, 10, 5, 5)
	intersects := b1.Intersects(b2)
	assert.True(t, intersects)
}

func TestNewZeroBox(t *testing.T) {
	b := NewZero[int]()
	assert.Equal(t, b.Area(), 0)
}

func TestIntersectionBox1(t *testing.T) {
	b1, _ := NewWithXYWidthHeight(0, 0, 10, 10)
	b2, _ := NewWithXYWidthHeight(5, 5, 10, 10)
	expected, _ := NewWithXYWidthHeight(5, 5, 5, 5)
	output, ok := b1.IntersectionBox(b2)
	assert.True(t, ok)
	assert.Equal(t, expected, output)
}

func TestIntersectionBox2(t *testing.T) {
	b1, _ := NewWithXYWidthHeight(0, 0, 10, 10)
	b2, _ := NewWithXYWidthHeight(30, 30, 10, 10)
	_, ok := b1.IntersectionBox(b2)
	assert.False(t, ok)
}

func TestToZeroXY1(t *testing.T) {
	b1, _ := NewWithXYWidthHeight(10, 10, 10, 10)
	b2 := b1.ToZeroXY()
	assert.Equal(t, b2.MinX(), 0)
	assert.Equal(t, b2.MinY(), 0)
	assert.Equal(t, b2.Width(), 10)
	assert.Equal(t, b2.Height(), 10)
}

func TestPercentOverlapping1(t *testing.T) {
	b1, _ := NewWithXYWidthHeight(0, 0, 10, 10)
	b2, _ := NewWithXYWidthHeight(20, 20, 10, 10)
	overlapping := b1.PercentOverlapping(b2)
	assert.Equal(t, overlapping, 0.0)
}

func TestPercentOverlapping2(t *testing.T) {
	b1, _ := NewWithXYWidthHeight(0, 0, 10, 10)
	b2, _ := NewWithXYWidthHeight(5, 5, 10, 10)
	overlapping := b1.PercentOverlapping(b2)
	assert.Equal(t, overlapping, 0.25)
}

func TestMidX1(t *testing.T) {
	b1, _ := NewWithXYWidthHeight(0, 0, 10, 10)
	assert.Equal(t, b1.MidX(), 5)
	assert.Equal(t, b1.MidY(), 5)
}
