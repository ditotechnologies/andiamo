/*
 * Copyright (c) Dito Technologies LLC. All rights reserved.
 */

package slice

import "github.com/ditotechnologies/andiamo/optional"

func FirstOr[Elem any](slice []Elem, fallback Elem) Elem {
	if len(slice) > 0 {
		return slice[0]
	}
	return fallback
}

func First[Elem any](slice []Elem) optional.Optional[Elem] {
	if len(slice) > 0 {
		return optional.NewWithValue(slice[0])
	}
	return optional.NewEmpty[Elem]()
}
