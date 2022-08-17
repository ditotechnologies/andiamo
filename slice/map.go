package slice

func mapElemToCh[Elem any, Output any](elem Elem, idx int, mapFun func(Elem, int) Output) <-chan Output {
	ch := make(chan Output)
	go func() {
		defer close(ch)
		ch <- mapFun(elem, idx)
	}()
	return ch
}

// MapWithIndex Same as map, but gives you the index of the element in the mapping function
func MapWithIndex[Elem any, Output any](slice []Elem, mapFun func(Elem, int) Output) []Output {
	chs := make([]<-chan Output, len(slice))
	for idx, elem := range slice {
		chs[idx] = mapElemToCh(elem, idx, mapFun)
	}
	output := make([]Output, len(chs))
	for idx, ch := range chs {
		output[idx] = <-ch
	}
	return output
}

// Map Performs a function on each element of the slice
func Map[Elem any, Output any](slice []Elem, mapFun func(Elem) Output) []Output {
	upgradedFn := func(elem Elem, idx int) Output {
		return mapFun(elem)
	}
	return MapWithIndex(slice, upgradedFn)
}
