/*
 * Copyright (c) Dito Technologies LLC. All rights reserved.
 */

package box

import "math"

type Dimension interface {
	float64 | float32 | int
}

type Box[T Dimension] struct {
	x1 T
	y1 T
	x2 T
	y2 T
}

func (b *Box[T]) Area() T {
	return b.Width() * b.Height()
}

func (b *Box[T]) Width() T {
	return b.x2 - b.x1
}

func (b *Box[T]) Height() T {
	return b.y2 - b.y1
}

func (b *Box[T]) MinX() T {
	return b.x1
}

func (b *Box[T]) MidX() T {
	return (b.x1 + b.x2) / 2
}

func (b *Box[T]) MidY() T {
	return (b.y1 + b.y2) / 2
}

func (b *Box[T]) MinY() T {
	return b.y1
}

func (b *Box[T]) MaxX() T {
	return b.x2
}

func (b *Box[T]) MaxY() T {
	return b.y2
}

func (b *Box[T]) MoveInX(dx T) {
	b.x1 += dx
	b.x2 += dx
}

func (b *Box[T]) MoveInY(dy T) {
	b.y1 += dy
	b.y2 += dy
}

func (b Box[T]) ConstrainedTo(other Box[T]) (*Box[T], bool) {
	output := &Box[T]{
		x1: b.x1,
		x2: b.x2,
		y1: b.y1,
		y2: b.y2,
	}

	// first do the x movement
	if output.MinX() < other.MinX() {
		output.MoveInX(other.MinX() - b.MinX())
		// if we are past the max x, this is an error, cannot be constrained
		if output.MaxX() > other.MaxX() {
			return nil, false
		}
	}
	if output.MaxX() > other.MaxX() {
		output.MoveInX(-(output.MaxX() - other.MaxX()))
		// if we past the start, this is an error
		if output.MinX() < other.MinX() {
			return nil, false
		}
	}

	// now, do the y movement
	if output.MinY() < other.MinY() {
		output.MoveInY(other.MinY() - b.MinY())
		// if we are past the max y, this is an error, cannot be constrained
		if output.MaxY() > other.MaxY() {
			return nil, false
		}
	}
	if output.MaxY() > other.MaxY() {
		output.MoveInY(-(b.MaxY() - other.MaxY()))
		// if we past the start, this is an error
		if output.MinY() < other.MinY() {
			return nil, false
		}
	}

	return output, true
}

func (b Box[T]) Encompasses(other Box[T]) bool {
	return b.MinX() <= other.MinX() && b.MinY() <= other.MinY() && b.MaxX() >= other.MaxX() && b.MaxY() >= other.MaxX()
}

func (b Box[T]) Intersects(other Box[T]) bool {
	if b.Encompasses(other) || other.Encompasses(b) {
		return true
	}
	xIntersects := (b.MinX() <= other.MaxX() && other.MaxX() <= b.MaxX()) || (b.MinX() <= other.MinX() && other.MinX() <= b.MaxX()) || (other.MinX() < b.MinX() && other.MaxX() > b.MaxX())
	yIntersects := (b.MinY() <= other.MaxY() && other.MaxY() <= b.MaxY()) || (b.MinY() <= other.MinY() && other.MinY() <= b.MaxY()) || (other.MinY() < b.MinY() && other.MaxY() > b.MaxY())
	return xIntersects && yIntersects
}

func (b *Box[T]) ToZeroXY() Box[T] {
	return Box[T]{
		x1: 0,
		x2: b.x2 - b.x1,
		y1: 0,
		y2: b.y2 - b.y1,
	}
}

func (b Box[T]) IntersectionBox(other Box[T]) (*Box[T], bool) {
	if b == other || other.Encompasses(b) {
		output := b.ToZeroXY()
		return &output, true
	}

	if !b.Intersects(other) {
		return nil, false
	}

	x1 := other.MinX() - b.MinX()
	y1 := other.MinY() - b.MinY()

	x2 := x1 + other.Width()
	y2 := y1 + other.Height()

	if x1 < 0 {
		x1 = 0
	}
	if y1 < 0 {
		y1 = 0
	}

	if x2 > b.Width()-1 {
		x2 = b.Width()
	}
	if y2 > b.Height()-1 {
		y2 = b.Height()
	}

	return NewWithX1Y1X2Y2(x1, y1, x2, y2)

}

func (b Box[T]) PercentOverlapping(other Box[T]) float64 {
	if !b.Intersects(other) {
		return 0
	}

	otherX1 := float64(other.MinX())
	otherX2 := float64(other.MaxX())
	selfX1 := float64(b.MinX())
	selfX2 := float64(b.MaxX())

	dx := math.Min(
		math.Abs(selfX2-otherX1), math.Abs(selfX1-otherX2),
	)

	otherY1 := float64(other.MinY())
	otherY2 := float64(other.MaxY())
	selfY1 := float64(b.MinY())
	selfY2 := float64(b.MaxY())

	dy := math.Min(
		math.Abs(selfY2-otherY1), math.Abs(selfY1-otherY2),
	)

	area := float64(b.Area())

	return (dx * dy) / area

}

func NewZero[T Dimension]() Box[T] {
	return Box[T]{
		x1: 0,
		y1: 0,
		x2: 0,
		y2: 0,
	}
}

func NewWithX1Y1X2Y2[T Dimension](x1 T, y1 T, x2 T, y2 T) (*Box[T], bool) {
	if x1 > x2 || y1 > y2 {
		return nil, false
	}
	return &Box[T]{
		x1: x1,
		y1: y1,
		x2: x2,
		y2: y2,
	}, true
}

func NewWithXYWidthHeight[T Dimension](x1 T, y1 T, width T, height T) Box[T] {
	return Box[T]{
		x1: x1,
		y1: y1,
		x2: x1 + width,
		y2: y1 + height,
	}
}
