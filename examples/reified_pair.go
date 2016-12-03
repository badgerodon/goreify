package examples

type Pair_int8_int16 struct {
	// Fst is the first value
	Fst int8
	// Snd is the second value
	Snd int16
}

// Zip_int8_int16 takes in two slices and converts them into a slice of pairs
func Zip_int8_int16(xs []int8, ys []int16) []Pair_int8_int16 {
	mn := len(xs)
	if mn > len(ys) {
		mn = len(ys)
	}
	zs := make([]Pair_int8_int16, mn)
	for i := 0; i < mn; i++ {
		zs[i] = Pair_int8_int16{Fst: xs[i], Snd: ys[i]}
	}
	return zs
}
