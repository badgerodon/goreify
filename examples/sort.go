package examples

import (
	"sort"

	"github.com/badgerodon/goreify/generics"
)

// Sort sorts a list
func Sort(xs []generics.T1) {
	sort.Sort(generics.NewSorter(
		func() int {
			return len(xs)
		},
		func(i, j int) {
			xs[i], xs[j] = xs[j], xs[i]
		},
		func(i, j int) bool {
			return generics.Less(xs[i], xs[j])
		},
	))
}
