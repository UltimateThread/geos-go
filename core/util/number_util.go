package geos

import "math"

func EqualsWithTolerance(x1 float64, x2 float64, tolerance float64) bool {
	return math.Abs(x1-x2) <= tolerance
}
