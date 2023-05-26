/*
 * Copyright (c) Dito Technologies LLC. All rights reserved.
 */

package slice

func Batches[Elem any](input []Elem, batchSize int) [][]Elem {
	if batchSize <= 0 {
		panic("must have a batch size greater than 0")
	}
	output := make([][]Elem, 0)
	currentBatch := make([]Elem, 0)
	for _, elem := range input {
		if len(currentBatch) >= batchSize {
			output = append(output, currentBatch)
			currentBatch = make([]Elem, 0)
		}
		currentBatch = append(currentBatch, elem)
	}
	if len(currentBatch) > 0 {
		output = append(output, currentBatch)
	}
	return output
}
