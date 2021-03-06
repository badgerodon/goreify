package examples

import (
	sort "sort"

	generics "github.com/badgerodon/goreify/generics"
)

// SortInt8 sorts a list
func SortInt8(xs []int8) {
	sort.Sort(generics.NewSorter(
		func() int {
			return len(xs)
		},
		func(i, j int) {
			xs[i], xs[j] = xs[j], xs[i]
		},
		func(i, j int) bool {
			return xs[i] < xs[j]
		},
	))
}
