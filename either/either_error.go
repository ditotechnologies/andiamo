/*
 * Copyright (c) Dito Technologies LLC. All rights reserved.
 */

package either

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
