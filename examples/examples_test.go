package examples

import "testing"

func TestMergeInt(t *testing.T) {
	cs := make([]<-chan int, 5)
	for i := range cs {
		c := make(chan int)
		go func() {
			defer close(c)
			c <- 1
			c <- 1
		}()
		cs[i] = c
	}

	sum := 0
	for i := range MergeInt(cs...) {
		sum += i
	}
	if sum != 5*2 {
		t.Error("expected 1 to be sent through 5 * 2 times")
	}
}

func TestList(t *testing.T) {
	l := NewListFloat32()
	for i := 0; i < 10; i++ {
		l.Append(1.5)
	}

	sum := float32(0)
	for i := 0; i < l.Len(); i++ {
		sum += l.At(i)
	}
	if sum != 1.5*10 {
		t.Error("expected list to contain all elements")
	}
}
