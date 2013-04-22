// Package vector provides N-dimensional geometric vectors, where N is known at
// compile time.
package vector

import "math"

// TODO: Documentation
const Dimensionality = 3

// TODO: Documentation
type Vec [Dimensionality]float64

// Computes the dot product of two vectors. The dot product is defined as
// v[0]*o[0] + v[1]*o[1] + ... + v[n-1]*o[n-1].
func (v *Vec) Dot(o *Vec) float64 {
	var result float64

	for i, f := range *v {
		result += f * (*o)[i]
	}

	return result
}

// Adds the components of vector o to vector v, and returns a new vector
// containing the result. See AddInPlace for a more efficient method.
func (v *Vec) Add(o *Vec) *Vec {
	var result Vec

	for i, f := range *v {
		result[i] = f + (*o)[i]
	}

	return &result
}

// Adds the components of vector o to vector v, modifying the original vector.
// This function returns v for convenience.
func (v *Vec) AddInPlace(o *Vec) *Vec {
	for i, f := range *o {
		(*v)[i] += f
	}

	return v
}

// Subtracts the components of vector o from vector v and returns a new vector
// containing the result. See SubInPlace for a more efficient method.
func (v *Vec) Sub(o *Vec) *Vec {
	var result Vec

	for i, f := range *v {
		result[i] = f - (*o)[i]
	}

	return &result
}

// Subtracts the components of vector o from vector v, modifying the original
// vector. This function returns v for convenience.
func (v *Vec) SubInPlace(o *Vec) *Vec {
	for i, f := range *o {
		(*v)[i] -= f
	}

	return v
}

// Computes the magnitude of a vector. Magnitude is defined as the square root
// of the sum of the squares of the components of the vector. Use LenSq if you
// are comparing two magnitudes, as it does not require a square root.
func (v *Vec) Len() float64 {
	return math.Sqrt(v.LenSq())
}

// Computes the sum of the squares of the components of the vector. The square
// root of this value is the magnitude of the vector.
func (v *Vec) LenSq() float64 {
	var result float64

	for _, f := range *v {
		result += f * f
	}

	return result
}

// Computes the distance between two vectors. Distance is defined as the
// magnitude of the difference of the two vectors. Use DistSq if you are
// comparing two distances.
func (v *Vec) Dist(o *Vec) float64 {
	return math.Sqrt(v.DistSq(o))
}

// Computes the square of the distance between two vectors. Distance is defined
// as the magnitude of the difference of the two vectors.
func (v *Vec) DistSq(o *Vec) float64 {
	var result float64

	for i, f := range *v {
		f -= (*o)[i]
		result += f * f
	}

	return result
}

// Returns a normalized copy of the vector. That is, a vector with the same
// direction but with a magnitude of 1. Behavior is undefined on vectors with
// a magnitude of 0. NormInPlace is a cheaper alternative if the original vector
// is no longer needed.
func (v *Vec) Norm() *Vec {
	var result Vec

	m := v.Len()
	for i, f := range *v {
		result[i] = f / m
	}

	return &result
}

// Normalizes the vector in place and returns it for convenience. That is, it
// changes v to have a magnitude of 1 while keeping the direction constant.
// Behavior is undefined on vectors with a magnitude of 0.
func (v *Vec) NormInPlace() *Vec {
	m := v.Len()
	for i := range *v {
		(*v)[i] /= m
	}

	return v
}
