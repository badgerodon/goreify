package generics

import "sort"

//go:generate bash -c "go run ../generator/builtins/*.go > builtins.gen.go"

type (
	// Any is any type
	Any interface{}

	// T1 is a generic type
	T1 interface{}
)

type sorter struct {
	len  func() int
	swap func(i, j int)
	less func(i, j int) bool
}

// NewSorter creates a new sorter from functions
func NewSorter(
	len func() int,
	swap func(i, j int),
	less func(i, j int) bool,
) sort.Interface {
	return &sorter{
		len:  len,
		swap: swap,
		less: less,
	}
}

func (s *sorter) Len() int {
	return s.len()
}
func (s *sorter) Swap(i, j int) {
	s.swap(i, j)
}
func (s *sorter) Less(i, j int) bool {
	return s.less(i, j)
}
