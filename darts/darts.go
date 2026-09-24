// Package darts provides functions for calculating dart scores.
package darts

// IsPointInCircle reports whether a point is inside a circle
// centered at the origin.
func IsPointInCircle(x, y float64, radius float64) bool {
	return x*x+y*y <= radius*radius
}

// Score calculates the score for a dart at the given coordinates.
func Score(x, y float64) int {
	if IsPointInCircle(x, y, 1) {
		return 10
	}

	if IsPointInCircle(x, y, 5) {
		return 5
	}

	if IsPointInCircle(x, y, 10) {
		return 1
	}

	return 0
}
