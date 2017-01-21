package examples

type

// A PairInt8Int16 is a pair of values
PairInt8Int16 struct {
	// Fst is the first value
	Fst int8
	// Snd is the second value
	Snd int16
}

// ZipInt8Int16 takes in two slices and converts them into a slice of pairs
func ZipInt8Int16(xs []int8, ys []int16) []PairInt8Int16 {
	mn := len(xs)
	if mn > len(ys) {
		mn = len(ys)
	}
	zs := make([]PairInt8Int16, mn)
	for i := 0; i < mn; i++ {
		zs[i] = PairInt8Int16{Fst: xs[i], Snd: ys[i]}
	}
	return zs
}
