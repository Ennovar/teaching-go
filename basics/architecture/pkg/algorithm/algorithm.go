// This package contains various algorithms
// that assist the functionality of this
// program.
package algorithm

import "math"

const (
	pi = 3.14
	tau = 6.28
)

func AreaOfACircle(radius float64) float64 {
	return pi * math.Pow(radius, 2)
}