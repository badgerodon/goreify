package main

// A FunctionStack is a list of functions
type FunctionStack []func()

// Append appends the function to the stack
func (s *FunctionStack) Append(f func()) {
	*s = append(*s, f)
}

// Run runs the functions in reverse order
func (s *FunctionStack) Run() {
	for i := len(*s) - 1; i >= 0; i-- {
		(*s)[i]()
	}
	*s = (*s)[:0]
}
