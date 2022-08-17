/*
 * Copyright (c) Dito Technologies LLC. All rights reserved.
 */

package channel

func ToSlice[T any](ch <-chan T) []T {
	output := make([]T, 0)
	for elem := range ch {
		output = append(output, elem)
	}
	return output
}
