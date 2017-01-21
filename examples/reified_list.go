package examples

type

// A ListInt32 is a list of elements
ListInt32 struct {
	elements []int32
}

type

// A ListFloat32 is a list of elements
ListFloat32 struct {
	elements []float32
}

// NewListInt32 creates a new List
func NewListInt32() *ListInt32 {
	return &ListInt32{}
}

// NewListFloat32 creates a new List
func NewListFloat32() *ListFloat32 {
	return &ListFloat32{}
}

// Append appends elements to the list
func (l *ListInt32) Append(els ...int32) {
	l.elements = append(l.elements, els...)
}

// Append appends elements to the list
func (l *ListFloat32) Append(els ...float32) {
	l.elements = append(l.elements, els...)
}

// At returns the element at i
func (l *ListInt32) At(i int) int32 {
	return l.elements[i]
}

// At returns the element at i
func (l *ListFloat32) At(i int) float32 {
	return l.elements[i]
}

// Len returns the length of the list
func (l ListInt32) Len() int {
	return len(l.elements)
}

// Len returns the length of the list
func (l ListFloat32) Len() int {
	return len(l.elements)
}

// Slice slices the list
func (l *ListInt32) Slice(from, to int) *ListInt32 {
	return &ListInt32{
		elements: l.elements[from:to],
	}
}

// Slice slices the list
func (l *ListFloat32) Slice(from, to int) *ListFloat32 {
	return &ListFloat32{
		elements: l.elements[from:to],
	}
}
