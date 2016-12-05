package examples

// DiffInt finds the difference between two series
func DiffInt(xs, ys []int) []int {
	sz := len(xs)
	if len(ys) < sz {
		sz = len(ys)
	}

	zs := make([]int, sz)
	for i := 0; i < sz; i++ {
		zs[i] = xs[i] - ys[i]
	}
	return zs
}
