package examples

import "github.com/badgerodon/goreify/generics"

// Sum adds numbers
func Sum(xs []generics.T1) generics.T1 {
	var total generics.T1
	for _, x := range xs {
		total = generics.Add(total, x)
	}
	return total
}

// Diff finds the difference between two series
func Diff(xs, ys []generics.T1) []generics.T1 {
	sz := len(xs)
	if len(ys) < sz {
		sz = len(ys)
	}

	zs := make([]generics.T1, sz)
	for i := 0; i < sz; i++ {
		zs[i] = generics.Subtract(xs[i], ys[i])
	}
	return zs
}
