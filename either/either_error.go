/*
 * Copyright (c) Dito Technologies LLC. All rights reserved.
 */

package either

func FunctionResultOrError[T any](fn func() (T, error)) Either[T, error] {
	t, err := fn()
	if err != nil {
		return NewRight[T, error](err)
	}
	return NewLeft[T, error](t)
}

func CollectOrError[T any](eithers []Either[T, error]) ([]T, error) {
	output := make([]T, len(eithers))
	for idx, either := range eithers {
		if either.IsLeft() {
			output[idx] = either.Left()
		} else {
			return nil, either.Right()
		}
	}
	return output, nil
}
