package examples

// Diff_int finds the difference between two series
// reify:
//   types:
//     T1: numeric
func Diff_int(xs, ys []int) []int {
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
