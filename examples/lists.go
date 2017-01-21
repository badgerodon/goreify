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

// A List is a list of elements
type List struct {
	elements []generics.T1
}

// NewList creates a new List
func NewList() *List {
	return &List{}
}

// Append appends elements to the list
func (l *List) Append(els ...generics.T1) {
	l.elements = append(l.elements, els...)
}

// At returns the element at i
func (l *List) At(i int) generics.T1 {
	return l.elements[i]
}

// Len returns the length of the list
func (l List) Len() int {
	return len(l.elements)
}

// Slice slices the list
func (l *List) Slice(from, to int) *List {
	return &List{
		elements: l.elements[from:to],
	}
}
