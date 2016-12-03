package examples

import "github.com/badgerodon/goreify/generics"

// A Pair is a pair of values
type Pair struct {
	// Fst is the first value
	Fst generics.T1
	// Snd is the second value
	Snd generics.T2
}

// Zip takes in two slices and converts them into a slice of pairs
func Zip(xs []generics.T1, ys []generics.T2) []Pair {
	mn := len(xs)
	if mn > len(ys) {
		mn = len(ys)
	}
	zs := make([]Pair, mn)
	for i := 0; i < mn; i++ {
		zs[i] = Pair{Fst: xs[i], Snd: ys[i]}
	}
	return zs
}
