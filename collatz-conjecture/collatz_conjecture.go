// Package collatzconjecture provides a function to return the number of steps it
// takes to reach 1 according to the rules of the Collatz Conjecture.
package collatzconjecture

import "errors"

// CollatzConjecture return the number of steps it takes to reach 1
// according to the rules of the Collatz Conjecture.
func CollatzConjecture(n int) (int, error) {
	if n <= 0 {
		return 0, errors.New("input must be a positive integer")
	}

	steps := 0

	for n != 1 {
		if n%2 == 0 {
			n /= 2
		} else {
			n = n*3 + 1
		}

		steps++
	}

	return steps, nil
}
