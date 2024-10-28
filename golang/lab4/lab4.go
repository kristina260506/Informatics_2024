package lab4

import (
	"math"
)
func Colcul(a, b, x float64) float64 {
	tanBx := math.Tan(b * x)
	cotAx := 1 / math.Tan(a * x)

	y := a + (tanBx*tanBx)/b + (cotAx*cotAx)

	return y
}