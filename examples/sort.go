package examples

import (
	"sort"

	"github.com/badgerodon/goreify/generics"
)

type sorter struct {
	len  func() int
	swap func(i, j int)
	less func(i, j int) bool
}

func (s sorter) Len() int {
	return s.len()
}
func (s sorter) Swap(i, j int) {
	s.swap(i, j)
}
func (s sorter) Less(i, j int) bool {
	return s.less(i, j)
}

func Sort(xs []generics.T1) {
	sort.Sort(sorter{
		len: func() int {
			return len(xs)
		},
		swap: func(i, j int) {
			xs[i], xs[j] = xs[j], xs[i]
		},
		less: func(i, j int) bool {
			return generics.Less(xs[i], xs[j])
		},
	})
}
