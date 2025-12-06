package fetch

import "slices"

func insertSorted[S ~[]E, E any](s S, e E, cmp func(a, b E) int) S {
	l := len(s)
	first, last := 0, l
	for first < last {
		h := int(uint(first+last) >> 1)
		if cmp(s[h], e) > 0 {
			last = h
		} else {
			first = h + 1
		}
	}
	return slices.Insert(s, first, e)
}
