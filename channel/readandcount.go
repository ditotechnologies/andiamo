/*
 * Copyright (c) Dito Technologies LLC. All rights reserved.
 */

package channel

func ReadAndCount[T any](ch <-chan T) int {
	count := 0
	for range ch {
		count += 1
	}
	return count
}
