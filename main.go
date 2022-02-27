package goarea

import "math"

/*
Pi is the ratio of the circumference of any
circle to the diameter of that circle.
*/
const Pi = 3.1416

/* Circle calculates the area of a circle. */
func Circle(radius float64) float64 {
	return math.Pow(radius, 2) * Pi
}

/* Rectangle calculates the area of a rectangle. */
func Rectangle(base, height float64) float64 {
	return base * height
}

// func _EqTriangle(base, height float64) float64 {
// 	return (base * height) / 2
// }
