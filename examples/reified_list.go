package examples

type List_int32 struct {
	elements []int32
}

type List_float32 struct {
	elements []float32
}

// NewList_int32 creates a new List
func NewList_int32() *List_int32 {
	return &List_int32{}
}

// NewList_float32 creates a new List
func NewList_float32() *List_float32 {
	return &List_float32{}
}

// Append appends elements to the list
func (l *List_int32) Append(els ...int32) {
	l.elements = append(l.elements, els...)
}

// Append appends elements to the list
func (l *List_float32) Append(els ...float32) {
	l.elements = append(l.elements, els...)
}

// Len returns the length of the list
func (l List_int32) Len() int {
	return len(l.elements)
}

// Len returns the length of the list
func (l List_float32) Len() int {
	return len(l.elements)
}

// Slice slices the list
func (l *List_int32) Slice(from, to int) *List_int32 {
	return &List_int32{
		elements: l.elements[from:to],
	}
}

// Slice slices the list
func (l *List_float32) Slice(from, to int) *List_float32 {
	return &List_float32{
		elements: l.elements[from:to],
	}
}
