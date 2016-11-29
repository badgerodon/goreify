package examples

func Sum_int(xs []int) int {
	var total int
	for _, x := range xs {
		total = total + x
	}
	return total
}

func Sum_int8(xs []int8) int8 {
	var total int8
	for _, x := range xs {
		total = total + x
	}
	return total
}

func Sum_int16(xs []int16) int16 {
	var total int16
	for _, x := range xs {
		total = total + x
	}
	return total
}

func Sum_int32(xs []int32) int32 {
	var total int32
	for _, x := range xs {
		total = total + x
	}
	return total
}

func Sum_int64(xs []int64) int64 {
	var total int64
	for _, x := range xs {
		total = total + x
	}
	return total
}

func Sum_uint(xs []uint) uint {
	var total uint
	for _, x := range xs {
		total = total + x
	}
	return total
}

func Sum_uint8(xs []uint8) uint8 {
	var total uint8
	for _, x := range xs {
		total = total + x
	}
	return total
}

func Sum_uint16(xs []uint16) uint16 {
	var total uint16
	for _, x := range xs {
		total = total + x
	}
	return total
}

func Sum_uint32(xs []uint32) uint32 {
	var total uint32
	for _, x := range xs {
		total = total + x
	}
	return total
}

func Sum_uint64(xs []uint64) uint64 {
	var total uint64
	for _, x := range xs {
		total = total + x
	}
	return total
}

func Sum_float32(xs []float32) float32 {
	var total float32
	for _, x := range xs {
		total = total + x
	}
	return total
}

func Sum_float64(xs []float64) float64 {
	var total float64
	for _, x := range xs {
		total = total + x
	}
	return total
}

func Sum_complex64(xs []complex64) complex64 {
	var total complex64
	for _, x := range xs {
		total = total + x
	}
	return total
}

func Sum_complex128(xs []complex128) complex128 {
	var total complex128
	for _, x := range xs {
		total = total + x
	}
	return total
}

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

func Diff_int8(xs, ys []int8) []int8 {
	sz := len(xs)
	if len(ys) < sz {
		sz = len(ys)
	}

	zs := make([]int8, sz)
	for i := 0; i < sz; i++ {
		zs[i] = xs[i] - ys[i]
	}
	return zs
}

func Diff_int16(xs, ys []int16) []int16 {
	sz := len(xs)
	if len(ys) < sz {
		sz = len(ys)
	}

	zs := make([]int16, sz)
	for i := 0; i < sz; i++ {
		zs[i] = xs[i] - ys[i]
	}
	return zs
}

func Diff_int32(xs, ys []int32) []int32 {
	sz := len(xs)
	if len(ys) < sz {
		sz = len(ys)
	}

	zs := make([]int32, sz)
	for i := 0; i < sz; i++ {
		zs[i] = xs[i] - ys[i]
	}
	return zs
}

func Diff_int64(xs, ys []int64) []int64 {
	sz := len(xs)
	if len(ys) < sz {
		sz = len(ys)
	}

	zs := make([]int64, sz)
	for i := 0; i < sz; i++ {
		zs[i] = xs[i] - ys[i]
	}
	return zs
}

func Diff_uint(xs, ys []uint) []uint {
	sz := len(xs)
	if len(ys) < sz {
		sz = len(ys)
	}

	zs := make([]uint, sz)
	for i := 0; i < sz; i++ {
		zs[i] = xs[i] - ys[i]
	}
	return zs
}

func Diff_uint8(xs, ys []uint8) []uint8 {
	sz := len(xs)
	if len(ys) < sz {
		sz = len(ys)
	}

	zs := make([]uint8, sz)
	for i := 0; i < sz; i++ {
		zs[i] = xs[i] - ys[i]
	}
	return zs
}

func Diff_uint16(xs, ys []uint16) []uint16 {
	sz := len(xs)
	if len(ys) < sz {
		sz = len(ys)
	}

	zs := make([]uint16, sz)
	for i := 0; i < sz; i++ {
		zs[i] = xs[i] - ys[i]
	}
	return zs
}

func Diff_uint32(xs, ys []uint32) []uint32 {
	sz := len(xs)
	if len(ys) < sz {
		sz = len(ys)
	}

	zs := make([]uint32, sz)
	for i := 0; i < sz; i++ {
		zs[i] = xs[i] - ys[i]
	}
	return zs
}

func Diff_uint64(xs, ys []uint64) []uint64 {
	sz := len(xs)
	if len(ys) < sz {
		sz = len(ys)
	}

	zs := make([]uint64, sz)
	for i := 0; i < sz; i++ {
		zs[i] = xs[i] - ys[i]
	}
	return zs
}

func Diff_float32(xs, ys []float32) []float32 {
	sz := len(xs)
	if len(ys) < sz {
		sz = len(ys)
	}

	zs := make([]float32, sz)
	for i := 0; i < sz; i++ {
		zs[i] = xs[i] - ys[i]
	}
	return zs
}

func Diff_float64(xs, ys []float64) []float64 {
	sz := len(xs)
	if len(ys) < sz {
		sz = len(ys)
	}

	zs := make([]float64, sz)
	for i := 0; i < sz; i++ {
		zs[i] = xs[i] - ys[i]
	}
	return zs
}

func Diff_complex64(xs, ys []complex64) []complex64 {
	sz := len(xs)
	if len(ys) < sz {
		sz = len(ys)
	}

	zs := make([]complex64, sz)
	for i := 0; i < sz; i++ {
		zs[i] = xs[i] - ys[i]
	}
	return zs
}

func Diff_complex128(xs, ys []complex128) []complex128 {
	sz := len(xs)
	if len(ys) < sz {
		sz = len(ys)
	}

	zs := make([]complex128, sz)
	for i := 0; i < sz; i++ {
		zs[i] = xs[i] - ys[i]
	}
	return zs
}
