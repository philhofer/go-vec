package vec

import "math"
import "sort"

/*
Discreet Data Operations




*/

//Struct for bivariate data (xs, ys)
type BiVariateData struct {
	Xs []float64
	Ys []float64
	isSorted bool
}

func MakeBiVariateData(xs []float64, ys []float64) BiVariateData {
	out := BiVariateData{xs, ys, false}
	out.Sort()
	return out
}

func (b BiVariateData) Len() int{
	return len(b.Xs)
}

func (b BiVariateData) Less(i, j int) bool {
	if b.Xs[i] < b.Xs[j] {
		return true
	} else {
		return false
	}
}

func (b BiVariateData) Swap(i, j int) {
	b.Xs[i], b.Xs[j] = b.Xs[j], b.Xs[i]
	b.Ys[i], b.Ys[j] = b.Ys[j], b.Ys[i]
	return
}

func (b BiVariateData) Sort() {
	//edge case 1
	if b.isSorted {
		return
	}

	//edge case 2
	if sort.Float64sAreSorted(b.Xs) {
		b.isSorted = true
		return
	}

	sort.Sort(b)

	b.isSorted = true
}

func (b BiVariateData) findNearestX(x float64) int {
	if !b.isSorted {b.Sort()}
	
	//edge cases
	if x > b.Xs[len(b.Xs)] { return len(b.Xs)-1 }
	if x < b.Xs[0] { return 0 }

	out := 0
	for b.Xs[out] < x {
		out++
	}
	
	if math.Abs(b.Xs[out]-x) > math.Abs(b.Xs[out+1]) {
		out++
	}

	return out
	
}

/* 
func CubicSpline(data BiVariateData) Interpolation {}

Returns a CubicSplineInterpolation interface on 'data'



*/

/*
func DiscreetConvolve(dat []float64, conv []float64) []float64 {}

Returns a discreet convolution of 'conv' on 'dat'


*/
