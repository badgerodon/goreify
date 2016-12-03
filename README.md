# goreify

`goreify` implements a form of generics for go. It works by taking in a function
or type which uses dynamic functions and types and converting it into a function
or type which uses concrete (reified) types.

For example, this generic function:

    func Sum(xs []generics.T1) generics.T1 {
    	var total generics.T1
    	for _, x := range xs {
    		total = generics.Add(total, x)
    	}
    	return total
    }

Would be transformed into this reified version:

    func Sum_int(xs []int) int {
    	var total int
    	for _, x := range xs {
    		total = total + x
    	}
    	return total
    }

The original, generic version is completely valid and runnable Go code. It's
just less efficient because it relies on reflection, and a bit clunky to work
with because we can't use operators.

More examples can be found below.

## Installation

    go get github.com/badgerodon/goreify

## Usage

`goreify` takes two parameters, the input type/function, and the reified types
you'd like to generate:

    goreify path/to/your/package.YourFunction T1Type T2Type ...

For example:

    goreify github.com/badgerodon/goreify/examples.Sum numeric

The special `numeric` pseudo-type is replaced with integers, floats and complex
types.

In practice you should use `go:generate` lines in your code:

    //go:generate goreify github.com/badgerodon/goreify/examples.Sum numeric

## Supported Features

Global functions:

    //go:generate goreify github.com/badgerodon/goreify/examples.Diff int

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

    // yielding

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

Types:

    //go:generate goreify github.com/badgerodon/goreify/examples.Pair int8 int16

    // A Pair is a pair of values
    type Pair struct {
    	// Fst is the first value
    	Fst generics.T1
    	// Snd is the second value
    	Snd generics.T2
    }

    // yielding

    type Pair_int8_int16 struct {
    	// Fst is the first value
    	Fst int8
    	// Snd is the second value
    	Snd int16
    }

Methods on types:

    func (l *List) Append(els ...generics.T1) {
    	l.elements = append(l.elements, els...)
    }

    // yielding

    func (l *List_int32) Append(els ...int32) {
    	l.elements = append(l.elements, els...)
    }

Global functions using a type:

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

I haven't investigated Channels yet.

## Other Libraries

I found [`gengen`](https://github.com/joeshaw/gengen) after implementing most of
this functionality. It follows a similar approach, though doesn't take it quite
as far.
