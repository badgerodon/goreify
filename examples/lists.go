package examples

import "github.com/badgerodon/goreify/generics"

// Contains returns whether or not the list contains the element
func Contains(haystack []generics.T1, needle generics.T1) bool {
	for _, el := range haystack {
		if generics.Equal(el, needle) {
			return true
		}
	}
	return false
}

type List struct {
	elements []generics.T1
}

func NewList() *List {
	return &List{}
}

func (l *List) Append(els ...generics.T1) {
	l.elements = append(l.elements, els...)
}

func (l *List) Slice(from, to int) *List {
	n := NewList()
	n.Append(l.elements[from:to]...)
	return n
}
